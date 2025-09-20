package i18nx

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"google.golang.org/grpc/metadata"

	_ "embed"
)

var (
	HeaderSelectedLanguage = "x-selected-language"
	HeaderAcceptLanguage   = "accept-language"
)

var AllowHeaders = []string{
	http.CanonicalHeaderKey(HeaderSelectedLanguage),
}

var FallbackTag = language.English

type I18N struct {
	logger     *slog.Logger
	tmplOption string
	cl         *catalog.Builder
	matcher    language.Matcher
	languages  []language.Tag
	printers   sync.Map // key: language.Tag, value: *message.Printer
	templates  sync.Map // key: tmplKey, value: *template.Template
}

type tmplKey struct {
	tag language.Tag
	key message.Reference
}

//go:embed embed/default.csv
var defaultCatalogCSV string

// New creates a new I18N instance with the default catalog. If the override
// io.Reader is not nil, it is used to override the default catalog.
func New(overrides ...io.Reader) (*I18N, error) {
	overrides = append([]io.Reader{strings.NewReader(defaultCatalogCSV)}, overrides...)

	cl := catalog.NewBuilder(catalog.Fallback(FallbackTag))
	for _, override := range overrides {
		msgs, err := parseCSV(override)
		if err != nil {
			return nil, err
		}
		for _, msg := range msgs {
			if err := cl.SetString(msg.tag, msg.key, msg.value); err != nil {
				return nil, errors.Wrapf(err, "failed to set message %q for language %q", msg.key, msg.tag)
			}
		}
	}

	languages := cl.Languages()
	matcher := language.NewMatcher(languages, language.PreferSameScript(true))
	return &I18N{
		cl:         cl,
		matcher:    matcher,
		languages:  languages,
		logger:     slog.Default(),
		tmplOption: "missingkey=default",
	}, nil
}

func (b *I18N) WithLogger(logger *slog.Logger) *I18N {
	b.logger = logger
	return b
}

// WithTemplateOption sets the text/template option string used when parsing templates.
// Example: "missingkey=zero", "missingkey=default", "missingkey=error", or any valid text/template option.
// Returns receiver for chaining.
func (b *I18N) WithTemplateOption(opt string) *I18N {
	b.tmplOption = opt
	return b
}

// MatchStrings returns the best matching language tag for the given language strings.
func (b *I18N) MatchStrings(lang ...string) language.Tag {
	_, index := language.MatchStrings(b.matcher, lang...)
	return b.languages[index]
}

// getPrinter returns a cached printer for the language tag.
func (b *I18N) getPrinter(tag language.Tag) *message.Printer {
	if v, ok := b.printers.Load(tag); ok {
		return v.(*message.Printer)
	}
	p := message.NewPrinter(tag, message.Catalog(b.cl))
	b.printers.Store(tag, p)
	return p
}

// Sprintf prints a localized message.
// - Positional args: behaves like fmt.Sprintf through x/text/message.
// - Named arg (single map or struct): render with text/template using the localized message as template.
func (b *I18N) Sprintf(tag language.Tag, key message.Reference, args ...any) (xres string) {
	defer func() {
		if xres == "" && b.logger != nil {
			b.logger.Warn("i18n message is empty", "tag", tag.String(), "key", key, "args", args)
		}
	}()
	// Named-arg via map or struct: try gotpl rendering; fallback to positional on error.
	if len(args) == 1 && (isMap(args[0]) || isStruct(args[0])) {
		if out, err := b.renderTemplate(tag, key, args[0]); err == nil {
			return out
		} else if b.logger != nil {
			b.logger.Debug("i18n template render failed; fallback",
				"tag", tag.String(), "keyType", fmt.Sprintf("%T", key), "err", err)
		}
	}
	return b.getPrinter(tag).Sprintf(key, args...)
}

// renderTemplate renders using any message.Reference key by first resolving its localized text as template source.
func (b *I18N) renderTemplate(tag language.Tag, key message.Reference, data any) (string, error) {
	// Use (tag, key) directly as cache key since message.Reference is comparable
	cacheKey := tmplKey{tag: tag, key: key}

	var tmpl *template.Template
	if v, ok := b.templates.Load(cacheKey); ok {
		tmpl = v.(*template.Template)
	} else {
		// Get the actual template source for compilation
		src := b.getPrinter(tag).Sprintf(key)
		if src == "" {
			return "", errors.New("empty template source")
		}

		var err error
		tmpl, err = template.New("tmpl").Option(b.tmplOption).Parse(src)
		if err != nil {
			return "", err
		}
		b.templates.Store(cacheKey, tmpl)
	}

	var sb strings.Builder
	if err := tmpl.Execute(&sb, data); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (b *I18N) LanguageFromContext(ctx context.Context) language.Tag {
	var selected string
	var accept string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		selected = strings.Join(md.Get(HeaderSelectedLanguage), ",")
		accept = strings.Join(md.Get(HeaderAcceptLanguage), ",")
	}
	return b.MatchStrings(selected, accept)
}

type csvMessage struct {
	tag   language.Tag
	key   string
	value string
}

func parseCSV(r io.Reader) ([]*csvMessage, error) {
	reader := csv.NewReader(r)
	reader.LazyQuotes = true

	headers, err := reader.Read()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(headers) < 2 || headers[0] != "key" {
		return nil, errors.New("CSV header must start with 'key' followed by language codes")
	}

	langTags := make([]language.Tag, len(headers)-1)
	for i := 1; i < len(headers); i++ {
		langTag, err := language.Parse(headers[i])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid language code %q in CSV header", headers[i])
		}
		langTags[i-1] = langTag
	}

	var messages []*csvMessage
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.Wrap(err, "failed to read CSV record")
		}

		key := strings.TrimSpace(record[0])
		for i := 1; i < len(record); i++ {
			text := record[i]
			messages = append(messages, &csvMessage{
				tag:   langTags[i-1],
				key:   key,
				value: text,
			})
		}
	}
	return messages, nil
}

func isMap(v any) bool {
	_, ok := v.(map[string]any)
	return ok
}

func isStruct(v any) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	return rv.IsValid() && rv.Kind() == reflect.Struct
}

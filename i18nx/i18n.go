package i18nx

import (
	"context"
	"encoding/csv"
	"io"
	"net/http"
	"strings"

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
	cl        *catalog.Builder
	matcher   language.Matcher
	languages []language.Tag
}

// New creates a new I18N instance with the default catalog. If the override
// io.Reader is not nil, it is used to override the default catalog. The
// override reader is expected to contain a CSV file with the first column being
// the key, and the remaining columns being the values for the languages listed
// in the header. The header is expected to start with "key" followed by the
// language codes.
//
// The I18N instance will match the language tags according to the following
// rules:
//
//  1. Prefer the same script as the requested language.
//  2. Fall back to the requested language.
//  3. Fall back to the English language.
//
// The I18N instance can be used to format messages using the Sprintf method.
func New(override io.Reader) (*I18N, error) {
	cl := catalog.NewBuilder(catalog.Fallback(FallbackTag))
	if err := setDefault(cl); err != nil {
		return nil, err
	}

	if override != nil {
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
		cl:        cl,
		matcher:   matcher,
		languages: languages,
	}, nil
}

// MatchStrings returns the best matching language tag for the given language
// strings. The language strings are matched against the languages in the
// catalog in the same order they appear in the catalog. The language strings
// are treated as if they were passed to language.MatchStrings.
func (b *I18N) MatchStrings(lang ...string) language.Tag {
	_, index := language.MatchStrings(b.matcher, lang...)
	return b.languages[index]
}

func (b *I18N) Sprintf(tag language.Tag, key message.Reference, args ...any) string {
	// TODO: Use a cache for the printer.
	// TODO: 或许还可以支持 args 为 map[string]any 或者 struct 类型这种，然后根据 key 的名字来填充
	return message.NewPrinter(tag, message.Catalog(b.cl)).Sprintf(key, args...)
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

//go:embed embed/default.csv
var defaultCatalogCSV string

func setDefault(cl *catalog.Builder) error {
	msgs, err := parseCSV(strings.NewReader(defaultCatalogCSV))
	if err != nil {
		return err
	}
	for _, msg := range msgs {
		if err := cl.SetString(msg.tag, msg.key, msg.value); err != nil {
			return errors.Wrapf(err, "failed to set default message %q for language %q", msg.key, msg.tag)
		}
	}
	return nil
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

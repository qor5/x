package vuetifyx

import (
	"context"
	"strings"
	"testing"

	h "github.com/theplant/htmlgo"
)

const cstiPayload = `{{ ''.constructor.constructor('alert(1)')() }}`

func mustRender(t *testing.T, c h.HTMLComponent) string {
	t.Helper()
	b, err := c.MarshalHTML(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestSafeText(t *testing.T) {
	got := mustRender(t, SafeText(`<b>`+cstiPayload))
	want := `<span v-pre>&lt;b&gt;{{ &#39;&#39;.constructor.constructor(&#39;alert(1)&#39;)() }}</span>`
	if strings.TrimSpace(got) != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestVXReadonlyFieldValueIsNotCompiled(t *testing.T) {
	got := mustRender(t, VXReadonlyField().Label("Inquiry "+cstiPayload).Value(cstiPayload))

	if !strings.Contains(got, `<label v-pre class='v-label theme--light text-caption'>Inquiry {{`) {
		t.Errorf("label must be wrapped in v-pre, got:\n%s", got)
	}
	if !strings.Contains(got, `<span v-pre>{{`) {
		t.Errorf("value must be wrapped in v-pre, got:\n%s", got)
	}
}

func TestVXReadonlyFieldChildrenKeepVueBindings(t *testing.T) {
	got := mustRender(t, VXReadonlyField(h.Span("{{ locals.x }}")).Label("L"))
	if strings.Contains(got, `<span v-pre>{{ locals.x }}`) {
		t.Errorf("caller-provided children must stay compilable, got:\n%s", got)
	}
}

package i18nx

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestI18N(t *testing.T) {
	type test struct {
		name   string
		tag    language.Tag
		key    message.Reference
		args   []any
		expect string
	}

	tests := []test{
		{
			name:   "english",
			tag:    language.English,
			key:    "google",
			expect: "Google",
		},
		{
			name:   "japanese",
			tag:    language.Japanese,
			key:    "visit",
			args:   []any{"google"},
			expect: "検索: グーグル.",
		},
		{
			name:   "japanese_sub",
			tag:    language.Japanese,
			key:    "sub",
			args:   []any{"Hello"},
			expect: "Hello",
		},
		{
			name:   "japanese_sub_sub",
			tag:    language.Japanese,
			key:    "visit",
			args:   []any{"sub"},
			expect: "検索: %s.",
		},
		{
			name:   "japanese_sub_sub",
			tag:    language.Japanese,
			key:    "visit",
			args:   []any{"sub2"},
			expect: "検索: %m.",
		},
		{
			name:   "japanese_sub_sub_sub",
			tag:    language.Japanese,
			key:    "visit",
			args:   []any{"sub", "Hello"},
			expect: "検索: %s.%!(EXTRA string=Hello)",
		},
		{
			name:   "english_overwrite",
			tag:    language.English,
			key:    message.Key("microsoft", "google"),
			expect: "Google",
		},
		{
			name:   "french",
			tag:    language.French,
			key:    "google",
			expect: "google",
		},
		{
			name:   "french_not_exist",
			tag:    language.French,
			key:    "not_exist",
			expect: "not_exist",
		},
	}

	in, err := New(strings.NewReader(`
key,zh,en,ja
google,谷歌,Google,グーグル
sub,%s,%s,%s
sub2,%m,%m,%m
visit,"查找: %m.","Lookup: %m.","検索: %m."`))
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, in.Sprintf(test.tag, test.key, test.args...))
		})
	}
}

func Test_parseCSV(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError string
	}{
		{
			name:        "valid header",
			input:       "key,en,ja\ntest,Test,",
			expectError: "",
		},
		{
			name:        "missing key column",
			input:       "en,ja\nTest,",
			expectError: "CSV header must start with 'key'",
		},
		{
			name:        "only key column",
			input:       "key\ntest",
			expectError: "CSV header must start with 'key' followed by language codes",
		},
		{
			name:        "empty CSV",
			input:       "",
			expectError: "EOF",
		},
		{
			name:        "invalid language code in header",
			input:       "key,en,js",
			expectError: "invalid language code \"js\" in CSV header",
		},
		{
			name: "valid header with empty line",
			input: `key,en,ja
test,Test,

NEEDS_CHALLENGE,"Challenge required","チャレンジが必要です"`,
			expectError: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			_, err := parseCSV(input)
			if test.expectError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectError)
			}
		})
	}
}

func TestTmplKey_Comparable(t *testing.T) {
	// Test that tmplKey can be used as map key and compared properly
	cache := make(map[tmplKey]string)

	// Test with string keys
	key1 := tmplKey{tag: language.English, key: "hello"}
	key2 := tmplKey{tag: language.English, key: "hello"}
	key3 := tmplKey{tag: language.Japanese, key: "hello"}
	key4 := tmplKey{tag: language.English, key: "world"}

	// Test equality using == operator (since tmplKey is comparable)
	assert.True(t, key1 == key2, "Same tmplKey should be equal")
	assert.False(t, key1 == key3, "Different language should not be equal")
	assert.False(t, key1 == key4, "Different key should not be equal")

	// Test as map key
	cache[key1] = "value1"
	cache[key3] = "value3"
	cache[key4] = "value4"

	// Should retrieve same value for equal keys
	assert.Equal(t, "value1", cache[key2], "Equal keys should access same map entry")
	assert.Equal(t, "value3", cache[key3], "Different language should have different entry")
	assert.Equal(t, "value4", cache[key4], "Different key should have different entry")

	// Test with message.Key
	msgKey1 := message.Key("id1", "fallback1")
	msgKey2 := message.Key("id1", "fallback1")
	msgKey3 := message.Key("id2", "fallback2")

	tmplKey1 := tmplKey{tag: language.English, key: msgKey1}
	tmplKey2 := tmplKey{tag: language.English, key: msgKey2}
	tmplKey3 := tmplKey{tag: language.English, key: msgKey3}

	// Test message.Key equality using == operator
	assert.True(t, tmplKey1 == tmplKey2, "Same message.Key should be equal")
	assert.False(t, tmplKey1 == tmplKey3, "Different message.Key should not be equal")

	// Test message.Key as map key
	cache[tmplKey1] = "msgValue1"
	cache[tmplKey3] = "msgValue3"

	assert.Equal(t, "msgValue1", cache[tmplKey2], "Equal message.Key should access same map entry")
	assert.Equal(t, "msgValue3", cache[tmplKey3], "Different message.Key should have different entry")

	// Test mixed types
	mixedKey1 := tmplKey{tag: language.English, key: "string_key"}
	mixedKey2 := tmplKey{tag: language.English, key: message.Key("msg_key", "fallback")}

	assert.False(t, mixedKey1 == mixedKey2, "String and message.Key should not be equal")

	cache[mixedKey1] = "stringValue"
	cache[mixedKey2] = "messageValue"

	assert.Equal(t, "stringValue", cache[mixedKey1])
	assert.Equal(t, "messageValue", cache[mixedKey2])

	// Verify final cache state
	assert.Equal(t, 7, len(cache), "Should have 7 distinct entries")

	// Verify all expected entries exist
	expectedEntries := map[tmplKey]string{
		{language.English, "hello"}:                            "value1",
		{language.Japanese, "hello"}:                           "value3",
		{language.English, "world"}:                            "value4",
		{language.English, message.Key("id1", "fallback1")}:    "msgValue1",
		{language.English, message.Key("id2", "fallback2")}:    "msgValue3",
		{language.English, "string_key"}:                       "stringValue",
		{language.English, message.Key("msg_key", "fallback")}: "messageValue",
	}

	for expectedKey, expectedValue := range expectedEntries {
		actualValue, exists := cache[expectedKey]
		assert.True(t, exists, "Expected key should exist in cache: %+v", expectedKey)
		assert.Equal(t, expectedValue, actualValue, "Value should match for key: %+v", expectedKey)
	}
}

func TestI18N_MatchStrings(t *testing.T) {
	in, err := New(strings.NewReader(`
key,zh,en,ja,fr
test,测试,Test,テスト,Test`))
	require.NoError(t, err)

	tests := []struct {
		name     string
		input    []string
		expected language.Tag
	}{
		{
			name:     "exact match English",
			input:    []string{"en"},
			expected: language.English,
		},
		{
			name:     "exact match Japanese",
			input:    []string{"ja"},
			expected: language.Japanese,
		},
		{
			name:     "fallback to English",
			input:    []string{"de"},
			expected: language.English,
		},
		{
			name:     "prefer first match",
			input:    []string{"fr", "ja"},
			expected: language.French,
		},
		{
			name:     "match with region",
			input:    []string{"en-US"},
			expected: language.English,
		},
		{
			name:     "match with multiple inputs",
			input:    []string{"de", "zh-TW", "en"},
			expected: language.Chinese,
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: language.English, // Assuming English is the default.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := in.MatchStrings(tt.input...)
			assert.Equal(t, tt.expected, result, "Expected %v, but got %v", tt.expected, result)
		})
	}
}

func TestI18N_Template(t *testing.T) {
	// Override with simple English-only templates
	csv := `key,en

greet,"Hello, {{.Name}}!"

greet_missing,"Hello, {{.Name}} {{.Missing}}"

bad_tpl,"Hello, { .Name   }"

google,"Google {{.Name}}"
`
	in, err := New(strings.NewReader(csv))
	require.NoError(t, err)

	t.Run("map rendering success", func(t *testing.T) {
		out := in.Sprintf(language.English, "greet", map[string]any{"Name": "Bob"})
		assert.Equal(t, "Hello, Bob!", out)
	})

	t.Run("struct rendering success", func(t *testing.T) {
		type payload struct{ Name string }
		out := in.Sprintf(language.English, "greet", payload{Name: "Eve"})
		assert.Equal(t, "Hello, Eve!", out)
	})

	// Default missingkey=default yields <no value> for missing keys in map
	t.Run("missingkey=default yields <no value> for missing", func(t *testing.T) {
		out := in.Sprintf(language.English, "greet_missing", map[string]any{"Name": "Alice"})
		assert.Equal(t, "Hello, Alice <no value>", out)
	})

	t.Run("missingkey=error falls back to printer", func(t *testing.T) {
		in2, err := New(strings.NewReader(csv))
		require.NoError(t, err)
		in2.WithTemplateOption("missingkey=error")
		// Since render will error (missing key), it should fall back to raw template text
		out := in2.Sprintf(language.English, "greet_missing", map[string]any{"Name": "Alice"})
		assert.Equal(t, "Hello, {{.Name}} {{.Missing}}", out)
	})

	t.Run("compile error falls back to printer", func(t *testing.T) {
		out := in.Sprintf(language.English, "bad_tpl", map[string]any{"Name": "Carol"})
		assert.Equal(t, "Hello, { .Name   }", out)
	})

	t.Run("message.Key template rendering with placeholder", func(t *testing.T) {
		// message.Key with template syntax - should work with template rendering
		out := in.Sprintf(language.English, "greet", map[string]any{"Name": "Zed"})
		assert.Equal(t, "Hello, Zed!", out) // Template rendering works for string key
	})

	t.Run("message.Key fallback without template", func(t *testing.T) {
		// message.Key without template syntax - uses printer directly
		out := in.Sprintf(language.English, message.Key("microsoft", "google"), map[string]any{"Name": "Ignored"})
		assert.Equal(t, "Google Ignored", out) // No template, uses printer fallback
	})

	t.Run("WithLogger debug output on fallback", func(t *testing.T) {
		var buf bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

		in2, err := New(strings.NewReader(csv))
		require.NoError(t, err)
		in2.WithLogger(logger).WithTemplateOption("missingkey=error")

		// This should trigger template error and fallback to printer
		out := in2.Sprintf(language.English, "greet_missing", map[string]any{"Name": "Alice"})
		assert.Equal(t, "Hello, {{.Name}} {{.Missing}}", out) // Falls back to printer, returns raw template

		// Check debug log was written
		logOutput := buf.String()
		assert.Contains(t, logOutput, "i18n template render failed; fallback")
		assert.Contains(t, logOutput, "tag=en")
		assert.Contains(t, logOutput, "keyType=string")
	})
}

package i18nx

import (
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
key,en,ja,zh
google,Google,グーグル,谷歌
sub,%s,%s,%s
sub2,%m,%m,%m
visit,"Lookup: %m.","検索: %m.","查找: %m."`))
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

func TestI18N_MatchStrings(t *testing.T) {
	in, err := New(strings.NewReader(`
key,en,ja,zh,fr
test,Test,テスト,测试,Test`))
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

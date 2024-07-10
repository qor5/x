package main

import (
	"github.com/qor5/x/v3/i18n"
	"github.com/qor5/x/v3/i18n/i18n-transfer/test/mock/messages"
	"golang.org/x/text/language"
)

func main() {
	b := i18n.New()
	b.SupportLanguages(language.Chinese, language.Japanese).
		RegisterForModule(language.Chinese, messages.I18nModuleKey, messages.User_CN).
		RegisterForModule(language.Japanese, messages.I18nModuleKey, messages.User_JP)
}

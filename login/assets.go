package login

import (
	"embed"
)

//go:embed assets
var assetsFS embed.FS

var assetsPathPrefix = "/auth/assets/"
var (
	StyleCSSURL = assetsPathPrefix + "style.css"
	ZxcvbnJSURL = assetsPathPrefix + "zxcvbn.js"
)

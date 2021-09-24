module.exports = {
	runtimeCompiler: true,
	productionSourceMap: false,
	devServer: {
		port: 3080
	},
	configureWebpack: {
		output: {
			libraryExport: 'default'
		},
		externals: { vue: "Vue", vuetify: "Vuetify" },
	}
}

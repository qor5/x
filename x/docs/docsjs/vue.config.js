module.exports = {
	runtimeCompiler: true,
	productionSourceMap: false,
	devServer: {
		port: 3300
	},
	configureWebpack: {
		output: {
			libraryExport: 'default'
		},
		externals: { vue: "Vue" }
	}
}

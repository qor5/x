module.exports = {
    runtimeCompiler: true,
    productionSourceMap: false,
	devServer: {
		port: 3050
	},
	configureWebpack: {
		output: {
			libraryExport: 'default'
		},
		externals: {vue: "Vue"},
	}
}

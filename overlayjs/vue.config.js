module.exports = {
	runtimeCompiler: true,
	devServer: {
		port: 3050
	},
	configureWebpack: {
		output: {
			libraryExport: 'default'
		},
	}
}

// @snippet_begin(TipTapVueConfig)
module.exports = {
	runtimeCompiler: true,
	productionSourceMap: false,
	devServer: {
		port: 3500
	},
	configureWebpack: {
		output: {
			libraryExport: 'default'
		},
		externals: { vue: "Vue" },
	},
	chainWebpack: config => {
		const svgRule = config.module.rule('svg')
		svgRule.uses.clear()
		svgRule
			.use('vue-svg-loader')
			.loader('vue-svg-loader')
	}
}
// @snippet_end

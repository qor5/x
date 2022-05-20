// @snippet_begin(TipTapVueConfig)
const {defineConfig} = require('@vue/cli-service');
module.exports = defineConfig({
	transpileDependencies: true,
	runtimeCompiler: true,
	productionSourceMap: false,
	devServer: {
		port: 3500,
	},
	configureWebpack: {
		output: {
			libraryExport: 'default',
		},
		externals: {vue: 'Vue'},
	},
	chainWebpack: config => {
		const svgRule = config.module.rule('svg').clear();
		svgRule.
				test(/\.(svg)(\?.*)?$/).
				use('babel-loader').
				loader('babel-loader').
				end().
				use('vue-svg-loader').
				loader('vue-svg-loader');
	},
});
// @snippet_end

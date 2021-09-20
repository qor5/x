require('esbuild').buildSync({
  entryPoints: ['src/index.js'],
  bundle: true,
  loader: {},
  minify: true,
  sourcemap: true,
  target: ['es2020'],
  splitting: true,
  outdir: 'public',
  format: 'esm',
})

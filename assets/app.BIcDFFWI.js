import{R as p}from"./chunks/theme.Drj7qB1V.js";import{P as s,as as i,at as u,au as c,av as l,aw as f,ax as d,ay as m,az as h,aA as A,aB as g,d as v,u as w,n as y,v as C,aC as P,aD as R,aE as E,a4 as b}from"./chunks/framework.CAvCEoxM.js";function r(e){if(e.extends){const a=r(e.extends);return{...a,...e,async enhanceApp(t){a.enhanceApp&&await a.enhanceApp(t),e.enhanceApp&&await e.enhanceApp(t)}}}return e}const n=r(p),S=v({name:"VitePressApp",setup(){const{site:e,lang:a,dir:t}=w();return y(()=>{C(()=>{document.documentElement.lang=a.value,document.documentElement.dir=t.value})}),e.value.router.prefetchLinks&&P(),R(),E(),n.setup&&n.setup(),()=>b(n.Layout)}});async function T(){globalThis.__VITEPRESS__=!0;const e=x(),a=D();a.provide(u,e);const t=c(e.route);return a.provide(l,t),a.component("Content",f),a.component("ClientOnly",d),Object.defineProperties(a.config.globalProperties,{$frontmatter:{get(){return t.frontmatter.value}},$params:{get(){return t.page.value.params}}}),n.enhanceApp&&await n.enhanceApp({app:a,router:e,siteData:m}),{app:a,router:e,data:t}}function D(){return g(S)}function x(){let e=s;return h(a=>{let t=A(a),o=null;return t&&(e&&(t=t.replace(/\.js$/,".lean.js")),o=import(t)),s&&(e=!1),o},n.NotFound)}s&&T().then(({app:e,router:a,data:t})=>{a.go().then(()=>{i(a.route,t.site),e.mount("#app")})});export{T as createApp};

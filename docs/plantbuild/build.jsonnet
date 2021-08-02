local c = import 'dc.jsonnet';
local dc = c {
  dockerRegistry: 'gcr.io',
};

dc.build_apps_image('sunfmin/sunfmin', [
  {name: 'goplaid-docs', dockerfile: './docs/Dockerfile', context: '../'},
])

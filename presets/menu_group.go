package presets

type MenuGroupBuilder struct {
	name   string
	icon   string
	models []*ModelBuilder
}

func (b *MenuGroupBuilder) Icon(v string) (r *MenuGroupBuilder) {
	b.icon = v
	return b
}

func (b *MenuGroupBuilder) AppendModels(ms ...*ModelBuilder) (r *MenuGroupBuilder) {
	b.models = append(b.models, ms...)
	return b
}

func (b *MenuGroupBuilder) Models(ms ...*ModelBuilder) (r *MenuGroupBuilder) {
	b.models = ms
	return b
}

type MenuGroups struct {
	menuGroups []*MenuGroupBuilder
}

func (g *MenuGroups) MenuGroup(name string) (r *MenuGroupBuilder) {
	for _, mg := range g.menuGroups {
		if mg.name == name {
			return mg
		}
	}
	r = &MenuGroupBuilder{name: name}
	g.menuGroups = append(g.menuGroups, r)
	return
}

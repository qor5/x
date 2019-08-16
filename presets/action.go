package presets

type BulkActionBuilder struct {
	NameLabel
	updateFunc BulkActionUpdateFunc
	compFunc   BulkComponentFunc
}

func (b *ListingBuilder) BulkAction(name string) (r *BulkActionBuilder) {
	builder := b.getBulkAction(name)
	if builder != nil {
		return builder
	}
	r = &BulkActionBuilder{}
	r.name = name
	b.bulkActions = append(b.bulkActions, r)
	return
}

func (b *ListingBuilder) getBulkAction(name string) *BulkActionBuilder {
	for _, f := range b.bulkActions {
		if f.name == name {
			return f
		}
	}
	return nil
}

func (b *BulkActionBuilder) UpdateFunc(v BulkActionUpdateFunc) (r *BulkActionBuilder) {
	b.updateFunc = v
	return b
}

func (b *BulkActionBuilder) Label(v string) (r *BulkActionBuilder) {
	b.label = v
	return b
}

func (b *BulkActionBuilder) ComponentFunc(v BulkComponentFunc) (r *BulkActionBuilder) {
	b.compFunc = v
	return b
}

type ActionBuilder struct {
	name       string
	updateFunc UpdateFunc
	compFunc   FieldComponentFunc
}

func (b *DetailingBuilder) Action(name string) (r *ActionBuilder) {
	for _, f := range b.actions {
		if f.name == name {
			return f
		}
	}
	r = &ActionBuilder{name: name}
	b.actions = append(b.actions, r)
	return
}

func (b *ActionBuilder) UpdateFunc(v UpdateFunc) (r *ActionBuilder) {
	b.updateFunc = v
	return b
}

func (b *ActionBuilder) ComponentFunc(v FieldComponentFunc) (r *ActionBuilder) {
	b.compFunc = v
	return b
}

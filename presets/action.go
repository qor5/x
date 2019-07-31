package presets

type BulkActionBuilder struct {
	name       string
	updateFunc BulkActionUpdateFunc
	compFunc   ComponentFunc
}

func (b *ListingBuilder) BulkAction(name string) (r *BulkActionBuilder) {
	for _, f := range b.bulkActions {
		if f.name == name {
			return f
		}
	}
	r = &BulkActionBuilder{name: name}
	b.bulkActions = append(b.bulkActions, r)
	return
}

func (b *BulkActionBuilder) UpdateFunc(v BulkActionUpdateFunc) (r *BulkActionBuilder) {
	b.updateFunc = v
	return b
}

func (b *BulkActionBuilder) ComponentFunc(v ComponentFunc) (r *BulkActionBuilder) {
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

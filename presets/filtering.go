package presets

func (b *ListingBuilder) FilterDataFunc(v FilterDataFunc) {
	b.filterDataFunc = v
}

func (b *ListingBuilder) FilterTabsFunc(v FilterTabsFunc) {
	b.filterTabsFunc = v
}

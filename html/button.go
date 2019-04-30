package html

type ButtonBuilder struct {
	HTMLTagBuilder
}

func Button(label string) (r *ButtonBuilder) {
	tag := Tag("button").Text(label)
	r = &ButtonBuilder{
		HTMLTagBuilder: *tag,
	}
	return
}

func (b *ButtonBuilder) Type(v string) (r *ButtonBuilder) {
	b.Attr("type", v)
	return b
}

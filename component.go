package pagui

type Component interface {
}

type SchemaComponent interface {
	MarshalSchema(ctx *EventContext) ([]byte, error)
}

type HTMLComponent interface {
	MarshalHTML(ctx *EventContext) ([]byte, error)
}

type RawHTML string

func (s RawHTML) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type RawSchema string

func (s RawSchema) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type ComponentFunc func(ctx *EventContext) (r []byte, err error)

func (f ComponentFunc) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

func (f ComponentFunc) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

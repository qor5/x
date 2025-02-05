package vuetifyx

import (
	"context"
	"fmt"
	"strings"

	"github.com/samber/lo"
	h "github.com/theplant/htmlgo"
)

type VXTiptapEditorBuilder struct {
	tag        *h.HTMLTagBuilder
	extensions []*VXTiptapEditorExtension
}

func VXTiptapEditor() (r *VXTiptapEditorBuilder) {
	r = &VXTiptapEditorBuilder{
		tag: h.Tag("vx-tiptap-editor"),
	}
	return
}

type VXTiptapEditorExtension struct {
	Name    string         `json:"name"`
	Options map[string]any `json:"options"`
}

func (b *VXTiptapEditorBuilder) Extensions(v []*VXTiptapEditorExtension) (r *VXTiptapEditorBuilder) {
	b.extensions = v
	return b
}

func (b *VXTiptapEditorBuilder) Label(v string) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":label", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) Attr(vs ...any) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTiptapEditorBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTiptapEditorBuilder) Disabled(v bool) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXTiptapEditorBuilder) Readonly(v bool) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXTiptapEditorBuilder) Value(v string) (r *VXTiptapEditorBuilder) {
	b.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) MarkdownTheme(v string) (r *VXTiptapEditorBuilder) {
	b.Attr(":markdown-theme", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if len(b.extensions) > 0 {
		var imageGlueOnClick string
		imageGlue, exists := lo.Find(b.extensions, func(item *VXTiptapEditorExtension) bool {
			return item.Name == "ImageGlue"
		})
		if exists && imageGlue.Options != nil {
			var ok bool
			imageGlueOnClick, ok = imageGlue.Options["onClick"].(string)
			if !ok {
				return nil, fmt.Errorf("imageGlue.onClick is not a string")
			}
			imageGlue.Options["onClick"] = "__imageGluePlaceholder__"
		}
		jsonString := h.JSONString(b.extensions)
		if imageGlueOnClick != "" {
			jsonString = strings.ReplaceAll(jsonString, `"__imageGluePlaceholder__"`, imageGlueOnClick)
		}
		b.tag.Attr(":extensions", jsonString)
	}
	return b.tag.MarshalHTML(ctx)
}

func TiptapSlackLikeExtensions() []*VXTiptapEditorExtension {
	return []*VXTiptapEditorExtension{
		{
			Name: "BaseKit",
			Options: map[string]any{
				"placeholder": map[string]any{
					"placeholder": "Jot something down...",
				},
			},
		},
		{
			Name: "Bold",
		},
		{
			Name: "Italic",
		},
		{
			Name: "Strike",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Link",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "OrderedList",
		},
		{
			Name: "BulletList",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Blockquote",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Code",
		},
		{
			Name: "CodeBlock",
			Options: map[string]any{
				"divider": true,
			},
		},
	}
}

func TiptapExtensions() []*VXTiptapEditorExtension {
	return []*VXTiptapEditorExtension{
		{
			Name: "BaseKit",
			Options: map[string]any{
				"placeholder": map[string]any{
					"placeholder": "Enter some text...",
				},
			},
		},
		{
			Name: "Bold",
		},
		{
			Name: "Italic",
		},
		{
			Name: "Underline",
		},
		{
			Name: "Strike",
		},
		{
			Name: "Code",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Heading",
		},
		{
			Name: "TextAlign",
			Options: map[string]any{
				"types": []string{"heading", "paragraph", "image"},
			},
		},
		{
			Name: "FontFamily",
		},
		{
			Name: "FontSize",
		},
		{
			Name: "Color",
		},
		{
			Name: "Highlight",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "SubAndSuperScript",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "BulletList",
		},
		{
			Name: "OrderedList",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "TaskList",
		},
		{
			Name: "Indent",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Link",
		},
		{
			Name: "Image",
		},
		{
			Name: "Video",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Table",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Blockquote",
		},
		{
			Name: "HorizontalRule",
		},
		{
			Name: "CodeBlock",
			Options: map[string]any{
				"divider": true,
			},
		},
		{
			Name: "Clear",
		},
		{
			Name: "History",
			Options: map[string]any{
				"divider": true,
			},
		},
		// {
		// 	Name: "Fullscreen",
		// },
	}
}

# Vuetifyx

Vuetifyx is a high-level wrapper for Vuetify components, designed to better align with QOR5's UI patterns and usage conventions.

## Go Usage

```go
import (
    "github.com/qor5/x/v3/vuetifyx"
    "github.com/qor5/web"
)

// Create a button
vuetifyx.Button().
    Text("Submit").
    Color("primary").
    Attr("@click", web.Plaid().
        Query(func(ctx *web.EventContext) {
            // Handle click event
        }))

// Create a text field
vuetifyx.TextField().
    Label("Username").
    Attr("@change", web.Plaid().
        Query(func(ctx *web.EventContext) {
            // Handle input event
            value := ctx.R.FormValue("value")
        }))

// Create a select
vuetifyx.Select().
    Items([]string{"Option 1", "Option 2"}).
    Attr("@change", web.Plaid().
        Query(func(ctx *web.EventContext) {
            // Handle selection event
            selected := ctx.R.FormValue("value")
        }))
```

## Development Guide

### Prerequisites

- Go 1.22.5 or later
- Node.js 16 or later
- pnpm

### Development Commands

1. Start development server:

```bash
pnpm run dev
```

For real-time component preview with hot reload. Access via http://localhost:5173.

2. Component development mode:

```bash
pnpm run comp
```

Focused development environment for individual components:

- Auto-generates component templates
- Isolated development environment
- Real-time component preview

3. Build for production:

```bash
pnpm run build
```

Builds components for production:

- Generates ES and UMD format components
- Generates type declarations
- Optimizes and minifies code
- Outputs to `dist` directory

## Documentation

For complete component documentation, visit our [documentation site](./vuetifyxjs/docs).

## License

This project is part of QOR5 and is licensed under the MIT License.

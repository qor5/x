## The Mission

To empower Go developers to become fullstack developers without JavaScript skills. means be able to do a web application
without coding JavaScript for interactions. But have to code CSS.

JavaScript interactions are encapsulated within Go [htmlgo](https://github.com/theplant/htmlgo) components.

## Documentation

Online address: <https://docs.goplaid.dev>

Run documentation server in local: `./dev.sh`

To write a document with executable example

1. Add a code example under `docs/examples`
2. Register this example into router in `docs/mux.go`
3. Call `utils.Demo` with demo path and file path
4. Add a Slug to the `Doc` NOTE: cannot be same with other docs
5. Add a Title to the `Doc`

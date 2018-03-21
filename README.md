# structgen

Generate [Go](https://golang.org) Structs out of a given [JSON Schema](http://json-schema.org/)

This generator creates at least `map[string]*Object` fields and more objects with deeper nesting in the schema tree. Also considering the more important `PatternProperties` for e.g. a docker-compose schema.

## Status

Some kind of alpha. Issues / MRs are welcome.

### Known issues

* Resolve more `interface{}` types
* run some kind `go fmt` for the generated output

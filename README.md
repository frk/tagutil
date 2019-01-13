### tagutil



Package `tagutil` provides a couple helper methods for dealing with Go's struct tags.

This package follows the convention outlined in the Go's reflect package
documentation at: [reflect.StructTag](https://golang.org/pkg/reflect/#StructTag), ie *"By convention,
tag strings are a concatenation of optionally space-separated key:"value" pairs." [...]*

Additionally the "value" can be a comma-separated list of items, in which case
the first item is regarded as the "main" value and all of the subsequent items
are considered as "options" of that pair.

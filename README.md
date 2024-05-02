# mod2blob

This package aims to quickly generate [Bloblang](https://www.benthos.dev/docs/guides/bloblang/about/) [functions](https://www.benthos.dev/docs/guides/bloblang/functions] and [methods](https://www.benthos.dev/docs/guides/bloblang/methods] for use in [Benthos](https://www.benthos.dev). The tool, mod2blob, will accept a package argument `-package` (or env PACKAGE) that specifies the name of a golang module, such as `math` or `hbollon/go-edlib. The tool will then parse all exported functions and generate bloblang functions and methods from that module.

For example, if we want to create a bloblang plugin for the golang math package, we would run the following:

```
mod2blob -package math
```

This would generate math_functions.go and math_methods.go which an be compiled into Benthos as a plugin.

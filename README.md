# mod2blob

**This is considered pre-alpha**

## Overview
This package aims to quickly generate [Bloblang](https://www.benthos.dev/docs/guides/bloblang/about/) [functions](https://www.benthos.dev/docs/guides/bloblang/functions) and [methods](https://www.benthos.dev/docs/guides/bloblang/methods) for use in [Benthos](https://www.benthos.dev). The tool, mod2blob, will accept a package argument `-package` (or env PACKAGE) that specifies the name of a golang module, such as `math` or `hbollon/go-edlib`. The tool will then parse all exported functions and generate bloblang functions and methods from that module.

## Dependencies

* Golang (tested with 1.22.2)

## Usage

Generate code from go standard library package:
```bash
mod2blob -package strings
```

Generate code from package on github:
```bash
mod2blob -package github.com/hbollon/go-edlib
```

Note: When specifying packages from remote repositories, the module will be cloned into $GOPATH/src.  You must have GOPATH set to a location that is writable.


## Example

To generate a bloblang plugin from the math package, we would run the following command:
```go
mod2blob -package strings
```

This would generate math.go which can be compiled into Benthos as a bloblang plugin.

For example, here is the auto-generated math.go:

```go
package bloblang

import (
	"math"

	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	var err error

	objectAbsSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Abs returns the absolute value of x.
	err = bloblang.RegisterFunctionV2("abs", objectAbsSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Abs(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

      // ....code clipped
}
```

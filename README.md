# mod2blob

**This is still a work in progress**

This package aims to quickly generate [Bloblang](https://www.benthos.dev/docs/guides/bloblang/about/) [functions](https://www.benthos.dev/docs/guides/bloblang/functions) and [methods](https://www.benthos.dev/docs/guides/bloblang/methods) for use in [Benthos](https://www.benthos.dev). The tool, mod2blob, will accept a package argument `-package` (or env PACKAGE) that specifies the name of a golang module, such as `math` or `hbollon/go-edlib`. The tool will then parse all exported functions and generate bloblang functions and methods from that module.

For example, if we want to create a bloblang plugin for the golang math package, we would run the following:

```
mod2blob -package math
```

This would generate math_function.go and math_method.go which an be compiled into Benthos as a plugin.

For example, here is the auto-generated math_function.go:

```go

package bloblang

import (
	"log"

	"math"
	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	var (
		err error
	)

        
	objectAtan2Spec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("y")).Param(bloblang.NewFloat64Param("x"))

	err = bloblang.RegisterFunctionV2("Atan2", objectAtan2Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Atan2(y, x), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectCopysignSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("f")).Param(bloblang.NewFloat64Param("sign"))

	err = bloblang.RegisterFunctionV2("Copysign", objectCopysignSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			f, err := args.GetFloat64("f")
			if err != nil {
				return nil, err
			}
			
			sign, err := args.GetFloat64("sign")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Copysign(f, sign), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectDimSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Dim", objectDimSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Dim(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectFMASpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y")).Param(bloblang.NewFloat64Param("z"))

	err = bloblang.RegisterFunctionV2("FMA", objectFMASpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			
			z, err := args.GetFloat64("z")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.FMA(x, y, z), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectHypotSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("p")).Param(bloblang.NewFloat64Param("q"))

	err = bloblang.RegisterFunctionV2("Hypot", objectHypotSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			p, err := args.GetFloat64("p")
			if err != nil {
				return nil, err
			}
			
			q, err := args.GetFloat64("q")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Hypot(p, q), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectIsInfSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("f")).Param(bloblang.NewInt64Param("sign"))

	err = bloblang.RegisterFunctionV2("IsInf", objectIsInfSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			f, err := args.GetFloat64("f")
			if err != nil {
				return nil, err
			}
			
			sign, err := args.GetInt64("sign")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.IsInf(f, sign), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectLdexpSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("frac")).Param(bloblang.NewInt64Param("exp"))

	err = bloblang.RegisterFunctionV2("Ldexp", objectLdexpSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			frac, err := args.GetFloat64("frac")
			if err != nil {
				return nil, err
			}
			
			exp, err := args.GetInt64("exp")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Ldexp(frac, exp), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectMaxSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Max", objectMaxSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Max(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectMinSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Min", objectMinSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Min(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectModSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Mod", objectModSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Mod(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectNextafterSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Nextafter", objectNextafterSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Nextafter(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectPowSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Pow", objectPowSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Pow(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
	objectRemainderSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))

	err = bloblang.RegisterFunctionV2("Remainder", objectRemainderSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}
			
			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}
			return func() (interface{}, error) {
				return math.Remainder(x, y), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	
}
```

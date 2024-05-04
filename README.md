# mod2blob

**This is considered pre-alpha**

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

			ya := float64(y)

			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Atan2(ya, xa), nil
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

			fa := float64(f)

			sign, err := args.GetFloat64("sign")
			if err != nil {
				return nil, err
			}

			signa := float64(sign)
			return func() (interface{}, error) {
				return math.Copysign(fa, signa), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Dim(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)

			z, err := args.GetFloat64("z")
			if err != nil {
				return nil, err
			}

			za := float64(z)
			return func() (interface{}, error) {
				return math.FMA(xa, ya, za), nil
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

			pa := float64(p)

			q, err := args.GetFloat64("q")
			if err != nil {
				return nil, err
			}

			qa := float64(q)
			return func() (interface{}, error) {
				return math.Hypot(pa, qa), nil
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

			fa := float64(f)

			sign, err := args.GetInt64("sign")
			if err != nil {
				return nil, err
			}

			signa := int(sign)
			return func() (interface{}, error) {
				return math.IsInf(fa, signa), nil
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

			fraca := float64(frac)

			exp, err := args.GetInt64("exp")
			if err != nil {
				return nil, err
			}

			expa := int(exp)
			return func() (interface{}, error) {
				return math.Ldexp(fraca, expa), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Max(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Min(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Mod(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Nextafter(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Pow(xa, ya), nil
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

			xa := float64(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float64(y)
			return func() (interface{}, error) {
				return math.Remainder(xa, ya), nil
			}, nil
		})

	if err != nil {
		panic(err)
	}

}
```

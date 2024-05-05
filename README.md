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

	objectAcosSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Acos returns the arccosine, in radians, of x.
	err = bloblang.RegisterFunctionV2("acos", objectAcosSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Acos(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectAsinSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Asin returns the arcsine, in radians, of x.
	err = bloblang.RegisterFunctionV2("asin", objectAsinSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Asin(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectAsinhSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Asinh returns the inverse hyperbolic sine of x.
	err = bloblang.RegisterFunctionV2("asinh", objectAsinhSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Asinh(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectCbrtSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Cbrt returns the cube root of x.
	err = bloblang.RegisterFunctionV2("cbrt", objectCbrtSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Cbrt(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectCopysignSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("f")).Param(bloblang.NewFloat64Param("sign"))
	// Copysign returns a value with the magnitude of f and the sign of sign.
	err = bloblang.RegisterFunctionV2("copysign", objectCopysignSpec,
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
	// Dim returns the maximum of x-y or 0.
	err = bloblang.RegisterFunctionV2("dim", objectDimSpec,
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

	objectErfcSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Erfc returns the complementary error function of x.
	err = bloblang.RegisterFunctionV2("erfc", objectErfcSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Erfc(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectExpm1Spec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more
	err = bloblang.RegisterFunctionV2("expm1", objectExpm1Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Expm1(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectFMASpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y")).Param(bloblang.NewFloat64Param("z"))
	// FMA returns x * y &#43; z, computed with only one rounding. (That is, FMA
	err = bloblang.RegisterFunctionV2("fma", objectFMASpec,
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

	objectFloat32bitsSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("f"))
	// Float32bits returns the IEEE 754 binary representation of f,
	err = bloblang.RegisterFunctionV2("float32bits", objectFloat32bitsSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			f, err := args.GetFloat64("f")
			if err != nil {
				return nil, err
			}

			fa := float32(f)
			return func() (interface{}, error) {
				return math.Float32bits(fa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectFloat64bitsSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("f"))
	// Float64bits returns the IEEE 754 binary representation of f,
	err = bloblang.RegisterFunctionV2("float64bits", objectFloat64bitsSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			f, err := args.GetFloat64("f")
			if err != nil {
				return nil, err
			}

			fa := float64(f)
			return func() (interface{}, error) {
				return math.Float64bits(fa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectFloorSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Floor returns the greatest integer value less than or equal to x.
	err = bloblang.RegisterFunctionV2("floor", objectFloorSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Floor(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectGammaSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Gamma returns the Gamma function of x.
	err = bloblang.RegisterFunctionV2("gamma", objectGammaSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Gamma(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectHypotSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("p")).Param(bloblang.NewFloat64Param("q"))
	// Hypot returns Sqrt(p*p &#43; q*q), taking care to avoid unnecessary overflow and
	err = bloblang.RegisterFunctionV2("hypot", objectHypotSpec,
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

	objectJ0Spec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// J0 returns the order-zero Bessel function of the first kind.
	err = bloblang.RegisterFunctionV2("j0", objectJ0Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.J0(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectLgammaSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Lgamma returns the natural logarithm and sign (-1 or &#43;1) of Gamma(x).
	err = bloblang.RegisterFunctionV2("lgamma", objectLgammaSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Lgamma(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectMaxSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))
	// Max returns the larger of x or y.
	err = bloblang.RegisterFunctionV2("max", objectMaxSpec,
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

	objectNextafter32Spec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x")).Param(bloblang.NewFloat64Param("y"))
	// Nextafter32 returns the next representable float32 value after x towards y.
	err = bloblang.RegisterFunctionV2("nextafter32", objectNextafter32Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float32(x)

			y, err := args.GetFloat64("y")
			if err != nil {
				return nil, err
			}

			ya := float32(y)
			return func() (interface{}, error) {
				return math.Nextafter32(xa, ya), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectRoundSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Round returns the nearest integer, rounding half away from zero.
	err = bloblang.RegisterFunctionV2("round", objectRoundSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Round(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectSignbitSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Signbit reports whether x is negative or negative zero.
	err = bloblang.RegisterFunctionV2("signbit", objectSignbitSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Signbit(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectSincosSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Sincos returns Sin(x), Cos(x).
	err = bloblang.RegisterFunctionV2("sincos", objectSincosSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Sincos(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectSqrtSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Sqrt returns the square root of x.
	err = bloblang.RegisterFunctionV2("sqrt", objectSqrtSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Sqrt(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectTanSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Tan returns the tangent of the radian argument x.
	err = bloblang.RegisterFunctionV2("tan", objectTanSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Tan(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectTruncSpec := bloblang.NewPluginSpec().Param(bloblang.NewFloat64Param("x"))
	// Trunc returns the integer value of x.
	err = bloblang.RegisterFunctionV2("trunc", objectTruncSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			x, err := args.GetFloat64("x")
			if err != nil {
				return nil, err
			}

			xa := float64(x)
			return func() (interface{}, error) {
				return math.Trunc(xa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	objectInfSpec := bloblang.NewPluginSpec().Param(bloblang.NewInt64Param("sign"))
	// Inf returns positive infinity if sign &gt;= 0, negative infinity if sign &lt; 0.
	err = bloblang.RegisterFunctionV2("inf", objectInfSpec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			sign, err := args.GetInt64("sign")
			if err != nil {
				return nil, err
			}

			signa := int(sign)
			return func() (interface{}, error) {
				return math.Inf(signa), nil
			}, nil
		})
	if err != nil {
		panic(err)
	}
}
```

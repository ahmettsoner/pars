package functions

import (
	_math "parsdevkit.net/core/utils/math"
)

type MathFuncs struct{}

func (m MathFuncs) Abs(x interface{}) float64 {
	return _math.Abs(x)
}

func (m MathFuncs) Add(values ...interface{}) float64 {
	return _math.Add(values...)
}

func (m MathFuncs) Sub(x, y interface{}) float64 {
	return _math.Sub(x, y)
}

func (m MathFuncs) Pow(x, y interface{}) float64 {
	return _math.Pow(x, y)
}

func (m MathFuncs) Mul(values ...interface{}) float64 {
	return _math.Mul(values...)
}

func (m MathFuncs) Div(x, y interface{}) float64 {
	return _math.Div(x, y)
}

func (m MathFuncs) Max(x, y interface{}) float64 {
	return _math.Max(x, y)
}

func (m MathFuncs) Min(x, y interface{}) float64 {
	return _math.Min(x, y)
}

func (m MathFuncs) Mod(x, y interface{}) float64 {
	return _math.Mod(x, y)
}

func (m MathFuncs) Round(x, y interface{}) float64 {
	return _math.Round(x, y)
}

func (m MathFuncs) Floor(x, y interface{}) float64 {
	return _math.Floor(x, y)
}

func (f MathFuncs) IsInt(n interface{}) bool {
	return _math.IsInt(n)
}

func (f MathFuncs) IsFloat(n interface{}) bool {
	return _math.IsInt(n)
}
func (f MathFuncs) IsNum(n interface{}) bool {
	return _math.IsInt(n)
}

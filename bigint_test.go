package bigint

import (
	"database/sql/driver"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(200)
	c := big.NewInt(0)

	t.Run("abs", func(t *testing.T) {
		x := FromBigInt(a).Neg()
		y := FromBigInt(a)

		assert.Equal(t, y, x.Abs())
		assert.Equal(t, y, y.Abs())
	})

	t.Run("add", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 + 200 = 300
		assert.Equal(t, FromBigInt(big.NewInt(300)), x.Add(y))
	})

	t.Run("sub", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 -200 = -100
		assert.Equal(t, FromBigInt(big.NewInt(-100)), x.Sub(y))
	})

	t.Run("mul", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 * 200 = 20000
		assert.Equal(t, FromBigInt(big.NewInt(20000)), x.Mul(y))
	})

	t.Run("div", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 200 / 100 = 2
		assert.Equal(t, FromBigInt(big.NewInt(2)), y.Div(x))
	})

	t.Run("pow", func(t *testing.T) {
		x := FromInt64(10)
		y := FromInt64(6)
		// 10 ** 6 = 1 000 000
		assert.Equal(t, FromInt64(1000000), x.Pow(y))
	})

	t.Run("pow neg y", func(t *testing.T) {
		x := FromInt64(10)
		y := FromInt64(-5)
		// 10 ** 6 = 1 000 000
		assert.Equal(t, FromInt64(1), x.Pow(y))
	})

	t.Run("negation", func(t *testing.T) {
		x := FromBigInt(a)
		assert.Equal(t, FromBigInt(big.NewInt(-100)), x.Neg())
	})

	t.Run("compare", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, -1, x.Cmp(y))
		// 200 > 100
		assert.Equal(t, 1, y.Cmp(x))
		// 100 = 100
		assert.Equal(t, 0, x.Cmp(x))
	})

	t.Run("equal", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, false, x.Equal(y))
		// 200 > 100
		assert.Equal(t, false, y.Equal(x))
		// 100 = 100
		assert.Equal(t, true, x.Equal(x))
	})

	t.Run("greater than", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, false, x.GreaterThan(y))
		// 200 > 100
		assert.Equal(t, true, y.GreaterThan(x))
		// 100 = 100
		assert.Equal(t, false, x.GreaterThan(x))
	})

	t.Run("greater than or equal", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, false, x.GreaterThanOrEqual(y))
		// 200 > 100
		assert.Equal(t, true, y.GreaterThanOrEqual(x))
		// 100 = 100
		assert.Equal(t, true, x.GreaterThanOrEqual(x))
	})

	t.Run("less than", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, true, x.LessThan(y))
		// 200 > 100
		assert.Equal(t, false, y.LessThan(x))
		// 100 = 100
		assert.Equal(t, false, x.LessThan(x))
	})

	t.Run("less than or equal", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(b)
		// 100 < 200
		assert.Equal(t, true, x.LessThanOrEqual(y))
		// 200 > 100
		assert.Equal(t, false, y.LessThanOrEqual(x))
		// 100 = 100
		assert.Equal(t, true, x.LessThanOrEqual(x))
	})

	t.Run("is positive", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(new(big.Int).Neg(b))
		z := FromBigInt(c)

		assert.Equal(t, true, x.IsPositive())
		assert.Equal(t, false, y.IsPositive())
		assert.Equal(t, false, z.IsPositive())
	})

	t.Run("is negative", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(new(big.Int).Neg(b))
		z := FromBigInt(c)

		assert.Equal(t, false, x.IsNegative())
		assert.Equal(t, true, y.IsNegative())
		assert.Equal(t, false, z.IsNegative())
	})

	t.Run("is zero", func(t *testing.T) {
		x := FromBigInt(a)
		y := FromBigInt(new(big.Int).Neg(b))
		z := FromBigInt(c)

		assert.Equal(t, false, x.IsZero())
		assert.Equal(t, false, y.IsZero())
		assert.Equal(t, true, z.IsZero())
	})

	t.Run("int64", func(t *testing.T) {
		x := FromBigInt(a)
		assert.Equal(t, int64(-100), x.Neg().ToInt64())
	})

	t.Run("uint64", func(t *testing.T) {
		x := FromBigInt(a)
		assert.Equal(t, uint64(100), x.ToUInt64())
	})

	t.Run("toString", func(t *testing.T) {
		x := FromBigInt(a)
		assert.Equal(t, "100", x.String())
	})

	t.Run("fromString", func(t *testing.T) {
		x, err := FromString("100")
		assert.Nil(t, err)
		assert.Equal(t, "100", x.String())
	})

	t.Run("mustFromString", func(t *testing.T) {
		x := RequireFromString("100")
		assert.Equal(t, "100", x.String())
	})

	t.Run("fromInt64", func(t *testing.T) {
		x := FromInt64(100000000)
		assert.Equal(t, int64(100000000), x.ToInt64())
	})

	t.Run("value", func(t *testing.T) {
		x := FromInt64(10000000000)
		v, err := x.Value()
		assert.Nil(t, err)
		assert.True(t, true, driver.IsValue(v))
	})
}

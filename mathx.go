package mathx

import (
    "math"
)

type Number interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func Min[T Number](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

func Average[T Number](nums []T) float64 {
    if len(nums) == 0 {
        return 0
    }
    return float64(Sum(nums)) / float64(len(nums))
}

func Percent(val, total float64) float64 {
    if total == 0 {
        return 0
    }
    return (val / total) * 100
}

func RoundTo(val float64, precision int) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func Range(start, end, step int) []int {
    var res []int
    for i := start; i <= end; i += step {
        res = append(res, i)
    }
    return res
}

func Clamp[T Number](val, min, max T) T {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func Abs[T Number](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

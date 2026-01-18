package mathx

import "testing"

func TestMathX(t *testing.T) {
    if Max(10, 20) != 20 || Max(20, 10) != 20 {
        t.Error("Max failed")
    }

    if Min(10, 20) != 10 || Min(20, 10) != 10 {
        t.Error("Min failed")
    }

    if Sum([]int{1, 2, 3}) != 6 {
        t.Error("Sum failed")
    }

    if Average([]int{1, 2, 3}) != 2.0 || Average([]int{}) != 0 {
        t.Error("Average failed")
    }

    if Percent(50, 100) != 50.0 || Percent(50, 0) != 0 {
        t.Error("Percent failed")
    }

    if RoundTo(3.1415, 2) != 3.14 {
        t.Error("RoundTo failed")
    }

    if res := Range(1, 5, 2); len(res) != 3 || res[2] != 5 {
        t.Error("Range failed")
    }

    if Clamp(150, 0, 100) != 100 || Clamp(-10, 0, 100) != 0 || Clamp(50, 0, 100) != 50 {
        t.Error("Clamp failed")
    }

    if Abs(-50) != 50 || Abs(50) != 50 {
        t.Error("Abs failed")
    }
}

package mathx

import "testing"

func TestMathX(t *testing.T) {
    if Max(10, 20) != 20 {
        t.Error("Max failed")
    }
    
    nums := []int{1, 2, 3, 4}
    if Average(nums) != 2.5 {
        t.Error("Average failed")
    }
    
    if RoundTo(3.14159, 2) != 3.14 {
        t.Error("RoundTo failed")
    }
}

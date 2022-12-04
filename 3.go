package goLeet

func Min(x, y int) int {
	if x < y {
	  return x
	}
	return y
   }
   
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
// passes test cases but fails time limit
func maxArea(height []int) int {
    maxArea := 0
	l := len(height)
	for xI := 0; xI < l; xI++ {
		hI := height[xI]
		for xJ := xI+1; xJ < l; xJ++ {
			hJ := height[xJ]
			width := Abs(xJ-xI)
			height := Min(hI, hJ)
			area := width * height
			if area > maxArea {
				maxArea = area
			}

		}
	}
	return maxArea
}
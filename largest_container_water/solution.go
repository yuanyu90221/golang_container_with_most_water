package largest_container_water

func maxArea(height []int) int {
	lp := 0
	rp := len(height) - 1
	max := 0
	for lp < rp {
		w := rp - lp
		h := 0
		if height[rp] > height[lp] {
			h = height[lp]
			lp++
		} else {
			h = height[rp]
			rp--
		}
		area := h * w
		if max < area {
			max = area
		}
	}
	return max
}

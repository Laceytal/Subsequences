package functions

func FindMinFullAlphabet(a []int) int {
	n := len(a)
	cnt := make([]int, 27) // индексы 1..26
	unique := 0
	ans := n + 1
	l := 0

	for r := 0; r < n; r++ {
		x := a[r]
		if x >= 1 && x <= 26 {
			cnt[x]++
			if cnt[x] == 1 {
				unique++
			}
		}

		for unique == 26 && l <= r {
			windowLen := r - l + 1
			if windowLen < ans {
				ans = windowLen
			}
			y := a[l]
			if y >= 1 && y <= 26 {
				cnt[y]--
				if cnt[y] == 0 {
					unique--
				}
			}
			l++
		}
	}

	if ans <= n {
		return ans
	}
	return -1
}

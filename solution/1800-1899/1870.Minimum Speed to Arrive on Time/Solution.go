func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	const mx int = 1e7 + 1
	x := sort.Search(mx, func(s int) bool {
		if s == 0 {
			return false
		}
		var cost float64
		for _, v := range dist[:n-1] {
			cost += math.Ceil(float64(v) / float64(s))
		}
		cost += float64(dist[n-1]) / float64(s)
		return cost <= hour
	})
	if x < 0 || x == mx {
		return -1
	}
	return x
}
func fibonacci(n int) int {
	var m []int
	j := 2
	m = append(m, 1)
	m = append(m, 1)

	for j<=n{
	m = append(m,m[j-1] + m[j-2])
	j += 1
	}

	return m[n-1]

}





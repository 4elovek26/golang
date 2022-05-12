func minimumFromFour() int {
	a,b,c,d := 0,0,0,0
	fmt.Scan(&a,&b,&c,&d)
	min := a

	if min > b {
	min = b
	}

	if min > c {
	min = c
	}

	if min > d {
	min = d
	}

	return min
}





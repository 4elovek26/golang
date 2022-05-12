a:= 0
m := make(map[int]int)
for i:=0; i<10; i++{
	fmt.Scan(&a)
	if v, ok := m[a]; ok {
		fmt.Print(v, " ")
	}else{
		m[a] = work(a)
		fmt.Print(m[a], " ")
	}
}





// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task2(c chan string, value string) {
	for i := 0; i<5; i++{
		c<-value+" "
	}
}




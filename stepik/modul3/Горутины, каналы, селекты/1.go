// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task(c chan int, value int) {
	c<-value+1
}





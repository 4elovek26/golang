package main

import ("fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings")
type Battery struct {
	Energy string

}

func (batteryForTest Battery) String() string {
	count := strings.Count(batteryForTest.Energy, "0")
	newString := ""
	for i:= 0; i < count; i++{
		newString += " "
	}
	for ;count<10; count++{
		newString += "X"
	}
	return fmt.Sprintf("[%v]", newString)
}

func main() {
	var value string
	fmt.Scan(&value)
	batteryForTest := Battery{
		Energy: value,
	}
	// batteryForTest - не забывайте об имени
	// } Скобка, закрывающая функцию main() вам не видна, но она есть
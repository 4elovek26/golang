// Пакет и функция main уже объявлены.
// Выводить или вводить ничего не нужно!
func removeDuplicates(inputStream chan string, outputStream chan string) {
var sli string
var outsli string
isTrue := true
for value := range inputStream {
sli += value
}
for _, value := range sli {
for _, outvalue := range outsli {
if outvalue == value {
isTrue = false
continue

} else {
isTrue = true
continue
}
}
if isTrue {
outsli += string(value)
} else {
isTrue = true
}
}

outputStream <- outsli

defer close(outputStream)

}




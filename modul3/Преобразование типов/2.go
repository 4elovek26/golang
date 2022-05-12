// Пакет и функция main уже объявлены!
// Импортировать ничего не нужно!
// Удачи!
func adding(str1, str2 string) int64{
run1 := []rune(str1)
run2 := []rune(str2)
str1, str2 = "", ""
for _, v := range run1{
if v >= 48 && v <= 57 {
str1 += string(v)
}
}
for _, v := range run2{
if v >= 48 && v <= 57 {
str2 += string(v)
}
}
v1, _ := strconv.Atoi(str1)
v2, _ := strconv.Atoi(str2)
return int64(v1) + int64(v2)
}





// в задаче package и main уже объявлены+делаются какие-то проверки
var newUint uint8
var workArray [10]uint8
i := 0

fmt.Scan(&newUint)
for ; i<10; i++{
workArray[i] = newUint
if i+1 == 10 {break}
fmt.Scan(&newUint)
}

fIndex, sIndex := 0,0

for j := 0; j <3; j++{
fmt.Scan(&fIndex)
fmt.Scan(&sIndex)
workArray[fIndex], workArray[sIndex] = workArray[sIndex], workArray[fIndex]
}

for _, znach := range workArray{
fmt.Print(znach, " ")
}





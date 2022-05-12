
type weapon struct {
	On bool
	Ammo int
	Power int
}

func (w *weapon) Shoot() bool {
	if w.On == false {
	return false
	}
	if w.Ammo != 0 {
	w.Ammo--
	return true
	}
	return false
}

func (w *weapon) RideBike() bool {
	if w.On == false {
	return false
	}
	if w.Power != 0 {
	w.Power--
	return true
	}
	return false
}

func main() {

testStruct := new(weapon)
/*
 * Экземпляр созданной вами структуры необходимо передать в качестве
 * аргумента функции testStruct, которая выполнит проверку соблюдения
 * всех условий задания/
 */

// }
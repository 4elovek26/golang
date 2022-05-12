/*
 * Группировка городов по численности населения в тысячах человек
 * от 10 до 100, от 100 до 1000 и более 1000:
 * groupCity map[int][]string{
 *	 10: []string{...},
 *	 100: []string{...},
 *	 1000: []string{...},
 * }
 *
 * Население городов в тысячах человек:
 * cityPopulation map[string]int{...}
 */
inRange := false
for i := range cityPopulation {
for _, v := range groupCity[100] {
if i == v {
inRange = true
break
}
}
if inRange == false {
delete(cityPopulation, i)
}
inRange = false
}

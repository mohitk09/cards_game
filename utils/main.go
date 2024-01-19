package utils

func RetrieveCards(cards int32, suits int32) (a []int32) {
	for i := range make([]int32, cards*suits) {
		a = append(a, int32(i))
	}
	return
}

package utils

func retrieveCards(cards int, suits int) (a []int) {
	for i := range make([]int, cards*suits) {
		a = append(a, i)
	}
	return
}

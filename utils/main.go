package utils

import (
	"errors"
	"strings"

	"github.com/mohitk09/cards_game/constants"
)

func RetrieveCards(cards int32, suits int32) (a []int32) {
	for i := range make([]int32, cards*suits) {
		a = append(a, int32(i))
	}
	return
}

func ConvertCodeToID(code string) (id int32, err error) {
	var suit int32
	var value int32
	valueChar := code[0]
	suitChar := code[1]

	/* The following order suit order is maintained to calculate the ID
	Spades, Diamonds, Clubs and Hearts
	*/
	switch suitChar {
	case 'S':
		suit = 0
	case 'D':
		suit = 1
	case 'C':
		suit = 2
	case 'H':
		suit = 3
	default:
		return -1, errors.New("invalid suit")
	}

	switch valueChar {
	case 'A':
		value = 0
	case 'J':
		value = 10
	case 'Q':
		value = 11
	case 'K':
		value = 12
	case '2', '3', '4', '5', '6', '7', '8', '9':
		value = int32(valueChar-'0') - 1
	case '1': // This is a case when the value is 10, as we are only computing against first digit
		value = 9
	default:
		return -1, errors.New("invalid value")

	}

	/* The value would be in the range from 0 to 51 both inclusive
	Spades:-   [0, 12]
	Diamonds:  [13, 25]
	Clubs:-    [26, 38]
	Hearts:-   [39, 51]
	*/
	return suit*constants.Numbers + value, nil
}

func RetrieveSelectedCards(codeQuery string) (ids []int32, err error) {
	lisfOfCodes := strings.Split(codeQuery, ",")
	for _, code := range lisfOfCodes {
		codeID, err := ConvertCodeToID(code)
		if err != nil {
			return nil, err
		}
		ids = append(ids, codeID)
	}
	return
}

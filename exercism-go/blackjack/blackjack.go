package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	table := map[string]int{
		"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10,
		"king": 10, "joker": 0,
	}
	return table[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case sum == 21 && dealer < 10:
		return "W"
	case (sum >= 17 && sum <= 21) || (sum >= 12 && sum <= 16 && dealer < 7):
		return "S"
	case sum == 22:
		return "P"
	default:
		return "H"
	}
}

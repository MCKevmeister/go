package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "joker":
		return 0
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 10
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	case handValue == 22:
		return "P"
	case handValue == 21 && dealerValue <= 9:
		return "W"
	case handValue == 21 && dealerValue >= 10:
		return "S"
	case handValue >= 17:
		return "S"
	case handValue >= 12 && dealerValue < 7:
		return "S"
	}
	return "H"
}

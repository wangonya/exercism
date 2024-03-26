package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	firstTurn := ""
	myCardsValue := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		firstTurn = "P"
	case myCardsValue == 21 && dealerCard != "ace" && dealerCard != "ten" && dealerCard != "queen":
		firstTurn = "W"
	case myCardsValue == 21 && (dealerCard == "ace" || dealerCard == "ten" || dealerCard=="queen"):
		firstTurn = "S"
	case myCardsValue >= 17 && myCardsValue <= 20:
		firstTurn = "S"
	case (myCardsValue >= 12 && myCardsValue <= 16) && dealerCardValue >= 7:
		firstTurn = "H"
	case myCardsValue >= 12 && myCardsValue <= 16:
		firstTurn = "S"
	case myCardsValue <= 11:
		firstTurn = "H"
	}
	return firstTurn
}

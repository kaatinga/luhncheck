package luhncheck

import (
	"errors"
)

// Errors
var (
	IncorrectCardNumber       = errors.New("only the numerical characters are allowed")
	IncorrectCardNumberLength = errors.New("incorrect card number length")
	CheckedAlready            = errors.New("the card has been checked already")
)

// card models
type BankCard struct {
	Number  string
	Valid   bool
	Checked bool
}

// NewCard creates a new bank card
func NewCard(number string) (*BankCard, error) {

	// the maximum length is 19, the minimum is 13
	if len(number) < 13 || len(number) > 19 {
		return nil, IncorrectCardNumberLength
	}

	// only numbers are allowed
	for _, value := range number {
		if value < 48 && value > 57 {
			return nil, IncorrectCardNumber
		}
	}

	return &BankCard{Number: number}, nil
}

// Check checks if the card is valid
func (card *BankCard) Check() (bool, error) {
	if card.Checked {
		return false, CheckedAlready
	}

	card.Checked = true

	var cardNumberLength = len(card.Number)
	var i = cardNumberLength % 2

	var number byte
	var doubleNumbers = make(map[int]byte, cardNumberLength/2)

	// compose new slice with numbers
	for {

		// new map to check
		number = card.Number[i] - 48
		doubleNumbers[i] = number * 2
		if doubleNumbers[i] > 9 {
			doubleNumbers[i] = doubleNumbers[i] - 9
		}
		//log.Println("дубль для числа с индексом", i, "=", doubleNumbers[i])

		i = i + 2
		if i >= cardNumberLength {
			break
		}
	}

	// sum
	i = 0
	var sum byte
	var hasDouble bool
	for {
		_, hasDouble = doubleNumbers[i]
		if hasDouble {
			sum = sum + doubleNumbers[i]
		} else {
			sum = sum + card.Number[i]-48
		}

		i++
		if cardNumberLength == i {
			break
		}
	}

	if sum % 10 == 0 {
		card.Valid = true
		//log.Println("the card is valid!")
	}

	return card.Valid, nil
}

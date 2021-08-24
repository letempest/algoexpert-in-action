package main

import (
	"fmt"
	"strings"
)

func main() {
	phoneNumber := "1905"
	mnemonics := phoneNumberMnemonics(phoneNumber)
	fmt.Printf("%#v\n", mnemonics)
}

// Big O: O(4^n * n) time, worst case scenario is phone number filled with 7/9, then each digit invokes 4 recursive calls, so 4^n recursive calls in total;
// and for each recursive call, when hit the base case, a new mnemonic of length n is generated, this is O(n) time operation
// O(4^n * n) space, same reason, 4^n recursive calls, each call at the base case appends a new mnemonic of length n to the mnemonics set
func phoneNumberMnemonics(phoneNumber string) []string {
	DIGIT_LETTERS := map[string][]string{
		"0": {"0"},
		"1": {"1"},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	currentMnemonic := make([]string, len(phoneNumber))
	mnemonicsFound := []string{}

	phoneNumberMnemonicsHelper(0, phoneNumber, DIGIT_LETTERS, &currentMnemonic, &mnemonicsFound)
	return mnemonicsFound
}

func phoneNumberMnemonicsHelper(idx int, phoneNumber string, DIGIT_LETTERS map[string][]string, currentMnemonic *[]string, mnemonicsFound *[]string) {
	// base case
	if idx == len(phoneNumber) {
		mnemonic := strings.Join(*currentMnemonic, "")
		*mnemonicsFound = append(*mnemonicsFound, mnemonic)
	} else {
		digit := string(phoneNumber[idx])
		letters := DIGIT_LETTERS[digit]
		for _, letter := range letters {
			(*currentMnemonic)[idx] = letter
			phoneNumberMnemonicsHelper(idx+1, phoneNumber, DIGIT_LETTERS, currentMnemonic, mnemonicsFound)
		}
	}
}

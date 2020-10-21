package main

import "strings"

var replacer = strings.NewReplacer(
	"стат", "status",
	"адд все", "add .",
	"адд", "add",
	"добавь", "add",
	"бранч", "branch",
	"статус", "status",
	"рассол", "reset",
	"курва!", "push --force",
	"курва", "push",
	"сурово", "--hard",
	"мягко", "--soft",
	"тяни", "pull",
	// простая клавиатура
	"й", "q", "ц", "w", "у", "e", "к", "r", "е", "t", "н", "y", "г", "u", "ш", "i", "щ", "o", "з", "p",
	"ф", "a", "ы", "s", "в", "d", "а", "f", "п", "g", "р", "h", "о", "j", "л", "k", "д", "l",
	"я", "z", "ч", "x", "с", "c", "м", "v", "и", "b", "т", "n", "ь", "m",
)

func Repalce(s string) string {
	return replacer.Replace(s)
}

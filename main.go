package main

import (
	"fmt"
	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
	"os"
	"os/exec"
	"strings"
)

func main() {
	str := strings.Join(os.Args, " ")
	str = strings.ToLower(str)
	custom := translit.Map(
		map[string]string{
			"стат":    "status",
			"адд все": "add .",
			"адд":     "add",
			"добавь":  "add",
			"бранч":   "branch",
			"статус":  "status",
			"рассол":  "reset",
			"курва!":  "push --force",
			"курва":   "push",
			"сурово":  "--hard",
			"мягко":   "--soft",
			"тяни":    "pull",
			// простая клавиатура
			"й": "q", "ц": "w", "у": "e", "к": "r", "е": "t", "н": "y", "г": "u", "ш": "i", "щ": "o", "з": "p",
			"ф": "a", "ы": "s", "в": "d", "а": "f", "п": "g", "р": "h", "о": "j", "л": "k", "д": "l",
			"я": "z", "ч": "x", "с": "c", "м": "v", "и": "b", "т": "n", "ь": "m",
		},
	)

	tr := custom.Transformer()
	tr.Reset()
	s, _, err := transform.String(tr, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Що ти зробив:", s)
	args := strings.Split(s, " ")
	cmd := exec.Command("git", args[1:]...)
	path, _ := os.Getwd()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

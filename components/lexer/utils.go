package lexer

import "strconv"

func IsDigit(character *string) bool {
	_, err := strconv.Atoi(*character)
	return err == nil
}

package main

import (
	"fmt"
	"strconv"
)

func strToInt(ss string) int {
	ii, err := strconv.ParseInt(ss, 10, 64)
	if err != nil {
		ii = 0
	}
	return int(ii)
}

func rowToCol(rr int) string {
	rr--
	if rr < 26 {
		return string(rune(65 + rr))
	}
	dd := rr % 26
	rr = rr / 26
	return rowToCol(rr) + string(rune(65+dd))
}

func recurseToRow(cc []rune) int {
	if len(cc) == 1 {
		return int(cc[0]) - 64 // has a hidden +1 (we subtract 65 for 'A', then +1)
	}
	lx := len(cc)
	return (recurseToRow(cc[:lx-1]) * 26) + (int(cc[lx-1]) - 64)
}

func colToRow(cc string) int {
	rl := make([]rune, len(cc))
	for idx, ch := range cc {
		rl[idx] = ch
	}
	return recurseToRow(rl)
}

func main() {

	for {

		fmt.Print("Row: ")
		var imput string
		fmt.Scanln(&imput)
		num := strToInt(imput)
		fmt.Println(rowToCol(num))

		fmt.Print("Col: ")
		fmt.Scanln(&imput)
		fmt.Println(colToRow(imput))
		fmt.Println("")

	}
}

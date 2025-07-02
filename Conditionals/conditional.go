package main

import (
	"fmt"
	"strings"
)
func main() {
	password := "1234567890"

	if len(password) < 10 {
		println(`Password length must be atleast 10`)
	} else if strings.Compare(password, "1234567890") == 0 {
		fmt.Println(`Valid password`)
	}else {
		fmt.Println(`Invalid password`)
	}
}
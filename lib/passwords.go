package lib

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Ex4() {
	s := `password124`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Println(s)

	err = bcrypt.CompareHashAndPassword(bs, []byte(s+"1"))
	if err != nil {
		fmt.Println("Not Match!")
	}

	fmt.Println(string(bs))
}

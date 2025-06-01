package myreader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadExample() int {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Example Number: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err.Error())
	}

	example, err := strconv.Atoi(strings.Trim(text, "\n"))

	if err != nil {
		log.Fatal(err.Error())
	}

	return example
}

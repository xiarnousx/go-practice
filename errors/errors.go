package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
)

func Ex1(t string) {
	_, err := os.Open("no-file.txt")
	if err != nil {
		switch t {
		case "p":
			fmt.Println(err)
		case "l":
			log.Println(err)
		case "f":
			log.Fatalln(err)
		case "i":
			panic(err)
		}
	}
}

func Ex2() {
	sqrt := func(f float64) (float64, error) {
		if f < 0 {
			return 0, errors.New("negative value")
		}
		return math.Pow(f, 2), nil
	}

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

type norgateMath struct {
	lat  string
	long string
	err  error
}

func (n norgateMath) Error() string {
	return fmt.Sprintf("occured at %v, %v %v", n.lat, n.long, n.err)
}

func Ex3() {
	sqrt := func(f float64) (float64, error) {
		if f < 0 {
			return 0, norgateMath{lat: "50.222", long: "99.456", err: fmt.Errorf("negative sqrt %v", f)}
		}
		return math.Pow(f, 2), nil
	}

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Ex4() {
	p1 := struct {
		First string
		Last  string
	}{
		First: "X",
		Last:  "Y",
	}
	bs, err := toJson(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func toJson(a any) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("[ERROR]: %v", err)
	}
	return bs, nil
}

type errorRandom struct {
	random int
	err    error
}

func (e errorRandom) Error() string {
	return fmt.Sprintf("[ERROR] [%v]: %v", e.random, e.err)
}

func foo(e error) {
	log.Println(e)
}

func Ex5() {
	e1 := errorRandom{
		random: rand.IntN(100),
		err:    errors.New("traverse the universe"),
	}
	foo(e1)
}

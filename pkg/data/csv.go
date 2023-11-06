package data

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gabrielmusskopf/avl/pkg/types"
)

type PersonReader interface {
    Read(string) *[]types.Person
}

const (
    // csv fields
	CPF = iota
	RG
	Name
	BirthDate
	BirthCity
)

type CsvPersonReader struct {
}

func (r CsvPersonReader) Read(filename string) []*types.Person {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	people := make([]*types.Person, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // header

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ";")
		person := types.NewPerson(tokens[CPF], tokens[RG], tokens[Name], tokens[BirthDate], tokens[BirthCity])
		people = append(people, person)
	}

    return people
}

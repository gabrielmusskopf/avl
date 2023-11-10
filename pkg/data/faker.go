package data

import (
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

const DataPath string = "pkg/data/fake_data.csv"

func writeOrDie(f *os.File, content string) {
	b := []byte(content)
	read, err := f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	if read != len(b) {
		log.Fatalf("Not enough bytes written to data csv file")
	}
}

func HasData() bool {
    _, err := os.Open(DataPath)
    return err == nil
}

func Generate(n int) {
	gofakeit.Seed(0)
	f, err := os.Create(DataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writeOrDie(f, "cpf;rg;nome;data_nascimento;cidade_nascimento")

	for i := 0; i < n; i++ {
		cpf := gofakeit.Numerify("###########")
		rg := gofakeit.Numerify("#########")
		nome := gofakeit.Name()
		dataNascimento := gofakeit.Date().Format("02/01/2006")
		cidadeNascimento := gofakeit.City()

		writeOrDie(f, fmt.Sprintf("%s;%s;%s;%s;%s\n", cpf, rg, nome, dataNascimento, cidadeNascimento))
	}
}

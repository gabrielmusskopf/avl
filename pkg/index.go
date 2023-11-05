package avl

import "github.com/gabrielmusskopf/avl/pkg/types"


type Index struct {
	Names     *IndexedTree[types.String, *types.Person]
	CPF       *IndexedTree[types.String, *types.Person]
	BirthDate *IndexedTree[types.Date, *types.Person]
}

func BuildIndexes(people []*types.Person) *Index {
	indexed := &Index{}
	for _, person := range people {
		indexed.Names = indexed.Names.Add(types.String(person.Name), person)
		indexed.CPF = indexed.CPF.Add(types.String(person.CPF), person)
		indexed.BirthDate = indexed.BirthDate.Add(types.DateFromString(person.BirthDate), person)
	}
    return indexed
}

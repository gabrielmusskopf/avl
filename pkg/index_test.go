package avl

import (
	"testing"

	"github.com/gabrielmusskopf/avl/pkg/types"
)

var index *Index

func init() {
	people := []*types.Person{
		{Name: "Bar", CPF: "1", BirthDate: "01/01/2000"},
		{Name: "Foo", CPF: "2", BirthDate: "01/01/2001"},
		{Name: "Baz", CPF: "3", BirthDate: "01/01/2002"},
	}
	index = BuildIndexes(people)
}

func TestNameIndexTree(t *testing.T) {
	nameIndex := index.Names

	validate("Baz", nameIndex, t)
	validate("Bar", nameIndex.Left, t)
	validate("Foo", nameIndex.Right, t)
}

func TestCPFIndexTree(t *testing.T) {
	cpfIndex := index.CPF

	validate("2", cpfIndex, t)
	validate("1", cpfIndex.Left, t)
	validate("3", cpfIndex.Right, t)
}

func TestBirthIndexTree(t *testing.T) {
	birthIndex := index.BirthDate

	validate(types.DateFromString("01/01/2001"), birthIndex, t)
	validate(types.DateFromString("01/01/2000"), birthIndex.Left, t)
	validate(types.DateFromString("01/01/2002"), birthIndex.Right, t)
}

func validate[K Ordered[K], T any](a K, node *IndexedTree[K, T], t *testing.T) {
	if node.Key.Compare(a) != 0 {
		t.Errorf("got %v in root, but wanted %v", node.Key, a)
	}
}

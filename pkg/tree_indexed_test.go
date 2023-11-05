package avl

import (
	"sort"
	"strings"
	"testing"

	"github.com/gabrielmusskopf/avl/pkg/types"
)

type Int int

func (i Int) Compare(other Int) int {
	if int(i) < int(other) {
		return -1
	} else if int(i) > int(other) {
		return 1
	}
	return 0
}

func (i Int) Less(other Int) bool {
	return i.Compare(other) == -1
}

func TestIndexedTreeCreation(t *testing.T) {
	tree := newIndexedTree(Int(1), "first")

	if tree == nil {
		t.Errorf("%v should not be nil", t)
	}
}

func TestIndexedTreeInsertion(t *testing.T) {
	tree := newIndexedTree(Int(10), "first")
	tree = tree.Add(Int(5), "second")
	tree = tree.Add(Int(15), "thirth")

	if tree == nil {
		t.Errorf("%v should not be nil", t)
	}

	first := tree.Search(Int(10))
	if first == nil {
		t.Errorf("expected not nil, but recieved %v", first)
	}
	if first.Value != "first" {
		t.Errorf("expected '%s', but recieved %s", "first", first.Value)
	}

	second := tree.Search(Int(5))
	if second == nil {
		t.Errorf("expected not nil, but recieved %v", second)
	}
	if second.Value != "second" {
		t.Errorf("expected '%s', but recieved %s", "second", second.Value)
	}

	thirth := tree.Search(Int(15))
	if thirth == nil {
		t.Errorf("expected not nil, but recieved %v", thirth)
	}
	if thirth.Value != "thirth" {
		t.Errorf("expected '%s', but recieved %s", "thirth", thirth.Value)
	}
}

func TestIndexedTreeSearch(t *testing.T) {
	tree := newIndexedTree(Int(10), "first")
	tree = tree.Add(Int(5), "second")
	tree = tree.Add(Int(15), "thirth")

	if tree == nil {
		t.Errorf("%v should not be nil", t)
	}

	nodesMatched := tree.SearchAllBy(Int(10),
		func(k1, k2 Int) bool { return k1 == k2 },
		func(k1, k2 Int) int { return k1.Compare(k2) })

	if len(nodesMatched) != 1 {
		t.Errorf("expected length 1 but recieved %d", len(nodesMatched))
	}
	if nodesMatched[0].Value != "first" {
		t.Errorf("expected 'first' but recieved %s", nodesMatched[0].Value)
	}
}

func compareWithLength(s, other types.String) int {
	if len(s) > len(other) {
		s = s[:len(other)]
	} else {
		other = other[:len(s)]
	}
	return s.Compare(other)
}

func TestStringIndexedTreeSearch(t *testing.T) {
	tree := newIndexedTree(types.String("baz"), 10)
	tree = tree.Add(types.String("bar"), 5)
	tree = tree.Add(types.String("foo"), 15)

	if tree == nil {
		t.Errorf("%v should not be nil", t)
	}

	nodesMatched := tree.SearchAllBy(types.String("baz"),
		func(k1, k2 types.String) bool { return strings.HasPrefix(string(k1), string(k2)) },
		func(k1, k2 types.String) int { return compareWithLength(k1, k2) })

	if len(nodesMatched) != 1 {
		t.Errorf("expected length 1 but recieved %d", len(nodesMatched))
	}
	if nodesMatched[0].Value != 10 {
		t.Errorf("expected 'first' but recieved %d", nodesMatched[0].Value)
	}
}

func TestStringIndexedTreeSearchMultiValues(t *testing.T) {
	tree := newIndexedTree(types.String("baz"), 10)
	tree = tree.Add(types.String("bar"), 5)
	tree = tree.Add(types.String("foo"), 15)

	if tree == nil {
		t.Errorf("%v should not be nil", t)
	}

	nodesMatched := tree.SearchAllBy(types.String("ba"),
		func(k1, k2 types.String) bool { return strings.HasPrefix(string(k1), string(k2)) },
		func(k1, k2 types.String) int { return compareWithLength(k1, k2) })

	if len(nodesMatched) != 2 {
		t.Errorf("expected length 1 but recieved %d", len(nodesMatched))
	}

	sort.Slice(nodesMatched, func(i, j int) bool { return i > j })

	if nodesMatched[0].Value != 5 {
		t.Errorf("expected '5' but recieved %d", nodesMatched[0].Value)
	}
	if nodesMatched[1].Value != 10 {
		t.Errorf("expected '10' but recieved %d", nodesMatched[1].Value)
	}
}

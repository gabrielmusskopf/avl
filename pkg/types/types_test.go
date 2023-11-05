package types

import "testing"

func TestTypeStringLess(t *testing.T) {
    a := String("a")
    b := String("b")
    
    if !a.Less(b) {
		t.Error("got a > b, but wanted a < b")
    }
}

func TestTypeStringCompareLess(t *testing.T) {
    a := String("a")
    b := String("b")
    
    if a.Compare(b) != -1 {
		t.Error("got a > b, but wanted a < b")
    }
}

func TestTypeStringLessCompareWithLength(t *testing.T) {
    a := String("a")
    b := String("aaaa")

    if a.CompareWithLength(b) != 0 {
		t.Error("got a > b, but wanted a < b")
    }
}

func TestTypeStringGraterCompareWithLength(t *testing.T) {
    a := String("aaaa")
    b := String("a")

    if a.CompareWithLength(b) != 0 {
		t.Error("got a > b, but wanted a < b")
    }
}

func TestTypeStringEqualCompareWithLength(t *testing.T) {
    a := String("a")
    b := String("a")

    if a.CompareWithLength(b) != 0 {
		t.Error("got a > b, but wanted a < b")
    }
}

func TestTypeStringCompareEquals(t *testing.T) {
    a := String("a")
    b := String("a")
    
    if a.Compare(b) != 0 {
		t.Error("got a > b or a < b, but wanted a == b")
    }
}

func TestTypeDateLess(t *testing.T) {
    a := DateFromString("01/01/2000")
    b := DateFromString("01/01/2001")
    
    if !a.Less(b) {
		t.Errorf("got %v > %v, but wanted %v < %v", a, b, a, b)
    }
}

func TestTypeDateCompareLess(t *testing.T) {
    a := DateFromString("01/01/2000")
    b := DateFromString("01/01/2001")
    
    if a.Compare(b) != -1 {
		t.Errorf("got %v < %v, but wanted %v > %v", a, b, a, b)
    }
}

func TestTypeDateCompareEquals(t *testing.T) {
    a := DateFromString("01/01/2000")
    b := DateFromString("01/01/2000")
    
    if a.Compare(b) != 0 {
		t.Errorf("got %v != %v, but wanted %v == %v", a, b, a, b)
    }
}

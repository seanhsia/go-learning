package common_test

import (
	"go-learning/common"
	"testing"
)

type compositeKey struct {
	k1 string
	k2 int
}

func TestInitailize(t *testing.T) {
	set := &common.Set[string]{}
	set.New()

	if len(set.ToSlice()) != 0 || set.Len() != 0 {
		t.Errorf("The set is not clean after initialize: %v", set.ToSlice())
	}
}

func TestAdd(t *testing.T) {
	set := &common.Set[string]{}
	set.New()
	expected := "Apple"

	set.Add(expected)

	if set.ToSlice()[0] != expected {
		t.Errorf("Expect: %v The set has the wrong element: %v", expected, set.ToSlice()[0])
	}

}

func TestNotExist(t *testing.T) {
	set := &common.Set[string]{}
	set.New()
	nonexists := "Apple"

	if set.Exists(nonexists) {
		t.Errorf("Failed to detect nonexists element!")
	}

}

func TestExist(t *testing.T) {
	set := &common.Set[string]{}
	set.New()
	word := "Apple"
	set.Add(word)

	if !set.Exists(word) {
		t.Errorf("Word: %v does not exists in set: %v", word, set.ToSlice())
	}
}

func TestRemove(t *testing.T) {

	set := &common.Set[string]{}
	set.New()
	set2 := &common.Set[string]{}
	set2.New()
	word := "Apple"
	set.Add(word)

	set.Remove("")

	if set.Len() != 1 && set.ToSlice()[0] != word {
		t.Errorf("Removing an empty string should not effect the exists element %v", word)
	}

	set.Remove(word)

	if !set.Equal(set2) {
		t.Errorf("Removing word %v shoud got an empty set", word)
	}

}

func TestIntersect(t *testing.T) {
	s1 := &common.Set[string]{}
	s1.New()
	s1.Add("Apple")
	s1.Add("Banana")

	s2 := &common.Set[string]{}
	s2.New()
	s2.Add("Apple")
	s2.Add("Coconut")

	expected := &common.Set[string]{}
	expected.New()
	expected.Add("Apple")

	s1.Intersect(s2)

	if !s1.Equal(expected) {
		t.Errorf("Expected to be %v but got :%v", expected, s1)
	}
}

func TestUnion(t *testing.T) {
	s1 := &common.Set[string]{}
	s1.New()
	s1.Add("Apple")
	s1.Add("Banana")

	s2 := &common.Set[string]{}
	s2.New()
	s2.Add("Apple")
	s2.Add("Coconut")

	expected := common.ToSet[string]([]string{"Apple", "Banana", "Coconut"})

	s1.Union(s2)
	if !s1.Equal(expected) {
		t.Errorf("Expected to be %v but got :%v", expected, s1)
	}
}

func TestDiff(t *testing.T) {
	s1 := &common.Set[string]{}
	s1.New()
	s1.Add("Apple")
	s1.Add("Banana")

	s2 := &common.Set[string]{}
	s2.New()
	s2.Add("Apple")
	s2.Add("Coconut")

	expected := common.ToSet[string]([]string{"Banana"})

	s1.Diff(s2)
	if !s1.Equal(expected) {
		t.Errorf("Expected to be %v but got :%v", expected, s1)
	}
}

func TestStruct(t *testing.T) {
	slice := []compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}}
	set := common.ToSet[compositeKey](slice)
	expected := common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}, {k1: "C", k2: 28}})

	if set.Len() != len(slice) {
		t.Errorf("Expect the set have length: %v but got length: %v", len(slice), set.Len())
	}

	for _, key := range slice {
		if !set.Exists(key) {
			t.Errorf("Expected key: %v in set: %v", key, set)
		}
	}

	set.Add(compositeKey{k1: "C", k2: 28})

	if !set.Equal(expected) {
		t.Errorf("After adding element exepected %v, but got %v", expected, set)
	}

}

func TestTwoSets(t *testing.T) {
	type testcase struct {
		s1     *common.Set[compositeKey]
		s2     *common.Set[compositeKey]
		expect *common.Set[compositeKey]
	}

	testcases := []testcase{
		{
			s1:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}}),
			s2:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "C", k2: 28}}),
			expect: common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}, {k1: "C", k2: 28}}),
		},
		{
			s1:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}}),
			s2:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "C", k2: 28}}),
			expect: common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}}),
		},
		{
			s1:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "B", k2: 27}}),
			s2:     common.ToSet[compositeKey]([]compositeKey{{k1: "A", k2: 26}, {k1: "C", k2: 28}}),
			expect: common.ToSet[compositeKey]([]compositeKey{{k1: "B", k2: 27}}),
		},
	}

	for idx, testcase := range testcases {
		switch idx {
		case 0:
			testcase.s1.Union(testcase.s2)
			if !testcase.s1.Equal(testcase.expect) {
				t.Errorf("Expectation after union is %v but got %v", testcase.expect, testcase.s1)
			}
		case 1:
			testcase.s1.Intersect(testcase.s2)
			if !testcase.s1.Equal(testcase.expect) {
				t.Errorf("Expectation after union is %v but got %v", testcase.expect, testcase.s1)
			}
		case 2:
			testcase.s1.Diff(testcase.s2)
			if !testcase.s1.Equal(testcase.expect) {
				t.Errorf("Expectation after union is %v but got %v", testcase.expect, testcase.s1)
			}
		}
	}

}

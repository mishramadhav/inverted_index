package set_test

import (
	"testing"

	"github.com/mishramadhav/inverted_index/internal/set"
)

func TestSet(t *testing.T) {
	s := set.New[int]()
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}
	s.Add(1)
	if s.Size() != 1 {
		t.Errorf("expected size 1, got %d", s.Size())
	}
	if !s.Contains(1) {
		t.Error("expected set to contain 1")
	}
	if s.Contains(2) {
		t.Error("expected set to not contain 2")
	}
	s.Add(2)
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}
	if !s.Contains(2) {
		t.Error("expected set to contain 2")
	}
	s.Remove(1)
	if s.Size() != 1 {
		t.Errorf("expected size 1, got %d", s.Size())
	}
	if s.Contains(1) {
		t.Error("expected set to not contain 1")
	}
	if !s.Contains(2) {
		t.Error("expected set to contain 2")
	}
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}
	if s.Contains(2) {
		t.Error("expected set to not contain 2")
	}
	values := s.Values()
	if len(values) != 0 {
		t.Errorf("expected no values, got %v", values)
	}

	s2 := set.NewWithMaxSize[int](10)
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	values = s2.Values()
	if len(values) != 3 {
		t.Errorf("expected 3 values, got %v", values)
	}

	s3 := set.New[int]()
	s3.Add(4)
	s3.Add(5)
	s4 := s2.Union(s3)
	values = s4.Values()
	if len(values) != 5 {
		t.Errorf("expected 5 values, got %v, expected %v", values, []int{1, 2, 3, 4, 5})
	}

	s5 := set.New[int]()
	s5.Add(2)
	s5.Add(3)
	s6 := s2.Intersection(s5)
	values = s6.Values()
	if len(values) != 2 {
		t.Errorf("expected 2 value, got %v, expected %v", values, []int{2, 3})
	}
}

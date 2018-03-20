package deepcopy

import (
	"testing"
)

func TestCopyMap(t *testing.T) {
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	nn := Copy(m)
	if nn == nil {
		t.Fatal("should not return nil")
	}
	n, ok := nn.(map[string]int)
	if !ok {
		t.Fatalf("should return map[string]int but type is %T", n)
	}
	if len(m) != len(n) {
		t.Errorf("len(m) == %d && len(n) == %d", len(m), len(n))
	}
	testCopyMapKey(t, m, n, "1")
	testCopyMapKey(t, m, n, "2")
	testCopyMapKey(t, m, n, "3")
	n["3"] = 1
	n["1"] = 3
	testCopyMapKeyValue(t, m, "1", 1)
	testCopyMapKeyValue(t, m, "3", 3)
	testCopyMapKeyValue(t, n, "1", 3)
	testCopyMapKeyValue(t, n, "3", 1)
}

func testCopyMapKey(t *testing.T, m, n map[string]int, k string) {
	mv, ok := m[k]
	if !ok {
		t.Fatalf("map m has no key %v", k)
		return
	}
	nv, ok := n[k]
	if !ok {
		t.Fatalf("map m has no key %v", k)
		return
	}
	if mv != nv {
		t.Errorf("m[%[1]v] == %v && n[%[1]v] == %v", k, mv, nv)
	}
}

func testCopyMapKeyValue(t *testing.T, m map[string]int, k string, v interface{}) {
	mv, ok := m[k]
	if !ok {
		t.Fatalf("map m has no key %v", k)
		return
	}
	if mv != v {
		t.Errorf("m[%v] == %v != %v", k, mv, v)
	}
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3}
	copy := Copy(s)
	if copy == nil {
		t.Fatal("should not return nil")
	}
	ss, ok := copy.([]int)
	if !ok {
		t.Fatalf("should return []int but type is %T", ss)
	}
	if len(s) != len(ss) {
		t.Errorf("len(s) == %d && len(ss) == %d", len(s), len(ss))
	}
	for i := 0; i < len(s) && i < len(ss); i++ {
		if s[i] != ss[i] {
			t.Errorf("s[%[1]d] == %d && ss[%[1]d] == %d", i, s[i], ss[i])
		}
	}
	ss[2] = 4
	if s[2] == 4 {
		t.Errorf("s[2] == 4")
	}
}

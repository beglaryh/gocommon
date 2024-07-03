package hashmap

import (
	"github.com/beglaryh/gocommon/optional"
	"testing"
)

func TestHashMap(t *testing.T) {
	hmap := New[string, string]()
	hmap.Put("hello", "world")
	op := hmap.Get("hello")
	if op.IsEmpty() {
		t.Fail()
	}
	v, _ := op.Get()
	if v != "world" {
		t.Fail()
	}

	op = hmap.Compute("k1", func(k string, vop optional.Optional[string]) optional.Optional[string] {
		if vop.IsEmpty() {
			return optional.With[string]("empty")
		}
		v, _ := vop.Get()
		return optional.With[string](v + v)
	})

	v, _ = op.Get()
	if v != "empty" {
		t.Fail()
	}
	op = hmap.Compute("k1", func(k string, vop optional.Optional[string]) optional.Optional[string] {
		if vop.IsEmpty() {
			return optional.With[string]("empty")
		}
		v, _ := vop.Get()
		return optional.With[string](v + v)
	})

	v, _ = op.Get()
	if v != "emptyempty" {
		t.Fail()
	}

	op = hmap.Remove("hello")
	if hmap.Contains("hello") {
		t.Fail()
	}
}

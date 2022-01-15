package orderedMap

import (
	"fmt"
	"testing"
)

type (
	testableOM interface {
		Set(key, value string)
		Get(key string) (string, bool)
		Delete(key string)
		Range() <-chan string
	}
)

func testInsertion(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	obj.Set("space", "space-1")
	obj.Set("lift", "lift-1")
	obj.Set("face", "face-1")
	obj.Set("broker", "broker-1")
	want := "foo=goo-1;true;bar=bar-1;true;space=space-1;true;lift=lift-1;true;face=face-1;true;broker=broker-1;true;"
	var got string
	for key := range obj.Range() {
		val, ok := obj.Get(key)
		got += fmt.Sprintf("%s=%s;%t;", key, val, ok)
	}
	if got != want {
		t.Errorf("matching error\nwant: %s\ngot:  %s", want, got)
	}
}

func testDeletion(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	obj.Set("space", "space-1")
	obj.Set("lift", "lift-1")
	obj.Set("face", "face-1")
	obj.Delete("space")
	obj.Delete("foo")
	obj.Delete("not-exist")
	obj.Set("broker", "broker-1")
	want := "bar=bar-1;true;lift=lift-1;true;face=face-1;true;broker=broker-1;true;"
	var got string
	for key := range obj.Range() {
		val, ok := obj.Get(key)
		got += fmt.Sprintf("%s=%s;%t;", key, val, ok)
	}
	if got != want {
		t.Errorf("matching error\nwant: %s\ngot:  %s", want, got)
	}
}

func testInsertionAfterDeletion(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	obj.Set("space", "space-1")
	obj.Set("lift", "lift-1")
	obj.Set("face", "face-1")
	obj.Delete("space")
	obj.Set("space", "space-2")
	obj.Delete("foo")
	obj.Set("broker", "broker-1")
	want := "bar=bar-1;true;lift=lift-1;true;face=face-1;true;space=space-2;true;broker=broker-1;true;"
	var got string
	for key := range obj.Range() {
		val, ok := obj.Get(key)
		got += fmt.Sprintf("%s=%s;%t;", key, val, ok)
	}
	if got != want {
		t.Errorf("matching error\nwant: %s\ngot:  %s", want, got)
	}
}

func testUpdated(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	obj.Set("space", "space-1")
	obj.Set("lift", "lift-1")
	obj.Set("face", "face-1")
	obj.Set("broker", "broker-1")
	obj.Set("space", "space-2")
	want := "foo=goo-1;true;bar=bar-1;true;space=space-2;true;lift=lift-1;true;face=face-1;true;broker=broker-1;true;"
	var got string
	for key := range obj.Range() {
		val, ok := obj.Get(key)
		got += fmt.Sprintf("%s=%s;%t;", key, val, ok)
	}
	if got != want {
		t.Errorf("matching error\nwant: %s\ngot:  %s", want, got)
	}
}

func testGetNonExisting(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	want := "space=;false;bar=bar-1;true;"
	var got string
	val, ok := obj.Get("space")
	got += fmt.Sprintf("%s=%s;%t;", "space", val, ok)
	val, ok = obj.Get("bar")
	got += fmt.Sprintf("%s=%s;%t;", "bar", val, ok)

	if got != want {
		t.Errorf("matching error\nwant: %s\ngot:  %s", want, got)
	}
}

func testSpecialRangeWithDeletion(t *testing.T, obj testableOM) {
	obj.Set("foo", "goo-1")
	obj.Set("bar", "bar-1")
	obj.Set("space", "space-1")
	obj.Set("lift", "lift-1")
	obj.Set("face", "face-1")
	obj.Set("broker", "broker-1")
	for key := range obj.Range() {
		obj.Delete(key)
	}
	for key := range obj.Range() {
		t.Errorf("expected empty map, got: %s", key)
	}
}

func Test_OrderedMap(t *testing.T) {
	t.Run("testInsertion", func(t *testing.T) {
		// because of randomly generated keys in the regular map realization
		for n := 0; n < 10; n++ {
			var obj = New()
			testInsertion(t, obj)
		}
	})
	t.Run("testDeletion", func(t *testing.T) {
		var obj = New()
		testDeletion(t, obj)
	})
	t.Run("testInsertionAfterDeletion", func(t *testing.T) {
		var obj = New()
		testInsertionAfterDeletion(t, obj)
	})
	t.Run("testUpdated", func(t *testing.T) {
		var obj = New()
		testUpdated(t, obj)
	})
	t.Run("testGetNonExisting", func(t *testing.T) {
		var obj = New()
		testGetNonExisting(t, obj)
	})
	t.Run("testSpecialRangeWithDeletion", func(t *testing.T) {
		var obj = New()
		testSpecialRangeWithDeletion(t, obj)
	})
	t.Run("testOverhead", func(t *testing.T) {
		var obj = New()
		// fill data
		for n := 0; n < 100; n++ {
			obj.Set(fmt.Sprintf("key-%d", n), fmt.Sprintf("val-%d", n))
		}
		// delete data
		for n := 0; n < 99; n++ {
			obj.Delete(fmt.Sprintf("key-%d", n))
		}
		// just for control
		obj.Set("foo", "bar")
		deleted := float64(len(obj.ordering) - len(obj.data))
		load := float64(len(obj.ordering))
		if deleted/load > 0.25 {
			t.Errorf("deleted: %0.2f, load: %0.2f", deleted, load)
		}
	})
}

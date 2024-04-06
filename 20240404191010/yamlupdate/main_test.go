package main

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v3"
)

type Q struct {
	String        string
	StringSlice   []string
	StringMap     map[string]string
	QMap          map[string]Q
	QMapPointer   map[string]*Q
	QSlice        []Q
	QSlicePointer []*Q
	R             R
	QPointer      *Q
}

type R struct {
	String string
}

func Test1(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}
}

func Test1Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
  qpointer:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val2
  qpointer:
    string: val2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}
}

func Test2(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer == nil {
		t.Fatalf("expected value not to be nil")
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer: null
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer != nil {
		t.Fatalf("expected value to be nil")
	}
}

func Test2Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
  qpointer:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer.QPointer == nil {
		t.Fatalf("expected value not to be nil")
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val2
  qpointer: null
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer.QPointer != nil {
		t.Fatalf("expected value to be nil")
	}
}

func Test3(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	expected := "val2"
	got := q.QPointer.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test3Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qpointer:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qpointer:
    string: val2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	expected := "val2"
	got := q.QPointer.QPointer.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test4(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer: {}
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	expected := "val1"
	got := q.QPointer.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test4Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qpointer:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QPointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qpointer: {}
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QPointer)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	expected := "val1"
	got := q.QPointer.QPointer.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test5(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
string: val1
r:
  string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
string: null
r: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := "val1"
	got := q.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}

	got = q.R.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test5Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: val1
  r:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: null
  r: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := "val1"
	got := q.QPointer.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}

	got = q.QPointer.R.String
	if got != expected {
		t.Fatalf("wrong value, expected: %s got: %s", expected, got)
	}
}

func Test6(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
stringmap:
  k1: y1.1
  k2: y1.2
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
stringmap:
  k2: y2.2
  k3: y2.3
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
		"k2": "y2.2",
		"k3": "y2.3",
	}

	if diff := deep.Equal(q.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test6Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap:
    k1: y1.1
    k2: y1.2
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap:
    k2: y2.2
    k3: y2.3
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
		"k2": "y2.2",
		"k3": "y2.3",
	}

	if diff := deep.Equal(q.QPointer.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test7(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qmappointer:
  q1:
    string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qmappointer:
  q1: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]*Q{
		"q1": nil,
	}

	if diff := deep.Equal(q.QMapPointer, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test7Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmappointer:
    q1:
      string: val1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmappointer:
    q1: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]*Q{
		"q1": nil,
	}

	if diff := deep.Equal(q.QPointer.QMapPointer, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test8(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
stringmap:
  k1: y1.1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
stringmap:
  k1: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
	}

	if diff := deep.Equal(q.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test8Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap:
    k1: y1.1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap:
    k1: null
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
	}

	if diff := deep.Equal(q.QPointer.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test9(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
stringmap:
  k1: y1.1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
stringmap: {}
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
	}

	if diff := deep.Equal(q.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test9Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap:
    k1: y1.1
`), q); err != nil {
		t.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap: {}
`), q); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"k1": "y1.1",
	}

	if diff := deep.Equal(q.QPointer.StringMap, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test10(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	if q.StringMap != nil {
		t.Fatalf("expected map to be nil")
	}

	if q.QMap != nil {
		t.Fatalf("expected map to be nil")
	}

	if q.QMapPointer != nil {
		t.Fatalf("expected map to be nil")
	}
}

func Test10Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer.StringMap != nil {
		t.Fatalf("expected map to be nil")
	}

	if q.QPointer.QMap != nil {
		t.Fatalf("expected map to be nil")
	}

	if q.QPointer.QMapPointer != nil {
		t.Fatalf("expected map to be nil")
	}
}

func Test11(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
stringmap: {}
qmap: {}
qmappointer: {}
`), q); err != nil {
		t.Fatal(err)
	}

	if q.StringMap == nil {
		t.Fatalf("expected map to be initialized")
	}

	if q.QMap == nil {
		t.Fatalf("expected map to be initialized")
	}

	if q.QMapPointer == nil {
		t.Fatalf("expected map to be initialized")
	}
}

func Test11Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringmap: {}
  qmap: {}
  qmappointer: {}
`), q); err != nil {
		t.Fatal(err)
	}

	if q.QPointer.StringMap == nil {
		t.Fatalf("expected map to be initialized")
	}

	if q.QPointer.QMap == nil {
		t.Fatalf("expected map to be initialized")
	}

	if q.QPointer.QMapPointer == nil {
		t.Fatalf("expected map to be initialized")
	}
}

func Test12(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qmappointer:
  q1:
    string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QMapPointer["q1"])

	if err := yaml.Unmarshal([]byte(`---
qmappointer:
  q1:
    string: v2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QMapPointer["q1"])
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	if q.QMapPointer["q1"].String != "v2" {
		t.Fatalf("incorrect value")
	}
}

func Test12Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmappointer:
    q1:
      string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QMapPointer["q1"])

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmappointer:
    q1:
      string: v2
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QMapPointer["q1"])
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	if q.QPointer.QMapPointer["q1"].String != "v2" {
		t.Fatalf("incorrect value")
	}
}

func Test13(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qmap:
  q1:
    string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	v := q.QMap["q1"]
	addr := fmt.Sprintf("%p", &v)

	if err := yaml.Unmarshal([]byte(`---
qmap:
  q1:
    string: v2
`), q); err != nil {
		t.Fatal(err)
	}

	v = q.QMap["q1"]
	newAddr := fmt.Sprintf("%p", &v)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	if q.QMap["q1"].String != "v2" {
		t.Fatalf("incorrect value")
	}
}

func Test13Nested(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmap:
    q1:
      string: v1
`), q); err != nil {
		t.Fatal(err)
	}

	v := q.QPointer.QMap["q1"]
	addr := fmt.Sprintf("%p", &v)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qmap:
    q1:
      string: v2
`), q); err != nil {
		t.Fatal(err)
	}

	v = q.QPointer.QMap["q1"]
	newAddr := fmt.Sprintf("%p", &v)
	if addr != newAddr {
		t.Fatalf("pointers don't match: %s %s", addr, newAddr)
	}

	if q.QPointer.QMap["q1"].String != "v2" {
		t.Fatalf("incorrect value")
	}
}

func Test14_Primitive(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
stringslice:
- v1
- v2
- v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.StringSlice)

	if err := yaml.Unmarshal([]byte(`---
stringslice:
- v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.StringSlice)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []string{
		"v4",
	}

	if diff := deep.Equal(q.StringSlice, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test14Nested_Primitive(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringslice:
  - v1
  - v2
  - v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.StringSlice)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  stringslice:
  - v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.StringSlice)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []string{
		"v4",
	}

	if diff := deep.Equal(q.QPointer.StringSlice, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test14_Struct(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qslice:
- string: v1
- string: v2
- string: v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QSlice)

	if err := yaml.Unmarshal([]byte(`---
qslice:
- string: v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QSlice)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []Q{
		Q{String: "v4"},
	}

	if diff := deep.Equal(q.QSlice, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test14Nested_Struct(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qslice:
  - string: v1
  - string: v2
  - string: v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QSlice)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qslice:
  - string: v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QSlice)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []Q{
		Q{String: "v4"},
	}

	if diff := deep.Equal(q.QPointer.QSlice, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test14_Pointer(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qslicepointer:
- string: v1
- string: v2
- string: v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QSlicePointer)

	if err := yaml.Unmarshal([]byte(`---
qslicepointer:
- string: v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QSlicePointer)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []*Q{
		&Q{String: "v4"},
	}

	if diff := deep.Equal(q.QSlicePointer, expected); diff != nil {
		t.Fatal(diff)
	}
}

func Test14Nested_Pointer(t *testing.T) {
	q := &Q{}
	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qslicepointer:
  - string: v1
  - string: v2
  - string: v3
`), q); err != nil {
		t.Fatal(err)
	}

	addr := fmt.Sprintf("%p", q.QPointer.QSlicePointer)

	if err := yaml.Unmarshal([]byte(`---
qpointer:
  qslicepointer:
  - string: v4
`), q); err != nil {
		t.Fatal(err)
	}

	newAddr := fmt.Sprintf("%p", q.QPointer.QSlicePointer)
	if addr == newAddr {
		t.Fatalf("pointers should not match: %s", addr)
	}

	expected := []*Q{
		&Q{String: "v4"},
	}

	if diff := deep.Equal(q.QPointer.QSlicePointer, expected); diff != nil {
		t.Fatal(diff)
	}
}

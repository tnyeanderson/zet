package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var loremipsum []byte
var short = []byte("aaaaaaaaaaaa some other text aaaaa a naa aan aa repeated aaaaa")

func init() {
	b, err := os.ReadFile("testdata/lorem-ipsum.txt")
	if err != nil {
		panic(err)
	}
	loremipsum = b
}

func TestRemoveRepeatsWithRegex(t *testing.T) {
	expected := "aa some other text aa a naa aan aa repeated aa"
	b := removeRepeatsWithRegex(short, 'a', 2)
	if string(b) != expected {
		t.Fatal()
	}
}

func TestRemoveRepeatsWithReader(t *testing.T) {
	expected := "aa some other text aa a naa aan aa repeated aa"
	reader := bytes.NewReader(short)
	r := newRepeatRemover(reader, 'a', 2)
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != expected {
		t.Fatal()
	}
}

func TestRemoveRepeatsWithReader_NoOp(t *testing.T) {
	expected := "this text doesn't contain any repeats that would match"
	reader := bytes.NewReader([]byte(expected))
	r := newRepeatRemover(reader, 'a', 3)
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != expected {
		t.Fatal()
	}
}

func TestRemoveRepeatsWithRegex_LoremIpsum(t *testing.T) {
	b := removeRepeatsWithRegex(loremipsum, 'l', 1)
	// There are 31 times that "ll" appears in the text
	if len(b) != len(loremipsum)-31 {
		t.Fatal()
	}
}

func TestRemoveRepeatsWithReader_LoremIpsum(t *testing.T) {
	reader := bytes.NewReader(loremipsum)
	r := newRepeatRemover(reader, 'l', 1)
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	// There are 31 times that "ll" appears in the text
	if len(b) != len(loremipsum)-31 {
		t.Fatal()
	}
}

func TestRemoveRepeatsWithReader_Partial(t *testing.T) {
	expected := "aa some other text aa a naa aan aa repeated aa"
	reader := bytes.NewReader(short)
	r := newRepeatRemover(reader, 'a', 2)
	// Read the first 2 bytes, stopping in the middle of a repeat
	buf := make([]byte, 2)
	_, err := r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	// Read the rest of the data
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	got := string(buf) + string(b)
	if got != expected {
		t.Fatalf("expected: %s\n\ngot: %s\n", expected, got)
	}
}

func BenchmarkRemoveRepeatsWithRegex(b *testing.B) {
	for b.Loop() {
		removeRepeatsWithRegex(short, 'a', 4)
	}
}

func BenchmarkRemoveRepeatsWithReader(b *testing.B) {
	reader := bytes.NewReader(short)
	r := newRepeatRemover(reader, 'a', 4)
	for b.Loop() {
		io.ReadAll(r)
	}
}

func BenchmarkRemoveRepeatsWithRegex_LoremIpsum(b *testing.B) {
	for b.Loop() {
		removeRepeatsWithRegex(loremipsum, 'l', 1)
	}
}

func BenchmarkRemoveRepeatsWithReader_LoremIpsum(b *testing.B) {
	reader := bytes.NewReader(loremipsum)
	r := newRepeatRemover(reader, 'l', 1)
	for b.Loop() {
		io.ReadAll(r)
	}
}

func BenchmarkRemoveRepeatsWithRegex_LoremIpsum10x(b *testing.B) {
	for b.Loop() {
		removeRepeatsWithRegex(bytes.Repeat(loremipsum, 10), 'l', 1)
	}
}

func BenchmarkRemoveRepeatsWithReader_LoremIpsum10x(b *testing.B) {
	reader := bytes.NewReader(bytes.Repeat(loremipsum, 10))
	r := newRepeatRemover(reader, 'l', 1)
	for b.Loop() {
		io.ReadAll(r)
	}
}

func BenchmarkRemoveRepeatsWithRegex_LoremIpsum100x(b *testing.B) {
	for b.Loop() {
		removeRepeatsWithRegex(bytes.Repeat(loremipsum, 100), 'l', 1)
	}
}

func BenchmarkRemoveRepeatsWithReader_LoremIpsum100x(b *testing.B) {
	reader := bytes.NewReader(bytes.Repeat(loremipsum, 100))
	r := newRepeatRemover(reader, 'l', 1)
	for b.Loop() {
		io.ReadAll(r)
	}
}

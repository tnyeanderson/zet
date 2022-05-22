// This file is meant to demonstrate a question I have about embedded structs in go.
// I have one struct (ContainingStruct) which embeds another struct (EmbeddedStruct).
// The structs both implement an interface (Interfacer) which has two methods (Labeler and Printer).
// Labeler() is implemented uniquely for each struct.
// However, Printer() is only implemented by the EmbeddedStruct.
// Printer() prints the results of Labeler().
// If I call ContainingStruct.Printer(), it uses EmbeddedStruct.Labeler() for its result.
// This is because Printer() is a method on EmbeddedStruct only, so it calls the corresponding Labeler() signature
// Instead, I want it to use ContainingStruct.Labeler() for its result.
// Is this possible? Also see the QUESTION block below if more explanation is necessary!!
package main

import (
	"fmt"
)

// Interfacer is implemented by EmbeddedStruct and ContainingStruct
type Interfacer interface {
	Labeler() string
	Printer()
}

// EmbeddedStruct is embedded into ContainingStruct
type EmbeddedStruct struct{}

// ContainingStruct has EmbeddedStruct embedded in it
type ContainingStruct struct {
	EmbeddedStruct
}

// Printer is only implemented in the EmbeddedStruct
// It is "inherited" (not really) by ContainingStruct
func (e EmbeddedStruct) Printer() {
	fmt.Println(e.Labeler())
}

// Labeler is implemented in both EmbeddedStruct and ContainingStruct
func (e EmbeddedStruct) Labeler() string {
	return "Embedded Struct String!"
}

// Labeler is implemented in both EmbeddedStruct and ContainingStruct
func (c ContainingStruct) Labeler() string {
	return "Containing Struct String!"
}

// This calls Printer() on the provided Interfacer
func PrintResults(i Interfacer) {
	i.Printer()
}

func main() {
	// Create a ContainingStruct which will call Printer() on itself
	c := ContainingStruct{}

	// QUESTION:
	// I want this to print 'Containing Struct String!'
	// Since c is a ContainingStruct which implements Labeler
	// Instead it prints 'Embedded Struct String!'
	// I understand why this happens
	// But is what I want even possible in go?
	// What is the go way to do this??
	PrintResults(c)
}

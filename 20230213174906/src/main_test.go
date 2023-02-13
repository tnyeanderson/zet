package main

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshalStruct(b *testing.B) {
	result := []State{}
	json.Unmarshal(Bytes, &result)
}

func BenchmarkUnmarshalInterface(b *testing.B) {
	result := []interface{}{}
	json.Unmarshal(Bytes, &result)
}

func BenchmarkUnmarshalStructShort(b *testing.B) {
	result := []State{}
	json.Unmarshal(ShortBytes, &result)
}

func BenchmarkUnmarshalInterfaceShort(b *testing.B) {
	result := []interface{}{}
	json.Unmarshal(ShortBytes, &result)
}

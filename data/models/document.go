package models

// Compatibility between other providers like, mongodb, dinamoDB or other NoSQL providers
type Document []DocumentKey

type DocumentKey struct {
	Field string
	Value interface{}
}

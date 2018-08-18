package object

import "fmt"

type ObjectType string

const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return IntegerObj }
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string { return BooleanObj }
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

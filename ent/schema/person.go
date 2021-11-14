package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}

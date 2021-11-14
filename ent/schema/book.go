package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.Time("created_at"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Person.Type).Ref("books").Unique(),
	}
}

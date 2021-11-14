package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname"),
		field.String("email"),
		field.Int("scores"),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return nil
}

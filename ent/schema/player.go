package schema

import "entgo.io/ent"

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return nil
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return nil
}

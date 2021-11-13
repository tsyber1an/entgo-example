package main

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Player struct {
	ent.Schema
}

func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname"),
		field.String("email"),
		field.Int("scores"),
	}
}

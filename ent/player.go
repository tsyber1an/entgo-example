// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Funfun/entgo-example/ent/player"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Scores holds the value of the "scores" field.
	Scores int `json:"scores,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldID, player.FieldScores:
			values[i] = new(sql.NullInt64)
		case player.FieldNickname, player.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Player", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case player.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				pl.Nickname = value.String
			}
		case player.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pl.Email = value.String
			}
		case player.FieldScores:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field scores", values[i])
			} else if value.Valid {
				pl.Scores = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return (&PlayerClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", nickname=")
	builder.WriteString(pl.Nickname)
	builder.WriteString(", email=")
	builder.WriteString(pl.Email)
	builder.WriteString(", scores=")
	builder.WriteString(fmt.Sprintf("%v", pl.Scores))
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player

func (pl Players) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}

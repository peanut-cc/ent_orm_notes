package schema

import (
	"database/sql"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Amount is a custom Go type that's convertible to the basic float64 type.
type Amount float64

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").
			GoType(Amount(0)),
		field.String("name").
			Optional().
			GoType(&sql.NullString{}),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}

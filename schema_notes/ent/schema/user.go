package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

/*
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`json:"oid,omitempty"`),
	}
}

*/

/*
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
	}
}

*/

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("required_name"),
		field.String("optional_name").Optional(),
		field.String("nilable_name").Optional().Nillable(),
		field.String("nilable_name2").Optional().Nillable(),
		field.Int("age").Optional(),
		field.Int("age2").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

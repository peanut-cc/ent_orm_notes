package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("field1"),
		field.String("field2"),
		field.String("field3").Unique(),
		field.String("first_name"),
		field.String("last_name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("field1", "field2"),
		// unique index
		index.Fields("first_name", "last_name").Unique(),
	}
}

package schema

import (
	"errors"
	"regexp"
	"strings"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("group name must begin with uppercase")
				}
				return nil
			}),
		field.Int("age").Positive(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}

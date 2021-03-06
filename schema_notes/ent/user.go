// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/peanut-cc/ent_orm_notes/schema_notes/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RequiredName holds the value of the "required_name" field.
	RequiredName string `json:"required_name,omitempty"`
	// OptionalName holds the value of the "optional_name" field.
	OptionalName string `json:"optional_name,omitempty"`
	// NilableName holds the value of the "nilable_name" field.
	NilableName *string `json:"nilable_name,omitempty"`
	// NilableName2 holds the value of the "nilable_name2" field.
	NilableName2 *string `json:"nilable_name2,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Age2 holds the value of the "age2" field.
	Age2 *int `json:"age2,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // required_name
		&sql.NullString{}, // optional_name
		&sql.NullString{}, // nilable_name
		&sql.NullString{}, // nilable_name2
		&sql.NullInt64{},  // age
		&sql.NullInt64{},  // age2
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field required_name", values[0])
	} else if value.Valid {
		u.RequiredName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field optional_name", values[1])
	} else if value.Valid {
		u.OptionalName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nilable_name", values[2])
	} else if value.Valid {
		u.NilableName = new(string)
		*u.NilableName = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nilable_name2", values[3])
	} else if value.Valid {
		u.NilableName2 = new(string)
		*u.NilableName2 = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[4])
	} else if value.Valid {
		u.Age = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age2", values[5])
	} else if value.Valid {
		u.Age2 = new(int)
		*u.Age2 = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", required_name=")
	builder.WriteString(u.RequiredName)
	builder.WriteString(", optional_name=")
	builder.WriteString(u.OptionalName)
	if v := u.NilableName; v != nil {
		builder.WriteString(", nilable_name=")
		builder.WriteString(*v)
	}
	if v := u.NilableName2; v != nil {
		builder.WriteString(", nilable_name2=")
		builder.WriteString(*v)
	}
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	if v := u.Age2; v != nil {
		builder.WriteString(", age2=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}

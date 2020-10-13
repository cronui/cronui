package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(func() uuid.UUID {
				id, _ := uuid.NewRandom()
				return id
			}),
		field.String("username").
			Unique().
			MinLen(4).
			MaxLen(36),
		field.String("password").
			MaxLen(60),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at").Annotations(entsql.Default("now()")),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// create an inverse-edge called 'owner' of type 'User'
		// and reference it to the 'cars' edge (in User schema)
		// explicitly using the 'Ref' method
		edge.From("owner", User.Type).
			Ref("cars").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}

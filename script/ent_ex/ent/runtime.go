// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"example.com/ent_ex/ent/car"
	"example.com/ent_ex/ent/group"
	"example.com/ent_ex/ent/schema"
	"example.com/ent_ex/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescRegisteredAt is the schema descriptor for registered_at field.
	carDescRegisteredAt := carFields[2].Descriptor()
	// car.DefaultRegisteredAt holds the default value on creation for the registered_at field.
	car.DefaultRegisteredAt = carDescRegisteredAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
}

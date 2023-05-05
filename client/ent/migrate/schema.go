// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 8},
		{Name: "phone_number", Type: field.TypeString, Size: 10},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "license_id", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString},
		{Name: "membership_number", Type: field.TypeInt},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "password", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SUBSCRIBER", "ADMIN"}, Default: "SUBSCRIBER"},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// FlightsColumns holds the columns for the "flights" table.
	FlightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "depart_date", Type: field.TypeTime},
		{Name: "depart_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ON_TIME", "DELAYED", "CANCELLED", "SCHEDULED"}, Default: "SCHEDULED"},
		{Name: "available_slots", Type: field.TypeInt},
		{Name: "return_date", Type: field.TypeTime},
		{Name: "flight_plane", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FlightsTable holds the schema information for the "flights" table.
	FlightsTable = &schema.Table{
		Name:       "flights",
		Columns:    FlightsColumns,
		PrimaryKey: []*schema.Column{FlightsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CustomersTable,
		FlightsTable,
	}
)

func init() {
}

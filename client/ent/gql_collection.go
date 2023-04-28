// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer/ent/customer"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CustomerQuery) CollectFields(ctx context.Context, satisfies ...string) (*CustomerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CustomerQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(customer.Columns))
		selectedFields = []string{customer.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[customer.FieldName]; !ok {
				selectedFields = append(selectedFields, customer.FieldName)
				fieldSeen[customer.FieldName] = struct{}{}
			}
		case "phoneNumber":
			if _, ok := fieldSeen[customer.FieldPhoneNumber]; !ok {
				selectedFields = append(selectedFields, customer.FieldPhoneNumber)
				fieldSeen[customer.FieldPhoneNumber] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[customer.FieldEmail]; !ok {
				selectedFields = append(selectedFields, customer.FieldEmail)
				fieldSeen[customer.FieldEmail] = struct{}{}
			}
		case "licenseID":
			if _, ok := fieldSeen[customer.FieldLicenseID]; !ok {
				selectedFields = append(selectedFields, customer.FieldLicenseID)
				fieldSeen[customer.FieldLicenseID] = struct{}{}
			}
		case "address":
			if _, ok := fieldSeen[customer.FieldAddress]; !ok {
				selectedFields = append(selectedFields, customer.FieldAddress)
				fieldSeen[customer.FieldAddress] = struct{}{}
			}
		case "membershipNumber":
			if _, ok := fieldSeen[customer.FieldMembershipNumber]; !ok {
				selectedFields = append(selectedFields, customer.FieldMembershipNumber)
				fieldSeen[customer.FieldMembershipNumber] = struct{}{}
			}
		case "isActive":
			if _, ok := fieldSeen[customer.FieldIsActive]; !ok {
				selectedFields = append(selectedFields, customer.FieldIsActive)
				fieldSeen[customer.FieldIsActive] = struct{}{}
			}
		case "password":
			if _, ok := fieldSeen[customer.FieldPassword]; !ok {
				selectedFields = append(selectedFields, customer.FieldPassword)
				fieldSeen[customer.FieldPassword] = struct{}{}
			}
		case "createAt":
			if _, ok := fieldSeen[customer.FieldCreateAt]; !ok {
				selectedFields = append(selectedFields, customer.FieldCreateAt)
				fieldSeen[customer.FieldCreateAt] = struct{}{}
			}
		case "updateAt":
			if _, ok := fieldSeen[customer.FieldUpdateAt]; !ok {
				selectedFields = append(selectedFields, customer.FieldUpdateAt)
				fieldSeen[customer.FieldUpdateAt] = struct{}{}
			}
		case "role":
			if _, ok := fieldSeen[customer.FieldRole]; !ok {
				selectedFields = append(selectedFields, customer.FieldRole)
				fieldSeen[customer.FieldRole] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type customerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CustomerPaginateOption
}

func newCustomerPaginateArgs(rv map[string]interface{}) *customerPaginateArgs {
	args := &customerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}

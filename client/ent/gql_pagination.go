// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"customer/ent/booking"
	"customer/ent/customer"
	"customer/ent/flight"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// BookingEdge is the edge representation of Booking.
type BookingEdge struct {
	Node   *Booking `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// BookingConnection is the connection containing edges to Booking.
type BookingConnection struct {
	Edges      []*BookingEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (c *BookingConnection) build(nodes []*Booking, pager *bookingPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Booking
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Booking {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Booking {
			return nodes[i]
		}
	}
	c.Edges = make([]*BookingEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &BookingEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// BookingPaginateOption enables pagination customization.
type BookingPaginateOption func(*bookingPager) error

// WithBookingOrder configures pagination ordering.
func WithBookingOrder(order *BookingOrder) BookingPaginateOption {
	if order == nil {
		order = DefaultBookingOrder
	}
	o := *order
	return func(pager *bookingPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultBookingOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithBookingFilter configures pagination filter.
func WithBookingFilter(filter func(*BookingQuery) (*BookingQuery, error)) BookingPaginateOption {
	return func(pager *bookingPager) error {
		if filter == nil {
			return errors.New("BookingQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type bookingPager struct {
	reverse bool
	order   *BookingOrder
	filter  func(*BookingQuery) (*BookingQuery, error)
}

func newBookingPager(opts []BookingPaginateOption, reverse bool) (*bookingPager, error) {
	pager := &bookingPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultBookingOrder
	}
	return pager, nil
}

func (p *bookingPager) applyFilter(query *BookingQuery) (*BookingQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *bookingPager) toCursor(b *Booking) Cursor {
	return p.order.Field.toCursor(b)
}

func (p *bookingPager) applyCursors(query *BookingQuery, after, before *Cursor) (*BookingQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultBookingOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *bookingPager) applyOrder(query *BookingQuery) *BookingQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultBookingOrder.Field {
		query = query.Order(DefaultBookingOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *bookingPager) orderExpr(query *BookingQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultBookingOrder.Field {
			b.Comma().Ident(DefaultBookingOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Booking.
func (b *BookingQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...BookingPaginateOption,
) (*BookingConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newBookingPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if b, err = pager.applyFilter(b); err != nil {
		return nil, err
	}
	conn := &BookingConnection{Edges: []*BookingEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = b.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if b, err = pager.applyCursors(b, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		b.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := b.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	b = pager.applyOrder(b)
	nodes, err := b.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// BookingOrderField defines the ordering field of Booking.
type BookingOrderField struct {
	// Value extracts the ordering value from the given Booking.
	Value    func(*Booking) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) booking.OrderOption
	toCursor func(*Booking) Cursor
}

// BookingOrder defines the ordering of Booking.
type BookingOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *BookingOrderField `json:"field"`
}

// DefaultBookingOrder is the default ordering of Booking.
var DefaultBookingOrder = &BookingOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &BookingOrderField{
		Value: func(b *Booking) (ent.Value, error) {
			return b.ID, nil
		},
		column: booking.FieldID,
		toTerm: booking.ByID,
		toCursor: func(b *Booking) Cursor {
			return Cursor{ID: b.ID}
		},
	},
}

// ToEdge converts Booking into BookingEdge.
func (b *Booking) ToEdge(order *BookingOrder) *BookingEdge {
	if order == nil {
		order = DefaultBookingOrder
	}
	return &BookingEdge{
		Node:   b,
		Cursor: order.Field.toCursor(b),
	}
}

// CustomerEdge is the edge representation of Customer.
type CustomerEdge struct {
	Node   *Customer `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// CustomerConnection is the connection containing edges to Customer.
type CustomerConnection struct {
	Edges      []*CustomerEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *CustomerConnection) build(nodes []*Customer, pager *customerPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Customer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Customer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Customer {
			return nodes[i]
		}
	}
	c.Edges = make([]*CustomerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CustomerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CustomerPaginateOption enables pagination customization.
type CustomerPaginateOption func(*customerPager) error

// WithCustomerOrder configures pagination ordering.
func WithCustomerOrder(order *CustomerOrder) CustomerPaginateOption {
	if order == nil {
		order = DefaultCustomerOrder
	}
	o := *order
	return func(pager *customerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCustomerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCustomerFilter configures pagination filter.
func WithCustomerFilter(filter func(*CustomerQuery) (*CustomerQuery, error)) CustomerPaginateOption {
	return func(pager *customerPager) error {
		if filter == nil {
			return errors.New("CustomerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type customerPager struct {
	reverse bool
	order   *CustomerOrder
	filter  func(*CustomerQuery) (*CustomerQuery, error)
}

func newCustomerPager(opts []CustomerPaginateOption, reverse bool) (*customerPager, error) {
	pager := &customerPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCustomerOrder
	}
	return pager, nil
}

func (p *customerPager) applyFilter(query *CustomerQuery) (*CustomerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *customerPager) toCursor(c *Customer) Cursor {
	return p.order.Field.toCursor(c)
}

func (p *customerPager) applyCursors(query *CustomerQuery, after, before *Cursor) (*CustomerQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCustomerOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *customerPager) applyOrder(query *CustomerQuery) *CustomerQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCustomerOrder.Field {
		query = query.Order(DefaultCustomerOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *customerPager) orderExpr(query *CustomerQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCustomerOrder.Field {
			b.Comma().Ident(DefaultCustomerOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Customer.
func (c *CustomerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CustomerPaginateOption,
) (*CustomerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCustomerPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if c, err = pager.applyFilter(c); err != nil {
		return nil, err
	}
	conn := &CustomerConnection{Edges: []*CustomerEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = c.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if c, err = pager.applyCursors(c, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		c.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := c.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	c = pager.applyOrder(c)
	nodes, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// CustomerOrderField defines the ordering field of Customer.
type CustomerOrderField struct {
	// Value extracts the ordering value from the given Customer.
	Value    func(*Customer) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) customer.OrderOption
	toCursor func(*Customer) Cursor
}

// CustomerOrder defines the ordering of Customer.
type CustomerOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *CustomerOrderField `json:"field"`
}

// DefaultCustomerOrder is the default ordering of Customer.
var DefaultCustomerOrder = &CustomerOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CustomerOrderField{
		Value: func(c *Customer) (ent.Value, error) {
			return c.ID, nil
		},
		column: customer.FieldID,
		toTerm: customer.ByID,
		toCursor: func(c *Customer) Cursor {
			return Cursor{ID: c.ID}
		},
	},
}

// ToEdge converts Customer into CustomerEdge.
func (c *Customer) ToEdge(order *CustomerOrder) *CustomerEdge {
	if order == nil {
		order = DefaultCustomerOrder
	}
	return &CustomerEdge{
		Node:   c,
		Cursor: order.Field.toCursor(c),
	}
}

// FlightEdge is the edge representation of Flight.
type FlightEdge struct {
	Node   *Flight `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// FlightConnection is the connection containing edges to Flight.
type FlightConnection struct {
	Edges      []*FlightEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *FlightConnection) build(nodes []*Flight, pager *flightPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Flight
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Flight {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Flight {
			return nodes[i]
		}
	}
	c.Edges = make([]*FlightEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &FlightEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// FlightPaginateOption enables pagination customization.
type FlightPaginateOption func(*flightPager) error

// WithFlightOrder configures pagination ordering.
func WithFlightOrder(order *FlightOrder) FlightPaginateOption {
	if order == nil {
		order = DefaultFlightOrder
	}
	o := *order
	return func(pager *flightPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultFlightOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithFlightFilter configures pagination filter.
func WithFlightFilter(filter func(*FlightQuery) (*FlightQuery, error)) FlightPaginateOption {
	return func(pager *flightPager) error {
		if filter == nil {
			return errors.New("FlightQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type flightPager struct {
	reverse bool
	order   *FlightOrder
	filter  func(*FlightQuery) (*FlightQuery, error)
}

func newFlightPager(opts []FlightPaginateOption, reverse bool) (*flightPager, error) {
	pager := &flightPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultFlightOrder
	}
	return pager, nil
}

func (p *flightPager) applyFilter(query *FlightQuery) (*FlightQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *flightPager) toCursor(f *Flight) Cursor {
	return p.order.Field.toCursor(f)
}

func (p *flightPager) applyCursors(query *FlightQuery, after, before *Cursor) (*FlightQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultFlightOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *flightPager) applyOrder(query *FlightQuery) *FlightQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultFlightOrder.Field {
		query = query.Order(DefaultFlightOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *flightPager) orderExpr(query *FlightQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultFlightOrder.Field {
			b.Comma().Ident(DefaultFlightOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Flight.
func (f *FlightQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...FlightPaginateOption,
) (*FlightConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newFlightPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if f, err = pager.applyFilter(f); err != nil {
		return nil, err
	}
	conn := &FlightConnection{Edges: []*FlightEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = f.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if f, err = pager.applyCursors(f, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		f.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := f.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	f = pager.applyOrder(f)
	nodes, err := f.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// FlightOrderField defines the ordering field of Flight.
type FlightOrderField struct {
	// Value extracts the ordering value from the given Flight.
	Value    func(*Flight) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) flight.OrderOption
	toCursor func(*Flight) Cursor
}

// FlightOrder defines the ordering of Flight.
type FlightOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *FlightOrderField `json:"field"`
}

// DefaultFlightOrder is the default ordering of Flight.
var DefaultFlightOrder = &FlightOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &FlightOrderField{
		Value: func(f *Flight) (ent.Value, error) {
			return f.ID, nil
		},
		column: flight.FieldID,
		toTerm: flight.ByID,
		toCursor: func(f *Flight) Cursor {
			return Cursor{ID: f.ID}
		},
	},
}

// ToEdge converts Flight into FlightEdge.
func (f *Flight) ToEdge(order *FlightOrder) *FlightEdge {
	if order == nil {
		order = DefaultFlightOrder
	}
	return &FlightEdge{
		Node:   f,
		Cursor: order.Field.toCursor(f),
	}
}

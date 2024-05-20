// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
)

// XProviderCursorQuery is the builder for querying XProviderCursor entities.
type XProviderCursorQuery struct {
	config
	ctx        *QueryContext
	order      []xprovidercursor.OrderOption
	inters     []Interceptor
	predicates []predicate.XProviderCursor
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the XProviderCursorQuery builder.
func (xcq *XProviderCursorQuery) Where(ps ...predicate.XProviderCursor) *XProviderCursorQuery {
	xcq.predicates = append(xcq.predicates, ps...)
	return xcq
}

// Limit the number of records to be returned by this query.
func (xcq *XProviderCursorQuery) Limit(limit int) *XProviderCursorQuery {
	xcq.ctx.Limit = &limit
	return xcq
}

// Offset to start from.
func (xcq *XProviderCursorQuery) Offset(offset int) *XProviderCursorQuery {
	xcq.ctx.Offset = &offset
	return xcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (xcq *XProviderCursorQuery) Unique(unique bool) *XProviderCursorQuery {
	xcq.ctx.Unique = &unique
	return xcq
}

// Order specifies how the records should be ordered.
func (xcq *XProviderCursorQuery) Order(o ...xprovidercursor.OrderOption) *XProviderCursorQuery {
	xcq.order = append(xcq.order, o...)
	return xcq
}

// First returns the first XProviderCursor entity from the query.
// Returns a *NotFoundError when no XProviderCursor was found.
func (xcq *XProviderCursorQuery) First(ctx context.Context) (*XProviderCursor, error) {
	nodes, err := xcq.Limit(1).All(setContextOp(ctx, xcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{xprovidercursor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (xcq *XProviderCursorQuery) FirstX(ctx context.Context) *XProviderCursor {
	node, err := xcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first XProviderCursor ID from the query.
// Returns a *NotFoundError when no XProviderCursor ID was found.
func (xcq *XProviderCursorQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = xcq.Limit(1).IDs(setContextOp(ctx, xcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{xprovidercursor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (xcq *XProviderCursorQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := xcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single XProviderCursor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one XProviderCursor entity is found.
// Returns a *NotFoundError when no XProviderCursor entities are found.
func (xcq *XProviderCursorQuery) Only(ctx context.Context) (*XProviderCursor, error) {
	nodes, err := xcq.Limit(2).All(setContextOp(ctx, xcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{xprovidercursor.Label}
	default:
		return nil, &NotSingularError{xprovidercursor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (xcq *XProviderCursorQuery) OnlyX(ctx context.Context) *XProviderCursor {
	node, err := xcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only XProviderCursor ID in the query.
// Returns a *NotSingularError when more than one XProviderCursor ID is found.
// Returns a *NotFoundError when no entities are found.
func (xcq *XProviderCursorQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = xcq.Limit(2).IDs(setContextOp(ctx, xcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{xprovidercursor.Label}
	default:
		err = &NotSingularError{xprovidercursor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (xcq *XProviderCursorQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := xcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of XProviderCursors.
func (xcq *XProviderCursorQuery) All(ctx context.Context) ([]*XProviderCursor, error) {
	ctx = setContextOp(ctx, xcq.ctx, "All")
	if err := xcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*XProviderCursor, *XProviderCursorQuery]()
	return withInterceptors[[]*XProviderCursor](ctx, xcq, qr, xcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (xcq *XProviderCursorQuery) AllX(ctx context.Context) []*XProviderCursor {
	nodes, err := xcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of XProviderCursor IDs.
func (xcq *XProviderCursorQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if xcq.ctx.Unique == nil && xcq.path != nil {
		xcq.Unique(true)
	}
	ctx = setContextOp(ctx, xcq.ctx, "IDs")
	if err = xcq.Select(xprovidercursor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (xcq *XProviderCursorQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := xcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (xcq *XProviderCursorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, xcq.ctx, "Count")
	if err := xcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, xcq, querierCount[*XProviderCursorQuery](), xcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (xcq *XProviderCursorQuery) CountX(ctx context.Context) int {
	count, err := xcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (xcq *XProviderCursorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, xcq.ctx, "Exist")
	switch _, err := xcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (xcq *XProviderCursorQuery) ExistX(ctx context.Context) bool {
	exist, err := xcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the XProviderCursorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (xcq *XProviderCursorQuery) Clone() *XProviderCursorQuery {
	if xcq == nil {
		return nil
	}
	return &XProviderCursorQuery{
		config:     xcq.config,
		ctx:        xcq.ctx.Clone(),
		order:      append([]xprovidercursor.OrderOption{}, xcq.order...),
		inters:     append([]Interceptor{}, xcq.inters...),
		predicates: append([]predicate.XProviderCursor{}, xcq.predicates...),
		// clone intermediate query.
		sql:  xcq.sql.Clone(),
		path: xcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChainID uint64 `json:"chain_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.XProviderCursor.Query().
//		GroupBy(xprovidercursor.FieldChainID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (xcq *XProviderCursorQuery) GroupBy(field string, fields ...string) *XProviderCursorGroupBy {
	xcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &XProviderCursorGroupBy{build: xcq}
	grbuild.flds = &xcq.ctx.Fields
	grbuild.label = xprovidercursor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChainID uint64 `json:"chain_id,omitempty"`
//	}
//
//	client.XProviderCursor.Query().
//		Select(xprovidercursor.FieldChainID).
//		Scan(ctx, &v)
func (xcq *XProviderCursorQuery) Select(fields ...string) *XProviderCursorSelect {
	xcq.ctx.Fields = append(xcq.ctx.Fields, fields...)
	sbuild := &XProviderCursorSelect{XProviderCursorQuery: xcq}
	sbuild.label = xprovidercursor.Label
	sbuild.flds, sbuild.scan = &xcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a XProviderCursorSelect configured with the given aggregations.
func (xcq *XProviderCursorQuery) Aggregate(fns ...AggregateFunc) *XProviderCursorSelect {
	return xcq.Select().Aggregate(fns...)
}

func (xcq *XProviderCursorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range xcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, xcq); err != nil {
				return err
			}
		}
	}
	for _, f := range xcq.ctx.Fields {
		if !xprovidercursor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if xcq.path != nil {
		prev, err := xcq.path(ctx)
		if err != nil {
			return err
		}
		xcq.sql = prev
	}
	return nil
}

func (xcq *XProviderCursorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*XProviderCursor, error) {
	var (
		nodes = []*XProviderCursor{}
		_spec = xcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*XProviderCursor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &XProviderCursor{config: xcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, xcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (xcq *XProviderCursorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := xcq.querySpec()
	_spec.Node.Columns = xcq.ctx.Fields
	if len(xcq.ctx.Fields) > 0 {
		_spec.Unique = xcq.ctx.Unique != nil && *xcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, xcq.driver, _spec)
}

func (xcq *XProviderCursorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(xprovidercursor.Table, xprovidercursor.Columns, sqlgraph.NewFieldSpec(xprovidercursor.FieldID, field.TypeUUID))
	_spec.From = xcq.sql
	if unique := xcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if xcq.path != nil {
		_spec.Unique = true
	}
	if fields := xcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xprovidercursor.FieldID)
		for i := range fields {
			if fields[i] != xprovidercursor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := xcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := xcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := xcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := xcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (xcq *XProviderCursorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(xcq.driver.Dialect())
	t1 := builder.Table(xprovidercursor.Table)
	columns := xcq.ctx.Fields
	if len(columns) == 0 {
		columns = xprovidercursor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if xcq.sql != nil {
		selector = xcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if xcq.ctx.Unique != nil && *xcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range xcq.predicates {
		p(selector)
	}
	for _, p := range xcq.order {
		p(selector)
	}
	if offset := xcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := xcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// XProviderCursorGroupBy is the group-by builder for XProviderCursor entities.
type XProviderCursorGroupBy struct {
	selector
	build *XProviderCursorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (xcgb *XProviderCursorGroupBy) Aggregate(fns ...AggregateFunc) *XProviderCursorGroupBy {
	xcgb.fns = append(xcgb.fns, fns...)
	return xcgb
}

// Scan applies the selector query and scans the result into the given value.
func (xcgb *XProviderCursorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xcgb.build.ctx, "GroupBy")
	if err := xcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XProviderCursorQuery, *XProviderCursorGroupBy](ctx, xcgb.build, xcgb, xcgb.build.inters, v)
}

func (xcgb *XProviderCursorGroupBy) sqlScan(ctx context.Context, root *XProviderCursorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(xcgb.fns))
	for _, fn := range xcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*xcgb.flds)+len(xcgb.fns))
		for _, f := range *xcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*xcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// XProviderCursorSelect is the builder for selecting fields of XProviderCursor entities.
type XProviderCursorSelect struct {
	*XProviderCursorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (xcs *XProviderCursorSelect) Aggregate(fns ...AggregateFunc) *XProviderCursorSelect {
	xcs.fns = append(xcs.fns, fns...)
	return xcs
}

// Scan applies the selector query and scans the result into the given value.
func (xcs *XProviderCursorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, xcs.ctx, "Select")
	if err := xcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*XProviderCursorQuery, *XProviderCursorSelect](ctx, xcs.XProviderCursorQuery, xcs, xcs.inters, v)
}

func (xcs *XProviderCursorSelect) sqlScan(ctx context.Context, root *XProviderCursorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(xcs.fns))
	for _, fn := range xcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*xcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

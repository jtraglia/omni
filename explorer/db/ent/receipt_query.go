// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
)

// ReceiptQuery is the builder for querying Receipt entities.
type ReceiptQuery struct {
	config
	ctx        *QueryContext
	order      []receipt.OrderOption
	inters     []Interceptor
	predicates []predicate.Receipt
	withBlock  *BlockQuery
	withMsgs   *MsgQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReceiptQuery builder.
func (rq *ReceiptQuery) Where(ps ...predicate.Receipt) *ReceiptQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReceiptQuery) Limit(limit int) *ReceiptQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReceiptQuery) Offset(offset int) *ReceiptQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReceiptQuery) Unique(unique bool) *ReceiptQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReceiptQuery) Order(o ...receipt.OrderOption) *ReceiptQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryBlock chains the current query on the "block" edge.
func (rq *ReceiptQuery) QueryBlock() *BlockQuery {
	query := (&BlockClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, selector),
			sqlgraph.To(block.Table, block.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, receipt.BlockTable, receipt.BlockPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMsgs chains the current query on the "msgs" edge.
func (rq *ReceiptQuery) QueryMsgs() *MsgQuery {
	query := (&MsgClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(receipt.Table, receipt.FieldID, selector),
			sqlgraph.To(msg.Table, msg.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, receipt.MsgsTable, receipt.MsgsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Receipt entity from the query.
// Returns a *NotFoundError when no Receipt was found.
func (rq *ReceiptQuery) First(ctx context.Context) (*Receipt, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{receipt.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReceiptQuery) FirstX(ctx context.Context) *Receipt {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Receipt ID from the query.
// Returns a *NotFoundError when no Receipt ID was found.
func (rq *ReceiptQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{receipt.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReceiptQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Receipt entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Receipt entity is found.
// Returns a *NotFoundError when no Receipt entities are found.
func (rq *ReceiptQuery) Only(ctx context.Context) (*Receipt, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{receipt.Label}
	default:
		return nil, &NotSingularError{receipt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReceiptQuery) OnlyX(ctx context.Context) *Receipt {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Receipt ID in the query.
// Returns a *NotSingularError when more than one Receipt ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReceiptQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{receipt.Label}
	default:
		err = &NotSingularError{receipt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReceiptQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Receipts.
func (rq *ReceiptQuery) All(ctx context.Context) ([]*Receipt, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Receipt, *ReceiptQuery]()
	return withInterceptors[[]*Receipt](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReceiptQuery) AllX(ctx context.Context) []*Receipt {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Receipt IDs.
func (rq *ReceiptQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(receipt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReceiptQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReceiptQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReceiptQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReceiptQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReceiptQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReceiptQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReceiptQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReceiptQuery) Clone() *ReceiptQuery {
	if rq == nil {
		return nil
	}
	return &ReceiptQuery{
		config:     rq.config,
		ctx:        rq.ctx.Clone(),
		order:      append([]receipt.OrderOption{}, rq.order...),
		inters:     append([]Interceptor{}, rq.inters...),
		predicates: append([]predicate.Receipt{}, rq.predicates...),
		withBlock:  rq.withBlock.Clone(),
		withMsgs:   rq.withMsgs.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithBlock tells the query-builder to eager-load the nodes that are connected to
// the "block" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithBlock(opts ...func(*BlockQuery)) *ReceiptQuery {
	query := (&BlockClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withBlock = query
	return rq
}

// WithMsgs tells the query-builder to eager-load the nodes that are connected to
// the "msgs" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReceiptQuery) WithMsgs(opts ...func(*MsgQuery)) *ReceiptQuery {
	query := (&MsgClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withMsgs = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BlockHash []byte `json:"block_hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Receipt.Query().
//		GroupBy(receipt.FieldBlockHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReceiptQuery) GroupBy(field string, fields ...string) *ReceiptGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReceiptGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = receipt.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BlockHash []byte `json:"block_hash,omitempty"`
//	}
//
//	client.Receipt.Query().
//		Select(receipt.FieldBlockHash).
//		Scan(ctx, &v)
func (rq *ReceiptQuery) Select(fields ...string) *ReceiptSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReceiptSelect{ReceiptQuery: rq}
	sbuild.label = receipt.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReceiptSelect configured with the given aggregations.
func (rq *ReceiptQuery) Aggregate(fns ...AggregateFunc) *ReceiptSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReceiptQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !receipt.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReceiptQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Receipt, error) {
	var (
		nodes       = []*Receipt{}
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withBlock != nil,
			rq.withMsgs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Receipt).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Receipt{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withBlock; query != nil {
		if err := rq.loadBlock(ctx, query, nodes,
			func(n *Receipt) { n.Edges.Block = []*Block{} },
			func(n *Receipt, e *Block) { n.Edges.Block = append(n.Edges.Block, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withMsgs; query != nil {
		if err := rq.loadMsgs(ctx, query, nodes,
			func(n *Receipt) { n.Edges.Msgs = []*Msg{} },
			func(n *Receipt, e *Msg) { n.Edges.Msgs = append(n.Edges.Msgs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReceiptQuery) loadBlock(ctx context.Context, query *BlockQuery, nodes []*Receipt, init func(*Receipt), assign func(*Receipt, *Block)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Receipt)
	nids := make(map[int]map[*Receipt]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(receipt.BlockTable)
		s.Join(joinT).On(s.C(block.FieldID), joinT.C(receipt.BlockPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(receipt.BlockPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(receipt.BlockPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Receipt]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Block](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "block" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (rq *ReceiptQuery) loadMsgs(ctx context.Context, query *MsgQuery, nodes []*Receipt, init func(*Receipt), assign func(*Receipt, *Msg)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Receipt)
	nids := make(map[int]map[*Receipt]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(receipt.MsgsTable)
		s.Join(joinT).On(s.C(msg.FieldID), joinT.C(receipt.MsgsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(receipt.MsgsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(receipt.MsgsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Receipt]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Msg](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "msgs" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rq *ReceiptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReceiptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(receipt.Table, receipt.Columns, sqlgraph.NewFieldSpec(receipt.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, receipt.FieldID)
		for i := range fields {
			if fields[i] != receipt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReceiptQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(receipt.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = receipt.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReceiptGroupBy is the group-by builder for Receipt entities.
type ReceiptGroupBy struct {
	selector
	build *ReceiptQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReceiptGroupBy) Aggregate(fns ...AggregateFunc) *ReceiptGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReceiptGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReceiptQuery, *ReceiptGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReceiptGroupBy) sqlScan(ctx context.Context, root *ReceiptQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReceiptSelect is the builder for selecting fields of Receipt entities.
type ReceiptSelect struct {
	*ReceiptQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReceiptSelect) Aggregate(fns ...AggregateFunc) *ReceiptSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReceiptSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReceiptQuery, *ReceiptSelect](ctx, rs.ReceiptQuery, rs, rs.inters, v)
}

func (rs *ReceiptSelect) sqlScan(ctx context.Context, root *ReceiptQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

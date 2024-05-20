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

// BlockQuery is the builder for querying Block entities.
type BlockQuery struct {
	config
	ctx          *QueryContext
	order        []block.OrderOption
	inters       []Interceptor
	predicates   []predicate.Block
	withMsgs     *MsgQuery
	withReceipts *ReceiptQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlockQuery builder.
func (bq *BlockQuery) Where(ps ...predicate.Block) *BlockQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BlockQuery) Limit(limit int) *BlockQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BlockQuery) Offset(offset int) *BlockQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BlockQuery) Unique(unique bool) *BlockQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BlockQuery) Order(o ...block.OrderOption) *BlockQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryMsgs chains the current query on the "msgs" edge.
func (bq *BlockQuery) QueryMsgs() *MsgQuery {
	query := (&MsgClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(block.Table, block.FieldID, selector),
			sqlgraph.To(msg.Table, msg.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, block.MsgsTable, block.MsgsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceipts chains the current query on the "receipts" edge.
func (bq *BlockQuery) QueryReceipts() *ReceiptQuery {
	query := (&ReceiptClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(block.Table, block.FieldID, selector),
			sqlgraph.To(receipt.Table, receipt.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, block.ReceiptsTable, block.ReceiptsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Block entity from the query.
// Returns a *NotFoundError when no Block was found.
func (bq *BlockQuery) First(ctx context.Context) (*Block, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{block.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BlockQuery) FirstX(ctx context.Context) *Block {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Block ID from the query.
// Returns a *NotFoundError when no Block ID was found.
func (bq *BlockQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{block.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BlockQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Block entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Block entity is found.
// Returns a *NotFoundError when no Block entities are found.
func (bq *BlockQuery) Only(ctx context.Context) (*Block, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{block.Label}
	default:
		return nil, &NotSingularError{block.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BlockQuery) OnlyX(ctx context.Context) *Block {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Block ID in the query.
// Returns a *NotSingularError when more than one Block ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BlockQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{block.Label}
	default:
		err = &NotSingularError{block.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BlockQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Blocks.
func (bq *BlockQuery) All(ctx context.Context) ([]*Block, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Block, *BlockQuery]()
	return withInterceptors[[]*Block](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BlockQuery) AllX(ctx context.Context) []*Block {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Block IDs.
func (bq *BlockQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(block.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BlockQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BlockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BlockQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BlockQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BlockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BlockQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BlockQuery) Clone() *BlockQuery {
	if bq == nil {
		return nil
	}
	return &BlockQuery{
		config:       bq.config,
		ctx:          bq.ctx.Clone(),
		order:        append([]block.OrderOption{}, bq.order...),
		inters:       append([]Interceptor{}, bq.inters...),
		predicates:   append([]predicate.Block{}, bq.predicates...),
		withMsgs:     bq.withMsgs.Clone(),
		withReceipts: bq.withReceipts.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithMsgs tells the query-builder to eager-load the nodes that are connected to
// the "msgs" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockQuery) WithMsgs(opts ...func(*MsgQuery)) *BlockQuery {
	query := (&MsgClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withMsgs = query
	return bq
}

// WithReceipts tells the query-builder to eager-load the nodes that are connected to
// the "receipts" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BlockQuery) WithReceipts(opts ...func(*ReceiptQuery)) *BlockQuery {
	query := (&ReceiptClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withReceipts = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hash []byte `json:"hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Block.Query().
//		GroupBy(block.FieldHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BlockQuery) GroupBy(field string, fields ...string) *BlockGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BlockGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = block.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Hash []byte `json:"hash,omitempty"`
//	}
//
//	client.Block.Query().
//		Select(block.FieldHash).
//		Scan(ctx, &v)
func (bq *BlockQuery) Select(fields ...string) *BlockSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BlockSelect{BlockQuery: bq}
	sbuild.label = block.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlockSelect configured with the given aggregations.
func (bq *BlockQuery) Aggregate(fns ...AggregateFunc) *BlockSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BlockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !block.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BlockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Block, error) {
	var (
		nodes       = []*Block{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withMsgs != nil,
			bq.withReceipts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Block).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Block{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withMsgs; query != nil {
		if err := bq.loadMsgs(ctx, query, nodes,
			func(n *Block) { n.Edges.Msgs = []*Msg{} },
			func(n *Block, e *Msg) { n.Edges.Msgs = append(n.Edges.Msgs, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withReceipts; query != nil {
		if err := bq.loadReceipts(ctx, query, nodes,
			func(n *Block) { n.Edges.Receipts = []*Receipt{} },
			func(n *Block, e *Receipt) { n.Edges.Receipts = append(n.Edges.Receipts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BlockQuery) loadMsgs(ctx context.Context, query *MsgQuery, nodes []*Block, init func(*Block), assign func(*Block, *Msg)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Block)
	nids := make(map[int]map[*Block]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(block.MsgsTable)
		s.Join(joinT).On(s.C(msg.FieldID), joinT.C(block.MsgsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(block.MsgsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(block.MsgsPrimaryKey[0]))
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
					nids[inValue] = map[*Block]struct{}{byID[outValue]: {}}
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
func (bq *BlockQuery) loadReceipts(ctx context.Context, query *ReceiptQuery, nodes []*Block, init func(*Block), assign func(*Block, *Receipt)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Block)
	nids := make(map[int]map[*Block]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(block.ReceiptsTable)
		s.Join(joinT).On(s.C(receipt.FieldID), joinT.C(block.ReceiptsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(block.ReceiptsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(block.ReceiptsPrimaryKey[0]))
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
					nids[inValue] = map[*Block]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Receipt](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "receipts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (bq *BlockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BlockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(block.Table, block.Columns, sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, block.FieldID)
		for i := range fields {
			if fields[i] != block.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BlockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(block.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = block.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlockGroupBy is the group-by builder for Block entities.
type BlockGroupBy struct {
	selector
	build *BlockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlockGroupBy) Aggregate(fns ...AggregateFunc) *BlockGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BlockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlockQuery, *BlockGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BlockGroupBy) sqlScan(ctx context.Context, root *BlockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlockSelect is the builder for selecting fields of Block entities.
type BlockSelect struct {
	*BlockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BlockSelect) Aggregate(fns ...AggregateFunc) *BlockSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlockQuery, *BlockSelect](ctx, bs.BlockQuery, bs, bs.inters, v)
}

func (bs *BlockSelect) sqlScan(ctx context.Context, root *BlockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

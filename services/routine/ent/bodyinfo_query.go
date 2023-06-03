// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"routine/ent/bodyinfo"
	"routine/ent/predicate"
	"routine/ent/programrec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BodyInfoQuery is the builder for querying BodyInfo entities.
type BodyInfoQuery struct {
	config
	ctx            *QueryContext
	order          []bodyinfo.OrderOption
	inters         []Interceptor
	predicates     []predicate.BodyInfo
	withProgramRec *ProgramRecQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BodyInfoQuery builder.
func (biq *BodyInfoQuery) Where(ps ...predicate.BodyInfo) *BodyInfoQuery {
	biq.predicates = append(biq.predicates, ps...)
	return biq
}

// Limit the number of records to be returned by this query.
func (biq *BodyInfoQuery) Limit(limit int) *BodyInfoQuery {
	biq.ctx.Limit = &limit
	return biq
}

// Offset to start from.
func (biq *BodyInfoQuery) Offset(offset int) *BodyInfoQuery {
	biq.ctx.Offset = &offset
	return biq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (biq *BodyInfoQuery) Unique(unique bool) *BodyInfoQuery {
	biq.ctx.Unique = &unique
	return biq
}

// Order specifies how the records should be ordered.
func (biq *BodyInfoQuery) Order(o ...bodyinfo.OrderOption) *BodyInfoQuery {
	biq.order = append(biq.order, o...)
	return biq
}

// QueryProgramRec chains the current query on the "program_rec" edge.
func (biq *BodyInfoQuery) QueryProgramRec() *ProgramRecQuery {
	query := (&ProgramRecClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bodyinfo.Table, bodyinfo.FieldID, selector),
			sqlgraph.To(programrec.Table, programrec.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bodyinfo.ProgramRecTable, bodyinfo.ProgramRecColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BodyInfo entity from the query.
// Returns a *NotFoundError when no BodyInfo was found.
func (biq *BodyInfoQuery) First(ctx context.Context) (*BodyInfo, error) {
	nodes, err := biq.Limit(1).All(setContextOp(ctx, biq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bodyinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (biq *BodyInfoQuery) FirstX(ctx context.Context) *BodyInfo {
	node, err := biq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BodyInfo ID from the query.
// Returns a *NotFoundError when no BodyInfo ID was found.
func (biq *BodyInfoQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = biq.Limit(1).IDs(setContextOp(ctx, biq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bodyinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (biq *BodyInfoQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := biq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BodyInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BodyInfo entity is found.
// Returns a *NotFoundError when no BodyInfo entities are found.
func (biq *BodyInfoQuery) Only(ctx context.Context) (*BodyInfo, error) {
	nodes, err := biq.Limit(2).All(setContextOp(ctx, biq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bodyinfo.Label}
	default:
		return nil, &NotSingularError{bodyinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (biq *BodyInfoQuery) OnlyX(ctx context.Context) *BodyInfo {
	node, err := biq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BodyInfo ID in the query.
// Returns a *NotSingularError when more than one BodyInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (biq *BodyInfoQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = biq.Limit(2).IDs(setContextOp(ctx, biq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bodyinfo.Label}
	default:
		err = &NotSingularError{bodyinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (biq *BodyInfoQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := biq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BodyInfos.
func (biq *BodyInfoQuery) All(ctx context.Context) ([]*BodyInfo, error) {
	ctx = setContextOp(ctx, biq.ctx, "All")
	if err := biq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BodyInfo, *BodyInfoQuery]()
	return withInterceptors[[]*BodyInfo](ctx, biq, qr, biq.inters)
}

// AllX is like All, but panics if an error occurs.
func (biq *BodyInfoQuery) AllX(ctx context.Context) []*BodyInfo {
	nodes, err := biq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BodyInfo IDs.
func (biq *BodyInfoQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if biq.ctx.Unique == nil && biq.path != nil {
		biq.Unique(true)
	}
	ctx = setContextOp(ctx, biq.ctx, "IDs")
	if err = biq.Select(bodyinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (biq *BodyInfoQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := biq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (biq *BodyInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, biq.ctx, "Count")
	if err := biq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, biq, querierCount[*BodyInfoQuery](), biq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (biq *BodyInfoQuery) CountX(ctx context.Context) int {
	count, err := biq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (biq *BodyInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, biq.ctx, "Exist")
	switch _, err := biq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (biq *BodyInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := biq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BodyInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (biq *BodyInfoQuery) Clone() *BodyInfoQuery {
	if biq == nil {
		return nil
	}
	return &BodyInfoQuery{
		config:         biq.config,
		ctx:            biq.ctx.Clone(),
		order:          append([]bodyinfo.OrderOption{}, biq.order...),
		inters:         append([]Interceptor{}, biq.inters...),
		predicates:     append([]predicate.BodyInfo{}, biq.predicates...),
		withProgramRec: biq.withProgramRec.Clone(),
		// clone intermediate query.
		sql:  biq.sql.Clone(),
		path: biq.path,
	}
}

// WithProgramRec tells the query-builder to eager-load the nodes that are connected to
// the "program_rec" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BodyInfoQuery) WithProgramRec(opts ...func(*ProgramRecQuery)) *BodyInfoQuery {
	query := (&ProgramRecClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withProgramRec = query
	return biq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Author uint64 `json:"author,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BodyInfo.Query().
//		GroupBy(bodyinfo.FieldAuthor).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (biq *BodyInfoQuery) GroupBy(field string, fields ...string) *BodyInfoGroupBy {
	biq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BodyInfoGroupBy{build: biq}
	grbuild.flds = &biq.ctx.Fields
	grbuild.label = bodyinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Author uint64 `json:"author,omitempty"`
//	}
//
//	client.BodyInfo.Query().
//		Select(bodyinfo.FieldAuthor).
//		Scan(ctx, &v)
func (biq *BodyInfoQuery) Select(fields ...string) *BodyInfoSelect {
	biq.ctx.Fields = append(biq.ctx.Fields, fields...)
	sbuild := &BodyInfoSelect{BodyInfoQuery: biq}
	sbuild.label = bodyinfo.Label
	sbuild.flds, sbuild.scan = &biq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BodyInfoSelect configured with the given aggregations.
func (biq *BodyInfoQuery) Aggregate(fns ...AggregateFunc) *BodyInfoSelect {
	return biq.Select().Aggregate(fns...)
}

func (biq *BodyInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range biq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, biq); err != nil {
				return err
			}
		}
	}
	for _, f := range biq.ctx.Fields {
		if !bodyinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if biq.path != nil {
		prev, err := biq.path(ctx)
		if err != nil {
			return err
		}
		biq.sql = prev
	}
	return nil
}

func (biq *BodyInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BodyInfo, error) {
	var (
		nodes       = []*BodyInfo{}
		_spec       = biq.querySpec()
		loadedTypes = [1]bool{
			biq.withProgramRec != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BodyInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BodyInfo{config: biq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, biq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := biq.withProgramRec; query != nil {
		if err := biq.loadProgramRec(ctx, query, nodes, nil,
			func(n *BodyInfo, e *ProgramRec) { n.Edges.ProgramRec = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (biq *BodyInfoQuery) loadProgramRec(ctx context.Context, query *ProgramRecQuery, nodes []*BodyInfo, init func(*BodyInfo), assign func(*BodyInfo, *ProgramRec)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*BodyInfo)
	for i := range nodes {
		if nodes[i].ProgramRecID == nil {
			continue
		}
		fk := *nodes[i].ProgramRecID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(programrec.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "program_rec_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (biq *BodyInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := biq.querySpec()
	_spec.Node.Columns = biq.ctx.Fields
	if len(biq.ctx.Fields) > 0 {
		_spec.Unique = biq.ctx.Unique != nil && *biq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, biq.driver, _spec)
}

func (biq *BodyInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bodyinfo.Table, bodyinfo.Columns, sqlgraph.NewFieldSpec(bodyinfo.FieldID, field.TypeUint64))
	_spec.From = biq.sql
	if unique := biq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if biq.path != nil {
		_spec.Unique = true
	}
	if fields := biq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bodyinfo.FieldID)
		for i := range fields {
			if fields[i] != bodyinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if biq.withProgramRec != nil {
			_spec.Node.AddColumnOnce(bodyinfo.FieldProgramRecID)
		}
	}
	if ps := biq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := biq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := biq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := biq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (biq *BodyInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(biq.driver.Dialect())
	t1 := builder.Table(bodyinfo.Table)
	columns := biq.ctx.Fields
	if len(columns) == 0 {
		columns = bodyinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if biq.sql != nil {
		selector = biq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if biq.ctx.Unique != nil && *biq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range biq.predicates {
		p(selector)
	}
	for _, p := range biq.order {
		p(selector)
	}
	if offset := biq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := biq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BodyInfoGroupBy is the group-by builder for BodyInfo entities.
type BodyInfoGroupBy struct {
	selector
	build *BodyInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bigb *BodyInfoGroupBy) Aggregate(fns ...AggregateFunc) *BodyInfoGroupBy {
	bigb.fns = append(bigb.fns, fns...)
	return bigb
}

// Scan applies the selector query and scans the result into the given value.
func (bigb *BodyInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bigb.build.ctx, "GroupBy")
	if err := bigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BodyInfoQuery, *BodyInfoGroupBy](ctx, bigb.build, bigb, bigb.build.inters, v)
}

func (bigb *BodyInfoGroupBy) sqlScan(ctx context.Context, root *BodyInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bigb.fns))
	for _, fn := range bigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bigb.flds)+len(bigb.fns))
		for _, f := range *bigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BodyInfoSelect is the builder for selecting fields of BodyInfo entities.
type BodyInfoSelect struct {
	*BodyInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bis *BodyInfoSelect) Aggregate(fns ...AggregateFunc) *BodyInfoSelect {
	bis.fns = append(bis.fns, fns...)
	return bis
}

// Scan applies the selector query and scans the result into the given value.
func (bis *BodyInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bis.ctx, "Select")
	if err := bis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BodyInfoQuery, *BodyInfoSelect](ctx, bis.BodyInfoQuery, bis, bis.inters, v)
}

func (bis *BodyInfoSelect) sqlScan(ctx context.Context, root *BodyInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bis.fns))
	for _, fn := range bis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"routine/ent/act"
	"routine/ent/dailyroutine"
	"routine/ent/predicate"
	"routine/ent/routineact"
	"routine/ent/routineactrec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActQuery is the builder for querying RoutineAct entities.
type RoutineActQuery struct {
	config
	ctx                *QueryContext
	order              []routineact.OrderOption
	inters             []Interceptor
	predicates         []predicate.RoutineAct
	withAct            *ActQuery
	withDailyRoutine   *DailyRoutineQuery
	withRoutineActRecs *RoutineActRecQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoutineActQuery builder.
func (raq *RoutineActQuery) Where(ps ...predicate.RoutineAct) *RoutineActQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit the number of records to be returned by this query.
func (raq *RoutineActQuery) Limit(limit int) *RoutineActQuery {
	raq.ctx.Limit = &limit
	return raq
}

// Offset to start from.
func (raq *RoutineActQuery) Offset(offset int) *RoutineActQuery {
	raq.ctx.Offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *RoutineActQuery) Unique(unique bool) *RoutineActQuery {
	raq.ctx.Unique = &unique
	return raq
}

// Order specifies how the records should be ordered.
func (raq *RoutineActQuery) Order(o ...routineact.OrderOption) *RoutineActQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryAct chains the current query on the "act" edge.
func (raq *RoutineActQuery) QueryAct() *ActQuery {
	query := (&ActClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineact.Table, routineact.FieldID, selector),
			sqlgraph.To(act.Table, act.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, routineact.ActTable, routineact.ActColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDailyRoutine chains the current query on the "daily_routine" edge.
func (raq *RoutineActQuery) QueryDailyRoutine() *DailyRoutineQuery {
	query := (&DailyRoutineClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineact.Table, routineact.FieldID, selector),
			sqlgraph.To(dailyroutine.Table, dailyroutine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, routineact.DailyRoutineTable, routineact.DailyRoutineColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutineActRecs chains the current query on the "routine_act_recs" edge.
func (raq *RoutineActQuery) QueryRoutineActRecs() *RoutineActRecQuery {
	query := (&RoutineActRecClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routineact.Table, routineact.FieldID, selector),
			sqlgraph.To(routineactrec.Table, routineactrec.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, routineact.RoutineActRecsTable, routineact.RoutineActRecsColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoutineAct entity from the query.
// Returns a *NotFoundError when no RoutineAct was found.
func (raq *RoutineActQuery) First(ctx context.Context) (*RoutineAct, error) {
	nodes, err := raq.Limit(1).All(setContextOp(ctx, raq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{routineact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *RoutineActQuery) FirstX(ctx context.Context) *RoutineAct {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoutineAct ID from the query.
// Returns a *NotFoundError when no RoutineAct ID was found.
func (raq *RoutineActQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = raq.Limit(1).IDs(setContextOp(ctx, raq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{routineact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *RoutineActQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoutineAct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoutineAct entity is found.
// Returns a *NotFoundError when no RoutineAct entities are found.
func (raq *RoutineActQuery) Only(ctx context.Context) (*RoutineAct, error) {
	nodes, err := raq.Limit(2).All(setContextOp(ctx, raq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{routineact.Label}
	default:
		return nil, &NotSingularError{routineact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *RoutineActQuery) OnlyX(ctx context.Context) *RoutineAct {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoutineAct ID in the query.
// Returns a *NotSingularError when more than one RoutineAct ID is found.
// Returns a *NotFoundError when no entities are found.
func (raq *RoutineActQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = raq.Limit(2).IDs(setContextOp(ctx, raq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{routineact.Label}
	default:
		err = &NotSingularError{routineact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *RoutineActQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoutineActs.
func (raq *RoutineActQuery) All(ctx context.Context) ([]*RoutineAct, error) {
	ctx = setContextOp(ctx, raq.ctx, "All")
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoutineAct, *RoutineActQuery]()
	return withInterceptors[[]*RoutineAct](ctx, raq, qr, raq.inters)
}

// AllX is like All, but panics if an error occurs.
func (raq *RoutineActQuery) AllX(ctx context.Context) []*RoutineAct {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoutineAct IDs.
func (raq *RoutineActQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if raq.ctx.Unique == nil && raq.path != nil {
		raq.Unique(true)
	}
	ctx = setContextOp(ctx, raq.ctx, "IDs")
	if err = raq.Select(routineact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *RoutineActQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *RoutineActQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, raq.ctx, "Count")
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, raq, querierCount[*RoutineActQuery](), raq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (raq *RoutineActQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *RoutineActQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, raq.ctx, "Exist")
	switch _, err := raq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *RoutineActQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoutineActQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *RoutineActQuery) Clone() *RoutineActQuery {
	if raq == nil {
		return nil
	}
	return &RoutineActQuery{
		config:             raq.config,
		ctx:                raq.ctx.Clone(),
		order:              append([]routineact.OrderOption{}, raq.order...),
		inters:             append([]Interceptor{}, raq.inters...),
		predicates:         append([]predicate.RoutineAct{}, raq.predicates...),
		withAct:            raq.withAct.Clone(),
		withDailyRoutine:   raq.withDailyRoutine.Clone(),
		withRoutineActRecs: raq.withRoutineActRecs.Clone(),
		// clone intermediate query.
		sql:  raq.sql.Clone(),
		path: raq.path,
	}
}

// WithAct tells the query-builder to eager-load the nodes that are connected to
// the "act" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RoutineActQuery) WithAct(opts ...func(*ActQuery)) *RoutineActQuery {
	query := (&ActClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withAct = query
	return raq
}

// WithDailyRoutine tells the query-builder to eager-load the nodes that are connected to
// the "daily_routine" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RoutineActQuery) WithDailyRoutine(opts ...func(*DailyRoutineQuery)) *RoutineActQuery {
	query := (&DailyRoutineClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withDailyRoutine = query
	return raq
}

// WithRoutineActRecs tells the query-builder to eager-load the nodes that are connected to
// the "routine_act_recs" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RoutineActQuery) WithRoutineActRecs(opts ...func(*RoutineActRecQuery)) *RoutineActQuery {
	query := (&RoutineActRecClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withRoutineActRecs = query
	return raq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ActID uint64 `json:"act_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoutineAct.Query().
//		GroupBy(routineact.FieldActID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (raq *RoutineActQuery) GroupBy(field string, fields ...string) *RoutineActGroupBy {
	raq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoutineActGroupBy{build: raq}
	grbuild.flds = &raq.ctx.Fields
	grbuild.label = routineact.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ActID uint64 `json:"act_id,omitempty"`
//	}
//
//	client.RoutineAct.Query().
//		Select(routineact.FieldActID).
//		Scan(ctx, &v)
func (raq *RoutineActQuery) Select(fields ...string) *RoutineActSelect {
	raq.ctx.Fields = append(raq.ctx.Fields, fields...)
	sbuild := &RoutineActSelect{RoutineActQuery: raq}
	sbuild.label = routineact.Label
	sbuild.flds, sbuild.scan = &raq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoutineActSelect configured with the given aggregations.
func (raq *RoutineActQuery) Aggregate(fns ...AggregateFunc) *RoutineActSelect {
	return raq.Select().Aggregate(fns...)
}

func (raq *RoutineActQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range raq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, raq); err != nil {
				return err
			}
		}
	}
	for _, f := range raq.ctx.Fields {
		if !routineact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *RoutineActQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoutineAct, error) {
	var (
		nodes       = []*RoutineAct{}
		_spec       = raq.querySpec()
		loadedTypes = [3]bool{
			raq.withAct != nil,
			raq.withDailyRoutine != nil,
			raq.withRoutineActRecs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoutineAct).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoutineAct{config: raq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := raq.withAct; query != nil {
		if err := raq.loadAct(ctx, query, nodes, nil,
			func(n *RoutineAct, e *Act) { n.Edges.Act = e }); err != nil {
			return nil, err
		}
	}
	if query := raq.withDailyRoutine; query != nil {
		if err := raq.loadDailyRoutine(ctx, query, nodes, nil,
			func(n *RoutineAct, e *DailyRoutine) { n.Edges.DailyRoutine = e }); err != nil {
			return nil, err
		}
	}
	if query := raq.withRoutineActRecs; query != nil {
		if err := raq.loadRoutineActRecs(ctx, query, nodes,
			func(n *RoutineAct) { n.Edges.RoutineActRecs = []*RoutineActRec{} },
			func(n *RoutineAct, e *RoutineActRec) { n.Edges.RoutineActRecs = append(n.Edges.RoutineActRecs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (raq *RoutineActQuery) loadAct(ctx context.Context, query *ActQuery, nodes []*RoutineAct, init func(*RoutineAct), assign func(*RoutineAct, *Act)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*RoutineAct)
	for i := range nodes {
		fk := nodes[i].ActID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(act.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "act_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (raq *RoutineActQuery) loadDailyRoutine(ctx context.Context, query *DailyRoutineQuery, nodes []*RoutineAct, init func(*RoutineAct), assign func(*RoutineAct, *DailyRoutine)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*RoutineAct)
	for i := range nodes {
		fk := nodes[i].DailyRoutineID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dailyroutine.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "daily_routine_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (raq *RoutineActQuery) loadRoutineActRecs(ctx context.Context, query *RoutineActRecQuery, nodes []*RoutineAct, init func(*RoutineAct), assign func(*RoutineAct, *RoutineActRec)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*RoutineAct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(routineactrec.FieldRoutineActID)
	}
	query.Where(predicate.RoutineActRec(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(routineact.RoutineActRecsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoutineActID
		if fk == nil {
			return fmt.Errorf(`foreign-key "routine_act_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "routine_act_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (raq *RoutineActQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	_spec.Node.Columns = raq.ctx.Fields
	if len(raq.ctx.Fields) > 0 {
		_spec.Unique = raq.ctx.Unique != nil && *raq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *RoutineActQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(routineact.Table, routineact.Columns, sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64))
	_spec.From = raq.sql
	if unique := raq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if raq.path != nil {
		_spec.Unique = true
	}
	if fields := raq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routineact.FieldID)
		for i := range fields {
			if fields[i] != routineact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if raq.withAct != nil {
			_spec.Node.AddColumnOnce(routineact.FieldActID)
		}
		if raq.withDailyRoutine != nil {
			_spec.Node.AddColumnOnce(routineact.FieldDailyRoutineID)
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *RoutineActQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(routineact.Table)
	columns := raq.ctx.Fields
	if len(columns) == 0 {
		columns = routineact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if raq.ctx.Unique != nil && *raq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoutineActGroupBy is the group-by builder for RoutineAct entities.
type RoutineActGroupBy struct {
	selector
	build *RoutineActQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *RoutineActGroupBy) Aggregate(fns ...AggregateFunc) *RoutineActGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the selector query and scans the result into the given value.
func (ragb *RoutineActGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ragb.build.ctx, "GroupBy")
	if err := ragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineActQuery, *RoutineActGroupBy](ctx, ragb.build, ragb, ragb.build.inters, v)
}

func (ragb *RoutineActGroupBy) sqlScan(ctx context.Context, root *RoutineActQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ragb.flds)+len(ragb.fns))
		for _, f := range *ragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoutineActSelect is the builder for selecting fields of RoutineAct entities.
type RoutineActSelect struct {
	*RoutineActQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ras *RoutineActSelect) Aggregate(fns ...AggregateFunc) *RoutineActSelect {
	ras.fns = append(ras.fns, fns...)
	return ras
}

// Scan applies the selector query and scans the result into the given value.
func (ras *RoutineActSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ras.ctx, "Select")
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoutineActQuery, *RoutineActSelect](ctx, ras.RoutineActQuery, ras, ras.inters, v)
}

func (ras *RoutineActSelect) sqlScan(ctx context.Context, root *RoutineActQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ras.fns))
	for _, fn := range ras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

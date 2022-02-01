// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bongo/ent/brand"
	"bongo/ent/predicate"
	"bongo/ent/sellerproduct"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BrandQuery is the builder for querying Brand entities.
type BrandQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Brand
	// eager-loading edges.
	withBrand *SellerProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BrandQuery builder.
func (bq *BrandQuery) Where(ps ...predicate.Brand) *BrandQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BrandQuery) Limit(limit int) *BrandQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BrandQuery) Offset(offset int) *BrandQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BrandQuery) Unique(unique bool) *BrandQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BrandQuery) Order(o ...OrderFunc) *BrandQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryBrand chains the current query on the "brand" edge.
func (bq *BrandQuery) QueryBrand() *SellerProductQuery {
	query := &SellerProductQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(brand.Table, brand.FieldID, selector),
			sqlgraph.To(sellerproduct.Table, sellerproduct.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, brand.BrandTable, brand.BrandColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Brand entity from the query.
// Returns a *NotFoundError when no Brand was found.
func (bq *BrandQuery) First(ctx context.Context) (*Brand, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{brand.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BrandQuery) FirstX(ctx context.Context) *Brand {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Brand ID from the query.
// Returns a *NotFoundError when no Brand ID was found.
func (bq *BrandQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{brand.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BrandQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Brand entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Brand entity is not found.
// Returns a *NotFoundError when no Brand entities are found.
func (bq *BrandQuery) Only(ctx context.Context) (*Brand, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{brand.Label}
	default:
		return nil, &NotSingularError{brand.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BrandQuery) OnlyX(ctx context.Context) *Brand {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Brand ID in the query.
// Returns a *NotSingularError when exactly one Brand ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bq *BrandQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = &NotSingularError{brand.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BrandQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Brands.
func (bq *BrandQuery) All(ctx context.Context) ([]*Brand, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BrandQuery) AllX(ctx context.Context) []*Brand {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Brand IDs.
func (bq *BrandQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(brand.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BrandQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BrandQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BrandQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BrandQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BrandQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BrandQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BrandQuery) Clone() *BrandQuery {
	if bq == nil {
		return nil
	}
	return &BrandQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		predicates: append([]predicate.Brand{}, bq.predicates...),
		withBrand:  bq.withBrand.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithBrand tells the query-builder to eager-load the nodes that are connected to
// the "brand" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BrandQuery) WithBrand(opts ...func(*SellerProductQuery)) *BrandQuery {
	query := &SellerProductQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withBrand = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Brand.Query().
//		GroupBy(brand.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BrandQuery) GroupBy(field string, fields ...string) *BrandGroupBy {
	group := &BrandGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Brand.Query().
//		Select(brand.FieldName).
//		Scan(ctx, &v)
//
func (bq *BrandQuery) Select(fields ...string) *BrandSelect {
	bq.fields = append(bq.fields, fields...)
	return &BrandSelect{BrandQuery: bq}
}

func (bq *BrandQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !brand.ValidColumn(f) {
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

func (bq *BrandQuery) sqlAll(ctx context.Context) ([]*Brand, error) {
	var (
		nodes       = []*Brand{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withBrand != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Brand{config: bq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withBrand; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Brand)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Brand = []*SellerProduct{}
		}
		query.withFKs = true
		query.Where(predicate.SellerProduct(func(s *sql.Selector) {
			s.Where(sql.InValues(brand.BrandColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.brand_brand
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "brand_brand" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "brand_brand" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Brand = append(node.Edges.Brand, n)
		}
	}

	return nodes, nil
}

func (bq *BrandQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BrandQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bq *BrandQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   brand.Table,
			Columns: brand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: brand.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, brand.FieldID)
		for i := range fields {
			if fields[i] != brand.FieldID {
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
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
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

func (bq *BrandQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(brand.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = brand.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BrandGroupBy is the group-by builder for Brand entities.
type BrandGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BrandGroupBy) Aggregate(fns ...AggregateFunc) *BrandGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BrandGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BrandGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BrandGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BrandGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bgb *BrandGroupBy) StringX(ctx context.Context) string {
	v, err := bgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BrandGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BrandGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bgb *BrandGroupBy) IntX(ctx context.Context) int {
	v, err := bgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BrandGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BrandGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bgb *BrandGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BrandGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BrandGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BrandGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bgb *BrandGroupBy) BoolX(ctx context.Context) bool {
	v, err := bgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BrandGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bgb.fields {
		if !brand.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BrandGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BrandSelect is the builder for selecting fields of Brand entities.
type BrandSelect struct {
	*BrandQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BrandSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BrandQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BrandSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BrandSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BrandSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bs *BrandSelect) StringX(ctx context.Context) string {
	v, err := bs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BrandSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BrandSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bs *BrandSelect) IntX(ctx context.Context) int {
	v, err := bs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BrandSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BrandSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bs *BrandSelect) Float64X(ctx context.Context) float64 {
	v, err := bs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BrandSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BrandSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bs *BrandSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{brand.Label}
	default:
		err = fmt.Errorf("ent: BrandSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bs *BrandSelect) BoolX(ctx context.Context) bool {
	v, err := bs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BrandSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
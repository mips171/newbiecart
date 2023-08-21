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
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/ent/productcategory"
)

// ProductCategoryQuery is the builder for querying ProductCategory entities.
type ProductCategoryQuery struct {
	config
	ctx          *QueryContext
	order        []productcategory.OrderOption
	inters       []Interceptor
	predicates   []predicate.ProductCategory
	withProducts *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductCategoryQuery builder.
func (pcq *ProductCategoryQuery) Where(ps ...predicate.ProductCategory) *ProductCategoryQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *ProductCategoryQuery) Limit(limit int) *ProductCategoryQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *ProductCategoryQuery) Offset(offset int) *ProductCategoryQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *ProductCategoryQuery) Unique(unique bool) *ProductCategoryQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *ProductCategoryQuery) Order(o ...productcategory.OrderOption) *ProductCategoryQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryProducts chains the current query on the "products" edge.
func (pcq *ProductCategoryQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productcategory.Table, productcategory.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productcategory.ProductsTable, productcategory.ProductsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductCategory entity from the query.
// Returns a *NotFoundError when no ProductCategory was found.
func (pcq *ProductCategoryQuery) First(ctx context.Context) (*ProductCategory, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *ProductCategoryQuery) FirstX(ctx context.Context) *ProductCategory {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductCategory ID from the query.
// Returns a *NotFoundError when no ProductCategory ID was found.
func (pcq *ProductCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *ProductCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductCategory entity is found.
// Returns a *NotFoundError when no ProductCategory entities are found.
func (pcq *ProductCategoryQuery) Only(ctx context.Context) (*ProductCategory, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productcategory.Label}
	default:
		return nil, &NotSingularError{productcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *ProductCategoryQuery) OnlyX(ctx context.Context) *ProductCategory {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductCategory ID in the query.
// Returns a *NotSingularError when more than one ProductCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *ProductCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productcategory.Label}
	default:
		err = &NotSingularError{productcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *ProductCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductCategories.
func (pcq *ProductCategoryQuery) All(ctx context.Context) ([]*ProductCategory, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductCategory, *ProductCategoryQuery]()
	return withInterceptors[[]*ProductCategory](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *ProductCategoryQuery) AllX(ctx context.Context) []*ProductCategory {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductCategory IDs.
func (pcq *ProductCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(productcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *ProductCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *ProductCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*ProductCategoryQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *ProductCategoryQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *ProductCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *ProductCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *ProductCategoryQuery) Clone() *ProductCategoryQuery {
	if pcq == nil {
		return nil
	}
	return &ProductCategoryQuery{
		config:       pcq.config,
		ctx:          pcq.ctx.Clone(),
		order:        append([]productcategory.OrderOption{}, pcq.order...),
		inters:       append([]Interceptor{}, pcq.inters...),
		predicates:   append([]predicate.ProductCategory{}, pcq.predicates...),
		withProducts: pcq.withProducts.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *ProductCategoryQuery) WithProducts(opts ...func(*ProductQuery)) *ProductCategoryQuery {
	query := (&ProductClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withProducts = query
	return pcq
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
//	client.ProductCategory.Query().
//		GroupBy(productcategory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *ProductCategoryQuery) GroupBy(field string, fields ...string) *ProductCategoryGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductCategoryGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = productcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
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
//	client.ProductCategory.Query().
//		Select(productcategory.FieldName).
//		Scan(ctx, &v)
func (pcq *ProductCategoryQuery) Select(fields ...string) *ProductCategorySelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &ProductCategorySelect{ProductCategoryQuery: pcq}
	sbuild.label = productcategory.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductCategorySelect configured with the given aggregations.
func (pcq *ProductCategoryQuery) Aggregate(fns ...AggregateFunc) *ProductCategorySelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *ProductCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !productcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *ProductCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductCategory, error) {
	var (
		nodes       = []*ProductCategory{}
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductCategory{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withProducts; query != nil {
		if err := pcq.loadProducts(ctx, query, nodes,
			func(n *ProductCategory) { n.Edges.Products = []*Product{} },
			func(n *ProductCategory, e *Product) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *ProductCategoryQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*ProductCategory, init func(*ProductCategory), assign func(*ProductCategory, *Product)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ProductCategory)
	nids := make(map[int]map[*ProductCategory]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(productcategory.ProductsTable)
		s.Join(joinT).On(s.C(product.FieldID), joinT.C(productcategory.ProductsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(productcategory.ProductsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(productcategory.ProductsPrimaryKey[0]))
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
					nids[inValue] = map[*ProductCategory]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Product](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "products" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pcq *ProductCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *ProductCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productcategory.Table, productcategory.Columns, sqlgraph.NewFieldSpec(productcategory.FieldID, field.TypeInt))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcategory.FieldID)
		for i := range fields {
			if fields[i] != productcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *ProductCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(productcategory.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = productcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductCategoryGroupBy is the group-by builder for ProductCategory entities.
type ProductCategoryGroupBy struct {
	selector
	build *ProductCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *ProductCategoryGroupBy) Aggregate(fns ...AggregateFunc) *ProductCategoryGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *ProductCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductCategoryQuery, *ProductCategoryGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *ProductCategoryGroupBy) sqlScan(ctx context.Context, root *ProductCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductCategorySelect is the builder for selecting fields of ProductCategory entities.
type ProductCategorySelect struct {
	*ProductCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *ProductCategorySelect) Aggregate(fns ...AggregateFunc) *ProductCategorySelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *ProductCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductCategoryQuery, *ProductCategorySelect](ctx, pcs.ProductCategoryQuery, pcs, pcs.inters, v)
}

func (pcs *ProductCategorySelect) sqlScan(ctx context.Context, root *ProductCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
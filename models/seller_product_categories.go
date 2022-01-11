// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SellerProductCategory is an object representing the database table.
type SellerProductCategory struct {
	ID              int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt       null.Time  `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt       null.Time  `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt       null.Time  `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	SellerProductID null.Int64 `boil:"seller_product_id" json:"seller_product_id,omitempty" toml:"seller_product_id" yaml:"seller_product_id,omitempty"`
	CategoryID      null.Int64 `boil:"category_id" json:"category_id,omitempty" toml:"category_id" yaml:"category_id,omitempty"`

	R *sellerProductCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sellerProductCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SellerProductCategoryColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	SellerProductID string
	CategoryID      string
}{
	ID:              "id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	SellerProductID: "seller_product_id",
	CategoryID:      "category_id",
}

var SellerProductCategoryTableColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	SellerProductID string
	CategoryID      string
}{
	ID:              "seller_product_categories.id",
	CreatedAt:       "seller_product_categories.created_at",
	UpdatedAt:       "seller_product_categories.updated_at",
	DeletedAt:       "seller_product_categories.deleted_at",
	SellerProductID: "seller_product_categories.seller_product_id",
	CategoryID:      "seller_product_categories.category_id",
}

// Generated where

var SellerProductCategoryWhere = struct {
	ID              whereHelperint64
	CreatedAt       whereHelpernull_Time
	UpdatedAt       whereHelpernull_Time
	DeletedAt       whereHelpernull_Time
	SellerProductID whereHelpernull_Int64
	CategoryID      whereHelpernull_Int64
}{
	ID:              whereHelperint64{field: "\"seller_product_categories\".\"id\""},
	CreatedAt:       whereHelpernull_Time{field: "\"seller_product_categories\".\"created_at\""},
	UpdatedAt:       whereHelpernull_Time{field: "\"seller_product_categories\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"seller_product_categories\".\"deleted_at\""},
	SellerProductID: whereHelpernull_Int64{field: "\"seller_product_categories\".\"seller_product_id\""},
	CategoryID:      whereHelpernull_Int64{field: "\"seller_product_categories\".\"category_id\""},
}

// SellerProductCategoryRels is where relationship names are stored.
var SellerProductCategoryRels = struct {
	SellerProduct string
}{
	SellerProduct: "SellerProduct",
}

// sellerProductCategoryR is where relationships are stored.
type sellerProductCategoryR struct {
	SellerProduct *SellerProduct `boil:"SellerProduct" json:"SellerProduct" toml:"SellerProduct" yaml:"SellerProduct"`
}

// NewStruct creates a new relationship struct
func (*sellerProductCategoryR) NewStruct() *sellerProductCategoryR {
	return &sellerProductCategoryR{}
}

// sellerProductCategoryL is where Load methods for each relationship are stored.
type sellerProductCategoryL struct{}

var (
	sellerProductCategoryAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "seller_product_id", "category_id"}
	sellerProductCategoryColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "seller_product_id", "category_id"}
	sellerProductCategoryColumnsWithDefault    = []string{"id"}
	sellerProductCategoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// SellerProductCategorySlice is an alias for a slice of pointers to SellerProductCategory.
	// This should almost always be used instead of []SellerProductCategory.
	SellerProductCategorySlice []*SellerProductCategory
	// SellerProductCategoryHook is the signature for custom SellerProductCategory hook methods
	SellerProductCategoryHook func(context.Context, boil.ContextExecutor, *SellerProductCategory) error

	sellerProductCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sellerProductCategoryType                 = reflect.TypeOf(&SellerProductCategory{})
	sellerProductCategoryMapping              = queries.MakeStructMapping(sellerProductCategoryType)
	sellerProductCategoryPrimaryKeyMapping, _ = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, sellerProductCategoryPrimaryKeyColumns)
	sellerProductCategoryInsertCacheMut       sync.RWMutex
	sellerProductCategoryInsertCache          = make(map[string]insertCache)
	sellerProductCategoryUpdateCacheMut       sync.RWMutex
	sellerProductCategoryUpdateCache          = make(map[string]updateCache)
	sellerProductCategoryUpsertCacheMut       sync.RWMutex
	sellerProductCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sellerProductCategoryBeforeInsertHooks []SellerProductCategoryHook
var sellerProductCategoryBeforeUpdateHooks []SellerProductCategoryHook
var sellerProductCategoryBeforeDeleteHooks []SellerProductCategoryHook
var sellerProductCategoryBeforeUpsertHooks []SellerProductCategoryHook

var sellerProductCategoryAfterInsertHooks []SellerProductCategoryHook
var sellerProductCategoryAfterSelectHooks []SellerProductCategoryHook
var sellerProductCategoryAfterUpdateHooks []SellerProductCategoryHook
var sellerProductCategoryAfterDeleteHooks []SellerProductCategoryHook
var sellerProductCategoryAfterUpsertHooks []SellerProductCategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SellerProductCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SellerProductCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SellerProductCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SellerProductCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SellerProductCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SellerProductCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SellerProductCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SellerProductCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SellerProductCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSellerProductCategoryHook registers your hook function for all future operations.
func AddSellerProductCategoryHook(hookPoint boil.HookPoint, sellerProductCategoryHook SellerProductCategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		sellerProductCategoryBeforeInsertHooks = append(sellerProductCategoryBeforeInsertHooks, sellerProductCategoryHook)
	case boil.BeforeUpdateHook:
		sellerProductCategoryBeforeUpdateHooks = append(sellerProductCategoryBeforeUpdateHooks, sellerProductCategoryHook)
	case boil.BeforeDeleteHook:
		sellerProductCategoryBeforeDeleteHooks = append(sellerProductCategoryBeforeDeleteHooks, sellerProductCategoryHook)
	case boil.BeforeUpsertHook:
		sellerProductCategoryBeforeUpsertHooks = append(sellerProductCategoryBeforeUpsertHooks, sellerProductCategoryHook)
	case boil.AfterInsertHook:
		sellerProductCategoryAfterInsertHooks = append(sellerProductCategoryAfterInsertHooks, sellerProductCategoryHook)
	case boil.AfterSelectHook:
		sellerProductCategoryAfterSelectHooks = append(sellerProductCategoryAfterSelectHooks, sellerProductCategoryHook)
	case boil.AfterUpdateHook:
		sellerProductCategoryAfterUpdateHooks = append(sellerProductCategoryAfterUpdateHooks, sellerProductCategoryHook)
	case boil.AfterDeleteHook:
		sellerProductCategoryAfterDeleteHooks = append(sellerProductCategoryAfterDeleteHooks, sellerProductCategoryHook)
	case boil.AfterUpsertHook:
		sellerProductCategoryAfterUpsertHooks = append(sellerProductCategoryAfterUpsertHooks, sellerProductCategoryHook)
	}
}

// One returns a single sellerProductCategory record from the query.
func (q sellerProductCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SellerProductCategory, error) {
	o := &SellerProductCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for seller_product_categories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SellerProductCategory records from the query.
func (q sellerProductCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (SellerProductCategorySlice, error) {
	var o []*SellerProductCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SellerProductCategory slice")
	}

	if len(sellerProductCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SellerProductCategory records in the query.
func (q sellerProductCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count seller_product_categories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sellerProductCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if seller_product_categories exists")
	}

	return count > 0, nil
}

// SellerProduct pointed to by the foreign key.
func (o *SellerProductCategory) SellerProduct(mods ...qm.QueryMod) sellerProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SellerProductID),
	}

	queryMods = append(queryMods, mods...)

	query := SellerProducts(queryMods...)
	queries.SetFrom(query.Query, "\"seller_products\"")

	return query
}

// LoadSellerProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (sellerProductCategoryL) LoadSellerProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSellerProductCategory interface{}, mods queries.Applicator) error {
	var slice []*SellerProductCategory
	var object *SellerProductCategory

	if singular {
		object = maybeSellerProductCategory.(*SellerProductCategory)
	} else {
		slice = *maybeSellerProductCategory.(*[]*SellerProductCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &sellerProductCategoryR{}
		}
		if !queries.IsNil(object.SellerProductID) {
			args = append(args, object.SellerProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &sellerProductCategoryR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SellerProductID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.SellerProductID) {
				args = append(args, obj.SellerProductID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`seller_products`),
		qm.WhereIn(`seller_products.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SellerProduct")
	}

	var resultSlice []*SellerProduct
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SellerProduct")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for seller_products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for seller_products")
	}

	if len(sellerProductCategoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.SellerProduct = foreign
		if foreign.R == nil {
			foreign.R = &sellerProductR{}
		}
		foreign.R.SellerProductCategories = append(foreign.R.SellerProductCategories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SellerProductID, foreign.ID) {
				local.R.SellerProduct = foreign
				if foreign.R == nil {
					foreign.R = &sellerProductR{}
				}
				foreign.R.SellerProductCategories = append(foreign.R.SellerProductCategories, local)
				break
			}
		}
	}

	return nil
}

// SetSellerProduct of the sellerProductCategory to the related item.
// Sets o.R.SellerProduct to related.
// Adds o to related.R.SellerProductCategories.
func (o *SellerProductCategory) SetSellerProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SellerProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"seller_product_categories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"seller_product_id"}),
		strmangle.WhereClause("\"", "\"", 2, sellerProductCategoryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.SellerProductID, related.ID)
	if o.R == nil {
		o.R = &sellerProductCategoryR{
			SellerProduct: related,
		}
	} else {
		o.R.SellerProduct = related
	}

	if related.R == nil {
		related.R = &sellerProductR{
			SellerProductCategories: SellerProductCategorySlice{o},
		}
	} else {
		related.R.SellerProductCategories = append(related.R.SellerProductCategories, o)
	}

	return nil
}

// RemoveSellerProduct relationship.
// Sets o.R.SellerProduct to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *SellerProductCategory) RemoveSellerProduct(ctx context.Context, exec boil.ContextExecutor, related *SellerProduct) error {
	var err error

	queries.SetScanner(&o.SellerProductID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("seller_product_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.SellerProduct = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.SellerProductCategories {
		if queries.Equal(o.SellerProductID, ri.SellerProductID) {
			continue
		}

		ln := len(related.R.SellerProductCategories)
		if ln > 1 && i < ln-1 {
			related.R.SellerProductCategories[i] = related.R.SellerProductCategories[ln-1]
		}
		related.R.SellerProductCategories = related.R.SellerProductCategories[:ln-1]
		break
	}
	return nil
}

// SellerProductCategories retrieves all the records using an executor.
func SellerProductCategories(mods ...qm.QueryMod) sellerProductCategoryQuery {
	mods = append(mods, qm.From("\"seller_product_categories\""))
	return sellerProductCategoryQuery{NewQuery(mods...)}
}

// FindSellerProductCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSellerProductCategory(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*SellerProductCategory, error) {
	sellerProductCategoryObj := &SellerProductCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"seller_product_categories\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, sellerProductCategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from seller_product_categories")
	}

	if err = sellerProductCategoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sellerProductCategoryObj, err
	}

	return sellerProductCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SellerProductCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no seller_product_categories provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sellerProductCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sellerProductCategoryInsertCacheMut.RLock()
	cache, cached := sellerProductCategoryInsertCache[key]
	sellerProductCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sellerProductCategoryAllColumns,
			sellerProductCategoryColumnsWithDefault,
			sellerProductCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"seller_product_categories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"seller_product_categories\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into seller_product_categories")
	}

	if !cached {
		sellerProductCategoryInsertCacheMut.Lock()
		sellerProductCategoryInsertCache[key] = cache
		sellerProductCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SellerProductCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SellerProductCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sellerProductCategoryUpdateCacheMut.RLock()
	cache, cached := sellerProductCategoryUpdateCache[key]
	sellerProductCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sellerProductCategoryAllColumns,
			sellerProductCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update seller_product_categories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"seller_product_categories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, sellerProductCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, append(wl, sellerProductCategoryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update seller_product_categories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for seller_product_categories")
	}

	if !cached {
		sellerProductCategoryUpdateCacheMut.Lock()
		sellerProductCategoryUpdateCache[key] = cache
		sellerProductCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sellerProductCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for seller_product_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for seller_product_categories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SellerProductCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"seller_product_categories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, sellerProductCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sellerProductCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sellerProductCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SellerProductCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no seller_product_categories provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sellerProductCategoryColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	sellerProductCategoryUpsertCacheMut.RLock()
	cache, cached := sellerProductCategoryUpsertCache[key]
	sellerProductCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sellerProductCategoryAllColumns,
			sellerProductCategoryColumnsWithDefault,
			sellerProductCategoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			sellerProductCategoryAllColumns,
			sellerProductCategoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert seller_product_categories, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(sellerProductCategoryPrimaryKeyColumns))
			copy(conflict, sellerProductCategoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"seller_product_categories\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sellerProductCategoryType, sellerProductCategoryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert seller_product_categories")
	}

	if !cached {
		sellerProductCategoryUpsertCacheMut.Lock()
		sellerProductCategoryUpsertCache[key] = cache
		sellerProductCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SellerProductCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SellerProductCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SellerProductCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sellerProductCategoryPrimaryKeyMapping)
	sql := "DELETE FROM \"seller_product_categories\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from seller_product_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for seller_product_categories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sellerProductCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sellerProductCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from seller_product_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for seller_product_categories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SellerProductCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sellerProductCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"seller_product_categories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerProductCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sellerProductCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for seller_product_categories")
	}

	if len(sellerProductCategoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SellerProductCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSellerProductCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SellerProductCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SellerProductCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"seller_product_categories\".* FROM \"seller_product_categories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerProductCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SellerProductCategorySlice")
	}

	*o = slice

	return nil
}

// SellerProductCategoryExists checks if the SellerProductCategory row exists.
func SellerProductCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"seller_product_categories\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if seller_product_categories exists")
	}

	return exists, nil
}

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

// Brand is an object representing the database table.
type Brand struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *brandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L brandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BrandColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
}

var BrandTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "brands.id",
	CreatedAt: "brands.created_at",
	UpdatedAt: "brands.updated_at",
	DeletedAt: "brands.deleted_at",
	Name:      "brands.name",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BrandWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
	DeletedAt whereHelpernull_Time
	Name      whereHelperstring
}{
	ID:        whereHelperint64{field: "\"brands\".\"id\""},
	CreatedAt: whereHelpernull_Time{field: "\"brands\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"brands\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"brands\".\"deleted_at\""},
	Name:      whereHelperstring{field: "\"brands\".\"name\""},
}

// BrandRels is where relationship names are stored.
var BrandRels = struct {
	SellerProducts string
}{
	SellerProducts: "SellerProducts",
}

// brandR is where relationships are stored.
type brandR struct {
	SellerProducts SellerProductSlice `boil:"SellerProducts" json:"SellerProducts" toml:"SellerProducts" yaml:"SellerProducts"`
}

// NewStruct creates a new relationship struct
func (*brandR) NewStruct() *brandR {
	return &brandR{}
}

// brandL is where Load methods for each relationship are stored.
type brandL struct{}

var (
	brandAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "name"}
	brandColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "name"}
	brandColumnsWithDefault    = []string{"id"}
	brandPrimaryKeyColumns     = []string{"id"}
)

type (
	// BrandSlice is an alias for a slice of pointers to Brand.
	// This should almost always be used instead of []Brand.
	BrandSlice []*Brand
	// BrandHook is the signature for custom Brand hook methods
	BrandHook func(context.Context, boil.ContextExecutor, *Brand) error

	brandQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	brandType                 = reflect.TypeOf(&Brand{})
	brandMapping              = queries.MakeStructMapping(brandType)
	brandPrimaryKeyMapping, _ = queries.BindMapping(brandType, brandMapping, brandPrimaryKeyColumns)
	brandInsertCacheMut       sync.RWMutex
	brandInsertCache          = make(map[string]insertCache)
	brandUpdateCacheMut       sync.RWMutex
	brandUpdateCache          = make(map[string]updateCache)
	brandUpsertCacheMut       sync.RWMutex
	brandUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var brandBeforeInsertHooks []BrandHook
var brandBeforeUpdateHooks []BrandHook
var brandBeforeDeleteHooks []BrandHook
var brandBeforeUpsertHooks []BrandHook

var brandAfterInsertHooks []BrandHook
var brandAfterSelectHooks []BrandHook
var brandAfterUpdateHooks []BrandHook
var brandAfterDeleteHooks []BrandHook
var brandAfterUpsertHooks []BrandHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Brand) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Brand) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Brand) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Brand) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Brand) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Brand) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Brand) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Brand) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Brand) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range brandAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBrandHook registers your hook function for all future operations.
func AddBrandHook(hookPoint boil.HookPoint, brandHook BrandHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		brandBeforeInsertHooks = append(brandBeforeInsertHooks, brandHook)
	case boil.BeforeUpdateHook:
		brandBeforeUpdateHooks = append(brandBeforeUpdateHooks, brandHook)
	case boil.BeforeDeleteHook:
		brandBeforeDeleteHooks = append(brandBeforeDeleteHooks, brandHook)
	case boil.BeforeUpsertHook:
		brandBeforeUpsertHooks = append(brandBeforeUpsertHooks, brandHook)
	case boil.AfterInsertHook:
		brandAfterInsertHooks = append(brandAfterInsertHooks, brandHook)
	case boil.AfterSelectHook:
		brandAfterSelectHooks = append(brandAfterSelectHooks, brandHook)
	case boil.AfterUpdateHook:
		brandAfterUpdateHooks = append(brandAfterUpdateHooks, brandHook)
	case boil.AfterDeleteHook:
		brandAfterDeleteHooks = append(brandAfterDeleteHooks, brandHook)
	case boil.AfterUpsertHook:
		brandAfterUpsertHooks = append(brandAfterUpsertHooks, brandHook)
	}
}

// One returns a single brand record from the query.
func (q brandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Brand, error) {
	o := &Brand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for brands")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Brand records from the query.
func (q brandQuery) All(ctx context.Context, exec boil.ContextExecutor) (BrandSlice, error) {
	var o []*Brand

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Brand slice")
	}

	if len(brandAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Brand records in the query.
func (q brandQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count brands rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q brandQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if brands exists")
	}

	return count > 0, nil
}

// SellerProducts retrieves all the seller_product's SellerProducts with an executor.
func (o *Brand) SellerProducts(mods ...qm.QueryMod) sellerProductQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"seller_products\".\"brand_id\"=?", o.ID),
	)

	query := SellerProducts(queryMods...)
	queries.SetFrom(query.Query, "\"seller_products\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"seller_products\".*"})
	}

	return query
}

// LoadSellerProducts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (brandL) LoadSellerProducts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBrand interface{}, mods queries.Applicator) error {
	var slice []*Brand
	var object *Brand

	if singular {
		object = maybeBrand.(*Brand)
	} else {
		slice = *maybeBrand.(*[]*Brand)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &brandR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &brandR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`seller_products`),
		qm.WhereIn(`seller_products.brand_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load seller_products")
	}

	var resultSlice []*SellerProduct
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice seller_products")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on seller_products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for seller_products")
	}

	if len(sellerProductAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SellerProducts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sellerProductR{}
			}
			foreign.R.Brand = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.BrandID) {
				local.R.SellerProducts = append(local.R.SellerProducts, foreign)
				if foreign.R == nil {
					foreign.R = &sellerProductR{}
				}
				foreign.R.Brand = local
				break
			}
		}
	}

	return nil
}

// AddSellerProducts adds the given related objects to the existing relationships
// of the brand, optionally inserting them as new records.
// Appends related to o.R.SellerProducts.
// Sets related.R.Brand appropriately.
func (o *Brand) AddSellerProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SellerProduct) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.BrandID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"seller_products\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"brand_id"}),
				strmangle.WhereClause("\"", "\"", 2, sellerProductPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.BrandID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &brandR{
			SellerProducts: related,
		}
	} else {
		o.R.SellerProducts = append(o.R.SellerProducts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sellerProductR{
				Brand: o,
			}
		} else {
			rel.R.Brand = o
		}
	}
	return nil
}

// SetSellerProducts removes all previously related items of the
// brand replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Brand's SellerProducts accordingly.
// Replaces o.R.SellerProducts with related.
// Sets related.R.Brand's SellerProducts accordingly.
func (o *Brand) SetSellerProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SellerProduct) error {
	query := "update \"seller_products\" set \"brand_id\" = null where \"brand_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.SellerProducts {
			queries.SetScanner(&rel.BrandID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Brand = nil
		}

		o.R.SellerProducts = nil
	}
	return o.AddSellerProducts(ctx, exec, insert, related...)
}

// RemoveSellerProducts relationships from objects passed in.
// Removes related items from R.SellerProducts (uses pointer comparison, removal does not keep order)
// Sets related.R.Brand.
func (o *Brand) RemoveSellerProducts(ctx context.Context, exec boil.ContextExecutor, related ...*SellerProduct) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.BrandID, nil)
		if rel.R != nil {
			rel.R.Brand = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("brand_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.SellerProducts {
			if rel != ri {
				continue
			}

			ln := len(o.R.SellerProducts)
			if ln > 1 && i < ln-1 {
				o.R.SellerProducts[i] = o.R.SellerProducts[ln-1]
			}
			o.R.SellerProducts = o.R.SellerProducts[:ln-1]
			break
		}
	}

	return nil
}

// Brands retrieves all the records using an executor.
func Brands(mods ...qm.QueryMod) brandQuery {
	mods = append(mods, qm.From("\"brands\""))
	return brandQuery{NewQuery(mods...)}
}

// FindBrand retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBrand(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Brand, error) {
	brandObj := &Brand{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"brands\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, brandObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from brands")
	}

	if err = brandObj.doAfterSelectHooks(ctx, exec); err != nil {
		return brandObj, err
	}

	return brandObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Brand) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no brands provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(brandColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	brandInsertCacheMut.RLock()
	cache, cached := brandInsertCache[key]
	brandInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			brandAllColumns,
			brandColumnsWithDefault,
			brandColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(brandType, brandMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(brandType, brandMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"brands\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"brands\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into brands")
	}

	if !cached {
		brandInsertCacheMut.Lock()
		brandInsertCache[key] = cache
		brandInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Brand.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Brand) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	brandUpdateCacheMut.RLock()
	cache, cached := brandUpdateCache[key]
	brandUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			brandAllColumns,
			brandPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update brands, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"brands\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, brandPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(brandType, brandMapping, append(wl, brandPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update brands row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for brands")
	}

	if !cached {
		brandUpdateCacheMut.Lock()
		brandUpdateCache[key] = cache
		brandUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q brandQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for brands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for brands")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BrandSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"brands\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, brandPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in brand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all brand")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Brand) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no brands provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(brandColumnsWithDefault, o)

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

	brandUpsertCacheMut.RLock()
	cache, cached := brandUpsertCache[key]
	brandUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			brandAllColumns,
			brandColumnsWithDefault,
			brandColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			brandAllColumns,
			brandPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert brands, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(brandPrimaryKeyColumns))
			copy(conflict, brandPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"brands\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(brandType, brandMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(brandType, brandMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert brands")
	}

	if !cached {
		brandUpsertCacheMut.Lock()
		brandUpsertCache[key] = cache
		brandUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Brand record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Brand) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Brand provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), brandPrimaryKeyMapping)
	sql := "DELETE FROM \"brands\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from brands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for brands")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q brandQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no brandQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from brands")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for brands")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BrandSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(brandBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"brands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, brandPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from brand slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for brands")
	}

	if len(brandAfterDeleteHooks) != 0 {
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
func (o *Brand) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBrand(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BrandSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BrandSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brandPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"brands\".* FROM \"brands\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, brandPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BrandSlice")
	}

	*o = slice

	return nil
}

// BrandExists checks if the Brand row exists.
func BrandExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"brands\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if brands exists")
	}

	return exists, nil
}

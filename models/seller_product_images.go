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

// SellerProductImage is an object representing the database table.
type SellerProductImage struct {
	ID              int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt       null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt       null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt       null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	SellerProductID null.Int64  `boil:"seller_product_id" json:"seller_product_id,omitempty" toml:"seller_product_id" yaml:"seller_product_id,omitempty"`
	Display         null.Bool   `boil:"display" json:"display,omitempty" toml:"display" yaml:"display,omitempty"`
	Image           null.String `boil:"image" json:"image,omitempty" toml:"image" yaml:"image,omitempty"`

	R *sellerProductImageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sellerProductImageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SellerProductImageColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	SellerProductID string
	Display         string
	Image           string
}{
	ID:              "id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	SellerProductID: "seller_product_id",
	Display:         "display",
	Image:           "image",
}

var SellerProductImageTableColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	SellerProductID string
	Display         string
	Image           string
}{
	ID:              "seller_product_images.id",
	CreatedAt:       "seller_product_images.created_at",
	UpdatedAt:       "seller_product_images.updated_at",
	DeletedAt:       "seller_product_images.deleted_at",
	SellerProductID: "seller_product_images.seller_product_id",
	Display:         "seller_product_images.display",
	Image:           "seller_product_images.image",
}

// Generated where

var SellerProductImageWhere = struct {
	ID              whereHelperint64
	CreatedAt       whereHelpernull_Time
	UpdatedAt       whereHelpernull_Time
	DeletedAt       whereHelpernull_Time
	SellerProductID whereHelpernull_Int64
	Display         whereHelpernull_Bool
	Image           whereHelpernull_String
}{
	ID:              whereHelperint64{field: "\"seller_product_images\".\"id\""},
	CreatedAt:       whereHelpernull_Time{field: "\"seller_product_images\".\"created_at\""},
	UpdatedAt:       whereHelpernull_Time{field: "\"seller_product_images\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"seller_product_images\".\"deleted_at\""},
	SellerProductID: whereHelpernull_Int64{field: "\"seller_product_images\".\"seller_product_id\""},
	Display:         whereHelpernull_Bool{field: "\"seller_product_images\".\"display\""},
	Image:           whereHelpernull_String{field: "\"seller_product_images\".\"image\""},
}

// SellerProductImageRels is where relationship names are stored.
var SellerProductImageRels = struct {
	SellerProduct string
}{
	SellerProduct: "SellerProduct",
}

// sellerProductImageR is where relationships are stored.
type sellerProductImageR struct {
	SellerProduct *SellerProduct `boil:"SellerProduct" json:"SellerProduct" toml:"SellerProduct" yaml:"SellerProduct"`
}

// NewStruct creates a new relationship struct
func (*sellerProductImageR) NewStruct() *sellerProductImageR {
	return &sellerProductImageR{}
}

// sellerProductImageL is where Load methods for each relationship are stored.
type sellerProductImageL struct{}

var (
	sellerProductImageAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "seller_product_id", "display", "image"}
	sellerProductImageColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "seller_product_id", "image"}
	sellerProductImageColumnsWithDefault    = []string{"id", "display"}
	sellerProductImagePrimaryKeyColumns     = []string{"id"}
)

type (
	// SellerProductImageSlice is an alias for a slice of pointers to SellerProductImage.
	// This should almost always be used instead of []SellerProductImage.
	SellerProductImageSlice []*SellerProductImage
	// SellerProductImageHook is the signature for custom SellerProductImage hook methods
	SellerProductImageHook func(context.Context, boil.ContextExecutor, *SellerProductImage) error

	sellerProductImageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sellerProductImageType                 = reflect.TypeOf(&SellerProductImage{})
	sellerProductImageMapping              = queries.MakeStructMapping(sellerProductImageType)
	sellerProductImagePrimaryKeyMapping, _ = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, sellerProductImagePrimaryKeyColumns)
	sellerProductImageInsertCacheMut       sync.RWMutex
	sellerProductImageInsertCache          = make(map[string]insertCache)
	sellerProductImageUpdateCacheMut       sync.RWMutex
	sellerProductImageUpdateCache          = make(map[string]updateCache)
	sellerProductImageUpsertCacheMut       sync.RWMutex
	sellerProductImageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sellerProductImageBeforeInsertHooks []SellerProductImageHook
var sellerProductImageBeforeUpdateHooks []SellerProductImageHook
var sellerProductImageBeforeDeleteHooks []SellerProductImageHook
var sellerProductImageBeforeUpsertHooks []SellerProductImageHook

var sellerProductImageAfterInsertHooks []SellerProductImageHook
var sellerProductImageAfterSelectHooks []SellerProductImageHook
var sellerProductImageAfterUpdateHooks []SellerProductImageHook
var sellerProductImageAfterDeleteHooks []SellerProductImageHook
var sellerProductImageAfterUpsertHooks []SellerProductImageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SellerProductImage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SellerProductImage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SellerProductImage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SellerProductImage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SellerProductImage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SellerProductImage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SellerProductImage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SellerProductImage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SellerProductImage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sellerProductImageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSellerProductImageHook registers your hook function for all future operations.
func AddSellerProductImageHook(hookPoint boil.HookPoint, sellerProductImageHook SellerProductImageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		sellerProductImageBeforeInsertHooks = append(sellerProductImageBeforeInsertHooks, sellerProductImageHook)
	case boil.BeforeUpdateHook:
		sellerProductImageBeforeUpdateHooks = append(sellerProductImageBeforeUpdateHooks, sellerProductImageHook)
	case boil.BeforeDeleteHook:
		sellerProductImageBeforeDeleteHooks = append(sellerProductImageBeforeDeleteHooks, sellerProductImageHook)
	case boil.BeforeUpsertHook:
		sellerProductImageBeforeUpsertHooks = append(sellerProductImageBeforeUpsertHooks, sellerProductImageHook)
	case boil.AfterInsertHook:
		sellerProductImageAfterInsertHooks = append(sellerProductImageAfterInsertHooks, sellerProductImageHook)
	case boil.AfterSelectHook:
		sellerProductImageAfterSelectHooks = append(sellerProductImageAfterSelectHooks, sellerProductImageHook)
	case boil.AfterUpdateHook:
		sellerProductImageAfterUpdateHooks = append(sellerProductImageAfterUpdateHooks, sellerProductImageHook)
	case boil.AfterDeleteHook:
		sellerProductImageAfterDeleteHooks = append(sellerProductImageAfterDeleteHooks, sellerProductImageHook)
	case boil.AfterUpsertHook:
		sellerProductImageAfterUpsertHooks = append(sellerProductImageAfterUpsertHooks, sellerProductImageHook)
	}
}

// One returns a single sellerProductImage record from the query.
func (q sellerProductImageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SellerProductImage, error) {
	o := &SellerProductImage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for seller_product_images")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SellerProductImage records from the query.
func (q sellerProductImageQuery) All(ctx context.Context, exec boil.ContextExecutor) (SellerProductImageSlice, error) {
	var o []*SellerProductImage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SellerProductImage slice")
	}

	if len(sellerProductImageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SellerProductImage records in the query.
func (q sellerProductImageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count seller_product_images rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sellerProductImageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if seller_product_images exists")
	}

	return count > 0, nil
}

// SellerProduct pointed to by the foreign key.
func (o *SellerProductImage) SellerProduct(mods ...qm.QueryMod) sellerProductQuery {
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
func (sellerProductImageL) LoadSellerProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSellerProductImage interface{}, mods queries.Applicator) error {
	var slice []*SellerProductImage
	var object *SellerProductImage

	if singular {
		object = maybeSellerProductImage.(*SellerProductImage)
	} else {
		slice = *maybeSellerProductImage.(*[]*SellerProductImage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &sellerProductImageR{}
		}
		if !queries.IsNil(object.SellerProductID) {
			args = append(args, object.SellerProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &sellerProductImageR{}
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

	if len(sellerProductImageAfterSelectHooks) != 0 {
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
		foreign.R.SellerProductImages = append(foreign.R.SellerProductImages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SellerProductID, foreign.ID) {
				local.R.SellerProduct = foreign
				if foreign.R == nil {
					foreign.R = &sellerProductR{}
				}
				foreign.R.SellerProductImages = append(foreign.R.SellerProductImages, local)
				break
			}
		}
	}

	return nil
}

// SetSellerProduct of the sellerProductImage to the related item.
// Sets o.R.SellerProduct to related.
// Adds o to related.R.SellerProductImages.
func (o *SellerProductImage) SetSellerProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SellerProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"seller_product_images\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"seller_product_id"}),
		strmangle.WhereClause("\"", "\"", 2, sellerProductImagePrimaryKeyColumns),
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
		o.R = &sellerProductImageR{
			SellerProduct: related,
		}
	} else {
		o.R.SellerProduct = related
	}

	if related.R == nil {
		related.R = &sellerProductR{
			SellerProductImages: SellerProductImageSlice{o},
		}
	} else {
		related.R.SellerProductImages = append(related.R.SellerProductImages, o)
	}

	return nil
}

// RemoveSellerProduct relationship.
// Sets o.R.SellerProduct to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *SellerProductImage) RemoveSellerProduct(ctx context.Context, exec boil.ContextExecutor, related *SellerProduct) error {
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

	for i, ri := range related.R.SellerProductImages {
		if queries.Equal(o.SellerProductID, ri.SellerProductID) {
			continue
		}

		ln := len(related.R.SellerProductImages)
		if ln > 1 && i < ln-1 {
			related.R.SellerProductImages[i] = related.R.SellerProductImages[ln-1]
		}
		related.R.SellerProductImages = related.R.SellerProductImages[:ln-1]
		break
	}
	return nil
}

// SellerProductImages retrieves all the records using an executor.
func SellerProductImages(mods ...qm.QueryMod) sellerProductImageQuery {
	mods = append(mods, qm.From("\"seller_product_images\""))
	return sellerProductImageQuery{NewQuery(mods...)}
}

// FindSellerProductImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSellerProductImage(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*SellerProductImage, error) {
	sellerProductImageObj := &SellerProductImage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"seller_product_images\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, sellerProductImageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from seller_product_images")
	}

	if err = sellerProductImageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sellerProductImageObj, err
	}

	return sellerProductImageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SellerProductImage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no seller_product_images provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(sellerProductImageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sellerProductImageInsertCacheMut.RLock()
	cache, cached := sellerProductImageInsertCache[key]
	sellerProductImageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sellerProductImageAllColumns,
			sellerProductImageColumnsWithDefault,
			sellerProductImageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"seller_product_images\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"seller_product_images\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into seller_product_images")
	}

	if !cached {
		sellerProductImageInsertCacheMut.Lock()
		sellerProductImageInsertCache[key] = cache
		sellerProductImageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SellerProductImage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SellerProductImage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sellerProductImageUpdateCacheMut.RLock()
	cache, cached := sellerProductImageUpdateCache[key]
	sellerProductImageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sellerProductImageAllColumns,
			sellerProductImagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update seller_product_images, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"seller_product_images\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, sellerProductImagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, append(wl, sellerProductImagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update seller_product_images row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for seller_product_images")
	}

	if !cached {
		sellerProductImageUpdateCacheMut.Lock()
		sellerProductImageUpdateCache[key] = cache
		sellerProductImageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sellerProductImageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for seller_product_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for seller_product_images")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SellerProductImageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"seller_product_images\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, sellerProductImagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sellerProductImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sellerProductImage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SellerProductImage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no seller_product_images provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(sellerProductImageColumnsWithDefault, o)

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

	sellerProductImageUpsertCacheMut.RLock()
	cache, cached := sellerProductImageUpsertCache[key]
	sellerProductImageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sellerProductImageAllColumns,
			sellerProductImageColumnsWithDefault,
			sellerProductImageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			sellerProductImageAllColumns,
			sellerProductImagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert seller_product_images, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(sellerProductImagePrimaryKeyColumns))
			copy(conflict, sellerProductImagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"seller_product_images\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sellerProductImageType, sellerProductImageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert seller_product_images")
	}

	if !cached {
		sellerProductImageUpsertCacheMut.Lock()
		sellerProductImageUpsertCache[key] = cache
		sellerProductImageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SellerProductImage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SellerProductImage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SellerProductImage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sellerProductImagePrimaryKeyMapping)
	sql := "DELETE FROM \"seller_product_images\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from seller_product_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for seller_product_images")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sellerProductImageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sellerProductImageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from seller_product_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for seller_product_images")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SellerProductImageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sellerProductImageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"seller_product_images\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerProductImagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sellerProductImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for seller_product_images")
	}

	if len(sellerProductImageAfterDeleteHooks) != 0 {
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
func (o *SellerProductImage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSellerProductImage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SellerProductImageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SellerProductImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sellerProductImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"seller_product_images\".* FROM \"seller_product_images\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sellerProductImagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SellerProductImageSlice")
	}

	*o = slice

	return nil
}

// SellerProductImageExists checks if the SellerProductImage row exists.
func SellerProductImageExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"seller_product_images\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if seller_product_images exists")
	}

	return exists, nil
}

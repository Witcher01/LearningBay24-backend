// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testForums(t *testing.T) {
	t.Parallel()

	query := Forums()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testForumsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Forums().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ForumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Forums().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ForumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testForumsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ForumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Forum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ForumExists to return true, but got false.")
	}
}

func testForumsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	forumFound, err := FindForum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if forumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testForumsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Forums().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testForumsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Forums().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testForumsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	forumOne := &Forum{}
	forumTwo := &Forum{}
	if err = randomize.Struct(seed, forumOne, forumDBTypes, false, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}
	if err = randomize.Struct(seed, forumTwo, forumDBTypes, false, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = forumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = forumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Forums().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testForumsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	forumOne := &Forum{}
	forumTwo := &Forum{}
	if err = randomize.Struct(seed, forumOne, forumDBTypes, false, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}
	if err = randomize.Struct(seed, forumTwo, forumDBTypes, false, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = forumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = forumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func forumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func forumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Forum) error {
	*o = Forum{}
	return nil
}

func testForumsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Forum{}
	o := &Forum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, forumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Forum object: %s", err)
	}

	AddForumHook(boil.BeforeInsertHook, forumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	forumBeforeInsertHooks = []ForumHook{}

	AddForumHook(boil.AfterInsertHook, forumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	forumAfterInsertHooks = []ForumHook{}

	AddForumHook(boil.AfterSelectHook, forumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	forumAfterSelectHooks = []ForumHook{}

	AddForumHook(boil.BeforeUpdateHook, forumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	forumBeforeUpdateHooks = []ForumHook{}

	AddForumHook(boil.AfterUpdateHook, forumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	forumAfterUpdateHooks = []ForumHook{}

	AddForumHook(boil.BeforeDeleteHook, forumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	forumBeforeDeleteHooks = []ForumHook{}

	AddForumHook(boil.AfterDeleteHook, forumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	forumAfterDeleteHooks = []ForumHook{}

	AddForumHook(boil.BeforeUpsertHook, forumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	forumBeforeUpsertHooks = []ForumHook{}

	AddForumHook(boil.AfterUpsertHook, forumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	forumAfterUpsertHooks = []ForumHook{}
}

func testForumsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testForumsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(forumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testForumToManyCourses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Forum
	var b, c Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, courseDBTypes, false, courseColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, courseDBTypes, false, courseColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ForumID = a.ID
	c.ForumID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Courses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ForumID == b.ForumID {
			bFound = true
		}
		if v.ForumID == c.ForumID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ForumSlice{&a}
	if err = a.L.LoadCourses(ctx, tx, false, (*[]*Forum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Courses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Courses = nil
	if err = a.L.LoadCourses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Courses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testForumToManyForumEntries(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Forum
	var b, c ForumEntry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, forumEntryDBTypes, false, forumEntryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, forumEntryDBTypes, false, forumEntryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ForumID = a.ID
	c.ForumID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ForumEntries().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ForumID == b.ForumID {
			bFound = true
		}
		if v.ForumID == c.ForumID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ForumSlice{&a}
	if err = a.L.LoadForumEntries(ctx, tx, false, (*[]*Forum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ForumEntries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ForumEntries = nil
	if err = a.L.LoadForumEntries(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ForumEntries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testForumToManyAddOpCourses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Forum
	var b, c, d, e Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, forumDBTypes, false, strmangle.SetComplement(forumPrimaryKeyColumns, forumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Course{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Course{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCourses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ForumID {
			t.Error("foreign key was wrong value", a.ID, first.ForumID)
		}
		if a.ID != second.ForumID {
			t.Error("foreign key was wrong value", a.ID, second.ForumID)
		}

		if first.R.Forum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Forum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Courses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Courses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Courses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testForumToManyAddOpForumEntries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Forum
	var b, c, d, e ForumEntry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, forumDBTypes, false, strmangle.SetComplement(forumPrimaryKeyColumns, forumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ForumEntry{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, forumEntryDBTypes, false, strmangle.SetComplement(forumEntryPrimaryKeyColumns, forumEntryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ForumEntry{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddForumEntries(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ForumID {
			t.Error("foreign key was wrong value", a.ID, first.ForumID)
		}
		if a.ID != second.ForumID {
			t.Error("foreign key was wrong value", a.ID, second.ForumID)
		}

		if first.R.Forum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Forum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ForumEntries[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ForumEntries[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ForumEntries().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testForumsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testForumsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ForumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testForumsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Forums().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	forumDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_            = bytes.MinRead
)

func testForumsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(forumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(forumAllColumns) == len(forumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, forumDBTypes, true, forumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testForumsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(forumAllColumns) == len(forumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Forum{}
	if err = randomize.Struct(seed, o, forumDBTypes, true, forumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, forumDBTypes, true, forumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(forumAllColumns, forumPrimaryKeyColumns) {
		fields = forumAllColumns
	} else {
		fields = strmangle.SetComplement(
			forumAllColumns,
			forumPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ForumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testForumsUpsert(t *testing.T) {
	t.Parallel()

	if len(forumAllColumns) == len(forumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLForumUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Forum{}
	if err = randomize.Struct(seed, &o, forumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Forum: %s", err)
	}

	count, err := Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, forumDBTypes, false, forumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Forum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Forum: %s", err)
	}

	count, err = Forums().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
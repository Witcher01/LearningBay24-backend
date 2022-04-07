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

func testDirectories(t *testing.T) {
	t.Parallel()

	query := Directories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDirectoriesSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
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

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Directories().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DirectorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
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

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Directories().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DirectorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDirectoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DirectoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Directory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DirectoryExists to return true, but got false.")
	}
}

func testDirectoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	directoryFound, err := FindDirectory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if directoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDirectoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Directories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDirectoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Directories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDirectoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	directoryOne := &Directory{}
	directoryTwo := &Directory{}
	if err = randomize.Struct(seed, directoryOne, directoryDBTypes, false, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}
	if err = randomize.Struct(seed, directoryTwo, directoryDBTypes, false, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = directoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = directoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Directories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDirectoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	directoryOne := &Directory{}
	directoryTwo := &Directory{}
	if err = randomize.Struct(seed, directoryOne, directoryDBTypes, false, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}
	if err = randomize.Struct(seed, directoryTwo, directoryDBTypes, false, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = directoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = directoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func directoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func directoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Directory) error {
	*o = Directory{}
	return nil
}

func testDirectoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Directory{}
	o := &Directory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, directoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Directory object: %s", err)
	}

	AddDirectoryHook(boil.BeforeInsertHook, directoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	directoryBeforeInsertHooks = []DirectoryHook{}

	AddDirectoryHook(boil.AfterInsertHook, directoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	directoryAfterInsertHooks = []DirectoryHook{}

	AddDirectoryHook(boil.AfterSelectHook, directoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	directoryAfterSelectHooks = []DirectoryHook{}

	AddDirectoryHook(boil.BeforeUpdateHook, directoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	directoryBeforeUpdateHooks = []DirectoryHook{}

	AddDirectoryHook(boil.AfterUpdateHook, directoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	directoryAfterUpdateHooks = []DirectoryHook{}

	AddDirectoryHook(boil.BeforeDeleteHook, directoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	directoryBeforeDeleteHooks = []DirectoryHook{}

	AddDirectoryHook(boil.AfterDeleteHook, directoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	directoryAfterDeleteHooks = []DirectoryHook{}

	AddDirectoryHook(boil.BeforeUpsertHook, directoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	directoryBeforeUpsertHooks = []DirectoryHook{}

	AddDirectoryHook(boil.AfterUpsertHook, directoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	directoryAfterUpsertHooks = []DirectoryHook{}
}

func testDirectoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDirectoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(directoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDirectoryToManyFiles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Directory
	var b, c File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into `directory_has_files` (`directory_id`, `file_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `directory_has_files` (`directory_id`, `file_id`) values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Files().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DirectorySlice{&a}
	if err = a.L.LoadFiles(ctx, tx, false, (*[]*Directory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Files); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Files = nil
	if err = a.L.LoadFiles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Files); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDirectoryToManyAddOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Directory
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, directoryDBTypes, false, strmangle.SetComplement(directoryPrimaryKeyColumns, directoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*File{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*File{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFiles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Directories[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Directories[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Files[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Files[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Files().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDirectoryToManySetOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Directory
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, directoryDBTypes, false, strmangle.SetComplement(directoryPrimaryKeyColumns, directoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*File{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetFiles(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Files().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetFiles(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Files().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Directories) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Directories) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Directories[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Directories[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Files[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Files[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDirectoryToManyRemoveOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Directory
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, directoryDBTypes, false, strmangle.SetComplement(directoryPrimaryKeyColumns, directoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*File{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, fileDBTypes, false, strmangle.SetComplement(filePrimaryKeyColumns, fileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddFiles(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Files().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveFiles(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Files().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Directories) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Directories) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Directories[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Directories[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Files) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Files[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Files[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDirectoryToOneCourseUsingCourse(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Directory
	var foreign Course

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, directoryDBTypes, false, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, courseDBTypes, false, courseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Course struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CourseID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Course().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DirectorySlice{&local}
	if err = local.L.LoadCourse(ctx, tx, false, (*[]*Directory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Course = nil
	if err = local.L.LoadCourse(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Course == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDirectoryToOneSetOpCourseUsingCourse(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Directory
	var b, c Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, directoryDBTypes, false, strmangle.SetComplement(directoryPrimaryKeyColumns, directoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, courseDBTypes, false, strmangle.SetComplement(coursePrimaryKeyColumns, courseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Course{&b, &c} {
		err = a.SetCourse(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Course != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Directories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CourseID != x.ID {
			t.Error("foreign key was wrong value", a.CourseID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CourseID))
		reflect.Indirect(reflect.ValueOf(&a.CourseID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CourseID != x.ID {
			t.Error("foreign key was wrong value", a.CourseID, x.ID)
		}
	}
}

func testDirectoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
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

func testDirectoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DirectorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDirectoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Directories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	directoryDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `CourseID`: `int`, `VisibleFrom`: `timestamp`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                = bytes.MinRead
)

func testDirectoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(directoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(directoryAllColumns) == len(directoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDirectoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(directoryAllColumns) == len(directoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Directory{}
	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, directoryDBTypes, true, directoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(directoryAllColumns, directoryPrimaryKeyColumns) {
		fields = directoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			directoryAllColumns,
			directoryPrimaryKeyColumns,
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

	slice := DirectorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDirectoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(directoryAllColumns) == len(directoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDirectoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Directory{}
	if err = randomize.Struct(seed, &o, directoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Directory: %s", err)
	}

	count, err := Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, directoryDBTypes, false, directoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Directory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Directory: %s", err)
	}

	count, err = Directories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
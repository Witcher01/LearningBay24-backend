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

func testSubmissions(t *testing.T) {
	t.Parallel()

	query := Submissions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSubmissionsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
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

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Submissions().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubmissionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
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

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Submissions().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubmissionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubmissionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SubmissionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Submission exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SubmissionExists to return true, but got false.")
	}
}

func testSubmissionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	submissionFound, err := FindSubmission(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if submissionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSubmissionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Submissions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSubmissionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Submissions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSubmissionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	submissionOne := &Submission{}
	submissionTwo := &Submission{}
	if err = randomize.Struct(seed, submissionOne, submissionDBTypes, false, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}
	if err = randomize.Struct(seed, submissionTwo, submissionDBTypes, false, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = submissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = submissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Submissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSubmissionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	submissionOne := &Submission{}
	submissionTwo := &Submission{}
	if err = randomize.Struct(seed, submissionOne, submissionDBTypes, false, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}
	if err = randomize.Struct(seed, submissionTwo, submissionDBTypes, false, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = submissionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = submissionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func submissionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func submissionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Submission) error {
	*o = Submission{}
	return nil
}

func testSubmissionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Submission{}
	o := &Submission{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, submissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Submission object: %s", err)
	}

	AddSubmissionHook(boil.BeforeInsertHook, submissionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	submissionBeforeInsertHooks = []SubmissionHook{}

	AddSubmissionHook(boil.AfterInsertHook, submissionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	submissionAfterInsertHooks = []SubmissionHook{}

	AddSubmissionHook(boil.AfterSelectHook, submissionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	submissionAfterSelectHooks = []SubmissionHook{}

	AddSubmissionHook(boil.BeforeUpdateHook, submissionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	submissionBeforeUpdateHooks = []SubmissionHook{}

	AddSubmissionHook(boil.AfterUpdateHook, submissionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	submissionAfterUpdateHooks = []SubmissionHook{}

	AddSubmissionHook(boil.BeforeDeleteHook, submissionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	submissionBeforeDeleteHooks = []SubmissionHook{}

	AddSubmissionHook(boil.AfterDeleteHook, submissionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	submissionAfterDeleteHooks = []SubmissionHook{}

	AddSubmissionHook(boil.BeforeUpsertHook, submissionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	submissionBeforeUpsertHooks = []SubmissionHook{}

	AddSubmissionHook(boil.AfterUpsertHook, submissionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	submissionAfterUpsertHooks = []SubmissionHook{}
}

func testSubmissionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubmissionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(submissionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubmissionToManyFiles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
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

	_, err = tx.Exec("insert into `submission_has_files` (`submission_id`, `file_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `submission_has_files` (`submission_id`, `file_id`) values (?, ?)", a.ID, c.ID)
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

	slice := SubmissionSlice{&a}
	if err = a.L.LoadFiles(ctx, tx, false, (*[]*Submission)(&slice), nil); err != nil {
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

func testSubmissionToManyUserSubmissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c UserSubmission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userSubmissionDBTypes, false, userSubmissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userSubmissionDBTypes, false, userSubmissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SubmissionID = a.ID
	c.SubmissionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserSubmissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SubmissionID == b.SubmissionID {
			bFound = true
		}
		if v.SubmissionID == c.SubmissionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SubmissionSlice{&a}
	if err = a.L.LoadUserSubmissions(ctx, tx, false, (*[]*Submission)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubmissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserSubmissions = nil
	if err = a.L.LoadUserSubmissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubmissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSubmissionToManyAddOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, false, strmangle.SetComplement(submissionPrimaryKeyColumns, submissionColumnsWithoutDefault)...); err != nil {
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

		if first.R.Submissions[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Submissions[0] != &a {
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

func testSubmissionToManySetOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, false, strmangle.SetComplement(submissionPrimaryKeyColumns, submissionColumnsWithoutDefault)...); err != nil {
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
	// if len(b.R.Submissions) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Submissions) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Submissions[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Submissions[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Files[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Files[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testSubmissionToManyRemoveOpFiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c, d, e File

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, false, strmangle.SetComplement(submissionPrimaryKeyColumns, submissionColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.Submissions) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Submissions) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Submissions[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Submissions[0] != &a {
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

func testSubmissionToManyAddOpUserSubmissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c, d, e UserSubmission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, false, strmangle.SetComplement(submissionPrimaryKeyColumns, submissionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserSubmission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userSubmissionDBTypes, false, strmangle.SetComplement(userSubmissionPrimaryKeyColumns, userSubmissionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserSubmission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserSubmissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SubmissionID {
			t.Error("foreign key was wrong value", a.ID, first.SubmissionID)
		}
		if a.ID != second.SubmissionID {
			t.Error("foreign key was wrong value", a.ID, second.SubmissionID)
		}

		if first.R.Submission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Submission != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserSubmissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserSubmissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserSubmissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSubmissionToOneCourseUsingCourse(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Submission
	var foreign Course

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, submissionDBTypes, false, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
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

	slice := SubmissionSlice{&local}
	if err = local.L.LoadCourse(ctx, tx, false, (*[]*Submission)(&slice), nil); err != nil {
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

func testSubmissionToOneSetOpCourseUsingCourse(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Submission
	var b, c Course

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, submissionDBTypes, false, strmangle.SetComplement(submissionPrimaryKeyColumns, submissionColumnsWithoutDefault)...); err != nil {
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

		if x.R.Submissions[0] != &a {
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

func testSubmissionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
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

func testSubmissionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubmissionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSubmissionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Submissions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	submissionDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Deadline`: `timestamp`, `CourseID`: `int`, `MaxFilesize`: `int`, `VisibleFrom`: `timestamp`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `GradedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                 = bytes.MinRead
)

func testSubmissionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(submissionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(submissionAllColumns) == len(submissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSubmissionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(submissionAllColumns) == len(submissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Submission{}
	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, submissionDBTypes, true, submissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(submissionAllColumns, submissionPrimaryKeyColumns) {
		fields = submissionAllColumns
	} else {
		fields = strmangle.SetComplement(
			submissionAllColumns,
			submissionPrimaryKeyColumns,
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

	slice := SubmissionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSubmissionsUpsert(t *testing.T) {
	t.Parallel()

	if len(submissionAllColumns) == len(submissionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSubmissionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Submission{}
	if err = randomize.Struct(seed, &o, submissionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Submission: %s", err)
	}

	count, err := Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, submissionDBTypes, false, submissionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Submission struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Submission: %s", err)
	}

	count, err = Submissions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
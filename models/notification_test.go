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

func testNotifications(t *testing.T) {
	t.Parallel()

	query := Notifications()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNotificationsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
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

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Notifications().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
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

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Notifications().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNotificationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NotificationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Notification exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NotificationExists to return true, but got false.")
	}
}

func testNotificationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	notificationFound, err := FindNotification(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if notificationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNotificationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Notifications().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNotificationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Notifications().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNotificationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	notificationOne := &Notification{}
	notificationTwo := &Notification{}
	if err = randomize.Struct(seed, notificationOne, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTwo, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Notifications().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNotificationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	notificationOne := &Notification{}
	notificationTwo := &Notification{}
	if err = randomize.Struct(seed, notificationOne, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}
	if err = randomize.Struct(seed, notificationTwo, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = notificationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = notificationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func notificationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func notificationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Notification) error {
	*o = Notification{}
	return nil
}

func testNotificationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Notification{}
	o := &Notification{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, notificationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Notification object: %s", err)
	}

	AddNotificationHook(boil.BeforeInsertHook, notificationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	notificationBeforeInsertHooks = []NotificationHook{}

	AddNotificationHook(boil.AfterInsertHook, notificationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	notificationAfterInsertHooks = []NotificationHook{}

	AddNotificationHook(boil.AfterSelectHook, notificationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	notificationAfterSelectHooks = []NotificationHook{}

	AddNotificationHook(boil.BeforeUpdateHook, notificationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	notificationBeforeUpdateHooks = []NotificationHook{}

	AddNotificationHook(boil.AfterUpdateHook, notificationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	notificationAfterUpdateHooks = []NotificationHook{}

	AddNotificationHook(boil.BeforeDeleteHook, notificationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	notificationBeforeDeleteHooks = []NotificationHook{}

	AddNotificationHook(boil.AfterDeleteHook, notificationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	notificationAfterDeleteHooks = []NotificationHook{}

	AddNotificationHook(boil.BeforeUpsertHook, notificationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	notificationBeforeUpsertHooks = []NotificationHook{}

	AddNotificationHook(boil.AfterUpsertHook, notificationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	notificationAfterUpsertHooks = []NotificationHook{}
}

func testNotificationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(notificationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNotificationToOneUserUsingUserTo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Notification
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, notificationDBTypes, false, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserToID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UserTo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NotificationSlice{&local}
	if err = local.L.LoadUserTo(ctx, tx, false, (*[]*Notification)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserTo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UserTo = nil
	if err = local.L.LoadUserTo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UserTo == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNotificationToOneSetOpUserUsingUserTo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Notification
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, notificationDBTypes, false, strmangle.SetComplement(notificationPrimaryKeyColumns, notificationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUserTo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UserTo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserToNotifications[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserToID != x.ID {
			t.Error("foreign key was wrong value", a.UserToID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserToID))
		reflect.Indirect(reflect.ValueOf(&a.UserToID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserToID != x.ID {
			t.Error("foreign key was wrong value", a.UserToID, x.ID)
		}
	}
}

func testNotificationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
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

func testNotificationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NotificationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNotificationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Notifications().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	notificationDBTypes = map[string]string{`ID`: `int`, `Title`: `varchar`, `Body`: `varchar`, `URL`: `varchar`, `UserToID`: `int`, `TimeRead`: `timestamp`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                   = bytes.MinRead
)

func testNotificationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(notificationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(notificationAllColumns) == len(notificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNotificationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(notificationAllColumns) == len(notificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Notification{}
	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, notificationDBTypes, true, notificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(notificationAllColumns, notificationPrimaryKeyColumns) {
		fields = notificationAllColumns
	} else {
		fields = strmangle.SetComplement(
			notificationAllColumns,
			notificationPrimaryKeyColumns,
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

	slice := NotificationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNotificationsUpsert(t *testing.T) {
	t.Parallel()

	if len(notificationAllColumns) == len(notificationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLNotificationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Notification{}
	if err = randomize.Struct(seed, &o, notificationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Notification: %s", err)
	}

	count, err := Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, notificationDBTypes, false, notificationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Notification struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Notification: %s", err)
	}

	count, err = Notifications().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

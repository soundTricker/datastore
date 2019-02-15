package testsuite

import (
	"context"
	"testing"

	"go.mercari.io/datastore"
)

// Issue52FooA has field that struct with pointer.
type Issue52FooA struct {
	Bar *Issue52Bar `datastore:",flatten"`
}

// Issue52FooB has field that struct without pointer.
type Issue52FooB struct {
	Bar Issue52Bar `datastore:",flatten"`
}

// Issue52Bar is a field of struct about Issue52Foo*.
type Issue52Bar struct {
	Name string
}

func issue52(ctx context.Context, t *testing.T, client datastore.Client) {
	defer func() {
		err := client.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	{
		obj := &Issue52FooA{
			Bar: &Issue52Bar{
				Name: "Issue 52",
			},
		}
		key := client.NameKey("Issue52FooA", "1", nil)
		_, err := client.Put(ctx, key, obj)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		obj := &Issue52FooB{
			Bar: Issue52Bar{
				Name: "Issue 52",
			},
		}
		key := client.NameKey("Issue52FooB", "1", nil)
		_, err := client.Put(ctx, key, obj)
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		obj := &Issue52FooA{}
		key := client.NameKey("Issue52FooA", "1", nil)
		err := client.Get(ctx, key, obj)
		if err != nil {
			t.Fatal(err)
		}
		if obj.Bar == nil {
			t.Error("Issue52FooA.Bar is nil")
		} else if v := obj.Bar.Name; v != "Issue 52" {
			t.Errorf("unexpected: %v", v)
		}
	}
	{
		obj := &Issue52FooB{}
		key := client.NameKey("Issue52FooB", "1", nil)
		err := client.Get(ctx, key, obj)
		if err != nil {
			t.Fatal(err)
		}
		if v := obj.Bar.Name; v != "Issue 52" {
			t.Errorf("unexpected: %v", v)
		}
	}
}

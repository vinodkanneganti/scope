package report_test

import (
	"reflect"
	"testing"

	"github.com/weaveworks/scope/report"
	"github.com/weaveworks/scope/test"
)

func TestTables(t *testing.T) {
	want := map[string]string{
		"foo1": "bar1",
		"foo2": "bar2",
	}
	nmd := report.MakeNode("foo1")

	nmd = nmd.AddTable("foo_", want)
	have, truncated := nmd.ExtractTable("foo_")

	if truncated {
		t.Error("Table shouldn't had been truncated")
	}

	if !reflect.DeepEqual(want, have) {
		t.Error(test.Diff(want, have))
	}
}

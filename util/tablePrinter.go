package util

import (
	"github.com/jedib0t/go-pretty/table"
	v2 "github.com/kgysu/oc-wrapper/v2"
	"os"
)

func NewStandardTable(items []v2.OpenshiftItem) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Kind", "Name"})
	for i, item := range items {
		t.AppendRow(table.Row{i, item.GetKind(), item.GetName()})
	}
	t.AppendFooter(table.Row{"=", "Total", len(items)})
	return t
}

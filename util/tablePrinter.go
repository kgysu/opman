package util

import (
	"github.com/jedib0t/go-pretty/table"
	v2 "github.com/kgysu/oc-wrapper/v2"
	"os"
)

func NewStandardTable(items []v2.OpenshiftItem) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Kind", "Name", "Status"})
	for i, item := range items {
		t.AppendRow(table.Row{i, item.GetKind(), item.GetName(), item.GetStatusMessage()})
	}
	t.AppendFooter(table.Row{"=", "Total", "", len(items)})
	return t
}

func NewMonitorItemsTable(items []MonitorItem) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Status", "Kind", "Name"})
	for i, item := range items {
		t.AppendRow(table.Row{i, item.status, item.item.GetKind(), item.item.GetName()})
	}
	return t
}

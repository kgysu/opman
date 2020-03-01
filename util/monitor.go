package util

import (
	"fmt"
	v2 "github.com/kgysu/oc-wrapper/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/exec"
	"time"
)

type MonitorItem struct {
	status string
	item   v2.OpenshiftItem
}

func MonitorRemoteItems(name string, namespace string, types string, labelSelector string, limit int64) error {
	fmt.Println("Monitor: ")
	return monitor(name, namespace, parseTypes(types), labelSelector, limit)
}

func monitor(name string, namespace string, types string, labelSelector string, limit int64) error {
	var itemsBefore []v2.OpenshiftItem
	for i := 0; i < 240; i++ {
		// get project
		projectNow, err := v2.NewFromRemote(name, namespace, parseTypes(types), v1.ListOptions{
			LabelSelector: labelSelector,
			Limit:         limit,
		})
		if err != nil {
			return err
		}
		// find diff and create table
		monitorItems := differenceBeforeNowItems(itemsBefore, projectNow.GetItems())
		t := NewMonitorItemsTable(monitorItems)
		// clear
		err = clear()
		if err != nil {
			return err
		}
		// print
		t.Render()
		// swap before=now
		itemsBefore = projectNow.GetItems()
		// sleep
		time.Sleep(5 * time.Second)
	}
	fmt.Printf("timeout")
	return nil
}

func differenceBeforeNowItems(itemsBefore []v2.OpenshiftItem, itemsNow []v2.OpenshiftItem) []MonitorItem {
	var diff []MonitorItem
	// items removed
	for _, itemBef := range itemsBefore {
		status := ""
		if !isInList(itemBef, itemsNow) {
			status = "-"
		}
		diff = append(diff, MonitorItem{
			status: status,
			item:   itemBef,
		})
	}
	// items added
	for _, itemNow := range itemsNow {
		if !isInList(itemNow, itemsBefore) {
			diff = append(diff, MonitorItem{
				status: "+",
				item:   itemNow,
			})
		}
	}
	return diff
}

func isInList(item v2.OpenshiftItem, list []v2.OpenshiftItem) bool {
	for _, listItem := range list {
		if listItem.GetName() == item.GetName() && listItem.GetKind() == item.GetKind() {
			return true
		}
	}
	return false
}

func clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

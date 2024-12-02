package db

import (
	"service-sentinel/monitors"
	"testing"
)


func TestInitDB(t *testing.T) () {
	err := Init("test_db")
	if err != nil {
		t.Error("connecting to db caused err", err)
	}
}
func TestGetAllmonitorsEmpty(t *testing.T) () {
	monitors, err := GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}
	if len(monitors) != 0 {
		t.Error("Amount of monitors should be 0 but it's ", len(monitors))
	}
}

func TestAddHttpMonitor(t *testing.T) () {
	err := InsertMonitor(monitors.HttpMonitor{})
	if err != nil {
		t.Error("error in inserting monitor", err)
	}

	monitors, err := GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}
	if len(monitors) != 1 {
		t.Error("Amount of monitors should be 1 but it's ", len(monitors))
	}
}

func TestAddPingMonitor(t *testing.T) () {
	err := InsertMonitor(monitors.PingMonitor{})
	if err != nil {
		t.Error("error in inserting monitor", err)
	}

	monitors, err := GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}
	if len(monitors) != 2 {
		t.Error("Amount of monitors should be 2 but it's ", len(monitors))
	}
}

func TestGetMonitorByKey (t *testing.T) () {
	id := uint(1)
	httpMonitor, httpErr := GetMonitorByKey(id, monitors.HttpMonitor{})
	if httpErr != nil {
		t.Error("error in getting monitor with id ", id, httpErr)
	}
	if httpMonitor.BaseInfo.Model.ID != id {
		t.Errorf("ID should be %d but it's %d", id, httpMonitor.BaseInfo.Model.ID)
	}

	httpMonitor, httpErr = GetMonitorByKey(id + 1, monitors.HttpMonitor{})
	if httpErr == nil {
		t.Error("found monitor with key that shouldn't be in db ", id, httpErr)
	}
	if httpMonitor.BaseInfo.Model.ID != 0 {
		t.Errorf("ID should be %d but it's %d", 0, httpMonitor.BaseInfo.Model.ID)
	}

	pingMonitor, pingErr := GetMonitorByKey(id, monitors.PingMonitor{})
	if pingErr != nil {
		t.Error("error in getting monitor with id ", id, pingErr)
	}
	if pingMonitor.BaseInfo.Model.ID != id {
		t.Errorf("ID should be %d but it's %d", id, pingMonitor.BaseInfo.Model.ID)
	}
	
}

func TestDeleteByKey (t *testing.T) () {
	DeleteMonitorByKey(1, monitors.PingMonitor{})
	monitorsList, err := GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}

	if len(monitorsList) != 1 {
		t.Error("Amount of monitors should be 1 but it's ", len(monitorsList))
	}

	DeleteMonitorByKey(2, monitors.HttpMonitor{})
	monitorsList, err = GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}

	if len(monitorsList) != 1 {
		t.Error("Amount of monitors should be 1 but it's ", len(monitorsList))
	}

	DeleteMonitorByKey(1, monitors.HttpMonitor{})
	monitorsList, err = GetAllMonitors()
	if err != nil {
		t.Error("error in getting all monitors", err)
	}

	if len(monitorsList) != 0 {
		t.Error("Amount of monitors should be 0 but it's ", len(monitorsList))
	}

}

func TestEmptyDB (t *testing.T) () {
	err := gormDB.Migrator().DropTable(&monitors.PingMonitor{})
	if err != nil {
		t.Error("error in dropping table ping monitors", err)
	}
	err = gormDB.Migrator().DropTable(&monitors.HttpMonitor{})
	if err != nil {
		t.Error("error in dropping table ping monitors", err)
	}
}
package db

import (
	"errors"
	"service-sentinel/monitors"
)

func InsertMonitor (Monitor monitors.ServiceMonitor) (error) {
    switch s := Monitor.(type) {
    case monitors.HttpMonitor:
        if err := gormDB.Create(&s).Error; err != nil {
            return err
        }
    case monitors.PingMonitor:
        if err := gormDB.Create(&s).Error; err != nil {
            return err
        }
    default:
        return errors.New("unsupported Monitor type")
    }
    return nil
}

func GetMonitorByKey[T monitors.ServiceMonitor] (ID uint, searched T) (T, error) {
    if err := gormDB.First(&searched, ID).Error; err != nil {
        return searched, err
    }
    return searched, nil
}

func DeleteMonitorByKey[T monitors.ServiceMonitor] (ID uint, _ T) () {
    var item T
    gormDB.Delete(&item, ID)
}


func GetHttpMonitors () ([]monitors.HttpMonitor, error) {
	return GetGenericMonitors(monitors.HttpMonitor{})
}


func GetPingMonitors () ([]monitors.PingMonitor, error) {
	return GetGenericMonitors(monitors.PingMonitor{})
}

func GetAllMonitors() ([]monitors.ServiceMonitor, error) {
    httpmonitors, err := GetHttpMonitors()
    if err != nil {
        return nil, err
    }

    pingmonitors, err := GetPingMonitors()
    if err != nil {
        return nil, err
    }

    var allmonitors []monitors.ServiceMonitor
    for _, monitor := range httpmonitors {
        allmonitors = append(allmonitors, monitor)
    }
    for _, monitor := range pingmonitors {
        allmonitors = append(allmonitors, monitor)
    }

    return allmonitors, nil
}

func GetGenericMonitors[T monitors.ServiceMonitor](T) ([]T, error) {
    var items []T
    if err := gormDB.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}
package db

import (
	"errors"
	"service-sentinel/samplers"
)

func InsertMonitor (sampler samplers.ServiceSampler) (error) {
    switch s := sampler.(type) {
    case samplers.HttpSampler:
        if err := gormDB.Create(s).Error; err != nil {
            return err
        }
    case samplers.PingSampler:
        if err := gormDB.Create(s).Error; err != nil {
            return err
        }
    default:
        return errors.New("unsupported sampler type")
    }
    return nil
}


func GetHttpMonitors () ([]samplers.HttpSampler, error) {
	return GetMonitors(samplers.HttpSampler{})
}


func GetPingMonitors () ([]samplers.PingSampler, error) {
	return GetMonitors(samplers.PingSampler{})
}

func GetAllMonitors() ([]samplers.ServiceSampler, error) {
    httpMonitors, err := GetHttpMonitors()
    if err != nil {
        return nil, err
    }

    pingMonitors, err := GetPingMonitors()
    if err != nil {
        return nil, err
    }

    var allMonitors []samplers.ServiceSampler
    for _, monitor := range httpMonitors {
        allMonitors = append(allMonitors, monitor)
    }
    for _, monitor := range pingMonitors {
        allMonitors = append(allMonitors, monitor)
    }

    return allMonitors, nil
}

func GetMonitors[T samplers.ServiceSampler](model T) ([]T, error) {
    var items []T
    if err := gormDB.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}
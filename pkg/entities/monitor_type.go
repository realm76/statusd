package entities

import (
	"errors"
	"strings"
)

type MonitorType int

var ErrInvalidMonitorType = errors.New("invalid monitor type")

const (
	HttpMonitorType MonitorType = iota + 1
	HttpsMonitorType
)

func (m *MonitorType) UnmarshalYAML(data func(interface{}) error) error {
	var name string
	if err := data(&name); err != nil {
		return err
	}

	mt, err := GetMonitorType(name)
	if err != nil {
		return err
	}

	*m = mt

	return nil
}

func GetMonitorType(name string) (MonitorType, error) {
	switch strings.ToLower(name) {
	case "http":
		return 1, nil
	case "https":
		return 2, nil

	}

	return 0, ErrInvalidMonitorType
}

func (m *MonitorType) String() string {
	return [...]string{"http", "https", ""}[*m]
}

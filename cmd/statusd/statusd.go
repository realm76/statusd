package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/realm76/statusd/pkg/entities"
	"log"
	"os"
	"path"
)

func main() {
	configPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	stsdFile, err := os.ReadFile(path.Join(configPath, "statusd.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	var stsd StatusD

	if err := yaml.Unmarshal(stsdFile, &stsd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", stsd)
}

type MonitorSlice []any

type StatusD struct {
	Monitors MonitorSlice
}

func (mrs *MonitorSlice) UnmarshalYAML(data func(interface{}) error) error {
	var rawMonitors []any
	var monitors MonitorSlice

	if err := data(&rawMonitors); err != nil {
		return err
	}

	for _, rawMonitor := range rawMonitors {
		monitor := rawMonitor.(map[string]interface{})

		rawType := monitor["type"].(string)

		monitorType, err := entities.GetMonitorType(rawType)
		if err != nil {
			return err
		}

		switch monitorType {
		case entities.HttpMonitorType, entities.HttpsMonitorType:
			url, urlExists := monitor["url"].(string)
			if !urlExists {
				log.Fatalln("url does not exist")
			}

			threshold, thresholdExists := monitor["threshold"].(float64)
			if !thresholdExists {
				threshold = 30
			}

			m := entities.MonitorHttp{
				Secure: monitorType == entities.HttpsMonitorType,
				Monitor: entities.Monitor{
					Name: monitor["name"].(string),
					Type: entities.HttpMonitorType,
				},
				Url:       url,
				Threshold: threshold,
			}

			monitors = append(monitors, m)
			break
		default:
			panic("unhandled default case")
		}
	}

	*mrs = monitors

	return nil
}

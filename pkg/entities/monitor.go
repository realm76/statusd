package entities

type Monitor struct {
	Name string
	Type MonitorType
}

type MonitorHttp struct {
	Monitor

	Secure    bool
	Url       string
	Threshold float64
}

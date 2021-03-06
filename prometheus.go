package grafaneus

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Prometheus struct {
}

type MetricsNames struct {
	Status string   `json:status`
	Data   []string `json:data`
}

type Metric struct {
	Name        string
	Type        string
	Description string
}

var MetricsMetadata = map[string]Metric{
	"go_goroutines": {Name: "go_goroutines", Type: "gauge", Description: "Number of goroutines that currently exist for GoLang process."},
	"go_threads":    {Name: "go_threads", Type: "gauge", Description: "Number of OS threads that currently exist for GoLang process."},
	"process_open_fds":    {Name: "process_open_fds", Type: "gauge", Description: "Number of open file descriptors for process."},
}

func (m *MetricsNames) Names() ([]string) {
	return m.Data
}

func ListMetrics() ([]Metric, error) {
	resp, err := http.Get("http://localhost:9090/api/v1/label/__name__/values")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var metricsNames MetricsNames
	err = json.Unmarshal(body, &metricsNames)
	if err != nil {
		return nil, err
	}
	result := make([]Metric, len(metricsNames.Data))
	for i, metric := range metricsNames.Data {
		if val, ok := MetricsMetadata[metric]; ok {
			result[i] = val
		} else {
			result[i] = Metric{Name: metric}
		}
	}
	return result, nil
}
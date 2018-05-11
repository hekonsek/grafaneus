package grafaneus

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Grafana struct {
}

type DataSource struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (*Grafana) ListDataSources() ([]DataSource, error) {
	resp, err := http.Get("http://admin:admin@localhost:3000/api/datasources")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var metricsNames []DataSource
	err = json.Unmarshal(body, &metricsNames)
	if err != nil {
		return nil, err
	}
	return metricsNames, nil
}

func (*Grafana) CreateDataSource() error {
	template := `{"id":1,"orgId":1,"name":"prometheus","type":"prometheus","typeLogoUrl":"public/app/plugins/datasource/prometheus/img/prometheus_logo.svg","access":"proxy","url":"http://localhost:9090","password":"","user":"","database":"","basicAuth":false,"isDefault":true,"jsonData":{"httpMethod":"GET","keepCookies":[]},"readOnly":false}`
	_, err := http.Post("http://admin:admin@localhost:3000/api/datasources", "application/json", strings.NewReader(template))
	if err != nil {
		return err
	}
	return nil
}

func (grafana *Grafana) ensureDatabaseExists() {
	dataSources, _ := grafana.ListDataSources()
	if len(dataSources) == 0 {
		grafana.CreateDataSource()
	}
}

func (*Grafana) GenerateGraph(title string, expression string) (string) {
	template := `{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": 1,
  "links": [],
  "panels": [
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "datasource": "prometheus",
      "fill": 1,
      "gridPos": {
        "h": 9,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 5,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "$$hashKey": "object:597",
          "expr": "EXPRESSION",
          "format": "time_series",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeShift": null,
      "title": "PANEL_TITLE",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    }
  ],
  "schemaVersion": 16,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-15m",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ],
    "time_options": [
      "5m",
      "15m",
      "1h",
      "6h",
      "12h",
      "24h",
      "2d",
      "7d",
      "30d"
    ]
  },
  "timezone": "",
  "title": "Prometheus",
  "uid": "VvCacnnik",
  "version": 2
}`
	templateWithExpression := strings.Replace(template, "EXPRESSION", expression, 1)
	return strings.Replace(templateWithExpression, "PANEL_TITLE", title, 1)
}

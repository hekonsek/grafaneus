package grafana

import (
	"encoding/json"
	"math"
)

type Dashboard struct {

	model map[string]interface{}

}

func (dashboard *Dashboard) importModel(modelJson string) error {
	err := json.Unmarshal([]byte(modelJson), &dashboard.model)
	if err != nil {
		return err
	}
	return nil
}

func (dashboard *Dashboard) exportModel() (string, error) {
	json, err := json.Marshal(dashboard.model)
	return string(json), err
}

func (dashboard *Dashboard) addPanel(panel map[string]interface{}) error {
	panels := dashboard.model["panels"].([]interface{})

	var max float64 = 0
	for _, panelObject := range panels {
		xxx := panelObject.(map[string]interface{})
		max = math.Max(max, xxx["id"].(float64))
	}
	panel["id"] = max + 1

	dashboard.model["panels"] = append(panels, panel)
	return nil
}

func (dashboard *Dashboard) addPanelJson(panelJson string) error {
	var panel map[string]interface{}
	err := json.Unmarshal([]byte(panelJson), &panel)
	if err != nil {
		return err
	}
	return dashboard.addPanel(panel)
}
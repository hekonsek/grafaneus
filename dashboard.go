package grafaneus

import "encoding/json"

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

func (dashboard *Dashboard) addPanel(panelJson string) error {
	panels := dashboard.model["panels"].([]interface{})
	var panel map[string]interface{}
	err := json.Unmarshal([]byte(panelJson), &panel)
	if err != nil {
		return err
	}
	dashboard.model["panels"] = append(panels, panel)
	return nil
}
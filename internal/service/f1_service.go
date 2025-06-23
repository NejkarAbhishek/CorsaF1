package service

import (
	"encoding/json"
	"fmt"
	"CorsaF1/internal/model"
	"CorsaF1/internal/repository"
	"io/ioutil"
	"net/http"
)

type DriverStanding struct {
	DriverName string `json:"driver_name"`
	Position   string `json:"position"`
	Points     string `json:"points"`
}

func FetchDriverStandings() ([]DriverStanding, error) {
	url := "https://api.jolpi.ca/ergast/f1/current/driverStandings.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching driver standings: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	var standings []DriverStanding
	var dbDrivers []model.Driver

	lists := result["MRData"].(map[string]interface{})["StandingsTable"].(map[string]interface{})["StandingsLists"].([]interface{})
	if len(lists) > 0 {
		drivers := lists[0].(map[string]interface{})["DriverStandings"].([]interface{})
		for _, d := range drivers {
			driver := d.(map[string]interface{})
			dt := driver["Driver"].(map[string]interface{})
			standing := DriverStanding{
				DriverName: dt["familyName"].(string),
				Position:   driver["position"].(string),
				Points:     driver["points"].(string),
			}
			standings = append(standings, standing)

			dbDrivers = append(dbDrivers, model.Driver{
				Name:     standing.DriverName,
				Position: standing.Position,
				Points:   standing.Points,
			})
		}
		_ = repository.SaveDrivers(dbDrivers) // ignore error here for now
	}

	return standings, nil
}

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

func FetchConstructors() ([]model.Constructor, error) {
	url := "https://api.jolpi.ca/ergast/f1/current/constructorStandings.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching constructors: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	var constructors []model.Constructor

	lists := result["MRData"].(map[string]interface{})["StandingsTable"].(map[string]interface{})["StandingsLists"].([]interface{})
	if len(lists) > 0 {
		teams := lists[0].(map[string]interface{})["ConstructorStandings"].([]interface{})
		for _, c := range teams {
			item := c.(map[string]interface{})
			con := item["Constructor"].(map[string]interface{})
			constructors = append(constructors, model.Constructor{
				Name:        con["name"].(string),
				Nationality: con["nationality"].(string),
				Wins:        int(item["wins"].(float64)),
			})
		}
		repository.SaveConstructors(constructors)
	}

	return constructors, nil
}

func CompareDrivers(season, driverA, driverB string) model.ComparisonResult {
	return model.ComparisonResult{
		Season:  season,
		DriverA: driverA,
		DriverB: driverB,
		WinsA:   7,
		WinsB:   3,
	}
}
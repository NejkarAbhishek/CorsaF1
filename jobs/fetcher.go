package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchStandings() ([]byte, error) {
	url := "http://ergast.com/api/f1/current/driverStandings.json"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

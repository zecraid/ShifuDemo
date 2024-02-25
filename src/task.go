package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const URL = "deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
const TimeInterval = 10

func main() {
	for {
		_, err := getData()
		if err != nil {
			fmt.Printf("Error getting measurement: %v\n", err)
			continue
		}

		time.Sleep(time.Second * time.Duration(TimeInterval))
	}
}

func getData() ([]float64, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	valuesStr := strings.Fields(string(body))
	fmt.Println(valuesStr)
	return nil, nil
}

// main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const measurementURL = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
const pollingInterval = 10 // 10 seconds, customize as needed

func main() {
	for {
		measurement, err := getMeasurement()
		if err != nil {
			fmt.Printf("Error getting measurement: %v\n", err)
			continue
		}

		average := calculateAverage(measurement)
		fmt.Printf("Average Measurement: %f\n", average)

		time.Sleep(time.Second * time.Duration(pollingInterval))
	}
}

func getMeasurement() ([]float64, error) {
	resp, err := http.Get(measurementURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Split the response into individual values
	valuesStr := strings.Fields(string(body))
	fmt.Println(valuesStr)
	values := make([]float64, len(valuesStr))

	// Convert string values to float64
	for i, v := range valuesStr {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}

	return values, nil
}

func calculateAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}

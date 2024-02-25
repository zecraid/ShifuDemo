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

const URL = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
const timeInterval = 10

func main() {
	for {
		measurement, err := getData()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		average := average(measurement)
		fmt.Println("Average= ", average)

		time.Sleep(time.Second * time.Duration(timeInterval))
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
	values := make([]float64, len(valuesStr))

	for i, v := range valuesStr {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}

	return values, nil
}

func average(values []float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}

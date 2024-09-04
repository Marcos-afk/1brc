package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	minimumTemperature float64
	maximumTemperature float64
	sumOfTemperatures  float64
	count              int
}

func main() {
	measurements, error := os.Open("measurements.txt")
	if error != nil {
		fmt.Println("Erro ao abrir o arquivo: ", error)
	}

	defer measurements.Close()


	measurementsMap := make(map[string]Measurement)


	scanner := bufio.NewScanner(measurements);

	for scanner.Scan(){
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		temperature, error :=  strconv.ParseFloat(rawData[semicolon + 1:], 64)


		if error != nil {
			fmt.Println("Erro ao converter valor para Float: ", error)
			return
		}


		measurement, ok := measurementsMap[location]
		if !ok {
				measurement = Measurement {
					minimumTemperature: temperature,
					maximumTemperature: temperature,
					sumOfTemperatures: temperature,
					count: 1,
				}

		} else{

			measurement.minimumTemperature = min(measurement.minimumTemperature, temperature)
			measurement.maximumTemperature = max(measurement.maximumTemperature, temperature)
			measurement.sumOfTemperatures += temperature
			measurement.count += 1
			
		}

		measurementsMap[location] = measurement
	}

	locations := make([]string, 0, len(measurementsMap))

	for name := range measurementsMap {
		locations = append(locations, name)
	}
	
	sort.Strings(locations)

	for _, name := range locations {
		measurement := measurementsMap[name]

		fmt.Printf("{")
		fmt.Printf("%s=%.1f/%.1f/%.1f", 
		name, measurement.minimumTemperature, (measurement.sumOfTemperatures / float64(measurement.count)), measurement.maximumTemperature)
		fmt.Printf("}")
		fmt.Printf(",")
	}
}
package main

import (
	"fmt"
	"sort"
	"time"
)

// Process the data and return the statistics
func processData(meterData map[string]map[string]map[string][]meterReading) UtilityDataStatistics {
	var utilityStatistics UtilityDataStatistics
	utilityStatistics.lowestDailykWh = float64(100000)
	utilityStatistics.lowestMonthlykWh = float64(100000)
	utilityStatistics.lowestYearlykWh = float64(100000)
	utilityStatistics.highestDailykWh = float64(0)
	utilityStatistics.highestMonthlykWh = float64(0)
	utilityStatistics.highestYearlykWh = float64(0)

	utilityStatistics.hourlykWh = make(map[string]float64)
	utilityStatistics.dailykWh = make(map[string]float64)
	utilityStatistics.lowestHour = make(map[string]float64)
	utilityStatistics.highestHour = make(map[string]float64)

	// Loop for all the data we have
	for year, yearData := range meterData {
		yearlykWh := float64(0)
		fmt.Println("Processing Year:", year)
		for month, monthData := range yearData {
			monthlykWh := float64(0)
			fmt.Println("Processing Month:", month)
			for day, daydata := range monthData {
				fmt.Println("Processing Day:", day)
				dailykWh := float64(0)
				utilityStatistics.totalDays++
				for _, hourData := range daydata {

					// Check to see if the map exists, if not make it
					if _, ok := utilityStatistics.lowestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")]; !ok {
						utilityStatistics.lowestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] = 100
					}

					// Add the kWh usage to the totals
					utilityStatistics.totalkWh += hourData.value
					dailykWh += hourData.value

					utilityStatistics.dailykWh[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] += hourData.value

					// Track the totals by hour as well
					utilityStatistics.hourlykWh[time.Unix(hourData.dateTime, 0).Format("15:04:05 PM")] += hourData.value
					if hourData.value > utilityStatistics.highestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] {
						utilityStatistics.highestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] = hourData.value
					}

					if hourData.value < utilityStatistics.lowestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] && hourData.value > 0 {
						utilityStatistics.lowestHour[time.Unix(hourData.dateTime, 0).Format("2006/01/02")] = hourData.value
					}

					// Increase the data point total
					utilityStatistics.totalDataPoints++
				}
				if dailykWh < utilityStatistics.lowestDailykWh && dailykWh > 0 {
					utilityStatistics.lowestDailykWh = dailykWh
				}
				if dailykWh > utilityStatistics.highestDailykWh {
					utilityStatistics.highestDailykWh = dailykWh
				}
				monthlykWh += dailykWh
			}
			if monthlykWh < utilityStatistics.lowestMonthlykWh {
				utilityStatistics.lowestMonthlykWh = monthlykWh
			}
			if monthlykWh > utilityStatistics.highestMonthlykWh {
				utilityStatistics.highestMonthlykWh = monthlykWh
			}
			yearlykWh += monthlykWh
		}
		if yearlykWh < utilityStatistics.lowestYearlykWh {
			utilityStatistics.lowestYearlykWh = yearlykWh
		}
		if yearlykWh > utilityStatistics.highestYearlykWh {
			utilityStatistics.highestYearlykWh = yearlykWh
		}
	}

	return utilityStatistics
}

// Sort the keys for printing in order later
func sortKeys(m map[string]float64) []string {
	var returnKeys []string

	for x := range m {
		returnKeys = append(returnKeys, x)
	}

	sort.Strings(returnKeys)

	return returnKeys
}

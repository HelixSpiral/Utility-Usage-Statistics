package main

// UtilityDataStatistics holds statistics about the data we've processed
type UtilityDataStatistics struct {
	totalkWh float64 // Total kWh used

	lowestDailykWh   float64 // Lowest daily kWh found
	lowestMonthlykWh float64 // Lowest monthly kWh found
	lowestYearlykWh  float64 // Lowest yearly kWh found

	highestDailykWh   float64 // Highest daily kWh found
	highestMonthlykWh float64 // Highest monthly kWh found
	highestYearlykWh  float64 // Highest yearly kWh found

	totalDays       float64 // Total days tracked
	totalDataPoints float64 // Total data points used

	dailykWh  map[string]float64 // Total kWh for that day
	hourlykWh map[string]float64 // Hourly data

	lowestHour  map[string]float64 // Lowest kWh at each hour
	highestHour map[string]float64 // Highest kWh at each hour
}

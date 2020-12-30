package crops

import "time"

type productionFunction struct {
	harvestCost                           float64
	cumulativeMonthlyProductionCostsEarly []float64
	productionCostLessHarvest             float64 //sum monthly or find max of cumulative...
	lossFromLatePlanting                  float64
}

func NewProductionFunction(mc []float64, cs CropSchedule, hc float64, latePlantingLoss float64) productionFunction {
	pf := productionFunction{
		harvestCost:          hc,
		lossFromLatePlanting: latePlantingLoss,
	}
	cmc, pclh := cumulateMonthlyCosts(mc, cs.StartPlantingDate, cs.DaysToMaturity)
	pf.cumulativeMonthlyProductionCostsEarly = cmc
	pf.productionCostLessHarvest = pclh
	return pf
}
func isLeapYear(year int) bool {
	leapFlag := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapFlag = true
			} else {
				leapFlag = false
			}
		} else {
			leapFlag = true
		}
	} else {
		leapFlag = false
	}
	return leapFlag
}
func (p productionFunction) GetCumulativeMonthlyProductionCostsEarly() []float64 {
	return p.cumulativeMonthlyProductionCostsEarly
}
func cumulateMonthlyCosts(mc []float64, start time.Time, daysToMaturity int) ([]float64, float64) {
	//this process assumes days to maturity is less than 1 year.
	totalCosts := 0.0
	cmc := make([]float64, 12)
	daysInYear := 365
	if isLeapYear(start.Year()) {
		daysInYear += 1
	}
	if daysToMaturity > daysInYear {
		panic("abort! abort! we hit an artery!")
	}
	startMonth := start.Month() //iota "enum"
	startMonthIndex := int(startMonth) - 1
	counter := 0
	updated := false
	year := start.Year()
	for ok := true; ok; ok = daysToMaturity > 0 {
		//compute days in the current month https://yourbasic.org/golang/last-day-month-date/
		t := time.Date(year, time.Month(startMonthIndex+counter+1), 0, 0, 0, 0, 0, time.UTC)
		daysInMonth := t.Day() //subtract the days in the current month from days to maturity.
		if counter == 0 {
			if !updated {
				daysInMonth -= start.Day() //remove the start of the month where crops weren't planted
			}
		}
		daysToMaturity -= daysInMonth
		if startMonthIndex+counter > 11 {
			if !updated {
				startMonthIndex = 0
				counter = 0
				year++
				updated = true
			}
		}
		totalCosts += mc[startMonthIndex+counter]
		cmc[startMonthIndex+counter] = totalCosts
		counter++
	}
	return cmc, totalCosts
}
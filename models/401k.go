package models

import (
	"math"
)

type Calculator401kSimple struct {
	Income                 float64 `json:"income"`
	Bonus                  float64 `json:"bonus"`
	Percent                float64 `json:"percent"`
	MatchPercent           float64 `json:"match_percent,omitempty"`
	MaxYearlyContributions float64 `json:"max_yearly_contributions,omitempty"`
	YearlyTotal            float64 `json:"yearly_total,omitempty"`
	EmployeeTotal          float64 `json:"employee_total,omitempty"`
	EmployerTotal          float64 `json:"employer_total,omitempty"`
}

func (c *Calculator401kSimple) CalculatePercentToHitMaxContribution() {
	//	we want to find the percent you would need to contribute to your 401 to hit the max yearly contribution
	//	we know the max yearly contribution is 23,000
	c.Percent = math.Round(c.MaxYearlyContributions/(c.Income*(1+c.Bonus))*100) / 100
	c.EmployerTotal = (c.Income * (1 + c.Bonus) * c.MatchPercent)
	c.EmployeeTotal = (c.Income * (1 + c.Bonus) * c.Percent)
	return
}

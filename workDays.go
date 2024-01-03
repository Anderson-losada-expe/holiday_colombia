package holiday_colombia

import "time"

// CalculateWorkDay calculates the number of workdays in a month,

func (hu *HolidayUtil) CalculateWorkDay(month time.Month, weekendDays []time.Weekday) int {
	workdays := 0

	// Create a Holiday Util instance for the current year and the specified month
	huCurrentMonth := NewHolidayUtil(time.Now().Year())
	huCurrentMonth.setFixedHolidays() // Ajustar feriados fijos

	// Set days of the week as weekend
	weekendMap := make(map[time.Weekday]bool)
	for _, day := range weekendDays {
		weekendMap[day] = true
	}

	// Get the last day of the month
	lastDay := time.Date(time.Now().Year(), month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Iterate over the days of the month
	for day := 1; day <= lastDay; day++ {
		date := time.Date(time.Now().Year(), month, day, 0, 0, 0, 0, time.UTC)

		// Check if the day is a work day (not a holiday and not a weekend)
		if !huCurrentMonth.IsHoliday(int(month), day) && !weekendMap[date.Weekday()] {
			workdays++
		}
	}

	return workdays
}

// GetWorkDaysList returns a list of workdays in a month, considering holidays and custom weekend days.
func (hu *HolidayUtil) GetWorkDaysList(month time.Month, weekendDays []time.Weekday) []time.Time {
	var workdaysList []time.Time

	// Create a Holiday Util instance for the current year and the specified month
	huCurrentMonth := NewHolidayUtil(time.Now().Year())
	huCurrentMonth.setFixedHolidays() // Ajustar feriados fijos

	// Set days of the week as weekend
	weekendMap := make(map[time.Weekday]bool)
	for _, day := range weekendDays {
		weekendMap[day] = true
	}

	// Get the last day of the month
	lastDay := time.Date(time.Now().Year(), month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Iterate over the days of the month
	for day := 1; day <= lastDay; day++ {
		date := time.Date(time.Now().Year(), month, day, 0, 0, 0, 0, time.UTC)

		// Check if the day is a work day (not a holiday and not a weekend)
		if !huCurrentMonth.IsHoliday(int(month), day) && !weekendMap[date.Weekday()] {
			workdaysList = append(workdaysList, date)
		}
	}

	return workdaysList
}

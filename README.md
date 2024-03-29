# holiday_colombia

[![Build Status](https://travis-ci.org/Anderson-losada-expe/holiday_colombia.svg?branch=main)](https://travis-ci.org/Anderson-losada-expe/holiday_colombia)
[![Go Report Card](https://goreportcard.com/badge/github.com/Anderson-losada-expe/holiday_colombia)](https://goreportcard.com/report/github.com/Anderson-losada-expe/holiday_colombia)
[![GoDoc](https://godoc.org/github.com/Anderson-losada-expe/holiday_colombia?status.svg)](https://godoc.org/github.com/Anderson-losada-expe/holiday_colombia)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

`holiday_colombia` It is a library to manage holidays in Colombia. Provides functions to check if a day is a holiday,
get the next business day from a given date, and count the number of business days between two dates.

## Facility

```go
go get -u github.com/Anderson-losada-expe/holiday_colombia 
```

## Use

- Here's a quick example of how to use the library:

```go
package main

import (
	"fmt"
	"time"

	"github.com/Anderson-losada-expe/holiday_colombia"
)

func main() {
	// Crear una instancia de HolidayUtil para el año actual
	hu := holiday_colombia.NewHolidayUtil(time.Now().Year())

	// Verificar si una fecha específica es festiva
	isHoliday := hu.IsHoliday(1, 1) // Año Nuevo
	fmt.Printf("¿El 1 de enero es festivo? %v\n", isHoliday)

	// Obtener el próximo día hábil desde una fecha dada
	startDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	nextBusinessDay := hu.GetNextBusinessDay(startDate, 1)
	fmt.Printf("Próximo día hábil después de %v: %v\n", startDate, nextBusinessDay)

	// Contar la cantidad de días hábiles entre dos fechas
	dateInit := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC)
	businessDays := hu.CountBusinessDays(dateInit, dateEnd)
	fmt.Printf("Días hábiles entre %v y %v: %d\n", dateInit, dateEnd, businessDays)

	// Calcular el número de días laborables en un mes con días de fin de semana específicos
	month := time.April
	weekendDays := []time.Weekday{time.Saturday, time.Sunday}
	workdays := hu.CalculateWorkDay(month, weekendDays)
	fmt.Printf("Días laborables en %s: %d\n", month, workdays)

	// Obtener la lista de días laborables en un mes con días de fin de semana específicos
	workdaysList := hu.GetWorkDaysList(month, weekendDays)
	fmt.Printf("Lista de días laborables en %s: %v\n", month, workdaysList)
}
```

## Functions

#### `NewHolidayUtil(year int) *holiday_colombia.HolidayUtil`

- Creates a new HolidayUtil instance for the given year.

#### `IsHoliday(month, day int) bool`

- Indicates if a day is a holiday.

#### `GetNextBusinessDay(date time.Time, days int) time.Time`

- Returns the next business day from a given date and number of days.

#### `CountBusinessDays(dateInit, dateEnd time.Time) int`

- Counts the number of business days between two dates.

#### `CountBusinessDays(dateInit, dateEnd time.Time) int`

- Counts the number of business days between two dates.

#### `CalculateWorkDay(month time.Month, weekendDays []time.Weekday) int`

- Calculates the number of workdays in a month with specified weekend days.

#### `GetWorkDaysList(month time.Month, weekendDays []time.Weekday) []time.Time`

- Returns the list of workdays in a month with specified weekend days.

## License

This project is under the MIT license.
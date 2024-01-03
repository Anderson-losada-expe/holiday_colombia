package holiday_colombia

import (
	"fmt"
	"time"
)

// HolidayUtil representa la utilidad para gestionar días festivos.
type HolidayUtil struct {
	year        int
	easterMonth int
	easterDay   int
	holidays    map[string]bool
}

// NewHolidayUtil crea una nueva instancia de HolidayUtil para el año dado.
func NewHolidayUtil(year int) *HolidayUtil {
	hu := &HolidayUtil{
		year:     year,
		holidays: make(map[string]bool),
	}

	hu.calculateEaster()
	hu.setFixedHolidays()
	hu.adjustIndependenciaCartagena()
	hu.adjustDiaDeLaRaza()

	hu.calculateEmiliani(0, 6)          // Reyes magos 6 de enero
	hu.calculateEmiliani(2, 19)         // San Jose 19 de marzo
	hu.calculateEmiliani(5, 29)         // San Pedro y San Pablo 29 de junio
	hu.calculateEmiliani(7, 15)         // Asuncion 15 de agosto
	hu.calculateEmiliani(9, 12)         // Descubrimiento de America 12 de octubre
	hu.calculateEmiliani(10, 1)         // Todos los Santos 1 de noviembre
	hu.calculateEmiliani(11, 11)        // Independencia de Cartagena 11 de noviembre
	hu.calculateOtherHoliday(-3, false) // Jueves Santos
	hu.calculateOtherHoliday(-2, false) // Viernes Santo
	hu.calculateOtherHoliday(40, true)  // Ascension del Senor de Pascua
	hu.calculateOtherHoliday(60, true)  // Corpus Christi
	hu.calculateOtherHoliday(68, true)  // Sagrado Corazon

	return hu
}

func (hu *HolidayUtil) calculateEaster() {
	a := hu.year % 19
	b := hu.year / 100
	c := hu.year % 100
	d := b / 4
	e := b % 4
	g := (8*b + 13) / 25
	h := (19*a + b - d - g + 15) % 30
	j := c / 4
	k := c % 4
	m := (a + 11*h) / 319
	r := (2*e + 2*j - k - h + m + 32) % 7
	hu.easterMonth = (h - m + r + 90) / 25
	hu.easterDay = (h - m + r + hu.easterMonth + 19) % 32
	hu.easterMonth--

	hu.holidays["0:1"] = true   // Primero de Enero
	hu.holidays["5:1"] = true   // Día del trabajo 1 de mayo
	hu.holidays["7:20"] = true  // Independencia 20 de Julio
	hu.holidays["8:7"] = true   // Batalla de Boyacá 7 de agosto
	hu.holidays["12:8"] = true  // Maria Inmaculada 8 de diciembre
	hu.holidays["12:25"] = true // Navidad 25 de diciembre
	hu.holidays["11:2"] = true  // Dia de todos los santos
}

func (hu *HolidayUtil) setFixedHolidays() {
	holidays := map[string]bool{
		"0:1":   true, // Primero de Enero
		"5:1":   true, // Día del trabajo 1 de mayo
		"7:20":  true, // Independencia 20 de Julio
		"8:7":   true, // Batalla de Boyacá 7 de agosto
		"12:8":  true, // Maria Inmaculada 8 de diciembre
		"12:25": true, // Navidad 25 de diciembre
	}

	for key, value := range holidays {
		hu.holidays[key] = value
	}
}

func (hu *HolidayUtil) adjustIndependenciaCartagena() {
	// Ajusta el día de la Independencia de Cartagena al primer lunes después del 11 de noviembre
	independenciaCartagena := time.Date(hu.year, time.November, 11, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(independenciaCartagena.Weekday())

	// Verifica si el 11 de noviembre ya es un lunes
	if dayOfWeek != 1 {
		// Encuentra el próximo lunes después del 11 de noviembre
		daysToAdd := (8 - dayOfWeek) % 7
		independenciaCartagena = independenciaCartagena.Add(time.Duration(daysToAdd) * 24 * time.Hour)
	}

	hu.holidays[independenciaCartagena.Format("1:2")] = true
}

func (hu *HolidayUtil) adjustDiaDeLaRaza() {
	// Agrega el día de la Raza al primer lunes después del 12 de octubre
	diaDeLaRaza := time.Date(hu.year, time.October, 12, 0, 0, 0, 0, time.UTC)
	dayOfWeekRaza := int(diaDeLaRaza.Weekday())

	// Verifica si el 12 de octubre ya es un lunes
	if dayOfWeekRaza != 1 {
		// Encuentra el próximo lunes después del 12 de octubre
		daysToAddRaza := (8 - dayOfWeekRaza) % 7
		diaDeLaRaza = diaDeLaRaza.Add(time.Duration(daysToAddRaza) * 24 * time.Hour)
	}

	hu.holidays[diaDeLaRaza.Format("1:2")] = true
}

// calculateEmiliani mueve un día festivo a lunes según la ley de Emiliani.
func (hu *HolidayUtil) calculateEmiliani(month, day int) {
	date := time.Date(hu.year, time.Month(month+1), day, 0, 0, 0, 0, time.UTC)
	dayOfWeek := int(date.Weekday())
	switch dayOfWeek {
	case 1:
		// Si es lunes, no es necesario moverlo.
	case 2:
		date = date.Add(6 * 24 * time.Hour)
	case 3:
		date = date.Add(5 * 24 * time.Hour)
	case 4:
		date = date.Add(4 * 24 * time.Hour)
	case 5:
		date = date.Add(3 * 24 * time.Hour)
	case 6:
		date = date.Add(2 * 24 * time.Hour)
	case 7:
		date = date.Add(24 * time.Hour)
	}
	hu.holidays[date.Format("1:2")] = true
}

// calculateOtherHoliday calcula los días festivos según el día de Pascua.
func (hu *HolidayUtil) calculateOtherHoliday(days int, emiliani bool) {
	date := time.Date(hu.year, time.Month(hu.easterMonth+1), hu.easterDay, 0, 0, 0, 0, time.UTC)
	date = date.Add(time.Duration(days) * 24 * time.Hour)
	if emiliani {
		hu.calculateEmiliani(int(date.Month())-1, date.Day())
	} else {
		hu.holidays[date.Format("1:2")] = true
	}
}

// IsHoliday indica si un día es festivo.
func (hu *HolidayUtil) IsHoliday(month, day int) bool {
	return hu.holidays[fmt.Sprintf("%d:%d", month, day)]
}

// GetYear devuelve el año.
func (hu *HolidayUtil) GetYear() int {
	return hu.year
}

// GetNextBusinessDay devuelve el próximo día hábil desde una fecha dada y una cantidad de días.
func (hu *HolidayUtil) GetNextBusinessDay(date time.Time, days int) time.Time {
	for days > 0 {
		date = date.Add(24 * time.Hour)
		if date.Year() != hu.GetYear() {
			hu = NewHolidayUtil(date.Year())
		}
		dayOfWeek := int(date.Weekday())
		if dayOfWeek != 0 && dayOfWeek != 6 && !hu.IsHoliday(int(date.Month()), date.Day()) {
			days--
		}
	}
	return date
}

// CountBusinessDays cuenta la cantidad de días hábiles entre dos fechas.
func (hu *HolidayUtil) CountBusinessDays(dateInit, dateEnd time.Time) int {
	days := 0
	limitDay := dateEnd
	startDay := dateInit
	for startDay.Before(limitDay) {
		startDay = startDay.Add(24 * time.Hour)
		if startDay.Year() != hu.GetYear() {
			hu = NewHolidayUtil(startDay.Year())
		}
		dayOfWeek := int(startDay.Weekday())
		if dayOfWeek != 0 && dayOfWeek != 6 && !hu.IsHoliday(int(startDay.Month()), startDay.Day()) {
			days++
		}
	}
	return days
}

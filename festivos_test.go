package holiday_colombia

import (
	"fmt"
	"testing"
	"time"
)

func TestHolidayFunctions(t *testing.T) {
	// Casos de prueba para diferentes años
	years := []int{2024, 2025, 2026}

	for _, year := range years {
		t.Run(fmt.Sprintf("Year_%d", year), func(t *testing.T) {
			hu := NewHolidayUtil(year)

			// Verifica las fechas festivas
			runIsHolidayTest(t, hu, year)

			// Verifica GetNextBusinessDay
			runGetNextBusinessDayTest(t, hu, year)

			// Verifica CountBusinessDays
			runCountBusinessDaysTest(t, hu, year)

			// Agrega más tests según sea necesario
		})
	}
}

func getHolidaysForYear(year int) []struct{ month, day int } {
	switch year {
	case 2024:
		return []struct{ month, day int }{
			{0, 1},   // Año nuevo
			{1, 8},   // Reyes magos
			{3, 25},  // San Jose
			{3, 28},  // Dia Santo
			{3, 29},  // Dia Santo
			{5, 1},   // Día del trabajo
			{5, 13},  // Asencion
			{6, 3},   // Corpus Christi
			{6, 10},  // San Grado Corazon
			{7, 1},   // San Pedro y San Pablo
			{7, 20},  // Dia de la independencia
			{8, 7},   // Batalla de boyaca
			{8, 19},  // Asuncion de la virgen
			{10, 14}, // Dia de la raza
			{11, 4},  // Todos los Santos
			{11, 11}, // Todos los Santos
			{12, 25}, // Independencia cartagena
		}
	case 2025:
		return []struct{ month, day int }{
			{0, 1},   // Año nuevo
			{1, 6},   // Reyes magos
			{3, 24},  // San Jose
			{4, 17},  // Dia Santo
			{4, 18},  // Dia Santo
			{5, 1},   // Día del trabajo
			{6, 30},  // Corpus Christi
			{6, 2},   // Asencion
			{6, 23},  // San Grado Corazon
			{8, 7},   // Batalla de boyaca
			{8, 18},  // Asuncion de la virgen
			{10, 13}, // Dia de la raza
			{11, 17}, // Independencia cartagena
			{11, 3},  // Todos los Santos
			{12, 8},  // Maria inmaculada 8 de diciembre
			{12, 25}, // Navidad
		}
	case 2026:
		return []struct{ month, day int }{
			{0, 1},   // Año nuevo
			{1, 12},  // Reyes magos
			{3, 23},  // San Jose
			{4, 2},   // Dia Santo
			{4, 3},   // Dia Santo
			{5, 1},   // Día del trabajo
			{5, 18},  // Día Asencion
			{6, 8},   // Corpus Christi
			{6, 15},  // Asencion
			{6, 29},  // San Grado Corazon
			{7, 20},  // San Grado Corazon
			{8, 7},   // Batalla de boyaca
			{8, 17},  // Asuncion de la virgen
			{10, 12}, // Dia de la raza
			{11, 2},  // Independencia cartagena
			{11, 16}, // Todos los Santos
			{12, 8},  // Maria inmaculada 8 de diciembre
			{12, 25}, // Navidad
		}
	default:
		return nil
	}
}

func runIsHolidayTest(t *testing.T, hu *HolidayUtil, year int) {
	// Obtiene las fechas festivas para el año dado
	holidays := getHolidaysForYear(year)

	// Verifica que las fechas festivas estén marcadas correctamente
	for _, holiday := range holidays {
		if !hu.IsHoliday(holiday.month, holiday.day) {
			t.Errorf("[%d] Se esperaba que %d-%d sea festivo, pero no lo es", year, holiday.month+1, holiday.day)
		}
	}
}

func runGetNextBusinessDayTest(t *testing.T, hu *HolidayUtil, year int) {
	// Define una fecha de inicio y verifica el próximo día hábil
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedNextBusinessDay := time.Date(year, time.January, 2, 0, 0, 0, 0, time.UTC)

	nextBusinessDay := hu.GetNextBusinessDay(startDate, 1)

	if !nextBusinessDay.Equal(expectedNextBusinessDay) {
		t.Errorf("[%d] Se esperaba el próximo día hábil %v, pero obtuvo %v", year, expectedNextBusinessDay, nextBusinessDay)
	}
}

func runCountBusinessDaysTest(t *testing.T, hu *HolidayUtil, year int) {
	// Define dos fechas y verifica la cantidad de días hábiles entre ellas
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.January, 10, 0, 0, 0, 0, time.UTC)

	expectedBusinessDays := 6 // Días hábiles entre el 1 y el 10 de enero de {year}

	businessDays := hu.CountBusinessDays(startDate, endDate)

	if businessDays != expectedBusinessDays {
		t.Errorf("[%d] Se esperaba %d días hábiles, pero obtuvo %d", year, expectedBusinessDays, businessDays)
	}
}

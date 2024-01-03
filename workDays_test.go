package holiday_colombia

import (
	"reflect"
	"testing"
	"time"
)

func TestHolidayUtil_CalculateWorkDay(t *testing.T) {
	// Configuración de casos de prueba
	tests := []struct {
		name        string
		year        int
		month       time.Month
		weekendDays []time.Weekday
		want        int
	}{
		{
			name:        "January with no holidays",
			year:        2024,
			month:       time.January,
			weekendDays: []time.Weekday{time.Saturday, time.Sunday},
			want:        22,
		},
		{
			name:        "February with holidays",
			year:        2023,
			month:       time.February,
			weekendDays: []time.Weekday{time.Saturday, time.Sunday},
			want:        21,
		},
		{
			name:        "June with holidays and custom weekends",
			year:        2025,
			month:       time.June,
			weekendDays: []time.Weekday{time.Friday, time.Saturday},
			want:        19,
		},
		{
			name:        "February with holidays and custom weekends",
			year:        2024,
			month:       time.February,
			weekendDays: []time.Weekday{time.Friday, time.Saturday},
			want:        21,
		},
		{
			name:        "March with holidays and custom weekends",
			year:        2024,
			month:       time.March,
			weekendDays: []time.Weekday{time.Friday, time.Saturday},
			want:        19,
		},
		{
			name:        "March with holidays and custom weekends",
			year:        2024,
			month:       time.March,
			weekendDays: []time.Weekday{time.Sunday, time.Saturday},
			want:        18,
		},
	}

	// Ejecutar casos de prueba
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Crear una instancia de HolidayUtil para el año específico
			hu := NewHolidayUtil(tt.year)

			// Calcular días laborales según los parámetros del caso de prueba
			got := hu.CalculateWorkDay(tt.month, tt.weekendDays)

			// Verificar si el resultado coincide con el esperado
			if got != tt.want {
				t.Errorf("CalculateWorkDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHolidayUtil_GetWorkDaysList(t *testing.T) {
	type fields struct {
		year        int
		easterMonth int
		easterDay   int
		holidays    map[string]bool
	}
	type args struct {
		month       time.Month
		weekendDays []time.Weekday
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []time.Time
	}{
		{
			name: "Test with no holidays and weekends on Saturday and Sunday",
			fields: fields{
				year:        2024,
				easterMonth: 4,
				easterDay:   13,
				holidays:    map[string]bool{},
			},
			args: args{
				month:       time.April,
				weekendDays: []time.Weekday{time.Saturday, time.Sunday},
			},
			want: []time.Time{
				time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC),  // Wednesday
				time.Date(2024, time.April, 2, 0, 0, 0, 0, time.UTC),  // Thursday
				time.Date(2024, time.April, 3, 0, 0, 0, 0, time.UTC),  // Friday
				time.Date(2024, time.April, 4, 0, 0, 0, 0, time.UTC),  // Monday
				time.Date(2024, time.April, 5, 0, 0, 0, 0, time.UTC),  // Monday
				time.Date(2024, time.April, 8, 0, 0, 0, 0, time.UTC),  // Wednesday
				time.Date(2024, time.April, 9, 0, 0, 0, 0, time.UTC),  // Thursday
				time.Date(2024, time.April, 10, 0, 0, 0, 0, time.UTC), // Friday
				time.Date(2024, time.April, 11, 0, 0, 0, 0, time.UTC), // Friday
				time.Date(2024, time.April, 12, 0, 0, 0, 0, time.UTC), // Monday
				time.Date(2024, time.April, 15, 0, 0, 0, 0, time.UTC), // Wednesday
				time.Date(2024, time.April, 16, 0, 0, 0, 0, time.UTC), // Thursday
				time.Date(2024, time.April, 17, 0, 0, 0, 0, time.UTC), // Friday
				time.Date(2024, time.April, 18, 0, 0, 0, 0, time.UTC), // Monday
				time.Date(2024, time.April, 19, 0, 0, 0, 0, time.UTC), // Tuesday
				time.Date(2024, time.April, 22, 0, 0, 0, 0, time.UTC), // Wednesday
				time.Date(2024, time.April, 23, 0, 0, 0, 0, time.UTC), // Thursday
				time.Date(2024, time.April, 24, 0, 0, 0, 0, time.UTC), // Friday
				time.Date(2024, time.April, 25, 0, 0, 0, 0, time.UTC), // Monday
				time.Date(2024, time.April, 26, 0, 0, 0, 0, time.UTC), // Tuesday
				time.Date(2024, time.April, 29, 0, 0, 0, 0, time.UTC), // Wednesday
				time.Date(2024, time.April, 30, 0, 0, 0, 0, time.UTC), // Thursday
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hu := &HolidayUtil{
				year:        tt.fields.year,
				easterMonth: tt.fields.easterMonth,
				easterDay:   tt.fields.easterDay,
				holidays:    tt.fields.holidays,
			}
			if got := hu.GetWorkDaysList(tt.args.month, tt.args.weekendDays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkDaysList() = %v, want %v", got, tt.want)
			}
		})
	}
}

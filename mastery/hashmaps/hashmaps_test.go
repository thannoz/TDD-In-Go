package hashmaps

import (
	"reflect"
	"testing"
)

func TestGetCapitals(t *testing.T) {
	t.Parallel()

	// Tabelle tests
	testCases := []struct {
		country Country
		want    string
	}{
		{country: "Kongo DR", want: "Kinshasa"},
		{country: "Deutschland", want: "Berlin"},
		{country: "Angola", want: "Luanda"},
		{country: "Kenia", want: "Nairobi"},
		{country: "Senegal", want: "Dakar"},
		{country: "France", want: "Paris"},
		{country: "Sweden", want: "Stockholm"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.country), func(t *testing.T) {
			got := GetCapital(tc.country)
			if got != tc.want {
				t.Errorf("GetCapital(%q) = %q; want %q", tc.country, got, tc.want)
			}
		})
	}
}

func TestGetProductPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		product Product
		want    float64
	}{
		{product: "Banane", want: 2.34},
		{product: "Apfel", want: 5.34},
		{product: "Birne", want: 6.34},
		{product: "Kiwi", want: 3.64},
		{product: "Pfirsisch", want: 0},
	}

	for _, tc := range testCases {
		t.Run(string(tc.product), func(t *testing.T) {
			got := GetProductPrice(tc.product)
			if got != tc.want {
				t.Errorf("got %0.2f, want %0.2f", got, tc.want)
			}
		})
	}
}

func TestGetCapitalVersionTwo(t *testing.T) {
	t.Parallel()

	t.Run("This is an alternative version of the Test above", func(t *testing.T) {
		got := GetCapital("France")
		want := "Paris"
		if got != want {
			t.Errorf("Capital of %q, is not %q", got, want)
		}
	})
}

func TestDeleteFruitFromMap(t *testing.T) {
	t.Parallel()

	fruitToDelete := Fruit("Banane")
	DeleteFruitFromMap(fruitToDelete, &Fruits)

	_, ok := Fruits[fruitToDelete]
	if ok {
		t.Errorf("%v wurde nicht gelöscht.", fruitToDelete)
	}
}

func TestGetDaysOfTheMonth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		month Month
		want  int
	}{
		{month: Month("Januar"), want: 31},
		{month: Month("Februar"), want: 28},
		{month: Month("März"), want: 30},
		{month: Month("April"), want: 31},
		{month: Month("Mai"), want: 31},
		{month: Month("Juni"), want: 30},
		{month: Month("Juli"), want: 31},
		{month: Month("August"), want: 31},
		{month: Month("September"), want: 30},
		{month: Month("Oktober"), want: 31},
		{month: Month("November"), want: 30},
		{month: Month("Dezember"), want: 31},
	}

	for _, tc := range testCases {
		t.Run(string(tc.month), func(t *testing.T) {
			got := GetDaysOfTheMonth(tc.month)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAddAnimalLifeSpan(t *testing.T) {
	t.Parallel()

	tiere := map[Animal]string{
		Animal("Hund"):  "12-16 Jahre",
		Animal("Katze"): "12-30 Jahre",
		Animal("Maus"):  "1-6 Jahre",
	}

	testCases := []struct {
		animal   Animal
		lifeSpan string
	}{
		{animal: Animal("Hund"), lifeSpan: "12-16 Jahre"},
		{animal: Animal("Katze"), lifeSpan: "12-30 Jahre"},
		{animal: Animal("Rate"), lifeSpan: "12-30 Jahre"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.animal), func(t *testing.T) {
			got, err := AddAverageLifeSpanOfAnimal(tc.animal, tc.lifeSpan)
			if err != nil {
				t.Fatalf("Fehler erwartet: %v, aber kein Fehler aufgetreten", err)
			}

			for key := range got {
				if _, ok := tiere[key]; !ok {
					continue // continue if add key is not in tiere-map
				}
			}
		})
	}
}

func TestContainInAdressBook(t *testing.T) {
	t.Parallel()

	names := []string{
		"Carlos",
		"Carmelo",
		"Konzo",
		"Nzolani",
		"Melo",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			found := Contains(name)
			if !found {
				t.Logf("Benutzer '%s' nicht im Adressbuch gefunden", name)
			}
		})
	}
}

func TestNeighboringCountries(t *testing.T) {
	t.Parallel()

	// the definition of the struct-testcases depends on the function to be tested.
	// here we define our expectation in a struct
	// 1. we want to call the function with a country - the argument for the function
	// 2. we expect an array of neighboring countries
	// 3. we also specify if our expectations should fail or not by adding a boolean field
	testCases := []struct {
		country      string
		expected     []string
		expectingErr bool
	}{
		{
			"Deutschland",
			[]string{
				"Frankreich",
				"Belgien",
				"Schweiz",
				"Österreich",
				"Polen",
				"Tschechien",
				"Dänemark",
				"Niederlande",
			},
			false, // should not fail because "Deutschland" is in the map
		},
		{
			"Frankreich",
			[]string{
				"Deutschland",
				"Belgien",
				"Luxemburg",
				"Schweiz",
				"Italien",
				"Spanien",
				"Andorra",
				"Monaco",
			},
			false, // same as for "Deutschland"...
		},
		{
			"Unbekanntes Land",
			nil,
			true, // should fail because "Unbekanntes Land" is not in the map
		},
	}

	for _, tc := range testCases {
		t.Run(tc.country, func(t *testing.T) {
			neibours, err := PrintNeighbouringCountries(tc.country)
			if tc.expectingErr {
				if err == nil {
					t.Fatalf("expected an error for country %s, but got none", tc.country)
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error for country %s, but got %v", tc.country, err)
				}
				if len(neibours) != len(tc.expected) {
					t.Fatalf("for country %s, expected %d, neighbours, but got %d", tc.country, len(tc.expected), len(neibours))
				}
				// check if all expected neighbours are present
				expectedMap := make(map[string]bool)
				for _, n := range tc.expected {
					expectedMap[n] = true
				}
				for _, n := range neibours {
					if !expectedMap[n] {
						t.Fatalf("for country %s, unexpected neighbour found: %s", tc.country, n)
					}
				}
			}
		})
	}
}

func TestPrintAuthorsNames(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		bookName    Book
		expected    string
		expectedErr bool
	}{
		{"1984", "George Orwell", false},
		{"Schwarze Haut, weisse Masken", "Frantz Fanon", false},
		{"Unbekanntes Buch", "", true},
	}

	for _, tc := range testCases {
		author, err := PrintAuthorsName(tc.bookName)
		if tc.expectedErr { // should fail because no such book, hence true in expected struct
			if err == nil { // because theres no such book, there will be no author.
				// Hence we expect an error
				t.Errorf("expected an error for book %s, but got none", tc.bookName)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for book %s, but got %v", tc.bookName, err)
			}
			if author != tc.expected {
				t.Errorf("for book %s, expected author %s, but got %s", tc.bookName, tc.expected, author)
			}
		}
	}
}

func TestPrintStudentCourses(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		courseName  string
		studentList []string
		expectedErr bool
	}{
		{"Mathematik", []string{"Carlos", "Konzo", "Mwinda"}, false},
		{"Physik", []string{"Kalonji", "Selamawit"}, false},
		{"Informatik", []string{"Nzolani", "Carmelo", "Kyrie"}, false},
		{"Kunst", nil, true}, // should fail because no such subject in map
	}

	for _, tc := range testCases {
		t.Run(tc.courseName, func(t *testing.T) {
			t.Parallel()
			students, err := PrintStudentList(tc.courseName)
			if tc.expectedErr { // we are expecting this case to fail
				if err == nil { // function 'PrintStudentCourses' should return error but didn't return any
					t.Fatalf("expected an error for student '%s', but got none", tc.courseName)
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error for course %s, but got %v", tc.courseName, err)
				}
				if len(students) != len(tc.studentList) {
					t.Fatalf("for course %s, expected %d students, but got %d", tc.courseName, len(tc.studentList), len(students))
				}

				// check if slices are equal
				if !reflect.DeepEqual(students, tc.studentList) {
					t.Fatalf("for course %s, expected students %v, but got %v", tc.courseName, tc.studentList, students)
				}
			}
		})
	}
}

func TestPrintRestaurantMenu(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		restaurantName RestaurantName
		menu           map[string]float32
		expectedErr    bool
	}{
		{
			restaurantName: "Main Tower Restaurant & Lounge",
			menu: map[string]float32{
				"Kartoffelsuppe mit Würstchen": 6.70,
			},
			expectedErr: false,
		},
		{
			restaurantName: "The Ivory Club",
			menu: map[string]float32{
				"Rinderroulade mit Kartoffelpüree und Erbsen": 17.20,
			},
			expectedErr: false,
		},
		{
			restaurantName: "Holbeins",
			menu: map[string]float32{
				"Bratwurst mit Sauerkraut und Brot ": 9.80,
			},
			expectedErr: false,
		},
		{
			restaurantName: "Restaurant Villa Rothschild",
			menu: map[string]float32{
				"Wiener Schnitzel mit Kartoffelsalat ": 13.50,
			},
			expectedErr: false,
		},
		{
			restaurantName: "Apfelwein Wagner",
			menu: map[string]float32{
				"Sauerbraten mit Klößen und Rotkohl": 15.90,
			},
			expectedErr: false,
		},
		{
			restaurantName: "Unbekanntes Restaurant",
			menu:           nil,
			expectedErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(string(tc.restaurantName), func(t *testing.T) {
			t.Parallel()

			restaurantMenu, err := PrintRestaurantMenu(tc.restaurantName)
			if tc.expectedErr {
				if err == nil {
					t.Fatalf("expected an error for restaurant '%v', but got none", tc.restaurantName)
				}
			} else {
				if err != nil { // we are not expecting an error here but in case of an error, do this...
					t.Fatalf("did not expect an error for restaurant '%v', but got: %v", tc.restaurantName, err)
				}

				// check if maps are equal
				if !reflect.DeepEqual(tc.menu, restaurantMenu) {
					t.Fatalf("for restaurant %s, expected %v, but got %v", tc.restaurantName, tc.menu, restaurantMenu)
				}
			}
		})
	}
}

func TestVisitedCountries(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		username    string
		user        []VisitedCountry
		expectedErr bool
	}{
		{
			username: "Nzolani",
			user: []VisitedCountry{
				{Name: "Germany", Capital: "Berlin"},
				{Name: "USA", Capital: "Washington DC"}},
			expectedErr: false,
		},
		{
			username:    "Unknown",
			user:        nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.username, func(t *testing.T) {
			t.Parallel()

			countries, err := VisitedCountries(tc.username)
			if tc.expectedErr {
				if err == nil {
					t.Fatalf("expected an error for username '%s', but got got none", tc.username)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error for username '%s', but got '%v'", tc.username, err)
				}
				if !reflect.DeepEqual(tc.user, countries) {
					t.Fatalf("for username '%s', expected '%v', but got '%v'", tc.username, tc.user, countries)
				}
			}
		})
	}
}

func TestAirportDetails(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		airportName string
		details     map[FlightDestination]FlightsHour
		expectedErr bool
	}{
		{
			airportName: "Frankfurter Flughafen",
			details: map[FlightDestination]FlightsHour{
				"Berliner Flughafen": "12:30",
			},
			expectedErr: false,
		},
		{
			airportName: "Muenchener Flughafen",
			details: map[FlightDestination]FlightsHour{
				"Duesseldorfer Flughafen": "14:30",
			},
			expectedErr: false,
		},
		{
			airportName: "Flughafen Unbekannt",
			details:     nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.airportName, func(t *testing.T) {
			t.Parallel()
			details, err := AirportDetails(AirportName((tc.airportName)))
			if tc.expectedErr {
				if err == nil {
					t.Errorf("expected error for flight '%s' but got none", tc.airportName)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error for flight '%s', but got '%v'", tc.airportName, err)
				}
				if !reflect.DeepEqual(tc.details, details) {
					t.Fatalf("for flight '%s', expected '%v' but got '%v'", tc.airportName, tc.details, details)
				}
			}
		})
	}
}

func TestGetPaycheckDetails(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		CompanyName     string
		PaycheckDetails map[CompanyName]Paycheck
		expectedErr     bool
	}{
		{
			CompanyName: "Deutsche Börse Group",
			PaycheckDetails: map[CompanyName]Paycheck{
				"John Doe": 5000.45,
			},
			expectedErr: false,
		},
		{
			CompanyName: "Deutsche Bank AG",
			PaycheckDetails: map[CompanyName]Paycheck{
				"Jane Doe": 6000.50,
			},
			expectedErr: false,
		},
		{
			CompanyName: "Porsche AG",
			PaycheckDetails: map[CompanyName]Paycheck{
				"John Smith": 4500.55,
			},
			expectedErr: false,
		},
		{
			CompanyName: "Unbekanntes Unternehmen",
			PaycheckDetails: map[CompanyName]Paycheck{
				"Unbekanntes Unternehmen": 0.0,
			},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {

		var expectedChecks []Paycheck
		for _, expectedChecksDetails := range tc.PaycheckDetails {
			expectedChecks = append(expectedChecks, expectedChecksDetails)
		}

		t.Run(tc.CompanyName, func(t *testing.T) {
			paychecks, err := GetPaycheckDetails(CompanyName(tc.CompanyName))
			if tc.expectedErr {
				if err == nil {
					t.Errorf("expected an error for company '%s' but got none", tc.CompanyName)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error for company '%s' but got %v", tc.CompanyName, err)
				}
				if !reflect.DeepEqual(expectedChecks, paychecks) {
					t.Errorf("for company '%s' expected %v but got %v", tc.CompanyName, expectedChecks, paychecks)
				}
			}
		})
	}
}

func BenchmarkGetDaysOfTheMonth(b *testing.B) {
	months := []Month{
		"Januar",
		"Februar",
		"März",
		"April",
		"Mai",
		"Juni",
		"Juli",
		"August",
		"September",
		"Oktober",
		"November",
		"Dezember",
	}

	for i := 0; i < b.N; i++ {
		for _, month := range months {
			_ = GetDaysOfTheMonth(month)
		}
	}
}

func BenchmarkGetCapital(b *testing.B) {
	countries := []Country{"Kongo DR", "Deutschland", "Angola", "Kenia", "Senegal", "France"}

	for i := 0; i < b.N; i++ {
		for _, country := range countries {
			_ = GetCapital(country)
		}
	}
}

func BenchmarkGetProductPrice(b *testing.B) {
	products := []Product{"Banane, Apfel, Birne, Pfirsisch, Kiwi"}

	for i := 0; i < b.N; i++ {
		for _, product := range products {
			_ = GetProductPrice(product)
		}
	}
}

func BenchmarkDeleteFruitFromMap(b *testing.B) {
	fruits := []Fruit{"Banane", "Himberre", "Birne", "Kiwi", "Apfel"}

	for i := 0; i < b.N; i++ {
		for _, fruit := range fruits {
			DeleteFruitFromMap(fruit, &Fruits)
		}
	}
}

func BenchmarkNeighbouringCountries(b *testing.B) {
	countries := []string{"Deutschland", "Frankreich"}

	for i := 0; i < b.N; i++ {
		for _, country := range countries {
			PrintNeighbouringCountries(country)
		}
	}
}

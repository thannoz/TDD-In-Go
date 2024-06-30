package hashmaps

import (
	"errors"
	"fmt"
)

type Country string

func GetCapital(c Country) string {
	countries := map[Country]string{
		"Deutschland": "Berlin",
		"Kongo DR":    "Kinshasa",
		"Angola":      "Luanda",
		"Kenia":       "Nairobi",
		"Senegal":     "Dakar",
		"France":      "Paris",
		"Sweden":      "Stockholm",
	}
	return countries[c]
}

type Product string

var Products = map[Product]float64{
	Product("Banane"): 2.34,
	Product("Apfel"):  5.34,
	Product("Birne"):  6.34,
	Product("Kiwi"):   3.64,
}

func GetProductPrice(p Product) float64 {
	return Products[p]
}

type Fruit string

var Fruits = map[Fruit]string{
	Fruit("Banane"):   "gelb",
	Fruit("Himberre"): "rot",
	Fruit("Birne"):    "dunkel_rot",
	Fruit("Kiwi"):     "dunkel_gruen",
	Fruit("Apfel"):    "rot_gruen",
}

func DeleteFruitFromMap(f Fruit, fruits *map[Fruit]string) *map[Fruit]string {
	delete(*fruits, f)
	return fruits
}

type Month string

var Months = map[Month]int{
	Month("Januar"):    31,
	Month("Februar"):   28,
	Month("März"):      30,
	Month("April"):     31,
	Month("Mai"):       31,
	Month("Juni"):      30,
	Month("Juli"):      31,
	Month("August"):    31,
	Month("September"): 30,
	Month("Oktober"):   31,
	Month("November"):  30,
	Month("Dezember"):  31,
}

func GetDaysOfTheMonth(m Month) int {
	return Months[m]
}

type Animal string

var Animals = map[Animal]string{}

func AddAverageLifeSpanOfAnimal(animal Animal, lifeSpan string) (map[Animal]string, error) {
	if _, exists := Animals[animal]; exists {
		return nil, fmt.Errorf("Tier %s existiert bereits", animal)
	}

	Animals[animal] = lifeSpan
	return Animals, nil
}

var AdressBook = map[string]string{
	"Carlos":  "carlos.mwuana@hotmail.de",
	"Carmelo": "carmelo.mwinda@hotmail.de",
	"Nzolani": "nzolani.nzinga@hotmail.de",
	"Konzo":   "konzo.nzinga@hotmail.de",
}

func Contains(name string) bool {
	_, ok := AdressBook[name]
	return ok
}

var ClassRoom = map[string][]string{}

func AddStudentAndSubject(student string, subjects ...string) map[string][]string {
	if _, ok := ClassRoom[student]; !ok {
		ClassRoom[student] = []string{}
	}

	ClassRoom[student] = append(ClassRoom[student], subjects...)
	return ClassRoom
}

var NeighbouringCountries = map[string]map[string]bool{
	"Deutschland": {
		"Frankreich":  true,
		"Belgien":     true,
		"Niederlande": true,
		"Schweiz":     true,
		"Österreich":  true,
		"Polen":       true,
		"Tschechien":  true,
		"Dänemark":    true,
	},
	"Frankreich": {
		"Deutschland": true,
		"Belgien":     true,
		"Luxemburg":   true,
		"Schweiz":     true,
		"Italien":     true,
		"Spanien":     true,
		"Andorra":     true,
		"Monaco":      true,
	},
}

func PrintNeighbouringCountries(country string) ([]string, error) {
	var result []string
	if neighbours, exists := NeighbouringCountries[country]; exists {
		for neighbour := range neighbours {
			result = append(result, neighbour)
		}
		return result, nil
	}
	return nil, errors.New("country not found")
}

type Book string

var booksAuthors = map[Book]string{
	"Autobiografie Of Malcolm X":   "Alex Harley",
	"Schwarze Haut, weisse Masken": "Frantz Fanon",
	"Psychologie der Massen":       "Gustav Le Bon",
	"1984":                         "George Orwell",
	"Haben oder Sein":              "Erich Fromm",
}

func PrintAuthorsName(book Book) (string, error) {
	if author, exists := booksAuthors[book]; exists {
		return author, nil
	}
	return "", fmt.Errorf("author for '%s' not found", book)
}

var studentCourses = map[string][]string{
	"Mathematik": {"Carlos", "Konzo", "Mwinda"},
	"Physik":     {"Kalonji", "Selamawit"},
	"Informatik": {"Nzolani", "Carmelo", "Kyrie"},
}

func PrintStudentList(courseName string) ([]string, error) {
	studentList, exists := studentCourses[courseName]
	if !exists {
		return nil, fmt.Errorf("course '%s' not in map\n", courseName)
	}

	return studentList, nil
}

type RestaurantName string

var restaurantMenu = map[RestaurantName]map[string]float32{
	"Apfelwein Wagner": {
		"Sauerbraten mit Klößen und Rotkohl": 15.90,
	},
	"Restaurant Villa Rothschild": {
		"Wiener Schnitzel mit Kartoffelsalat ": 13.50,
	},
	"Holbeins": {
		"Bratwurst mit Sauerkraut und Brot ": 9.80,
	},
	"The Ivory Club": {
		"Rinderroulade mit Kartoffelpüree und Erbsen": 17.20,
	},
	"Main Tower Restaurant & Lounge": {
		"Kartoffelsuppe mit Würstchen": 6.70,
	},
	//"Unbekanntes Restaurant": nil,
}

func PrintRestaurantMenu(restaurantName RestaurantName) (map[string]float32, error) {
	restaurant, exists := restaurantMenu[restaurantName]
	if !exists {
		return nil, fmt.Errorf("restaurant '%s' not found", restaurantName)
	}

	return restaurant, nil
}

type VisitedCountry struct {
	Name    string
	Capital string
}

type User struct {
	Countries []VisitedCountry
}

var users = map[string]User{
	"Nzolani": {Countries: []VisitedCountry{
		{Name: "Germany", Capital: "Berlin"},
		{Name: "USA", Capital: "Washington DC"},
	}},
	"Konzo":   {Countries: []VisitedCountry{{Name: "Congo", Capital: "Kinshasa"}}},
	"Kalonji": {Countries: []VisitedCountry{{Name: "Angola", Capital: "Luanda"}}},
}

func VisitedCountries(username string) ([]VisitedCountry, error) {
	user, exists := users[username]
	if !exists {
		return nil, fmt.Errorf("username '%s' does not exist in map", username)
	}

	return user.Countries, nil
}

type (
	AirportName       string
	FlightDestination string
	FlightsHour       string
)

var airportInformation = map[AirportName]map[FlightDestination]FlightsHour{
	"Frankfurter Flughafen": {
		"Berliner Flughafen": "12:30",
	},
	"Muenchener Flughafen": {
		"Duesseldorfer Flughafen": "14:30",
	},
}

func AirportDetails(name AirportName) (map[FlightDestination]FlightsHour, error) {
	flightDetails, exists := airportInformation[name]
	if !exists {
		return nil, fmt.Errorf("Aiport '%s' does not exist", name)
	}

	return flightDetails, nil
}

type (
	CompanyName  string
	EmployeeName string
	Paycheck     float64
)

var companyDetails = map[CompanyName]map[EmployeeName]Paycheck{
	"Deutsche Börse Group": {
		"Jon Doe": 5000.45,
	},
	"Deutsche Bank AG": {
		"Jane Doe": 6000.50,
	},
	"Porsche AG": {
		"John Smith": 4500.55,
	},
}

func GetPaycheckDetails(cp CompanyName) ([]Paycheck, error) {
	salaryDetails, exists := companyDetails[cp]
	if !exists {
		return nil, fmt.Errorf("salary details for '%s' not found", cp)
	}

	var paycheckAmounts []Paycheck
	for _, salary := range salaryDetails {
		paycheckAmounts = append(paycheckAmounts, salary)
	}

	return paycheckAmounts, nil
}

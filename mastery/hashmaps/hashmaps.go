package hashmaps

import (
	"errors"
	"fmt"
)

type Country string

// 1. Erstelle eine Map, die den Namen eines Landes als Schlüssel und seine Hauptstadt als Wert enthält.
// Gib die Hauptstadt für das Land "Deutschland" aus.

// GetCountry returns the capital of a country
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

// 2. Erstelle eine Map, die den Namen eines Produkts als Schlüssel und seinen Preis als Wert enthält.
//Füge drei Produkte hinzu und gib den Preis für das Produkt mit dem Namen "Apfel" aus.

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

// 3. Erstelle eine Map, die den Namen einer Frucht als Schlüssel und ihre Farbe als Wert enthält.
// Entferne die Frucht "Banane" aus der Map.
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

// 4. Erstelle eine Map, die den Namen eines Monats als Schlüssel und
// die Anzahl der Tage in diesem
// Monat als Wert enthält. Gib die Anzahl der Tage im Monat "Februar" aus.
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

// 5. Erstelle eine Map, die den Namen eines Tieres als Schlüssel und seine durchschnittliche
// Lebensdauer als Wert enthält. Füge einige Tiere und ihre durchschnittliche Lebensdauer
// hinzu und gib die durchschnittliche Lebensdauer eines bestimmten Tieres aus.

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

// Intermediate:

// 1. Erstelle eine Map, die den Namen eines Benutzers als Schlüssel
// und seine E-Mail-Adresse als Wert enthält.
// Überprüfe, ob ein Benutzer mit dem Namen "Konzo" in der Map vorhanden ist.

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

// 2. Erstelle eine Map, die den Namen eines Studenten als Schlüssel und eine Liste
// seiner Fächer als Wert enthält. Füge einem Studenten "Max" das Fach "Mathematik"
// hinzu und gib seine Fächer aus.

var ClassRoom = map[string][]string{}

// AddStudentAndSubject assigns one or multiple subjects to a student
func AddStudentAndSubject(student string, subjects ...string) map[string][]string {
	if _, ok := ClassRoom[student]; !ok {
		ClassRoom[student] = []string{}
	}

	ClassRoom[student] = append(ClassRoom[student], subjects...)
	return ClassRoom
}

// 3. Erstelle eine Map, die den Namen eines Landes als Schlüssel und eine Map
//     mit seinen Nachbarländern als Wert enthält. Füge jedem Land einige Nachbarländer
//     hinzu und gib die Nachbarländer von "Deutschland" aus.

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

// PrintNeighbouringCountries returns neighboring countries of a given country
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

/*
 4. Erstelle eine Map, die den Namen eines Buches als Schlüssel und den Namen des
    Autors als Wert enthält. Füge einige Bücher und ihre Autoren hinzu und gib den
    Autor eines bestimmten Buches aus.
*/
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

/*
5. Erstelle eine Map, die den Namen eines Kurses als Schlüssel und die Liste der Schüler,
    die diesen Kurs besuchen, als Wert enthält. Füge Schüler zu verschiedenen Kursen hinzu
    und gib die Liste der Schüler für einen bestimmten Kurs aus. */

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

/*
 1. Erstelle eine Map, die den Namen eines Restaurants als Schlüssel und eine Map

mit den verfügbaren Gerichten und ihren Preisen als Wert enthält. Füge Gerichte
und Preise zu verschiedenen Restaurants hinzu und gib die Speisekarte eines bestimmten Restaurants aus.
*/
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

/*
 2. Erstelle eine Map, die den Namen eines Benutzers als Schlüssel und eine
    Map mit den von ihm besuchten Ländern und deren Hauptstädten als Wert enthält.
    Füge einem Benutzer Länder und Hauptstädte hinzu und gib die von einem bestimmten
    Benutzer besuchten Länder aus.
*/
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

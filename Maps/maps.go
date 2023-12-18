package maps


// make declaration of maps
var counters = make(map[string]string, 10)

// Using composotie literal
modelTomake := map[string]string{
	"prius": "toyota",
	"chevelle": "chevy",
}


carMake := modelTomake["prius"]
fmt.Println(carMake)

carMake2 := modelTomake["Maybach"]
fmt.Println(carMake2) // Returns a zero value as key is not in the map

if carMake3, ok := modelTomake["prius"]; ok {
	fmt.Printf(" car model \"prius\" has make %q", carMake3)
} else {
	fmt.Printf(" car model \"prius\" has make an unknown make")
}

//Adding key-value pairs or updating a key's value

modelTomake["prius"] = "subaru"
counters["pageHits"] = 10
fmt.Println(counters["pageHits"])

//Extracting values

for key, value := range modelTomake {
	fmt.Printf("The car model %q has make %q\n", key, value)
}
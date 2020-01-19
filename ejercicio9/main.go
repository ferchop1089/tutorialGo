package main

import restclient1 "tutorialGo/ejercicio9/client"

func main() {
	var rest restclient1.CountryRest
		rest = restclient1.NewCountryRest(rest)
	rest.GetCountry("1")
}

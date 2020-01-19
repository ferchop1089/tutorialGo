package restclient1

import "fmt"

type CountryAdapter struct {
	client interface{}
}

func NewCountryRest(client interface{}) Country {
	return CountryAdapter{client: client}
}

func (country CountryAdapter) GetCountry(ID string) {
	fmt.Println(ID)
}

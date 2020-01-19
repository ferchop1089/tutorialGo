package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hola mundo!")
	os.Setenv("ENV_TEST_GO", "Prueba de paquete os")
	os.Setenv("MELI_GET_COUNTRY_API", "https://api.mercadolibre.com/countries/ otra url")
	os.Setenv("GET_COUNTRY_API", "https://api.mercadolibre.com/countries/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	envValue := viper.Get("ENV_TEST_GO")
	fmt.Printf("El valor obtenido es1: %s\n", envValue)

	envValue = viper.Get("GET_COUNTRY_API")
	fmt.Printf("El valor obtenido es2: %s\n", envValue)

	envValue = os.Getenv("GET_COUNTRY_API")
	fmt.Printf("El valor obtenido es3: %s\n", envValue)

	// ------------------------------------------------- //
	viper.SetEnvPrefix("MELI")
	viper.BindEnv("GET_COUNTRY_API") // Enlaza el valor al prefijo. p.e "MELI_GET_COUNTRY_API"

	envValue = viper.GetString("GET_COUNTRY_API") // Concatena el prefijo definido arriba quedando p.e "MELI_GET_COUNTRY_API"
	fmt.Printf("El valor obtenido es4: %s\n", envValue)

	// ------------------------------------------------- //
	var str string = `{"message": "Country not found","error": "not_found","status": 404,"cause": []}`

	viper.SetConfigType("json")
	errorViper := viper.ReadConfig(bytes.NewBuffer([]byte(str)))
	if errorViper != nil {
		fmt.Println("Error:", errorViper)
	}
	envValue = viper.GetString("message")
	fmt.Printf("El valor obtenido del JSON es: %s\n", envValue)
}

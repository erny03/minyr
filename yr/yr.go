package yr

import (
	"fmt"
	"strconv"
	"strings"
	//"errors"
	"os"
	"bufio"
)

const (
	fahrFil = "kjevik-temp-fahr-20220318-20230318.csv"
	celsiusFil = "kjevik-temp-celsius-20220318-20230318.csv"

)


func averageTempCelsius() {

	file, err := os.Open(celsiusFil)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	count := 0

	scanner := bufio.NewReader(file)
	for scanner.Scan(){
		elements := strings.Fields(scanner.Text())

		num, err := strconv.Atoi(elements[3])
		if err != nil {
			panic(err)
		}
		sum += num
		count++
	}


	if count > 0 {
		return float64(sum) / float64(count)
	} else {
		fmt.Println("Ingen temperaturer funnet i filen.")
		return 0
	}
}



/* func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

        dividedString := strings.Split(line, ";")
	var err error
	
	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
	
	/*	
	return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
        */
// }


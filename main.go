package main

import (
	"os"
	"strconv"
	"fmt"
	"log"
	"bufio"
	"io"
	"strings"
	"github.com/erny03/minyr/conv"
	//"github.com/erny03/minyr/yr"
)	


const (
	celsiusFil = "kjevik-temp-celsius-20220318-20230318.csv"
	fahrFil = "kjevik-temp-fahr-20220318-20230318.csv"
)

func main() {

	_, error := os.Stat(fahrFil)

	//src, err := os.Open("table.csv")
	src, err := os.Open(celsiusFil)
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        	log.Println(src)

	fahrFile, err := os.Create(fahrFil)
	if err != nil {
		log.Fatal(err)
	}

	defer fahrFile.Close()

	writer := bufio.NewWriter(fahrFile)

	var totalCelsius float64
	var totalLinjer float64
	var input string
	var totalFahr float64

	fmt.Println("Skriv convert for aa oversette, skriv quit for aa avslutte.")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
	input = scanner.Text()


	if input != "convert" && input != "j" && input != "average" && input != "c" && input != "f" && input != "quit" && input != "n" {
		fmt.Println("Ikke en gylding kommando.")
	}

	if input == "convert" && !os.IsNotExist(error){
		fmt.Println("Filen eksisterer allerede. Vil du generere filen om igjen? (j/n).")
	}

	if input == "n" {
		fmt.Println("Nedlastning av ny fil stopped, program avsluttes.")
		break
	}
	if input == "j"{




	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
        bytesCount := 0

	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
		   // Her
		   //line := string(linebuf)
		   elementArray := strings.Split(string(linebuf), ";")
		   if len(elementArray) > 3 {
			 if elementArray[3] == "Lufttemperatur"{
			 	log.Println(string(linebuf))
				_, err = writer.WriteString(string(linebuf)+"\n")
				if err != nil {
					log.Fatal(err)
				}
			 }
		 	 if elementArray[3] != "Lufttemperatur" && elementArray[3] != ""{
				 celsius := elementArray[3]
				 celsiusFloat, err := strconv.ParseFloat(celsius, 64)
				 fahr := conv.CelsiusToFahrenheit(celsiusFloat)
				 fahrString:= strconv.FormatFloat(fahr, 'f',2, 64)
				 if err != nil{
					panic(err)
 				 }

				 log.Println(elementArray[0] + ";" + elementArray[1] + ";" + elementArray [2] + ";" + fahrString)
				 _, err = writer.WriteString(elementArray[0] + ";" + elementArray[1] + ";" + elementArray[2] + ";" + fahrString + "\n")
				 if err != nil {
					log.Fatal(err)
				 }
				 totalCelsius += celsiusFloat
				 totalLinjer ++
				 totalFahr += fahr

			}
			if elementArray[3] == ""{
				log.Println("Data er basert paa gyldig (per 18.03.2023) (CC BY 4.0) fra Meteorologisk insitutt (MET); Endringen er gjort av Erik Nygaard \n")
				_, err = writer.WriteString("Data er basert paa gyldig (per 18.03.2023) (CC BY 4.0) fra Meteorlogisk insitutt (MET); Endringen er gjort av Erik Nygaard\n")
				if err != nil{
					log.Fatal(err)
				}
				writer.Flush()
			}

	   	   }
                   linebuf = nil
		} else {
                   linebuf = append(linebuf, buffer[0])
		}
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}

		log.Println("For gjennomsnitts celsius temperatur skriv c, for fahrenheit skriv f")
		log.Println("For aa avslutte programmet skriv quit")

	} else if input == "f" {
		avgFahr := strconv.FormatFloat(totalFahr/totalLinjer, 'f', 2, 64)
		log.Println(avgFahr)
	} else if input == "c" {
		avgCelsius := strconv.FormatFloat(totalCelsius/totalLinjer, 'f', 2, 64)
		log.Println(avgCelsius)
	} else if input == "quit" {
		break
	}
}

}


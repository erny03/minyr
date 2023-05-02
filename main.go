package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"io"
	"strings"
	"github.com/erny03/minyr/conv"
)	

func main() {
	//src, err := os.Open("table.csv")
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        	log.Println(src)
        
	fahrFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer fahrFile.Close()

	writer := bufio.NewWriter(fahrFile)

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
			}
			if elementArray[3] == ""{
				log.Println("Data er basert paa gyldig (per 18.03.2023) (CC BY 4.0) fra Meteorologisk insitutt (MET); Endringen er gjort av Erik Nygaard \n")
				_, err = writer.WriteString("Data er basert paa gyldig (per 18.03.2023) (CC BY 4.0) fra Meteorlogisk insitutt (MET); Endringen er gjort av Erik Nygaard\n")
				if err != nil{
					log.Fatal(err)
				}
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
	//
	err = writer.Flush()
	if err != nil{
		log.Fatal(err)
	}
}

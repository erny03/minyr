package main

import (
	"os"
	"strconv"
	"log"
	"io"
	"strings"
	"github.com/erny03/minyr/conv"
)	

func main() {
	src, err := os.Open("table.csv")
	//src, err := os.Open("/home/janisg/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        log.Println(src)
        
	
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
		   elementArray := strings.Split(string(linebuf), ";")
		   if len(elementArray) > 3 {
			 if elementArray[3] == "Lufttemperatur"{
			 	log.Println(string(linebuf)) 
			 }
		 	 if elementArray[3] != "Lufttemperatur" && elementArray[3] != ""{
				 celsius := elementArray[3]
				 celsiusFloat, err := strconv.ParseFloat(celsius, 64)
				 fahr := conv.CelsiusToFahrenheit(celsiusFloat)
				 fahrString:= strconv.FormatFloat(fahr, 'f',2, 64)
				 if err != nil{
					panic(err)
 				 }

		        	 log.Println(elementArray[0] + ";" +  elementArray[1] + ";" + elementArray[2] + ";" + fahrString)
			 }
			if elementArray[3] == ""{
				log.Println(string(linebuf))
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

}


package yr

import (
	"os"
	"bufio"
	"testing"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
     type test struct {
	input string
	want string
     }
     tests := []test{
	     {input: "6", want: "42.8"},
	     {input: "0", want: "32.0"},
     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitString(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }
}
// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperatrmaaling i grader celsius
func TestCelsiusToFahrenheitLine(t *testing.T) {
     type test struct {
	input string
	want string
     }
     tests := []test{
	     {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	     {input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
	     {input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevfik;SN39040;08.03.2023 02:20;12.2"},

     }

     for _, tc := range tests {
	     got, _ := CelsiusToFahrenheitLine(tc.input)
	     if !(tc.want == got) {
		     t.Errorf("expected %s, got: %s", tc.want, got)
	     }
     }

}

func TestTotalLines(t *testing.T) {
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318,csv")

	if err != nil {
		t.Fatalf("error opening output file: %v", err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	totalLines := 0
	for scanner.Scan() {
		totalLines++
	}
	if totalLines != 16756 {
		t.Errorf("uforventet antall linjer i fahr fil: forventet: 16756, fikk %d", totalLines)
	}
}

func TestAverageCelsius(t *testing.T) {

	type test struct {
		input string
		want string
	}


	tests := []test{
		{input: "c", want: "8.56"},
	}
	for _, tc := range tests {
		got, _ := avgCelsius(tc.input)
		if !(tc.want == got) {
			t.Errorf("forventet %s, fikk: %s", tc.want, got)
		}
	}
}

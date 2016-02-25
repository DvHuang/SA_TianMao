package Positive_Negative
import (
	"strings"
	"os"
	"bufio"
	"log"
	"Lab_Tm_web/Global_Struct"
)



func Creat_EV (Ev_file string )(Ev_map map[string] Global_Struct.Emotional_vocabulary) {

	file, err := os.Open(Ev_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	Ev_map=make(map [string] Global_Struct.Emotional_vocabulary)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		if len(data) < 5 {
			continue
		}
		var linshi Global_Struct.Emotional_vocabulary
		linshi.Pos=data[1]
		linshi.M_E_classification=data[2]
		linshi.M_E_Strength=data[3]
		linshi.M_E_Polarity=data[4]

		Ev_map[data[0]]= linshi
	}

	return Ev_map
}
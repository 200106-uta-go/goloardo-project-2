package savefile

import (
	"fmt"
	"log"
	"os"

	"github.com/danish287/project-1/config"
)

//SaveFile saves requested horoscope as a text file under the /savedhoroscopes directory.
func SaveFile(reqHoroscope string) {
	exportFile, err := os.Create(config.FileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(exportFile, "%v\n", reqHoroscope)
	err = os.Rename(config.FileName, "../horoscope/savedhoroscopes/"+config.FileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your horoscoped has been saved as %s under the /savedhoroscopes directory.\n F", config.FileName)
}

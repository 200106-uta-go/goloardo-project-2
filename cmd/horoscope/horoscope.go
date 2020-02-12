package main

import (
	"fmt"
	"strings"

	"github.com/danish287/project-1/config"
	"github.com/danish287/project-1/internal/gethoroscope"
	"github.com/danish287/project-1/internal/savefile"
)

func main() {
	usrSunsign := config.UserSunsign
	reqDate := config.RequestedDate
	reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)
	fmt.Println(reqHoroscope)

	if !strings.Contains(reqHoroscope, "Please try") {
		fmt.Print("\nWould you like to save this horoscope as a text file? (y/n)\n")
		fmt.Scanln(&config.StoreHoroscope)

		if config.StoreHoroscope == "y" {
			savefile.SaveFile(reqHoroscope)
		}
	}
}

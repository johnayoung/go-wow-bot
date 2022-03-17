package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/johnayoung/go-wow-bot/utils"
	cp "github.com/otiai10/copy"
)

const pathToAddons = "C:/Program Files/Ascension Launcher/resources/client/Interface/AddOns/"

var addons = [3]string{"Ascended", "DataToColor", "LoadedReroll"}

func main() {
	Copy()
}

func Copy() {
	if _, err := os.Stat(pathToAddons); !os.IsNotExist(err) {
		fmt.Printf("Writing Addons to %s\n", pathToAddons)

		for _, addon := range addons {
			from := filepath.Join(utils.RootDir(), "addons", addon)
			to := filepath.Join(pathToAddons, addon)

			fmt.Printf("From: %s\n", from)
			fmt.Printf("To: %s\n", to)
			err := cp.Copy(from, to)
			fmt.Println(err) // nil
		}
	}
}

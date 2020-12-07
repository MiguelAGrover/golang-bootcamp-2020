package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"wizegolangapi/config"
	"wizegolangapi/domain/model"
	"wizegolangapi/infraestructure/datastore"
	"wizegolangapi/infraestructure/router"
	"wizegolangapi/registry"

	"github.com/labstack/echo"
)

func main() {

	config.ReadConfig()

	db := datastore.NewCSVDB(config.C.Dest.DigimonCSV)

	write(db)

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
}

// write Obtain data from an external API convert it to an array and save it into csv file
func write(db datastore.CSVDB) {
	if !db.FileExist() {
		resp, err := http.Get(config.C.Sources.DigimonAPI)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		var DigimonStructArray []model.Digimon
		json.Unmarshal(bodyBytes, &DigimonStructArray)

		var DigimonStringArray [][]string

		for _, digimon := range DigimonStructArray {
			var row []string
			row = append(row, digimon.Name)
			row = append(row, digimon.Level)
			row = append(row, digimon.Image)
			DigimonStringArray = append(DigimonStringArray, row)
		}

		db.WriteFullCSV(DigimonStringArray)
	}
}

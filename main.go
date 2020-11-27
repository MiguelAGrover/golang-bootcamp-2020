package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	// This is the methods used for the first delivery
	write(db)
	// read(csvdb)

	// From here it begins the clean architecture for the final delivery
	// db := datastore.NewDB(config.C.Sqlitedb.DBPath)
	//sqlDB, err := db.DB()
	// if err != nil {
	//  	log.Fatalln(err)
	// }

	// defer sqlDB.Close()

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
	resp, err := http.Get(config.C.Sources.DigimonAPI)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyBytes)

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
	// remember to flush!
	db.WriteFullCSV(DigimonStringArray)
}

// read Takes the information from a csv file, convert it to an array of Digimon structure and convert it to json.
func read(db datastore.CSVDB) {
	records, err := db.LoadCSV()
	var digimon model.Digimon
	var digimons []model.Digimon
	for _, rec := range records {
		digimon.Name = string(rec[0])
		digimon.Level = string(rec[1])
		digimon.Image = string(rec[2])
		digimons = append(digimons, digimon)
	}

	jsonData, err := json.Marshal(digimons)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(jsonData)
}

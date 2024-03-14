package main

import "strings"

// https://www.hamburg.de/postleitzahlen/
var blacklistedPostalCodes = []int{
	21035, 21037, // Allermöhe
	22111, 22115, 22117, 22119, // Billstedt
	22047, 22159, 22175, 22179, // Bramfeld
	22587,                             // Blankenese
	22457, 22523, 22525, 22527, 22547, // Eidelstedt
	21075, 21077, // Eißendorf
	21075, 21079, // Harburg
	21075, 21079, // Heimfeld
	22043, 22045, // Jenfeld
	21077,                      // Marmstorf
	21129,                      // Neuenfelde
	22453,                      //Niendorf
	22041,                      //Wandsbek
	22549,                      // Osdorf
	22391, 22393, 22395, 22399, //Poppenbüttel
}

func (app *Application) checkIfAdSuitable(ad Advertisement) bool {
	if contains(blacklistedPostalCodes, ad.PostalCode) {
		app.InfoLogger.Printf("Didnt add ad because postal code '%d' was in blacklist.", ad.PostalCode)
		return false
	}

	nameOK := app.checkName(strings.ToLower(ad.Name))
	if !nameOK {
		return false
	}

	//TODO #rooms, size in square meter, quarter-name?
	return true
}

func (app *Application) checkName(name string) bool {
	if strings.Contains(name, "tauschwohnung") {
		app.InfoLogger.Println("Didnt add ad because TAUSCHWOHNUNG")
		return false
	}

	if strings.Contains(name, "untermiete") {
		app.InfoLogger.Println("Didnt add ad because UNTERMIETE")
		return false
	}

	return true
}

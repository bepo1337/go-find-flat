package main

import "strings"

// https://www.hamburg.de/postleitzahlen/
var blacklistedPostalCodes = []int{
	21035, 21037, // Allermöhe
	22111, 22115, 22117, 22119, // Billstedt
	22047, 22159, 22175, 22177, 22179, 22309, // Bramfeld
	22587,                             // Blankenese
	22457, 22523, 22525, 22527, 22547, // Eidelstedt
	21075, 21077, // Eißendorf
	21075, 21079, // Harburg
	21075, 21079, // Heimfeld
	22043, 22045, // Jenfeld
	21077,                      // Marmstorf
	21129,                      // Neuenfelde
	22453, 22455, 22457, 22459, //Niendorf
	22041,                      //Wandsbek
	22549,                      // Osdorf
	22391, 22393, 22395, 22399, //Poppenbüttel
	21147, 21148, 21149, //Hausbruch
	22309,                             // Steilshoop
	22143, 22145, 22147, 22149, 22359, // Rahlstedt
	21029,        //Altengamme/Bergedorf
	22335,        //Fuhlsbüttel
	22419, 22417, //Langenhorn
	22339,        //Hummelsbüttel
	22049,        // Dulsberg
	22589,        //Blankenese
	22113,        // Billbrook
	22589,        //Iserbrook
	22559,        //Rissen
	22765,        //Ottensen
	21033, 21031, //Lohbrügge
	22609,                      //Otmatschen
	22397,                      //Duvenstedt
	22761,                      //Bahrenfeld
	22309, 22335, 22337, 22391, //Ohlsdorf
	20535, 20537, //Hamm
}

func (app *Application) checkIfAdSuitable(ad Advertisement) bool {
	if contains(blacklistedPostalCodes, ad.PostalCode) {
		app.InfoLogger.Printf("Didnt add ad because postal code '%d' was in blacklist.", ad.PostalCode)
		return false
	}

	nameOK := app.checkName(ad.Name)
	if !nameOK {
		return false
	}

	//TODO #rooms, size in square meter, quarter-name?
	return true
}

func (app *Application) checkName(name string) bool {
	name = strings.ToLower(name)

	if nameContainsOne(name, "tausch") {
		app.InfoLogger.Printf("Didnt add ad because TAUSCHWOHNUNG: %s\n", name)
		return false
	}

	if nameContainsOne(name, "untermiete", "untervermietung", "zwischenmiete") {
		app.InfoLogger.Printf("Didnt add ad because UNTERMIETE: %s\n", name)
		return false
	}

	if nameContainsOne(name, "wg-zimmer", "wg zimmer") {
		app.InfoLogger.Printf("Didnt add ad because WG-Zimmer: %s\n", name)
		return false
	}

	if nameContainsOne(name, "wg-zimmer", "wg zimmer") {
		app.InfoLogger.Printf("Didnt add ad because WG-Zimmer: %s\n", name)
		return false
	}

	if nameContainsOne(name, "1-zimmer", "1 zimmer", "1,5-zimmer", "1.5-zimmer", "1,5 zimmer", "1.5 zimmer",
		"1 1/2") {
		app.InfoLogger.Printf("Didnt add ad because <2 rooms: %s\n", name)
		return false
	}

	if nameContainsOne(name, "wir suchen", "paar sucht") {
		app.InfoLogger.Printf("Didnt add ad because its not an offer: %s\n", name)
		return false
	}

	return true
}

func nameContainsOne(name string, words ...string) bool {
	for _, word := range words {
		if strings.Contains(name, word) {
			return true
		}
	}
	return false
}

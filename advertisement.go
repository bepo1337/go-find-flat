package main

type Advertisement struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Quarter    string `json:"quarter"`
	PostalCode int    `json:"postalCode"`
}

func (a Advertisement) Compare(b Advertisement) int {
	if a.URL == b.URL {
		return 0
	} else {
		return 1
	}
}

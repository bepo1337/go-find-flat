package main

type Advertisement struct {
	Name       string
	URL        string
	Quarter    string
	PostalCode int
}

func (a Advertisement) Compare(b Advertisement) int {
	if a.URL == b.URL {
		return 0
	} else {
		return 1
	}
}

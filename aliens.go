package main

type alien struct {
	name   string
	atCity *city
	dead   bool
}

var aliens []*alien

var deadAliens int

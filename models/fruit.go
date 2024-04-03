package models

import "time"

type Fruit struct {
	Id     string
	Name   string
	Months []time.Month
}

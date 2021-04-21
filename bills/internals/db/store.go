package db

import "bills/internals/models/v1"

var categories = []models.DisplayCategory{
	{Title: "Airtime", CreatedOn: "", IsActive: true}, 
	{Title: "Data plan", CreatedOn: "", IsActive: true}, 
	{Title: "Vehicle licence", CreatedOn: "", IsActive: false},
}

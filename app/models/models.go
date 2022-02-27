package models

type Author struct {
	ID  		uint
	firstName  string
	lastName   string
	books      Book
	
}

type Book struct {
	ID    uint
	title  string
	author  string
	price  float32
}

type User struct {
	ID			uint
	firstName	string
	lastName    string
	age        	int
	region     	string
}
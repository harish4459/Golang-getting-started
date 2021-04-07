package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t\t%q\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{

	Book{
		ID:            1,
		Title:         "paperNumber1",
		Author:        "Sunny",
		YearPublished: 1994,
	},

	Book{
		ID:            2,
		Title:         "paperNumber2",
		Author:        "Sunny2",
		YearPublished: 1995,
	},
	Book{
		ID:            3,
		Title:         "paperNumber3",
		Author:        "Sunny3",
		YearPublished: 1996,
	},
	Book{
		ID:            4,
		Title:         "paperNumber4",
		Author:        "Sunny4",
		YearPublished: 1997,
	},
	Book{
		ID:            5,
		Title:         "paperNumber5",
		Author:        "Sunny5",
		YearPublished: 1998,
	},
	Book{
		ID:            6,
		Title:         "paperNumber6",
		Author:        "Sunny6",
		YearPublished: 1999,
	},
	Book{
		ID:            7,
		Title:         "paperNumber7",
		Author:        "Sunny7",
		YearPublished: 2000,
	},
	Book{
		ID:            8,
		Title:         "paperNumber8",
		Author:        "Sunny",
		YearPublished: 2001,
	},
	Book{
		ID:            9,
		Title:         "paperNumber9",
		Author:        "Sunny9",
		YearPublished: 2002,
	},
	Book{
		ID:            10,
		Title:         "paperNumber10",
		Author:        "Sunny10",
		YearPublished: 2003,
	},
}

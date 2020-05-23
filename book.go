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
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Alchemist",
		Author:        "Paulo Coelho",
		YearPublished: 2001,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	Book{
		ID:            6,
		Title:         "Don Quixote",
		Author:        "Miguel de Cervantes",
		YearPublished: 1931,
	},
	Book{
		ID:            7,
		Title:         "Moby Dick",
		Author:        "Herman Melville",
		YearPublished: 1954,
	},
	Book{
		ID:            8,
		Title:         "Hamlet",
		Author:        "William Shakespeare",
		YearPublished: 1599,
	},
	Book{
		ID:            9,
		Title:         "The Divine Comedy",
		Author:        "Dante Alighieri",
		YearPublished: 1472,
	},
	Book{
		ID:            10,
		Title:         "Anna Karenina",
		Author:        "Leo Tolstoy",
		YearPublished: 1875,
	},
}

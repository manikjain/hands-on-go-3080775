// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	author
	title string
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	library := library{}
	// add 2 books to the library by the same author
	library["SS"] = []book{
		{
			title: "Raag Darbari",
			author: author{
				name: "Srilal Shukl",
			},
		},
		{
			title: "Agyaatvaas",
			author: author{
				name: "Srilal Shukl",
			},
		},
	}

	// add 1 book to the library by a different author
	library.addBook(book{
		title: "The Small Bachelor",
		author: author{
			name: "PG Wodehouse",
		},
	})

	// dump the library with spew
	// spew.Dump(library)
	// fmt.Println(library)

	// lookup books by known author in the library
	books := library.lookupByAuthor("SS")
	spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title + " by " + b.author.name)
		// fmt.Println(library.lookupByAuthor("SS")[0].title + " by " + library.lookupByAuthor("SS")[0].author.name)
	}
}

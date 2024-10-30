package main

import (
	"fmt"
	"strings"
)

type Book struct {
	Title        string
	Author       string
	Year         int
	IsCheckedOut bool
}

var library []Book

func addBook(title, author string, year int) {
	book := Book{
		Title:        title,
		Author:       author,
		Year:         year,
		IsCheckedOut: false,
	}
	library = append(library, book)
	fmt.Printf("Added \"%s\" by %s.\n", title, author)
}

func displayBooks() {
	if len(library) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	for _, book := range library {
		status := "Available"
		if book.IsCheckedOut {
			status = "Checked Out"
		}
		fmt.Printf("%s by %s (%d) - %s\n", book.Title, book.Author, book.Year, status)
	}
}

func checkOutBook(title string) {
	for i := range library {
		if strings.ToLower(library[i].Title) == strings.ToLower(title) {
			if !library[i].IsCheckedOut {
				library[i].IsCheckedOut = true
				fmt.Printf("Checked out \"%s\".\n", title)
			} else {
				fmt.Printf("\"%s\" is already checked out.\n", title)
			}
			return
		}
	}
	fmt.Printf("Book titled \"%s\" not found.\n", title)
}

func returnBook(title string) {
	for i := range library {
		if strings.ToLower(library[i].Title) == strings.ToLower(title) {
			if library[i].IsCheckedOut {
				library[i].IsCheckedOut = false
				fmt.Printf("Returned \"%s\".\n", title)
			} else {
				fmt.Printf("\"%s\" was not checked out.\n", title)
			}
			return
		}
	}
	fmt.Printf("Book titled \"%s\" not found.\n", title)
}

func main() {
	for {
		fmt.Println("\n1. Add a book")
		fmt.Println("2. Display all books")
		fmt.Println("3. Check out a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title, author string
			var year int
			fmt.Print("Enter book title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter book author: ")
			fmt.Scanln(&author)
			fmt.Print("Enter publication year: ")
			fmt.Scan(&year)
			addBook(title, author, year)
		case 2:
			displayBooks()
		case 3:
			var title string
			fmt.Print("Enter book title to check out: ")
			fmt.Scanln(&title)
			checkOutBook(title)
		case 4:
			var title string
			fmt.Print("Enter book title to return: ")
			fmt.Scanln(&title)
			returnBook(title)
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

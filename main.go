package main

import "fmt"

type bookmarkType = map[string]string

func main() {
	fmt.Println("*** Bookmarks App ***")
	bookmarks := bookmarkType{} // how much memory do we need?

Menu:
	for {
		userAction := getMenu()

		switch userAction {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			removeBookmark(bookmarks)
		case 4:
			fmt.Println("Application closed")
			break Menu
		}
	}
}

func getMenu() int {
	var userAction int

	fmt.Println("Choose the menu option")
	fmt.Println("1. Show all bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Remove a bookmark")
	fmt.Println("4. Exit the application")
	fmt.Scan(&userAction)
	return userAction
}

func showBookmarks(bookmarks bookmarkType) {
	if len(bookmarks) == 0 {
		fmt.Println("You have no bookmarks yet")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks bookmarkType) {
	var newNameBookmark string
	var newAdressBookmark string

	fmt.Println("Please input bookmark name")
	fmt.Scan(&newNameBookmark)
	fmt.Println("Please input bookmark adress")
	fmt.Scan(&newAdressBookmark)
	bookmarks[newNameBookmark] = newAdressBookmark
}

func removeBookmark(bookmarks bookmarkType) {
	var bookmarkToDelete string
	fmt.Println("Please input bookmark name")
	fmt.Scan(&bookmarkToDelete)
	delete(bookmarks, bookmarkToDelete)
	fmt.Println(bookmarkToDelete, "has been removed")
}

package main

import "fmt"

func main() {
	fmt.Println("*** Bookmarks App ***")
	bookmarks := map[string]string{}

	for {
		userAction := getMenu()

		switch userAction {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = removeBookmark(bookmarks)
		case 4:
			fmt.Println("Application closed")
			break
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

func showBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("You have no bookmarks yet")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newNameBookmark string
	var newAdressBookmark string

	fmt.Println("Please input bookmark name")
	fmt.Scan(&newNameBookmark)
	fmt.Println("Please input bookmark adress")
	fmt.Scan(&newAdressBookmark)
	bookmarks[newNameBookmark] = newAdressBookmark
	return bookmarks
}

func removeBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkToDelete string
	fmt.Println("Please input bookmark name")
	fmt.Scan(&bookmarkToDelete)
	delete(bookmarks, bookmarkToDelete)
	fmt.Println(bookmarkToDelete, "has been removed")
	return bookmarks

}

package main

import "fmt"

func main() {

	actions := map[int]string{
		1: "Посмотреть закладки",
		2: "Добавить закладку",
		3: "Удалить закладку",
		4: "Выход",
	}

	bookmarks := map[string]string{
		"yandex": "https://yandex.ru/",
		"google": "https://google.com/",
	}

	for keyMenu, valueMenu := range actions {
		fmt.Println(keyMenu, valueMenu)
	}

	var userAction int8
	var userBookmarkName string
	var userBookmarkAdress string

	fmt.Scan(&userAction)

	switch actions {
	case 1:
		fmt.Println(bookmarks)
	case 2:
		fmt.Println("Please input bookmark name")
		fmt.Scan(&userBookmarkName)
		fmt.Println("Please input bookmark adress")
		fmt.Scan(&userBookmarkName)
		bookmarks[userBookmarkName] = userBookmarkAdress
	case 3:
		fmt.Println("Please input bookmark name")
		fmt.Scan(&userBookmarkName)
		delete(bookmarks, userBookmarkName)
		fmt.Println(userBookmarkName, "has been deleted")
		fmt.Println("Bookmark has", bookmarks)

	}
}

package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	fmt.Println("Менеджер закладок")
	bookmarks := bookmarkMap{}

	isContinue := true

	for isContinue == true {
		isContinue = doCommandMenu(bookmarks)
	}
}

func printCommandMenu() {
	fmt.Println("Выберите команду:")
	fmt.Println("1. Просмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
}

func doCommandMenu(bookmarks bookmarkMap) bool {
	printCommandMenu()

	command := 0
	fmt.Scan(&command)

	switch command {
	case 1:
		printBookmarks(bookmarks)
	case 2:
		addBookmark(bookmarks,
			inputKeyBookmark(),
			inputValueBookmark())
	case 3:
		deleteBookmark(bookmarks,
			inputKeyBookmark())
	case 4:
		return false
	default:
		fmt.Println("Команда не распознана")
		fmt.Println("")
		return true
	}

	return true
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
	}
	for key, value := range bookmarks {
		fmt.Println(key,
			": ",
			value)
	}
}

func addBookmark(bookmarks bookmarkMap,
	newBookmarkKey string,
	newBookmarkValue string) {

	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap,
	newBookmarkKey string) {

	delete(bookmarks,
		newBookmarkKey)
}

func inputKeyBookmark() string {
	fmt.Print("Введите название закладки: ")
	var key string
	fmt.Scan(&key)

	return key
}

func inputValueBookmark() string {
	fmt.Print("Введите ссылку: ")
	var value string
	fmt.Scan(&value)

	return value
}

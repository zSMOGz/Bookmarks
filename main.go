package main

import "fmt"

func main() {
	fmt.Println("Менеджер закладок")
	bookmarks := map[string]string{}

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

func doCommandMenu(bookmarks map[string]string) bool {
	printCommandMenu()

	command := 0
	fmt.Scan(&command)

	switch command {
	case 1:
		showBookmarks(bookmarks)
	case 2:
		addBookmark(bookmarks)
	case 3:
		deleteBookmark(bookmarks)
	case 4:
		return false
	default:
		fmt.Println("Команда не распознана")
		fmt.Println("")
		return true
	}

	return true
}

func showBookmarks(bookmarks map[string]string) {
	fmt.Println(bookmarks)
}

func addBookmark(bookmarks map[string]string) {
	bookmarks[inputKeyBookmark()] = inputValueBookmark()
}

func deleteBookmark(bookmarks map[string]string) {
	delete(bookmarks, inputKeyBookmark())
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

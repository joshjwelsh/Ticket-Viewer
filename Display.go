package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

const (
	INVALID_VALUE int = 99
	MAX_PAGE_SIZE int = 5
)

const (
	MENU_OPT_1 int = 1
	MENU_OPT_2 int = 2
	MENU_OPT_3 int = 3
)

func number(start int, end int) []int {
	nums := make([]int, MAX_PAGE_SIZE)
	for i := start + 1; i <= end; i++ {

		nums[(i-1)%MAX_PAGE_SIZE] = i
	}
	return nums
}

func display(page *Page) {
	funcMap := template.FuncMap{
		"number": number,
	}
	t, err := template.New("page.tmpl").Funcs(funcMap).ParseFiles("templates/page.tmpl")
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}
	err = t.ExecuteTemplate(os.Stdout, "page.tmpl", *page)
	if err != nil {
		log.Fatalf("Page template failed: %v", err)
	}
}

// ASCII encoding for cls
func clear() {
	fmt.Print("\033[H\033[2J")
}

func contViewAll() (int, error) {

	var input string
	fmt.Scanln(&input)
	if input == "1" {

		return strconv.Atoi(input)
	} else if input == "2" {

		return strconv.Atoi(input)
	} else if input == "3" {

		return strconv.Atoi(input)
	} else {
		return INVALID_VALUE, nil
	}
}

func Display(ticket []Ticket) {
	for {
		clear()
		menu(CreateMainMenu())
		choice, err := contMainMenu()
		if err != nil {
			log.Fatalf("Continue from main menu failed while returning a value: %v", err)
		}
		if choice == MENU_OPT_1 {
			clear()
			size := len(ticket)
			if size <= MAX_PAGE_SIZE {

				display(NewPage(ticket))
			} else {
				paginate(ticket)
			}
		}
		if choice == MENU_OPT_2 {

		}
		if choice == MENU_OPT_3 {
			break
		}

	}
}

func paginate(ticket []Ticket) {

	page := NewPage(ticket)
	for {
		display(page)
		menu(CreateViewAllMenu())
		choice, err := contViewAll()
		if err != nil {
			log.Fatalf("Continue from view all menu failed while returning a value: %v", err)
		}
		if choice == MENU_OPT_1 {
			page.PageForward()
		}
		if choice == MENU_OPT_2 {
			page.PageBack()
		}
		if choice == MENU_OPT_3 {
			return

		}
	}

}
func contMainMenu() (int, error) {

	var input string
	fmt.Scanln(&input)
	if input == "1" {
		return strconv.Atoi(input)
	} else if input == "2" {
		return strconv.Atoi(input)
	} else if input == "3" {
		return strconv.Atoi(input)
	} else {
		return INVALID_VALUE, nil
	}
}

func menu(m Menu) {
	t, err := template.ParseFiles("templates/menu.tmpl")
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}
	t.Execute(os.Stdout, m)
}

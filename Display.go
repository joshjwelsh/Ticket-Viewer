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

// handles display updates
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
			paginate(ticket)
		}
		if choice == MENU_OPT_2 {
			clear()
			selectOne(ticket)

		}
		if choice == MENU_OPT_3 {
			break
		}

	}
	fmt.Println("Exiting ticket viewer...Thanks for using it!")
}
func displayOne(ticket Ticket) {
	t, err := template.ParseFiles("templates/ticket.tmpl")
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}
	t.Execute(os.Stdout, ticket)

}

func selectOne(ticket []Ticket) {
	for {
		choice := selectOnePrompt(ticket)
		clear()
		page := NewPage(ticket)
		page.All = false
		page.Select = choice
		display(page)
		menu(CreateSelectMenu())
		choice, err := contViewSelect()
		if err != nil {
			log.Fatalf("Continue from view selected ticket failed while returning a value: %v", err)
		}
		if choice == MENU_OPT_3 {
			break
		}
	}

}

func selectOnePrompt(ticket []Ticket) int {
	fmt.Println("Enter a ticket number: ")
	var input string
	fmt.Scanln(&input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Invalid ticket selection. Enter a value between %v and %v.", 1, len(ticket))
	}
	if choice <= 0 {
		log.Fatalf("Invalid ticket selection. Enter a value between %v and %v.", 1, len(ticket))
	}
	if choice > len(ticket) {
		log.Fatalf("Invalid ticket selection. Enter a value between %v and %v, inclusively.", 1, len(ticket))
	}
	return choice

}

// lists all tickets to stdout
func display(page *Page) {
	funcMap := template.FuncMap{
		// template func to create ticket number
		"number": func(start int, end int) []int {
			nums := make([]int, MAX_PAGE_SIZE)
			for i := start + 1; i <= end; i++ {

				nums[(i-1)%MAX_PAGE_SIZE] = i
			}
			return nums
		},
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

//
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

// Get input from user
func contViewSelect() (int, error) {
	var input string
	fmt.Scanln(&input)
	if input == "" {
		return strconv.Atoi(input)
	} else if input == "2" {
		return INVALID_VALUE, nil
	} else if input == "3" {
		return strconv.Atoi(input)
	} else {
		return INVALID_VALUE, nil
	}
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

// ASCII encoding for clear
func clear() {
	const CLEAR_SCREEN string = "\033[H\033[2J"
	fmt.Print(CLEAR_SCREEN)
}

// pass menu template a menu struct and execute
func menu(m Menu) {
	t, err := template.ParseFiles("templates/menu.tmpl")
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}
	t.Execute(os.Stdout, m)
}

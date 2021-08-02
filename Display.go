package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
)

// handles display updates
func Display(stdin io.Reader, ticket []Ticket) error {
	for {
		clear()
		menu(CreateMainMenu())
		reader := CreateDevice(stdin)

		choice, err := cont(reader)
		if err != nil {
			return fmt.Errorf("Display crashed because of value returned by contMainMenu(): %v", err)
		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_1 {
			clear()
			reader = CreateDevice(stdin)
			page := NewPage(ticket)()
			paginate(&reader, page)
		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_2 {
			clear()
			reader = CreateDevice(stdin)
			err := selectOne(reader, ticket)
			if err != nil {
				return fmt.Errorf("Display crashed because of value returned by selectOne(): %v", err)
			}

		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_3 {
			fmt.Println("Exiting ticket viewer...Thanks for using it!")
			break
		}

	}
	return nil
}

func selectOne(reader ReadDevice, ticket []Ticket) error {

	for {

		choice, err := selectOnePrompt(reader, ticket)
		if err != nil {
			return fmt.Errorf("Error in selectOne: %v", err)
		}
		clear()
		page := NewPage(ticket)()
		page.All = false
		if ok := page.SelectOne(choice); ok == false {
			continue
		}
		display(page)
		menu(CreateSelectMenu())
		opt, err := cont(reader)
		if err != nil {
			return fmt.Errorf("Continue from view selected ticket failed while returning a value: %v", err)
		}
		if strings.TrimRight(opt, "\n") == MENU_OPT_3 {
			break
		}
	}
	return nil

}

func selectOnePrompt(reader ReadDevice, ticket []Ticket) (int, error) {
	var choice int
	for {
		fmt.Println("Enter a ticket number: ")
		// var input int
		err := reader.GetInput()
		log.Println("reader val:", reader.Input)
		if err != nil {
			return 0, fmt.Errorf("selectOnePrompt( ReadDevice, [] Ticket) returned an error: %v.", err)
		}
		// input, err :=

		choice, err = strconv.Atoi(strings.TrimRight(reader.Input, "\n"))
		if err != nil {
			fmt.Printf("Invalid input for selecting ticket.\n")
		}

		if choice <= 0 {
			fmt.Printf("Invalid ticket selection. Enter a value between %v and %v.", 1, len(ticket))
		}
		if choice > len(ticket) {
			fmt.Printf("Invalid ticket selection. Enter a value between %v and %v, inclusively.", 1, len(ticket))
		}
		break

	}
	return choice, nil
}

// lists all tickets to stdout
func display(page *Page) error {
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
		return fmt.Errorf("Error with template: %v", err)
	}
	err = t.ExecuteTemplate(os.Stdout, "page.tmpl", *page)
	if err != nil {
		return fmt.Errorf("Page template failed: %v", err)
	}
	return nil
}

// handle listing tickets on pages
func paginate(reader *ReadDevice, page *Page) {

	// page := NewPage(ticket)()
	for {

		display(page)

		menu(CreateViewAllMenu())
		choice, err := cont(*reader)

		if err != nil {
			log.Fatalf("Continue from view all menu failed while returning a value: %v", err)
		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_1 {
			page.PageForward()
		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_2 {
			page.PageBack()
		}
		if strings.TrimRight(choice, "\n") == MENU_OPT_3 {
			break

		}
	}
	return

}
func cont(reader ReadDevice) (string, error) {

	err := reader.GetInput()
	
	if err != nil {
		return "", fmt.Errorf("cont(ReadDevice) got error from reader: %v", err)
	}
	return reader.Input, nil
}

// clear terminal
func clear() bool {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

// pass menu template a menu struct and execute
func menu(m Menu) {
	t, err := template.ParseFiles("templates/menu.tmpl")
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}
	t.Execute(os.Stdout, m)
}

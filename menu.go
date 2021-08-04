package main

import (
	"log"
	"os"
	"text/template"
)

// -------------------------------------- Represents menu ------------------------------------

type Menu struct {
	Main    bool
	ContAll bool
	ViewOne bool
	Opt1    string
	Opt2    string
	Opt3    string
}

func CreateMainMenu() Menu {
	return Menu{
		Main: true,
		Opt1: "1)\tList all Tickets",
		Opt2: "2)\tView a ticket",
		Opt3: "3)\tExit",
	}
}

func CreateViewAllMenu() Menu {
	return Menu{
		ContAll: true,
		Opt1:    "1)\tNext Page",
		Opt2:    "2)\tPrevious Page",
		Opt3:    "3)\tReturn to main menu",
	}
}

func CreateSelectMenu() Menu {
	return Menu{
		ViewOne: true,
		Opt1:    "1)\tSelect another ticket",
		Opt3:    "3)\tReturn to main menu",
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

// -------------------------------------------------------------------------------------------------------------------

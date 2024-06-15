package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	showMenu()
}

/*
sortowanie uczniow
*/
func application() {
	fmt.Println(len(students))
	if len(students) <= 10 {
		var tempname, tempsurname, tempaddress string
		var tempage int
		fmt.Println("jak się nazywasz? (imie nazwisko)")
		fmt.Scanln(&tempname, &tempsurname)
		fmt.Println("jak masz adres?")
		fmt.Scanln(&tempaddress)
		fmt.Println("jaki jest twoj wiek?")
		for {
			var err error
			var input string

			fmt.Scanln(&input)

			tempage, err = strconv.Atoi(input)
			if err == nil && tempage > 0 && tempage < 100 {
				break
			}
			fmt.Println("prosze o wprowadzenie wieku (1-99)")
		}

		pkt := 0.0
		pkt += float64(examValidate("podaj wynik z egz POL")) * 0.35
		pkt += float64(examValidate("podaj wynik z egz MAT")) * 0.35
		pkt += float64(examValidate("podaj wynik z egz ANG")) * 0.30
		pkt += ifNumber("twoja ocena z polskiego")
		pkt += ifNumber("twoja ocena z przedmiotu nr 1")
		pkt += ifNumber("twoja ocena z przedmiotu nr 2")
		pkt += ifNumber("twoja ocena z przedmiotu nr 3")
		students = append(students, student{name: tempname, surname: tempsurname, address: tempaddress, age: tempage, points: int(pkt)})
	} else {
		fmt.Println("przekroczono max liczbe uczniow w szkole")
	}
}

func ifNumber(tekst string) float64 {
	ocena := [7]float64{0, 0, 2, 8, 14, 17, 18}
	var o int
	var input string
	var err error
	for {
		fmt.Println(tekst)
		fmt.Scanln(&input)
		o, err = strconv.Atoi(input)
		if o >= 0 && o <= 6 {
			if err == nil {
				return ocena[o]
			}
		}
		fmt.Println("prosze o wprowadzenie liczby calkowitej od 1 do 6")
	}
}

func examValidate(tekst string) int {
	var temppkt int
	fmt.Println(tekst)
	var input string
	var err error
	for {
		fmt.Scanln(&input)
		temppkt, err = strconv.Atoi(input)
		if temppkt > 0 && temppkt < 101 {
			if err == nil {
				return temppkt
			}
		}
		fmt.Println("prosze o wprowadzenie liczby calkowitej od 0 do 100")
	}
}

func showStudents() {
	for _, student := range students {
		fmt.Printf("imie: %v, nazwisko: %v, adres: %v, wiek: %v, punkty: %v\n", student.name, student.surname, student.address, student.age, student.points)
	}
	keyboard.Open()
	keyboard.GetKey()
	defer keyboard.Close()
}

func showMenu() {
	for {
		clearScreen()
		fmt.Println("***********")
		fmt.Println("[L]ista osob")
		fmt.Println("[N]abor")
		fmt.Println("[W]yjdz")
		fmt.Println("***********")
		fmt.Printf("gdzie chcesz isc??? N/L/W\n")
		var i string
		fmt.Scanln(&i)

		if i == "w" || i == "W" {
			break
		}
		if i == "L" || i == "l" {
			showStudents()
		}
		if i == "N" || i == "n" {
			application()
		}

	}
}

type student struct {
	name    string
	surname string
	address string
	age     int
	points  int
}

var students = []student{
	{
		name:    "Jan",
		surname: "Kowalski",
		address: "Warszawa",
		age:     20,
		points:  150,
	},
	{
		name:    "Anna",
		surname: "Nowak",
		address: "Kraków",
		age:     22,
		points:  180,
	},
	{
		name:    "Piotr",
		surname: "Wiśniewski",
		address: "Gdańsk",
		age:     21,
		points:  170,
	},
	{
		name:    "Katarzyna",
		surname: "Wójcik",
		address: "Poznań",
		age:     19,
		points:  160,
	},
	{
		name:    "Michał",
		surname: "Lewandowski",
		address: "Wrocław",
		age:     23,
		points:  140,
	},
}

func clearScreen() {
	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

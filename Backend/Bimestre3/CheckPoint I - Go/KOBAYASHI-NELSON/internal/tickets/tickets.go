package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	manha = "MORNING"
	tarde = "AFTERNOOM"
	noite = "NIGHT"
	madrugada = "COCKCROW"
)

type Ticket struct {
	id int
	nome string
	email string
	destino string
	partida string
	preco float64
}

/* 
Requirement 1:
Performs total count of one day per country 
*/

func GetTotalTickets(destino string) (int, error) {
	total := 0
	
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file")
	}

	tickets := strings.Split(string(res), "\n")

	for i := 0; i < len(tickets); i++ {
		atributos := strings.Split(tickets[i], ",")

		if atributos[3] == destino {
			total++
		}
	}

	return total, nil
	
}

/* 
Requirement 2:
Calculate the total trips made during a period of the day 
*/

func GetCountByHour(partida string) (string, error) {
	total := 0

	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file")
	}

	tickets := strings.Split(string(res), "\n")

	for i := 0; i < len(tickets); i++{
		atributos := strings.Split(tickets[i], ",")
		horario, err := strconv.Atoi(strings.Split(atributos[4], ":")[0])
		if err != nil {
			panic("err")
		}
		
		switch partida {
		case "manha":
			if horario >= 07 && horario <= 12 {
				total++
			}
		case "tarde":
			if horario >= 13 && horario <= 18 {
				total++
			}
		case "noite":
			if horario >= 19 && horario <= 23 {
				total++
			}
		case "madrugada":
			if horario >= 00 && horario <= 06 {
				total++
			}
		}
	}
	message := fmt.Sprintf("Total number of passengers during the period %s: %d", partida, total)
	return message, nil
}


/* 
Requirement 3:
Calculate the average number of people traveling to a given country.
*/

func AvgDistinctDestinations() (string) {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic("Could not read file.")
	}

	tickets := strings.Split(string(res), "\n")
	distinctDestinations := []string{}

	arrayContains := make(map[string]bool)

	for i := 0; i < len(tickets); i++ {
		attributes := strings.Split(tickets[i], ",")
		destination := attributes[3]

		if _, value := arrayContains[destination]; !value {
			arrayContains[destination] = true
			distinctDestinations = append(distinctDestinations, destination)
		}		
	}

	totalDistinctDestinations := len(distinctDestinations)


	totalTickets := len(tickets)
	var averageTicketsPerDestinations int
	if totalDistinctDestinations != 0 {
		averageTicketsPerDestinations = totalTickets / totalDistinctDestinations
	}
	
	message := fmt.Sprintf("Average tickets per destinations: %d", averageTicketsPerDestinations)

	return  message
}


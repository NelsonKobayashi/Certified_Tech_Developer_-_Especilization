package main

import (
	"fmt"
	"github.com/NelsonKobayashi/Certified_Tech_Developer_-_Especilization/Backend/Bimestre3/CheckPoint I - Go/KOBAYASHI-NELSON/internal/tickets"	
)

func main() {
	totalTickets, _ := tickets.GetTotalTickets("Argentina")
	fmt.Println(totalTickets)

	totalPerHour, _ := tickets.GetCountByHour("noite")
	fmt.Println("noite: ", totalPerHour)

	avgDistinctDestinations, _ := tickets.AvgDistinctDestinations()
	fmt.Println("Average of distinct destinations: ", avgDistinctDestinations)
	
}




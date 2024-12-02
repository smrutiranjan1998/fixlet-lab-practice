package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Gixlet struct {
	SiteID                string
	FxiletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

var gixlets []Gixlet

func loadCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		relevantComputerCount, _ := strconv.Atoi(record[4])
		gixlets = append(gixlets, Gixlet{
			SiteID:                record[0],
			FxiletID:              record[1],
			Name:                  record[2],
			Criticality:           record[3],
			RelevantComputerCount: relevantComputerCount,
		})
	}

	return nil
}

func listEntries() {
	fmt.Println("Listing all entries:")
	fmt.Println("SiteID, FxiletID, Name, Criticality, RelevantComputerCount")
	for _, gixlet := range gixlets {
		fmt.Printf("%s, %s, %s, %s, %d\n", gixlet.SiteID, gixlet.FxiletID, gixlet.Name, gixlet.Criticality, gixlet.RelevantComputerCount)
	}
}

func queryEntries(name string) {
	fmt.Printf("Querying entries for Name: %s\n", name)
	for _, gixlet := range gixlets {
		if gixlet.Name == name {
			fmt.Printf("%s, %s, %s, %s, %d\n", gixlet.SiteID, gixlet.FxiletID, gixlet.Name, gixlet.Criticality, gixlet.RelevantComputerCount)
		}
	}
}

func sortEntries(by string) {
	switch by {
	case "SiteID":
		sort.Slice(gixlets, func(i, j int) bool { return gixlets[i].SiteID < gixlets[j].SiteID })
	case "Criticality":
		sort.Slice(gixlets, func(i, j int) bool { return gixlets[i].Criticality < gixlets[j].Criticality })
	case "RelevantComputerCount":
		sort.Slice(gixlets, func(i, j int) bool { return gixlets[i].RelevantComputerCount < gixlets[j].RelevantComputerCount })
	}
	fmt.Println("Entries sorted by", by)
	listEntries()
}

func addEntry(siteID, fxiletID, name, criticality string, relevantComputerCount int) {
	gixlets = append(gixlets, Gixlet{
		SiteID:                siteID,
		FxiletID:              fxiletID,
		Name:                  name,
		Criticality:           criticality,
		RelevantComputerCount: relevantComputerCount,
	})
	fmt.Println("Entry added successfully!")
}

func deleteEntry(fxiletID string) {
	for i, gixlet := range gixlets {
		if gixlet.FxiletID == fxiletID {
			gixlets = append(gixlets[:i], gixlets[i+1:]...)
			fmt.Println("Entry deleted successfully!")
			return
		}
	}
	fmt.Println("Entry not found!")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go")
		return
	}

	filePath := os.Args[1]

	if err := loadCSV(filePath); err != nil {
		fmt.Printf("Error loading CSV: %v\n", err)
		return
	}

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. List entries")
		fmt.Println("2. Query entries by Name")
		fmt.Println("3. Sort entries")
		fmt.Println("4. Add an entry")
		fmt.Println("5. Delete an entry")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			listEntries()
		case 2:
			fmt.Print("Enter Name to query: ")
			var name string
			fmt.Scan(&name)
			queryEntries(name)
		case 3:
			fmt.Print("Enter field to sort by (SiteID, Criticality, RelevantComputerCount): ")
			var field string
			fmt.Scan(&field)
			sortEntries(field)
		case 4:
			fmt.Print("Enter SiteID, FxiletID, Name, Criticality, RelevantComputerCount: ")
			var siteID, fxiletID, name, criticality string
			var relevantComputerCount int
			fmt.Scan(&siteID, &fxiletID, &name, &criticality, &relevantComputerCount)
			addEntry(siteID, fxiletID, name, criticality, relevantComputerCount)
		case 5:
			fmt.Print("Enter FxiletID to delete: ")
			var fxiletID string
			fmt.Scan(&fxiletID)
			deleteEntry(fxiletID)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

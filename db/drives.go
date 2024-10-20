package db

import "time"

func parseDateTime(dateTimestr string) time.Time {
	dateTime, _ := time.Parse("2006-01-02 15:05", dateTimestr)
	return dateTime
}

type Drive struct {
	Name          string
	Address       string
	Customer      Customer
	ScheduledDate time.Time
}

var ScheduledDrives = []Drive{
	{
		Name:    "E-Waste Drive A",
		Address: "123 Green Street, Pune",
		Customer: Customer{
			Name:  "Shreyas Mali",
			Email: "shreyas@email.com",
			Phone: "0000000000",
		},
		ScheduledDate: parseDateTime("2024-11-01 10:00"),
	},
	{
		Name:          "E-Waste Drive B",
		Address:       "456 Blue Road, Mumbai",
		Customer: Customer{
			Name:  "Siddesh Narshinkar",
			Email: "siddesh@email.com",
			Phone: "0000000000",
		},
		ScheduledDate: parseDateTime("2024-12-15 14:30"),
	},
	{
		Name:          "E-Waste Drive C",
		Address:       "789 Red Lane, Bangalore",
		Customer: Customer{
			Name:  "Samiksha Pandit",
			Email: "smiksha@email.com",
			Phone: "0000000000",
		},
		ScheduledDate: parseDateTime("2025-01-10 09:45"),
	},
}

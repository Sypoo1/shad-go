//go:build !solution

package hotelbusiness

import (
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	load := []Load{}
	if len(guests) == 0 {
		return load
	}
	events := make(map[int]int)
	for _, guest := range guests {
		events[guest.CheckInDate]++
		events[guest.CheckOutDate]--
	}

	dates := []int{}
	for date :=  range events {
		dates = append(dates, date)
	}
	slices.Sort(dates)
	cnt := 0
	for _, date := range dates {
		if events[date] == 0 {
			continue
		}
		cnt += events[date]
		load = append(load, Load{date, cnt})
	}


	return load
}

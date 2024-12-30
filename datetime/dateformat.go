package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Formatted time is %s\n", start.Format(time.RFC3339))
	//TODO: custom formatted date time and use it in date functions.
	time.Sleep(2 * time.Second)
	fmt.Printf("time elapsed is %s\n", time.Now().Sub(start))
	fmt.Printf("time after 10 days is %s\n", start.Add(time.Hour*24*10))
	//Date function to receive yyyy-mm-dd H:i:s format
	t, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic("error loading location")
	}
	fmt.Printf("Location is %s\n", t)
	year1982 := time.Date(1982, 4, 13, 1, 20, 0, 0, t)
	year2024 := time.Date(2024, 4, 13, 1, 20, 0, 0, t)
	fmt.Printf("is Date after %s = %b\n", year1982.Format(time.UnixDate), year2024.After(year1982))
	fmt.Printf("DateOnly %s\n", year1982.Format(time.DateOnly))
	fmt.Printf("DateTime is %s\n", year1982.Format(time.DateTime))
	fmt.Printf("Kitchen (i.e. time only) is %s\n", year1982.Format(time.Kitchen))
	year, month, date := year1982.Date()
	fmt.Printf("fetch date\nYear:%d\nMonth:%d\nDay: %d\n", year, month, date)
	fmt.Printf("Age is %v\n", time.Since(year1982).Hours()/24/365)
}

package main

import (
	"fmt"
	"time"
	"log"
)

func main(){
	now := time.Now()
	timeStr := now.Format(time.RFC1123)
    fmt.Println(timeStr)
	fmt.Println(now)

	//Sun 19 Sep 2021 15:42:00 MSK

	currentTimeStr := "2021-09-19T15:59:41+03:00"
    
    layout := time.RFC3339 
	fmt.Println("после парсинга")
	fmt.Println(layout)
    currentTime, err := time.Parse(layout, currentTimeStr)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("после ошибки")
    fmt.Println(currentTime)

	

    fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
    fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
    fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))

	truncTime := now.Truncate(24 * time.Hour)
    fmt.Printf("Округление времени до начала дня: %v\n", truncTime)


	birthday := time.Date(1993, 11, 26, 0, 0, 0, 0, time.Local)
	birthday_100 := time.Date(2093, 11, 26, 0, 0, 0, 0, time.Local)
    days := birthday_100.Sub(birthday)
    fmt.Println(days.Hours() / 24)
	// Андрей родился 26 ноября 1993 года.

	
	
}
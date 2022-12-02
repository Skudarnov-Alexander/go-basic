package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Video struct {
    Id    string
    Title string
    Views int64
}

type (
    VideoB struct {
        Id   string
        Name string
        Tags Tags
    }

    Tags []string
) 

type Trend struct{
    T time.Time
    Count int
}
func TrendingCount(db *sql.DB) ([]Trend, error) {
    // проверяем на всякий случай
    if db == nil {
        return nil, errors.New("you haven`t open the database connection")
    }
    rows, err := db.Query("SELECT trending_date, COUNT(trending_date) FROM videos GROUP BY trending_date")
    if err != nil {
        return nil, err
    }
    trends := make([]Trend, 0)
    date := new(string)
    for rows.Next() {
        trend := Trend{}
        err = rows.Scan(date, &trend.Count)
        if err != nil {
            return nil, err
        }
        if trend.T , err = time.Parse("06.02.01", *date); err != nil {
            return nil, err
        }
        trends = append(trends, trend)
    }
    return trends, nil
} 

func (tags *Tags) Value() (driver.Value, error) {
    // преобразуем []string в string
    if len(*tags) == 0 {
        return "", nil
    }
    return strings.Join(*tags, "|"), nil
} 

func (tags *Tags) Scan(value interface{}) error {
    if value == nil {
        *tags = Tags{}
        return nil
    }

    sv, err := driver.String.ConvertValue(value)
    if err != nil {
        return fmt.Errorf("cannot scan value. %w", err)
    }

    v, ok := sv.(string)
    if !ok {
        return errors.New("cannot scan value. cannot convert value to string")
    }

    *tags = strings.Split(v, "|")

    return nil
} 

const limit = 20 

func QueryVideos(ctx context.Context, db *sql.DB, limit int) ([]Video, error) {
    videos := make([]Video, 0, limit)

    rows, err := db.QueryContext(ctx, "SELECT video_id, title, views from videos ORDER BY views LIMIT ?", limit)
    if err != nil {
        return nil, err
    }

    // обязательно закрываем перед возвратом функции
    defer rows.Close()

    // пробегаем по всем записям
    for rows.Next() {
        var v Video
        err = rows.Scan(&v.Id, &v.Title, &v.Views)
        if err != nil {
            return nil, err
        }

        videos = append(videos, v)
    }

    // проверяем на ошибки
    err = rows.Err()
    if err != nil {
        return nil, err
    }
    return videos, nil
}

func main() {
	fmt.Println("Старт программы")

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "../db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Инстанс sql.DB создан")

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("PING DB is OK")

	var id int64

	row := db.QueryRowContext(ctx, "SELECT COUNT(*) as count FROM videos")

	if err := row.Scan(&id); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

	var (
		s sql.NullString
		SomeID string = "9wRQljFNDW8"
	)

	err = db.QueryRow("SELECT title FROM videos WHERE video_id = ?", SomeID).Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	
	if s.Valid {
		fmt.Println(s.String)
	} else {
		fmt.Println("Title not set")
	}

	result, err := QueryVideos(ctx, db, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %+v\n", result)

	trends, err := TrendingCount(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("trends: %v\n", trends)

	






}

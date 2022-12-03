package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type (
	Video struct {
		Id,
		Title,
		Channel,
		Description,
		CategoryId string

		TrendingDate string
		PublishTime  time.Time

		Tags []Tag

		Views    int
		Likes    int
		Dislikes int
		Comments int

		CommentsDisabled    bool
		RatingDisabled      bool
		VideoErrorOrRemoved bool

		ThumbnailLink string
	}

	Tag struct {
		Name string
	}
)

func main() {
	// инициируем контекст
	ctx := context.Background()
	_ = ctx
	// открываем соединение с БД
	db, err := sql.Open("sqlite3", "../db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	// открываем csv-файл
	file, err := os.Open("../files/USvideos.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = createTable(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	// конструируем Reader из пакета encoding/csv
	// он умеет читать строки csv-файла
	csvReader := csv.NewReader(file)

	// читаем записи из файла в слайс []Video вспомогательной функцией

	videos, err := readVideos(csvReader)
	_ = videos
	if err != nil {
		log.Fatal(err)
	}

	// записываем []Video в базу данных в инициированном контексте
	// тоже вспомогательной функцией
	startTimer := time.Now()
	err = insertVideos(ctx, db, videos)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(startTimer)
	fmt.Printf("Время выполнения %d\n", duration.Milliseconds()) 
	fmt.Printf("Всего csv-записей %v\n", len(videos))

}

func createTable(ctx context.Context, db *sql.DB) error {
	quary := `
	CREATE TABLE videos(
		"video_id" TEXT,
		"trending_date" TEXT,
		"title" TEXT,
		"channel_title" TEXT,
		"category_id" INTEGER,
		"publish_time" TEXT,
		"tags" TEXT,
		"views" INTEGER,
		"likes" INTEGER,
		"dislikes" INTEGER,
		"comment_count" INTEGER,
		"thumbnail_link" TEXT,
		"comments_disabled" BOOLEAN,
		"ratings_disabled" BOOLEAN,
		"video_error_or_removed" BOOLEAN,
		"description" TEXT
	  ); `

	  _, err := db.ExecContext(ctx, quary)
	  if err != nil {
		return err
	  }

	  return nil

}

func readVideos(r *csv.Reader) ([]Video, error) {
	var videos []Video
	for {
		// csv.Reader за одну операцию Read() считывает одну csv-запись
		l, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// инициализируем целевую структуру,
		// в которую будем делать разбор csv-записи
		v := Video{
			Id:            l[0],
			TrendingDate:  l[1],
			Title:         l[2],
			Channel:       l[3],
			CategoryId:    l[4],
			ThumbnailLink: l[11],
			Description:   l[15],
		}
		// парсинг строковых записей в типизированные поля структуры
		if v.PublishTime, err = time.Parse(time.RFC3339, l[5]); err != nil {
			continue
		}
		tgs := strings.Split(l[6], " ")
		for _, tg := range tgs {
			v.Tags = append(v.Tags, Tag{tg})
		}
		if v.Views, err = strconv.Atoi(l[7]); err != nil {
			continue
		}
		if v.Likes, err = strconv.Atoi(l[8]); err != nil {
			continue
		}
		if v.Dislikes, err = strconv.Atoi(l[9]); err != nil {
			continue
		}
		if v.Comments, err = strconv.Atoi(l[10]); err != nil {
			continue
		}
		if v.CommentsDisabled, err = strconv.ParseBool(l[12]); err != nil {
			continue
		}
		if v.RatingDisabled, err = strconv.ParseBool(l[13]); err != nil {
			continue
		}

		if v.VideoErrorOrRemoved, err = strconv.ParseBool(l[14]); err != nil {
			continue
		}

		// добавляем полученную структуру в слайс
		videos = append(videos, v)
	}
	return videos, nil
}

func insertVideos(ctx context.Context, db *sql.DB, videos []Video) error {
    // шаг 1 — объявляем транзакцию
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    // шаг 1.1 — если возникает ошибка, откатываем изменения
    defer tx.Rollback()

    // шаг 2 — готовим инструкцию
    stmt, err := tx.PrepareContext(ctx, "INSERT INTO videos(title, description, views, likes) VALUES(?,?,?,?)")
    if err != nil {
        return err
    }
    // шаг 2.1 — не забываем закрыть инструкцию, когда она больше не нужна
    defer stmt.Close()

    for _, v := range videos {
        // шаг 3 — указываем, что каждое видео будет добавлено в транзакцию
        if _, err = stmt.ExecContext(ctx, v.Title, v.Description, v.Views, v.Likes); err != nil {
            return err
        }
    }
    // шаг 4 — сохраняем изменения
    return tx.Commit()
} 

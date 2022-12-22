package main

import (
	"fmt"
	"runtime"
	"time"
)

type Size struct {
	Width  uint32
	Height uint32
}

type Task struct {
	Filename string
	ToSize   Size
}

type Queue struct {
	ch chan *Task
}

func NewQueue() *Queue {
	q := Queue{}
	q.ch = make(chan *Task)
	return &q
}

func (q *Queue) Push(t *Task) {
	q.ch <- t
}

type Resizer struct {
}

func NewResizer() *Resizer {
	r := Resizer{}
	return &r
}

func (r *Resizer) Resize(filename string, toSize Size) error {
	// пропустим реализацию
	return nil
}

type Worker struct {
	id      int
	queue   *Queue
	resizer *Resizer
}

func NewWorker(id int, queue *Queue, resizer *Resizer) *Worker {
	w := Worker{
		id:      id,
		queue:   queue,
		resizer: resizer,
	}
	return &w
}

func (w *Worker) Loop() {
	for {
		t := <-w.queue.ch

		err := w.resizer.Resize(t.Filename, t.ToSize)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker #%d resized %s\n", w.id, t.Filename)
	}
}

func main() {
	queue := NewQueue()
	workers := make([]*Worker, 0, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		workers = append(workers, NewWorker(i, queue, NewResizer()))
	}

	for _, w := range workers {
		go w.Loop()
	}

	filenames := []string{"gopher.jpg", "test.png"}
	for _, f := range filenames {
		queue.Push(&Task{Filename: f, ToSize: Size{Width: 1024, Height: 1024}})
	}

	time.Sleep(1 * time.Second)
}

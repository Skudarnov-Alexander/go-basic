package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func healthCheck(url string, errCh chan<- error, wg *sync.WaitGroup, stopCh <-chan struct{}) {
    var defErr error
    defer func() {
        if defErr != nil {
            select {
            // первая горутина, поймавшая ошибку, сможет записать в канал
            case errCh <- defErr: 
            // остальные завершат работу, провалившись в этот case
            case <-stopCh: 
                log.Println("aborting", url)
            }
        }
        wg.Done()
    }()

    resp, err := http.Get(url)
    if err != nil {
        defErr = fmt.Errorf("healthcheck failed: %w", err)
        return
    }
    if resp.StatusCode != http.StatusOK {
        defErr = errors.New("healthcheck failed: status not ok")
        return
    }
}

func main() {
    wg := &sync.WaitGroup{}
    errCh := make(chan error)
    stopCh := make(chan struct{})
    // делаем сигнальный канал, но он будет работать иначе
    // горутины, которые нужно остановить, будут заблокированы на нём
    // если понадобится их завершить, просто вызовем close(stopCh)

    hostsToCheck := []string{
        "https://yan2dex.ru",
        "https://eda.yan5dex.ru",
        "https://lavka.yan4dex.ru",
    }
    for _, hostToCheck := range hostsToCheck {
        log.Println("checking", hostToCheck)
        wg.Add(1)
        go healthCheck(hostToCheck, errCh, wg, stopCh)
    }

    // в отдельной горутине ждём завершения всех healthCheck
    // после этого закрываем канал errCh — больше записей не будет
    go func() {
        wg.Wait()
        close(errCh)
    }()

    if err := <-errCh; err != nil {
        log.Println(err)
        close(stopCh)
		time.Sleep(2 * time.Second)
        return
    }

    log.Println("successful healthcheck")
	time.Sleep(2 * time.Second)
}
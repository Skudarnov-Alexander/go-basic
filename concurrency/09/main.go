package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		proofOfWorkNumber int64 = 1337
		probes            int64
		result            int64
	)

	var errFound = errors.New("found")

	g, ctx := errgroup.WithContext(context.Background())
	workers := 100

	for i := 0; i < workers; i++ {
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					seed := atomic.AddInt64(&probes, 1)
					source := rand.NewSource(seed)

					number := rand.New(source).Int63()
					if number%proofOfWorkNumber == 0 && number != 0 {
						atomic.StoreInt64(&result, number)
						return errFound
					}
				}

			}

			//return nil
		})
	}

	g.Wait()

	fmt.Printf("Found %v at %v probes\n", result, probes)

}

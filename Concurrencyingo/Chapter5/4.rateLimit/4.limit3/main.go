package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

// 多重限速
// 读文件同时受到api和文件读取限制
// 解析地址同时受到api和解析地址限制

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); l != nil {
			return err
		}
	}
	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

// 每次操作之间的间隔
func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

type APIConnection struct {
	networkLimit RateLimiter
	diskLimit    RateLimiter
	apiLimit     RateLimiter
}

func Open() *APIConnection {
	secondLimit := rate.NewLimiter(Per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(Per(10, time.Minute), 10)
	return &APIConnection{
		apiLimit:     MultiLimiter(secondLimit, minuteLimit),
		diskLimit:    MultiLimiter(rate.NewLimiter(rate.Limit(1), 1)),
		networkLimit: MultiLimiter(rate.NewLimiter(Per(3, time.Second), 3)),
	}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}
	// 假设我们在执行一些逻辑

	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}
	// 假设我们在执行一些逻辑

	return nil
}

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}

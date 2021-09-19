package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type myServer struct{}

func (server myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func signalRec() error {
	c := make(chan os.Signal)

	signal.Notify(c)
	s := <-c
	fmt.Println("catch system signal: ", s)
	switch s {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		return errors.New("system signal exits")
	default:
		fmt.Println("other signal")
	}
	return nil
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	var ms myServer
	s := http.Server{
		Addr:    ":9090",
		Handler: ms,
	}
	http.Handle("/", ms)

	group.Go(func() error {
		defer fmt.Println("server goroutine stop")
		fmt.Println("listen and server")
		return s.ListenAndServe()
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("shutdown server goroutine stop")
			return s.Shutdown(ctx)
		}
	})

	group.Go(func() error {
		err := signalRec()
		if err != nil {
			fmt.Println("signal goroutine stop")
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("all goroutines exit, error occurs: %v", err)
	}
}

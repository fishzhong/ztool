package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

type User struct {
	Name string
	Age  int
}

func main() {
	log.Println(">>>>> BEGIN")
	g, ctx := errgroup.WithContext(context.Background())
	fun2 := func() error {
		return doSomething2(ctx, 2)
	}
	g.Go(fun2)
	var user User
	fun4 := func() error {
		return doSomething4(ctx, &user)
	}
	g.Go(fun4)

	if err := g.Wait(); err != nil {
		log.Println(">>>>> (!!)Something goes wrong")
		fmt.Println(user)
		return
	}

	log.Println("Successfully done all jobs.")
	log.Println(">>>>> END")
}

func doSomething2(ctx context.Context, id int) error {
	log.Println(">>>>> START", id)
	var flag = make(chan error)
	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second)
			log.Println("worker2:", i)
		}
		log.Println("flag111")
		flag <- errors.New("测试粗无了")
		close(flag)
		log.Println("flag333")
	}()
	var err error
	select {
	case <-ctx.Done():
		//log.Println(ctx.Err().Error())
		err = fmt.Errorf("stopped")
	case err = <-flag:
		log.Println("flagErr", err)
	}
	if err != nil {
		log.Println("stopped by error!", err)
		return err
	}
	log.Println(">>>>> DONE", err)
	return nil
}

func doSomething4(ctx context.Context, user *User) error {
	user.Age = 4
	user.Name = "李四"
	log.Println(">>>>> START", 4)
	time.Sleep(10 * time.Second)
	log.Println(">>>>> ERROR HAPPENED", 4)
	return fmt.Errorf("error here")
}

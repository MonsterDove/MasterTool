package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
)

func main() {
	l := loga.ZapSugaredLogger()

	l.Info("starts")
	i := 0
	start := time.Now()
	last := time.Now()
	f := func() error {
		i++
		l.Debugw("apple", "index", i, "past", time.Since(last))
		last = time.Now()
		return errors.New("oh no")
	}
	f2 := func() error {
		i++
		l.Debugw("banana", "index", i, "past", time.Since(last))
		last = time.Now()
		return nil
	}

	// err := loga.FixedRetry(f, 10, 200*time.Millisecond)
	// err := loga.SimpleRetry(f)
	err := loga.BackOffRetry(f, 5, 200*time.Millisecond)
	l.Infow("done1", "total", time.Since(start), zap.Error(err))

	err = loga.BackOffRetry(f2, 5, 1000*time.Millisecond)
	l.Infow("done2", "total", time.Since(start), zap.Error(err))
}

func main2() {
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Printf("🌋: Hello World!\n💻: %s\n📂: %s\n⏰: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))
	fmt.Println(loga.Quote())

	name1 := "text1.txt"
	lines := []string{
		"√ Arts",
		"❤️",
	}

	fmt.Println("W1 Err:", loga.WriteLines(name1, lines))
	fmt.Println("W1 Err:", loga.WriteLines(name1, lines))

	name2 := "more.txt"
	fmt.Println("W2 Err:", loga.AppendLines(name2, lines))
	fmt.Println("W2 Err:", loga.AppendLines(name2, lines))
}

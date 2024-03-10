package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	fmt.Printf("Número de CPUs:       %d\n", runtime.NumCPU())
	fmt.Printf("Número de Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Sistema Operacional:  %s\n", runtime.GOOS)
	fmt.Printf("Arquitetura:   %s\n", runtime.GOARCH)
	fmt.Printf("Versão do Go:  %s\n", runtime.Version())
	fmt.Printf("Memória Usada: %v\n\n\n", mem.TotalAlloc)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo", 100)
		waitGroup.Done()

	}()

	go func() {
		escrever("Programando em Go", 2000)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	fmt.Println("----------------- FIM ----------------")
}

func escrever(texto string, milli uint16) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * time.Duration(milli))
	}
}

//https://www.youtube.com/watch?v=yNOe3STbtGE
package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			out1 <- "order processed"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "refund processed"
		}
	}()
	//siempre va a ir imprimiendo una order y un refund
	for {
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
		fmt.Println(<-out1)
		fmt.Println(<-out2)
	}
	//	var wg sync.WaitGroup
	//estamos diciendo que hay un proceso al que se debe esperar
	//	wg.Add(1)
	//go process("order")
	//process("order")
	//está tan ocupado con order que refund nunca se procesa
	//solución? ejercutar procesos en forma concurrente?
	//go tiene soporte para la concurrencia a través de las go routines
	// qué pasa si tenemos go process("order") y go process("refund")?
	//las llamadas se bloquean entre sí, necesitamos algo que obligue a las go routines a esperar por turnos
	//sincronizar
	//función anónima simil a los lambda en javascript(qué carajos? shit habrá que buscar) > agregar a TODO
	//se declara una función sin nombre que se ejecuta inmediatamente y al tener la keyword go delante
	//	go func() {
	//		process("order")
	//cuando el proceso termina le dice al waitgroup 	que está listo
	///		wg.Done()
	//	}()
	//go process("refund")
	//espera la señal del enter
	//fmt.Scanln()
	//block
	//	wg.Wait()
}

//declarar atributo channel y el tipo de dato que va a comunicar
func process(item string, out chan string) {
	//que se cierre solo después de terminar todas las instruccione
	defer close(out)
	for i := 1; true; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
		//enviar dato item al canal
		out <- item
	}
}

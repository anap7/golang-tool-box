package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func()  {
		for {
			//Delay de meio segundo
			time.Sleep(time.Millisecond * 500)
			channel1 <- "The Channel 01"
		}
	}()

	go func()  {
		for {
			//Delay de 2 segundos
			time.Sleep(time.Second * 2)
			channel2 <- "The Channel 02"
		}
	}()

	for {
		/*O select ele é igual ao switch, porém utilizado somente para concorrência,
		nesse caso estamos utilizando para priorizar as mensagens que chegam primeiro
		e expor. Isso impede que o canal1 espere o canal2 para receber a mensagem, tirando
		a dependecia do outro 
		
		O canal1 postará mensagens a cada meio segundo, o canal 2 a cada 2 segundos e para
		que não seja necessário o canal 1 esperar o canal 2 para continuar recebendo mensagens
		usamos o select para priorizar quem JÁ está recebendo as mensagens, isso deixando ambos
		os canais independentes*/
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
}

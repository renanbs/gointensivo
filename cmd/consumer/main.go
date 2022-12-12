package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/renanbs/gointensivo/internal/order/infra/database"
	"github.com/renanbs/gointensivo/internal/order/usecase"
	"github.com/renanbs/gointensivo/pkg/rabbitmq"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("Teste")
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery) // channel
	forever := make(chan bool)
	go rabbitmq.Consume(ch, out) // T2

	qtdWorkers := 50
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}

	<-forever
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	for msg := range deliveryMessage {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		//fmt.Println(outputDTO)
		//response, err := json.Marshal(outputDTO)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//log.Info(string(response))
		log.Infof("Worker %d has processed order %s\n", workerID, outputDTO.ID)
		time.Sleep(1 * time.Second)
	}
}

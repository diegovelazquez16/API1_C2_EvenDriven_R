package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"holamundo/historial/aplication/usecase"
	"holamundo/historial/domain/models"
)

type HistorialConsumer struct {
	HistorialUC *usecase.CreateHistorialUseCase
	conn        *amqp.Connection
	channel     *amqp.Channel
	queue       amqp.Queue
}

func NewHistorialConsumer(historialUC *usecase.CreateHistorialUseCase) (*HistorialConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"historial", 
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &HistorialConsumer{
		HistorialUC: historialUC,
		conn:        conn,
		channel:     ch,
		queue:       q,
	}, nil
}

func (c *HistorialConsumer) StartConsuming() {
	msgs, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al consumir mensajes de historial: %v", err)
	}

	go func() {
		for d := range msgs {
			var historial models.Historial
			err := json.Unmarshal(d.Body, &historial)
			if err != nil {
				log.Printf("Error al deserializar historial: %v", err)
				continue
			}

			log.Printf("📥 Historial recibido: %+v", historial)

			err = c.HistorialUC.Execute(&historial)
			if err != nil {
				log.Printf("❌ Error al guardar historial en BD: %v", err)
				continue
			}

			log.Printf("✅ Historial guardado en BD correctamente: %+v", historial)
		}
	}()

	log.Println("🕒 Esperando historial...")
	select {}
}

func (c *HistorialConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}

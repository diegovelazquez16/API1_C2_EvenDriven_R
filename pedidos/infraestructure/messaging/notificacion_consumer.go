package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Notificacion struct {
	PedidoID uint   `json:"pedido_id"`
	Mensaje  string `json:"mensaje"`
}

type NotificacionConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewNotificacionConsumer() (*NotificacionConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"notificaciones", 
		true,             
		false,            
		false,            
		false,            
		nil,              
	)
	if err != nil {
		return nil, err
	}

	return &NotificacionConsumer{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (nc *NotificacionConsumer) StartConsuming() {
	msgs, err := nc.channel.Consume(
		nc.queue.Name, 
		"",            
		true,          
		false,         
		false,         
		false,         
		nil,           
	)
	if err != nil {
		log.Fatalf("Error al consumir notificaciones: %v", err)
	}

	go func() {
		for msg := range msgs {
			var notif Notificacion
			if err := json.Unmarshal(msg.Body, &notif); err != nil {
				log.Printf("Error al deserializar la notificación: %v", err)
				continue
			}
			log.Printf("Notificación recibida: PedidoID: %d, Mensaje: %s", notif.PedidoID, notif.Mensaje)
		}
	}()

	log.Println("Consumidor de notificaciones iniciado. Esperando mensajes...")
}

func (nc *NotificacionConsumer) Close() {
	nc.channel.Close()
	nc.conn.Close()
}

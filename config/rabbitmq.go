package config

import (
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQConn abre e retorna uma conexão com o RabbitMQ
func RabbitMQConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Altere conforme necessário
	if err != nil {
		log.Fatalf("❌ Falha ao conectar ao RabbitMQ: %s", err)
	}
	log.Println("✅ Conectado ao RabbitMQ")
	return conn
}

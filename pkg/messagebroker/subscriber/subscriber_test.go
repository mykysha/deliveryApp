package subscriber_test

import (
	"testing"

	v1 "github.com/nndergunov/deliveryApp/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/pkg/messagebroker/subscriber"
	amqp "github.com/rabbitmq/amqp091-go"
)

const hostURL = "amqp://guest:guest@localhost:5672/"

func TestSubscriber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		postTopics []string
		postData   []string
	}{
		{
			name:       "order posted mock",
			postTopics: []string{"orders"},
			postData:   []string{"new order posted"},
		},
		{
			name:       "receiving from two topics",
			postTopics: []string{"topic1", "topic2"},
			postData:   []string{"data1", "data2"},
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			msgChan := make(chan []byte)

			// creating listener, main test subject
			listener, err := subscriber.NewEventSubscriber(hostURL, msgChan)
			if err != nil {
				t.Fatal(err)
			}

			for _, topic := range test.postTopics {
				err = listener.SubscribeToTopic(topic)
				if err != nil {
					t.Fatal(err)
				}
			}

			sender, err := amqp.Dial(hostURL)
			if err != nil {
				t.Fatal(err)
			}

			publishChan, err := sender.Channel()
			if err != nil {
				t.Fatal(err)
			}

			for id, topic := range test.postTopics {
				err = publishChan.ExchangeDeclare(topic, "fanout", true, false, false, false, nil)
				if err != nil {
					t.Fatal(err)
				}

				body, err := v1.Encode(test.postData[id])
				if err != nil {
					t.Fatal(err)
				}

				var pub amqp.Publishing

				pub.ContentType = "text/plain"
				pub.Body = body

				err = publishChan.Publish(topic, "", false, false, pub)
				if err != nil {
					t.Fatal(err)
				}

				// receiving message
				message := <-msgChan

				receiveData := new(string)

				err = v1.Decode(message, receiveData)
				if err != nil {
					t.Fatal(err)
				}

				if test.postData[id] != *receiveData {
					t.Errorf("Expected: %v; Got: %v", test.postData, receiveData)
				}
			}
		})
	}
}

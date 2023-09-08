package gooseneck

import mqtt "github.com/eclipse/paho.mqtt.golang"

const (
	BROKER = "BROKER"
	TOPIC  = "TOPIC"
)

func NewMessageQueueClient(broker string, clientId string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.OnConnect = func(client mqtt.Client) {
		Info().Str("broker", broker).Str("client", clientId).Msg("connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		Info().Str("broker", broker).Str("client", clientId).Msg("connection lost")
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

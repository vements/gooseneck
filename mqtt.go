package gooseneck

import mqtt "github.com/eclipse/paho.mqtt.golang"

const (
	MQTT_BROKER   = "MQTT_BROKER"
	MQTT_PASSWORD = "MQTT_PASSWORD"
	MQTT_TOPIC    = "MQTT_TOPIC"
	MQTT_USERNAME = "MQTT_USERNAME"
)

func NewMessageQueueClient(broker string, clientId string, username string, password string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
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

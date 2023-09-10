package gooseneck

import (
	"crypto/tls"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTT_BROKER   = "MQTT_BROKER"
	MQTT_PASSWORD = "MQTT_PASSWORD"
	MQTT_TOPIC    = "MQTT_TOPIC"
	MQTT_USERNAME = "MQTT_USERNAME"
)

func NewMessageQueueClient(broker string, clientId string, username string, password string) mqtt.Client {
	options := mqtt.NewClientOptions()
	client := mqtt.NewClient(
		options.
			AddBroker(broker).
			SetClientID(clientId).
			SetUsername(username).
			SetPassword(password).
			SetCleanSession(false).
			SetKeepAlive(time.Second * 30).
			SetOnConnectHandler(func(c mqtt.Client) {
				Info().Str("broker", broker).Str("client", clientId).Msg("connected")
			}).
			SetConnectionLostHandler(func(c mqtt.Client, err error) {
				Info().Str("broker", broker).Str("client", clientId).Err(err).Msg("connection lost")
			}).
			SetConnectionAttemptHandler(func(b *url.URL, cfg *tls.Config) *tls.Config {
				Info().Str("broker", broker).Str("client", clientId).Msg("connection attempt")
				return cfg
			}).
			SetReconnectingHandler(func(c mqtt.Client, o *mqtt.ClientOptions) {
				Info().Str("broker", broker).Str("client", clientId).Msg("re-connection attempt")
			}))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

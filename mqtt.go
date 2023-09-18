package gooseneck

import (
	"crypto/tls"
	"net/url"
	"os"
	"path/filepath"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTT_BROKER   = "MQTT_BROKER"
	MQTT_PASSWORD = "MQTT_PASSWORD"
	MQTT_TOPIC    = "MQTT_TOPIC"
	MQTT_USERNAME = "MQTT_USERNAME"
)

type MqttEnv struct {
	Env
}

func (e MqttEnv) Broker() string {
	return e.MustDefine(MQTT_BROKER)
}

func (e MqttEnv) ClientID() string {
	return filepath.Base(os.Args[0])
}

func (e MqttEnv) Topic() string {
	return e.MustDefine(MQTT_TOPIC)
}

func (e MqttEnv) Username() string {
	return e.MustDefine(MQTT_USERNAME)
}

func (e MqttEnv) Password() string {
	return e.MustDefine(MQTT_PASSWORD)
}

func NewMessageQueueClient(broker string, clientId string, username string, password string, defaultHandler mqtt.MessageHandler) mqtt.Client {
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
			}).
			SetDefaultPublishHandler(func(c mqtt.Client, m mqtt.Message) {
				Info().Str("broker", broker).Str("client", clientId).Str("topic", m.Topic()).Msg("default handler")
				defaultHandler(c, m)
			}))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

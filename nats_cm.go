package nats_cm

import (
	"os"

	"github.com/CarlosMore29/nats_cm/config/env"

	"github.com/nats-io/nats.go"
)

var NATS_URL string
var NATS_USER string
var NATS_PASSWORD string
var PORT string

func init() {
	envGLobals()
}

func GetNatsConnection() (conn *nats.Conn, connEncoder *nats.EncodedConn, errNats error) {
	conn, errNats = nats.Connect(NATS_URL, nats.UserInfo(NATS_USER, NATS_PASSWORD))
	if errNats == nil {
		connEncoder, errNats = getEncoderConn(conn)
	}
	return
}

func getEncoderConn(conn *nats.Conn) (connEncoder *nats.EncodedConn, errEncoder error) {
	connEncoder, errEncoder = nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	return
}

func envGLobals() {
	env.GetEnvFile()
	NATS_URL = os.Getenv("NATS_URL")
	NATS_USER = os.Getenv("NATS_USER")
	NATS_PASSWORD = os.Getenv("NATS_PASSWORD")
}

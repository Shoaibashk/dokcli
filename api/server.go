package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	Dok "github.com/shoaibashk/dokcli/server"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Server() {

	d := Dok.Dokcli.New(Dok.Dokcli{}, ":1212", true, true)
	d.Register()
	d.Middleware()
	d.Routing()
	go d.StartServer()

	stopCh, closeCh := Dok.CreateChannel()
	defer closeCh()
	log.Print("notified:", <-stopCh)

	d.Shutdown(context.Background())

}

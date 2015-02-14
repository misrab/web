/*
	Convenience methods for dealing with a websocket server
*/

package web

import (
	"log"

	"net/http"

	"github.com/gorilla/websocket"
)


/*
	Listens to text responses over a websocket and 
	passes them to a channel.

	Arguments:
		- url e.g. wss://live.stellar.org:9001
		- body e.g. `{ "command": "subscribe", "streams": [ "transactions" ] }`
*/
func ListenToSocket(url, subscribeBody, unsubscribeBody string, c chan []byte, quit chan bool) error {
	var defaultDialer *websocket.Dialer
	var header http.Header

	conn, _, err := defaultDialer.Dial(url, header)
	if err != nil { panic(err) }

	// log.Println("Connected to " + url + " websocket...")

	conn.WriteMessage(websocket.TextMessage, []byte(subscribeBody))

	go func() {
		for {
			_, p, err := conn.ReadMessage()

			if err != nil {
				log.Println("Websocket error:")
				log.Println(err.Error())

				close(c)

			    return err
			}

			c <- p
		}
	}()

	for {
		<- quit
		log.Println("Unsubscribing...")
		conn.WriteMessage(websocket.TextMessage, []byte(unsubscribeBody))
		close(c)
	}

	return nil
}
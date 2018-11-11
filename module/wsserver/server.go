package wsserver

import (
	"echo-basic/pkg/modules"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
)

var upgrader websocket.Upgrader

func Start() {

	var wsModul = modules.Register("websocket server", 2)

	addr := viper.GetString("websocket.addr")
	if addr == "" {
		return
	}

	upgrader = websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	key, cert := viper.GetString("websocket.key"), viper.GetString("websocket.cert")

	var erch = make(chan error, 1)

	http.HandleFunc("/", serve)

	if key != "" && cert != "" {
		go func() {
			erch <- http.ListenAndServeTLS(addr, cert, key, nil)
		}()

	} else {
		go func() {
			erch <- http.ListenAndServe(addr, nil)
		}()
	}

	log.Info("websocket server start at:%s", addr)

	select {
	case err := <-erch:
		log.Error("websocket server error:%v", err)
	case <-wsModul.Stop:
		log.Info("websocket server stoped")
		// TODO: 平滑关闭
		wsModul.StopComplete()
		return
	}
}

// TODO:: Emit Event
func serve(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

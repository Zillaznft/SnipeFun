package bot

import (
	"time"

	"GoSnipeFun/config"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/gorilla/websocket"
)

var wsEndpoint = "wss://pumpportal.fun/api/data"
var solClient = client.NewClient(config.RpcEndpoint)
var signerKeyPair, _ = types.AccountFromBase58(config.PrivateKey)
var watchedTokens = make(map[string]Record)
var conn *websocket.Conn

func StartBot(onlyManage bool) {
	for {
		var err error
		conn, _, err = websocket.DefaultDialer.Dial(wsEndpoint, nil)
		if err != nil {
			sendNotification(Notification{
				Message: "Lost connection with websocket, retrying in 3 seconds...",
				Type:    errorNotification,
			})
			time.Sleep(3 * time.Second)
			continue
		}
		defer conn.Close()

		watchedTokens, _ = parseFileToMemory()
		syncTokens()
		saveRecordsToFile(watchedTokens)

		searching = true
		if len(watchedTokens) >= config.MaxTrades {
			searching = false
		}
		subscribeToEvents(onlyManage)
	}
}

func StartCleaner() {
	syncTokens()
	liquidateAll()
}

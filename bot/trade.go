package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"GoSnipeFun/config"
	"github.com/blocto/solana-go-sdk/types"
)

func executeTradeRetry(trade Trade, retries int) (err error) {
	for i := 0; i < retries; i++ {
		err = executeTrade(trade)
		if err == nil {
			sendNotification(Notification{Message: trade.ToString(), Type: infoNotification})
			return
		}
		if err != nil {
			log.Printf("[%s - %s] retry %d: %s", trade.Type, trade.Token, i, err.Error())
		}
	}
	sendNotification(Notification{Message: trade.ToString(), Type: errorNotification})
	return err
}

func executeTrade(trade Trade) (err error) {
	if config.IsLocalRpc {
		return executeTradeLocal(trade)
	}
	bodyBytes, err := json.Marshal(buildTradePayload(trade))
	if err != nil {
		log.Printf("marshal error: %v", err)
		return
	}

	resp, err := http.Post(fmt.Sprintf("https://pumpportal.fun/api/trade?api-key=%s", config.ApiKey), "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("post error: %v", err)
		return
	}
	defer resp.Body.Close()

	var response map[string]any
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("decode error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error in %s trade: %v", trade.Type, response)
		err = fmt.Errorf("error in %s trade: %v", trade.Type, response)
	}
	return
}

func executeTradeLocal(trade Trade) error {
	bodyBytes, err := json.Marshal(buildTradePayloadLocal(trade))
	if err != nil {
		log.Printf("marshal error: %v", err)
		return err
	}

	resp, err := http.Post("https://pumpportal.fun/api/trade-local", "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("post error: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var response map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Printf("decode error: %v", err)
			return err
		}
		log.Printf("Error in %s trade: %v", trade.Type, response)
		return fmt.Errorf("error in %s trade: %v", trade.Type, response)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %v", err)
		return err
	}

	ctx := context.Background()
	tx, err := types.TransactionDeserialize(data)
	if err != nil {
		log.Printf("deserialize error: %v", err)
		return err
	}

	serialized, err := tx.Message.Serialize()
	if err != nil {
		log.Printf("serialize error: %v", err)
		return err
	}

	if err = tx.AddSignature(signerKeyPair.Sign(serialized)); err != nil {
		log.Printf("add signature error: %v", err)
	}

	signature, err := solClient.SendTransactionWithConfig(ctx, tx, txConfig)
	if err != nil {
		log.Printf("failed to send transaction: %v", err)
		return err
	}

	sendNotification(Notification{
		Message: fmt.Sprintf("Transaction successful: https://solscan.io/tx/%s", signature),
		Type:    infoNotification,
	})
	return nil
}

//Builders

func buildTradePayload(trade Trade) (body map[string]any) {
	body = map[string]any{
		"mint":             trade.Token,
		"pool":             "pump",
		"denominatedInSol": "false",
		"action":           trade.Type,
		"priorityFee":      config.GasFee,
		"slippage":         config.Slippage,
	}
	switch trade.Type {
	case sellType:
		body["amount"] = config.StopLossSellPcg
	case tpType:
		body["amount"] = config.TakeProfitSellPcg
		body["action"] = sellType
	case buyType:
		body["denominatedInSol"] = "true"
		body["amount"] = config.TradeSize
	}
	return body
}

func buildTradePayloadLocal(trade Trade) map[string]any {
	body := map[string]any{
		"publicKey":        config.PublicKey,
		"mint":             trade.Token,
		"pool":             "pump",
		"denominatedInSol": "true",
		"action":           trade.Type,
		"priorityFee":      config.GasFee,
		"slippage":         config.Slippage,
	}

	switch trade.Type {
	case sellType:
		body["amount"] = config.StopLossSellPcg
	case tpType:
		body["amount"] = config.TakeProfitSellPcg
	case buyType:
		body["amount"] = config.TradeSize
	}
	return body
}

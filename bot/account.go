package bot

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"

	"GoSnipeFun/config"
)

var tokenProgramID string = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"

var txConfig = client.SendTransactionConfig{
	PreflightCommitment: rpc.CommitmentConfirmed,
}

func syncTokens() {
	rpcResponse, err := solClient.RpcClient.GetTokenAccountsByOwnerWithConfig(
		context.Background(),
		config.PublicKey,
		rpc.GetTokenAccountsByOwnerConfigFilter{
			ProgramId: tokenProgramID,
		},
		rpc.GetTokenAccountsByOwnerConfig{
			Encoding: rpc.AccountEncodingJsonParsed,
		},
	)
	if err != nil {
		log.Printf("failed to get token accounts by owner: %v", err)
		return
	}

	mapTokens := make(map[string]Record)
	for _, account := range rpcResponse.Result.Value {
		parsedData, ok := account.Account.Data.(map[string]any)
		if !ok {
			log.Printf("failed to parse account data: %v", account.Account.Data)
			continue
		}
		if parsedData, ok = parsedData["parsed"].(map[string]any); !ok {
			log.Printf("failed to parse account data: %v", account.Account.Data)
			continue
		}

		info, ok := parsedData["info"].(map[string]any)
		if !ok {
			log.Printf("failed to parse account info: %v", parsedData["info"])
			continue
		}

		mint, ok := info["mint"].(string)
		if !ok {
			log.Printf("failed to parse mint: %v", info["mint"])
			continue
		}

		record := Record{
			Address: mint,
		}

		if watchedTokens[mint].Timestamp != 0 {
			record.Timestamp = watchedTokens[mint].Timestamp

		}
		if watchedTokens[mint].MarketCap != 0 {
			record.MarketCap = watchedTokens[mint].MarketCap
		}

		mapTokens[mint] = record
	}
	watchedTokens = mapTokens

	sendNotification(Notification{
		Message: fmt.Sprintf("**--- Sync Completed ---**\n"+
			"%s", strings.Join(getWatchedTokenKeys(), "\n")),
		Type: infoNotification,
	})
}

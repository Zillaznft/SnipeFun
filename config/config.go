package config

const (
	//Wallet & API keys
	ApiKey      string = "" // api key to get the data from the website
	PrivateKey  string = "" // private key to sign the transactions
	PublicKey   string = "" // public key to get the data from the website
	RpcEndpoint string = "" // rpc endpoint to connect to the solana

	//Discord webhooks
	NotifWebhookURL string = "" // webhook to send notifications
	EventWebhookURL string = "" // webhook to send events
	ErrorWebhookURL string = "" // webhook to send errors

	//Trade configs
	IsLocalRpc        bool    = true   // use local rpc or pump fun jto rpc
	Retries           int     = 3      // number of retries on sell/buy (only if tx fail)
	GasFee            float64 = 0.003  // gas fee to take priority
	TradeSize         float64 = 0.005  // amount of sol on each buy
	MaxTrades         int     = 10     // max trades simultaneously
	Slippage          int     = 15     // slippage for the trades
	StopLoss          bool    = true   // true to activate stop loss mechanism
	TakeProfit        bool    = true   // true to activate take profit mechanism
	TakeProfitPcg     float64 = 200    // percentage to take profit
	StopLossPcg       float64 = 80     // percentage to sell on loss
	StopLossSellPcg   string  = "100%" // percentage of the position to liquidate on stop loss
	TakeProfitSellPcg string  = "80%"  // percentage of position to liquidate on take profit
	ThresholdBuy      float64 = 1.15   // threshold to buy
	ThresholdSell     float64 = 1.6    // threshold to sell

	//Filters
	TwitterFilter     bool = true // filter to check on x (twitter)
	TelegramFilter    bool = true // filter to check on telegram
	WebsiteFilter     bool = true // filter to check on website
	ImageFilter       bool = true // filter to check on image
	DescriptionFilter int  = 33   // filter min length of description (0 to disable)

	//Bot configs
	FileName     string = "records.txt" // path to persist the data (only if stop loss or take profit true)
	StartingMode string = "bot"         // "bot" or "cleaner" or "manager"
	// Bot starts the bot with the configs
	// Cleaner liquidates all the tokens on the wallet
	// Manager only manages open positions and wallet watch for copy trading
)

var WalletsToWatch = map[string]*struct {
	AvoidFilters bool
	NewTokens    bool
	Buy          bool
	Sell         bool
}{
	"samp13": { // address to watch
		AvoidFilters: true,  // avoid filters for the events
		NewTokens:    false, // snipe new tokens
		Buy:          true,  // copy trade to buy the same tokens
		Sell:         true,  // copy trade to sell the same tokens
	},
	//...
}

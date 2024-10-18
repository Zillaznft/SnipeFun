
# SnipeFun Bot 🛠️

This project is a sniper bot for the Solana blockchain, built to integrate with the Pump Fun DApp. It supports fast trading operations through configurable parameters, allowing users to tailor the bot's behavior to their specific needs

---

## 🚀 installation

1. **clone the repository**  
   open a terminal and run the following commands:

   ```bash
   git clone https://github.com/NNull13/SnipeFun.git
   cd SnipeFun
   ```

2. **install dependencies**  
   ensure all necessary dependencies are installed by running:

   ```bash
   go mod tidy
   ```

---

## ⚙️ configuration

1. **edit `config.go`**  
   navigate to the `config/config.go` file within the project

2. **set your parameters**  
   update the following parameters according to your needs:

   - **api and keys**  
     configure the keys for the wallet and data access:

     ```go
     ApiKey      string = "" // api key from the website
     PrivateKey  string = "" // private key for transactions
     PublicKey   string = "" // public key for data access
     RpcEndpoint string = "" // solana rpc endpoint
     ```

   - **webhook settings**  
     configure the discord webhooks to receive notifications and alerts:

     ```go
     NotifWebhookURL string = "" // notifications
     EventWebhookURL string = "" // events
     ErrorWebhookURL string = "" // error reports
     ```

   - **trade configurations**  
     customize how the bot trades:

     ```go
     Retries     int = 3 // retries for failed transactions
     GasFee      float64 = 0.001 // gas priority fee
     TradeSize   float64 = 0.005 // sol per trade
     StopLoss    bool = true // activate stop loss
     TakeProfit  bool = true // activate take profit
     ```

   - **wallet monitoring**  
     you can add **multiple wallets** to monitor. each wallet will be observed for specific actions such as buying, selling, or sniping new tokens:

     ```go
     var WalletsToWatch = map[string]*struct {
         NewTokens bool
         Buy       bool
         Sell      bool
     }{
         "wallet1": {NewTokens: true, Buy: true, Sell: true},
         "wallet2": {NewTokens: false, Buy: true, Sell: false},
         // add as many wallets as you like
     }
     ```

---

## 🛠️ usage

1. **run the bot**  
   after configuring `config.go`, you can start the bot with:

   ```bash
   go run start.go
   ```

   the bot will begin monitoring wallets and executing trades based on your settings

---

## 👥 contribution

we welcome contributions to the project! follow these steps to get started:

1. **fork the repository**
2. **create a new branch**:

   ```bash
   git checkout -b feature/new-feature
   ```

3. **make your changes** and commit:

   ```bash
   git commit -am 'add new feature'
   ```

4. **push to your branch**:

   ```bash
   git push origin feature/new-feature
   ```

5. **open a pull request** and describe your changes

---

## 💖 Donations 

This project is open-source and non-profit. if you find it helpful and want to support its development, consider making a donation:

- **solana wallet:** `NoName13.sol`
- **ethereum wallet:** `NoName13.eth`
- **bitcoin wallet:** `noname13.btc`

---

### Disclaimer
This software is provided as-is, with no guarantees or warranties. the developer is not responsible for any damages or losses resulting from the use of this bot. users are solely responsible for their actions and must comply with relevant laws and financial regulations.

---

### crafted with ❤️ by NoName13
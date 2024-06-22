# SnipeFun Bot

This project is a Sniper Bot for the Solana blockchain, designed for the DApp Pump Fun. The bot allows you to perform fast trading operations based on configurable parameters.

## Installation

1. Clone the repository to your local machine:
    ```sh
    git clone https://github.com/NNull13/SnipeFun.git
    cd SnipeFun
    ```

2. Install the necessary dependencies:
    ```sh
    go mod tidy
    ```

## Configuration

1. Open the `config/config.go` file located in the the project.

2. Modify the parameters according to your needs. Make sure to enter your specific keys (burner)


## Usage

1. Once the parameters in `config.go` are set, you can start the bot by running the `start.go` file:
    ```sh
    go run start.go
    ```

## Contribution

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/new-feature`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

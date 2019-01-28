# Go-blockchain

this is a implementation of a simple blockchain in go, it follow the tutorial of [Tensor Programming](https://www.youtube.com/channel/UCYqCZOwHbnPwyjawKfE21wg) on Youtube

## Build

To launch clone this repo
```
git clone https://github.com/bz10bis/golang-blockchain.git
```
then to run simply type
```
go run main.go
```
or if you want an executable
```
go build main.go
```

## Usage

```
 getbalance -address ADDRESS - get the balance of the address
 createblockchain -address ADDRESS - create a blockchain and send reward to ADDRESS
 printblocks - Prints the content of the blockchain
 send -from FROM -to TO -amount AMOUNT - Send amount of coins
 createwallet - Create a new wallet
 listaddresses - List all addresses contain in wallet file
```

## Requirement

This project is tested on Windown 8.1
```
go version go1.11.5 windows/amd64
```

## To do 
* Finish the wallet module
* Add digital signature to sign transactions
* UTXO Persistence layer
* Add a Merkle Tree
* Add a networking system
* Writing test
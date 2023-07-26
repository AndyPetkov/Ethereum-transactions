package loader

import (
	"context"
	"mid-the-ethereum-fetcher-zxowaz/database"
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"mid-the-ethereum-fetcher-zxowaz/models"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadTransactions() error {
	logger.GetInstance().InfoLogger.Println("Starting the application...")
	client, err := ethclient.Dial(os.Getenv("ETH_NODE_URL"))
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}

	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return err
	}
	var transactions []models.Transaction
	for _, tx := range block.Transactions() {
		perTransaction := models.Transaction{}
		perTransaction.TransactionHash = tx.Hash().Hex()
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return err
		}
		perTransaction.TransactionStatus = int(receipt.Status)
		perTransaction.BlockHash = block.Hash().Hex()
		perTransaction.BlockNumber = int(block.NumberU64())
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return err
		}

		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			perTransaction.From = from.Hex()
		}
		perTransaction.To = tx.To().Hex()
		if perTransaction.To == "" || len(strings.TrimSpace(perTransaction.To)) == 0 {
			perTransaction.ContractAddress = address.Hex()
		}

		perTransaction.LogsCount = len(receipt.Logs)
		perTransaction.Input = string(tx.Data())
		perTransaction.Value = tx.Value().String()
		transactions = append(transactions, perTransaction)

	}

	err = Create(transactions)
	return err
}

func Create(transactions []models.Transaction) error {
	logger.GetInstance().InfoLogger.Println("Creating new transactions...")
	for index := range transactions {
		sqlStatement := `INSERT INTO transactions (transactionHash,transactionStatus,blockHash,blockNumber,to_,from_,contractAddress,logsCount,input,value_) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10)`
		_, err := database.Database.Exec(sqlStatement, transactions[index].TransactionHash, transactions[index].TransactionStatus,
			transactions[index].BlockHash, transactions[index].BlockNumber, transactions[index].To,
			transactions[index].From, transactions[index].ContractAddress, transactions[index].LogsCount,
			transactions[index].Input, transactions[index].Value)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return err
		}
		logger.GetInstance().InfoLogger.Println("Created new transaction")
	}
	return nil
}

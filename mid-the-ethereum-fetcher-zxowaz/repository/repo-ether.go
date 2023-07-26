package repository

import (
	"encoding/hex"
	"mid-the-ethereum-fetcher-zxowaz/database"
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"mid-the-ethereum-fetcher-zxowaz/models"
)

type TransactionRepo interface {
	GetAll() (models.Transactions, error)
	GetByRlphex(rlphex string) (interface{}, error)
}
type Transaction struct {
}

var BaseExecutorTransaction TransactionRepo = &Transaction{}

func NewRepoTransaction() TransactionRepo {
	return BaseExecutorTransaction
}

func (t *Transaction) GetByRlphex(rlphex string) (interface{}, error) {
	var transactionsSlice []models.Transaction
	transactions, err := t.GetAll()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	for _, tx := range transactions.Transactions {
		rawTxHex := hex.EncodeToString([]byte(tx.TransactionHash))
		if rawTxHex == rlphex {
			transactionsSlice = append(transactionsSlice, tx)
		}
	}
	var finalTransactions models.Transactions
	finalTransactions.Transactions = transactionsSlice
	logger.GetInstance().InfoLogger.Println("Got transactions by RLP")
	return finalTransactions, nil
}

func (t *Transaction) GetAll() (models.Transactions, error) {
	var transactions []models.Transaction
	data, err := database.Database.Query("SELECT * FROM transactions")
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return models.Transactions{}, err
	}
	count := 0
	for data.Next() {
		perTransaction := models.Transaction{}
		err = data.Scan(
			&perTransaction.TransactionHash, &perTransaction.TransactionStatus,
			&perTransaction.BlockHash, &perTransaction.BlockNumber, &perTransaction.To,
			&perTransaction.From, &perTransaction.ContractAddress, &perTransaction.LogsCount,
			&perTransaction.Input, &perTransaction.Value)
		if err != nil {
			logger.GetInstance().ErrorLogger.Println(err)
			return models.Transactions{}, err
		}
		transactions = append(transactions, perTransaction)
		count++
	}
	var finalTransactions models.Transactions
	finalTransactions.Transactions = transactions
	logger.GetInstance().InfoLogger.Println("Got all transactions")
	return finalTransactions, err
}

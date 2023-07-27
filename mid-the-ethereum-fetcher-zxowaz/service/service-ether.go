package service

import (
	"errors"
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"mid-the-ethereum-fetcher-zxowaz/repository"
)

type TransactionService interface {
	GetAll() (interface{}, error)
	GetByRlphex(rlphex string) (interface{}, error)
	ConfigureRepoTransaction(repository repository.TransactionRepo)
}
type transaction struct {
	Repo repository.TransactionRepo
}

var BaseExecutorTransaction TransactionService = &transaction{}

func NewServiceTransaction() TransactionService {
	return &transaction{repository.NewRepoTransaction()}
}

func (t *transaction) ConfigureRepoTransaction(repository repository.TransactionRepo) {
	t.Repo = repository
}

func (t *transaction) GetByRlphex(rlphex string) (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting transactions by rlphex...")
	if rlphex == " " {
		err := errors.New("you are missing rlphex parameter.")
		return nil, err
	}
	result, err := t.Repo.GetByRlphex(rlphex)
	if err != nil {
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got transactions by rlphex")
	return result, nil
}

func (t *transaction) GetAll() (interface{}, error) {
	logger.GetInstance().InfoLogger.Println("Getting all transactions...")
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("Got all transactions")
	return result, nil
}

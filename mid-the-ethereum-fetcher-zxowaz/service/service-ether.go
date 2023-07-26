package service

import (
	"errors"
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"mid-the-ethereum-fetcher-zxowaz/repository"
	"net/http"
)

type TransactionService interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByRlphex(w http.ResponseWriter, r *http.Request)
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

func (t *transaction) GetByRlphex(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting transactions by rlphex...")
	rlphex := r.URL.Query().Get(":rlphex")
	if rlphex == " " {
		err := errors.New("you are missing rlphex parameter.")
		writeResponse(w, err, nil, http.StatusBadRequest)
		return
	}
	result, err := t.Repo.GetByRlphex(rlphex)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Got transactions by rlphex")
}

func (t *transaction) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting all transactions...")
	result, err := t.Repo.GetAll()
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Got all transactions")
}

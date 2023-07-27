package controller

import (
	"mid-the-ethereum-fetcher-zxowaz/logger"
	"mid-the-ethereum-fetcher-zxowaz/service"
	"net/http"
)

type TransactionController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByRlphex(w http.ResponseWriter, r *http.Request)
	ConfigureServiceTransaction(service service.TransactionService)
}
type transaction struct {
	Service service.TransactionService
}

var BaseExecutorTransaction TransactionController = &transaction{}

func NewControllerTransaction() TransactionController {
	return &transaction{service.NewServiceTransaction()}
}

func (t *transaction) ConfigureServiceTransaction(service service.TransactionService) {
	t.Service = service
}

func (t *transaction) GetByRlphex(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting transactions by rlphex...")
	rlphex := r.URL.Query().Get(":rlphex")
	result, err := t.Service.GetByRlphex(rlphex)
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Got transactions by rlphex")
}

func (t *transaction) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.GetInstance().InfoLogger.Println("Getting all transactions...")
	result, err := t.Service.GetAll()
	if err != nil {
		writeResponse(w, err, result, http.StatusBadRequest)
		return
	}
	writeResponse(w, err, result, http.StatusOK)
	logger.GetInstance().InfoLogger.Println("Got all transactions")
}

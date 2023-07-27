package service

import (
	"errors"
	repomocks "mid-the-ethereum-fetcher-zxowaz/mocks"
	"mid-the-ethereum-fetcher-zxowaz/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("service-transaction", func() {
	var (
		actualTransactionsTest models.Transactions
		errTest                error
		errRlphexTest          error
		testTransaction        TransactionService
		mockTest               *repomocks.TransactionRepo
		expectedTransactions   models.Transactions
		//emptyTransactionTest models.Transactions
	)
	BeforeEach(func() {
		transactionsTest := []models.Transaction{
			{
				TransactionHash:   "0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524",
				TransactionStatus: 1,
				BlockHash:         "0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3",
				BlockNumber:       7976382,
				From:              "0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a",
				To:                "0x3664f6c1178e19bb775b597d6584caa3b88a1c35",
				ContractAddress:   "0x3664f6c1178e19bb775b597d6584caa3b88a1c35",
				LogsCount:         3,
				Input:             "0x",
				Value:             "0",
			},
			{
				TransactionHash:   "0x5a57e3051cb92e2d482515b07e7b3d1851722a74654657bd64a14c39ca3f9cf2",
				TransactionStatus: 0,
				BlockHash:         "0x92557f7e29c39cae6be013ffc817620fcd5233b68405cdfc6e0b5528261e81e5",
				BlockNumber:       7976373,
				From:              "0xf29a6c0f8ee500dc87d0d4eb8b26a6fac7a76767",
				To:                "0xb0428bf0d49eb5c2239a815b43e59e124b84e303",
				ContractAddress:   "",
				LogsCount:         0,
				Input:             "0x",
				Value:             "50000000000000000",
			},
		}
		//emptyTransactionTest = models.Transactions{Transactions: nil}
		actualTransactionsTest = models.Transactions{
			Transactions: transactionsTest,
		}
		expectedTransactions = models.Transactions{}
		errTest = errors.New("")
		errRlphexTest = errors.New("you are missing rlphex parameter.")
		mockTest = &repomocks.TransactionRepo{}
		testTransaction = NewServiceTransaction()
		testTransaction.ConfigureRepoTransaction(mockTest)
	})

	Describe("GetAll", func() {

		It("should return all transactions and nil when all passed parameters are valid", func() {
			mockTest.On("GetAll").Return(actualTransactionsTest, nil)
			expectedResult, expectedErr := testTransaction.GetAll()
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualTransactionsTest))
		},
		)
		It("should return nil and error", func() {
			mockTest.On("GetAll").Return(expectedTransactions, errTest)
			expectedResult, expectedErr := testTransaction.GetAll()
			Expect(expectedErr).ToNot(BeNil())
			Expect(expectedResult).To(BeNil())
		},
		)
	})

	Describe("GetByRlphex", func() {

		It("should return all transactions and nil when all passed parameters are valid", func() {
			mockTest.On("GetByRlphex", "f90110b8423078396232").Return(actualTransactionsTest, nil)
			expectedResult, expectedErr := testTransaction.GetByRlphex("f90110b8423078396232")
			Expect(expectedErr).To(BeNil())
			Expect(expectedResult).To(Equal(actualTransactionsTest))
		},
		)
		It("should return nil and error when wrong parameters are passed", func() {
			mockTest.On("GetByRlphex").Return(expectedTransactions, errTest)
			expectedResult, expectedErr := testTransaction.GetByRlphex(" ")
			Expect(expectedErr).To(Equal(errRlphexTest))
			Expect(expectedResult).To(BeNil())
		},
		)
	})

})

package transaction

import (
	"be_online_course/course"
	"be_online_course/payment"
	"fmt"
	"strconv"
)

type service struct {
	repository       Repository
	courseRepository course.Repository
	paymentService   payment.Service
}

type Service interface {
	// GetTransactionsByCourseID(input GetCourseTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionsInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]Transaction, error)
	GetUserByPaidStatus(UserID int) ([]Transaction, error)
}

func NewService(repository Repository, courseRepository course.Repository, paymentService payment.Service) *service {
	return &service{repository, courseRepository, paymentService}
}

// func (s *service) GetTransactionsByCampaignID(input GetCourseTransactionsInput) ([]Transaction, error) {
// 	// Retrieve the campaign using the provided campaign ID
// 	course, err := s.courseRepository.FindByID(input.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check if the course is not found
// 	if course == nil {
// 		return nil, errors.New("course not found")
// 	}

// 	// Retrieve transactions associated with the course
// 	transactions, err := s.repository.GetCampaignID(course.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionsInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CourseID = int(input.Course.ID)
	transaction.Amount = input.Course.Price
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}
	fmt.Println("sebelum payment url")
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User, input.Course)
	if err != nil {
		return newTransaction, err
	}
	fmt.Println("Payment Url", paymentURL, err)
	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionID, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transactionID)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	_, err = s.repository.Update(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetUserByPaidStatus(UserID int) ([]Transaction, error) {
	transactions, err := s.repository.FindUserByPaidStatus(UserID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

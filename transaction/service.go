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
	GetTransactionsByCourseID(input GetCourseTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionsInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]Transaction, error)
}

func NewService(repository Repository, courseRepository course.Repository, paymentService payment.Service) *service {
	return &service{repository, courseRepository, paymentService}
}

// func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
// 	campaign, err := s.courseRepository.FindByID(input.ID)
// 	if err != nil {
// 		return []Transaction{}, err
// 	}

// 	if campaign.UserID != input.User.ID {
// 		return []Transaction{}, errors.New("Not an owner of the campaign")
// 	}

// 	transactions, err := s.repository.GetCampaignID(input.ID)
// 	if err != nil {
// 		return transactions, err
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
	transaction.CourseID = input.CourseID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	paymentTranscation := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}
	fmt.Println("sebelum payment url")
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTranscation, input.User)
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
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
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

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	// course, err := s.courseRepository.FindByID(updatedTransaction.CourseID)
	// if err != nil {
	// 	return err
	// }

	// if updatedTransaction.Status == "paid" {
	// 	// course.BackerCount = campaign.BackerCount + 1
	// 	// campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

	// 	_, err := s.courseRepository.Update(campaign)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

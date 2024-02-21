package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCourseID(courseID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	GetByID(ID int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	FindAll() ([]Transaction, error)
	FindUserByPaidStatus(UserID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCourseID(courseID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Preload("User").Where("course = ?", courseID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("course.CourseImage", "course_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (r *repository) GetByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.First(&transaction, ID).Error
	if err != nil {
		return Transaction{}, err // Return a clear error if the record is not found
	}
	return transaction, nil
}

func (r *repository) FindAll() ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("Campaign").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) FindUserByPaidStatus(UserID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("Course.Chapter.Lessons").Where("status = ? AND user_id = ?", "paid", UserID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

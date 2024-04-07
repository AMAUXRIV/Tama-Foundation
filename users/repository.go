package users

import "gorm.io/gorm"

type Repository interface {
	Save(users Users) (Users, error)
	FindByEmail(email string) (Users, error)
	FindByID(ID int) (Users, error)
	Update(users Users) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) Save(users Users) (Users, error) {
	err := r.db.Create(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) FindByEmail(email string) (Users, error) {
	var user Users
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (Users, error) {
	var user Users
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(users Users) (Users, error) {
	err := r.db.Save(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

package repository

import (
	"errors"

	"github.com/RusseLHuang/consistent-hashing/consistent"
	"github.com/RusseLHuang/consistent-hashing/user/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db         map[string]*gorm.DB
	consistent consistent.Consistent
}

func (u UserRepository) CreateRepository(
	dbConnectionMap map[string]*gorm.DB,
	consistent consistent.Consistent,
) *UserRepository {
	return &UserRepository{
		db:         dbConnectionMap,
		consistent: consistent,
	}
}

func (u *UserRepository) Save(user *entity.User) {
	var conn *gorm.DB
	connectionKey := u.consistent.GetNearestKey(user.NationalID)
	conn = u.db[connectionKey]

	result := conn.Create(user)

	if result.Error != nil {
		panic(result.Error)
	}
}

func (u *UserRepository) AddBalance(
	user *entity.User,
	amount uint,
) error {
	var userEntity entity.User

	var conn *gorm.DB
	connectionKey := u.consistent.GetNearestKey(user.NationalID)
	conn = u.db[connectionKey]

	tx := conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(
		&userEntity,
		"id = ?",
		user.ID,
	)

	currentBalance := userEntity.Balance

	tx.Model(&entity.User{
		ID: user.ID,
	}).Update("balance", currentBalance+amount)

	tx.Commit()

	return nil
}

func (u *UserRepository) Transfer(
	fromUser *entity.User,
	toUser *entity.User,
	amount uint,
) error {
	userCondition := []uint{fromUser.ID, toUser.ID}

	var conn *gorm.DB
	connectionKey := u.consistent.GetNearestKey(fromUser.NationalID)
	conn = u.db[connectionKey]

	tx := conn.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var sourceUser entity.User
	var destinationUser entity.User

	var userEntity []entity.User

	tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(
		&userEntity,
		userCondition,
	)

	for i := 0; i < len(userEntity); i++ {
		if userEntity[i].ID == fromUser.ID {
			sourceUser = userEntity[i]
		} else if userEntity[i].ID == toUser.ID {
			destinationUser = userEntity[i]
		}
	}

	if sourceUser.Balance < amount {
		tx.Rollback()
		return errors.New("Balance is not enough to transfer")
	}

	sourceUser.Balance = sourceUser.Balance - amount
	destinationUser.Balance = destinationUser.Balance + amount

	tx.Model(&entity.User{
		ID: sourceUser.ID,
	}).Update("balance", sourceUser.Balance)

	tx.Model(&entity.User{
		ID: destinationUser.ID,
	}).Update("balance", destinationUser.Balance)

	tx.Commit()

	return nil
}

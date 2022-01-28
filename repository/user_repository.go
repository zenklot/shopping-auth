package repository

import (
	"context"
	"database/sql"
	"shopping/model/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Sign(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		panic(err)
	}
	password := string(bytes)

	SQL := "INSERT INTO user(username, password, email, phone, country, city, postcode, name, address) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, user.Username, password, user.Email, user.Phone, user.Country, user.City, user.Postcode, user.Name, user.Address)
	if err != nil {
		panic(err)
	}

	return user

}

func (repository *UserRepositoryImpl) Sign(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {

	SQL := "SELECT password FROM user where email = ?"
	result := tx.QueryRow(SQL, user.Email)

	var userData domain.User
	err := result.Scan(&userData.Password)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		panic(err)
	}

	return user
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	panic("not implemented") // TODO: Implement
}

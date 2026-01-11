package repository

import (
	"context"
	"database/sql"
	"nutech-test/internal/entity"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Create(user entity.User) error {
	
	_, err := ur.db.ExecContext(
		context.Background(), `
		INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)`, user.Email, user.Password, user.FirstName, user.LastName)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.db.QueryRowContext(context.Background(), `
	SELECT email, password FROM users WHERE email = $1`, email,
	).Scan(&user.Email, &user.Password)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}



//use pointer to reference the same struct when passing the parameter to the service layer
func (ur *UserRepo) UpdateUserByEmail(user *entity.User) (entity.User, error) {
	_, err := ur.db.ExecContext(context.Background(), `
	UPDATE users SET first_name = $1, last_name = $2 WHERE email = $3`,
	user.FirstName, user.LastName, user.Email)
	
	if err != nil {
		return entity.User{}, err
	}
	
	return *user, nil
}

//update profile image
func (ur *UserRepo) UpdateImageByEmail(user *entity.User) (entity.User, error) {
	_, err := ur.db.ExecContext(context.Background(), `
	UPDATE users SET profile_image = $1 WHERE email = $2`,
	user.ProfileImage, user.Email)
	
	if err != nil {
		return entity.User{}, err
	}
	
	return *user, nil
}

func (ur *UserRepo) ProfileGetByEmail(email string) (entity.User, error) {
	var user entity.User
	//cant scan profile_image because its null, maybe we can just set profile image default to " " so its not null?
	err := ur.db.QueryRowContext(context.Background(), `
	SELECT email, first_name, last_name, profile_image FROM users WHERE email = $1`, email,
	).Scan(&user.Email, &user.FirstName, &user.LastName, &user.ProfileImage)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *UserRepo) GetBalanceByEmail(email string) (entity.Saldo, error) {
	var saldo entity.Saldo
	err := ur.db.QueryRowContext(context.Background(), `
	SELECT balance FROM saldo WHERE user_email = $1`, 
    email).Scan(&saldo.Balance)
	
	if err != nil {
		return entity.Saldo{}, err
	}
	
	return saldo, nil
}

func (ur *UserRepo) UpdateBalanceByEmail(saldo *entity.Saldo) (entity.Saldo, error) {
	_, err := ur.db.ExecContext(context.Background(), `
	UPDATE saldo SET balance = $1 WHERE user_email = $2`, saldo.Balance, saldo.UserEmail)

	if err != nil {
		return entity.Saldo{}, err
	}

	return *saldo, nil
}

func (ur *UserRepo) CreateSaldoByEmail(saldo *entity.Saldo) error {
	_, err := ur.db.ExecContext(context.Background(), `
	INSERT INTO saldo (user_email, balance) VALUES ($1, $2)`, saldo.UserEmail, saldo.Balance)

	if err != nil {
		return err
	}

	return nil
}
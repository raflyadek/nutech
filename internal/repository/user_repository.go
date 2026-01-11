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
	SELECT * FROM users WHERE email = $1`, email)

	if err != nil {
		return entity.User{}, sql.ErrNoRows
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
	UPDATE users SET profile_iamge = $1 WHERE email = $2`,
	user.ProfileImage, user.Email)

	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}

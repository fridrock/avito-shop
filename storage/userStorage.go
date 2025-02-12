package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id             uuid.UUID `db:"id"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	Coins          int       `db:"coins"`
}

type UserStorage interface {
	FindUserByUsername(string) (User, error)
	SaveUser(User) (uuid.UUID, error)
	CheckEnoughCoins(int, uuid.UUID) bool
	GetUserById(uuid.UUID) (User, error)
}

type UserStorageImpl struct {
	db *sqlx.DB
}

func (as *UserStorageImpl) FindUserByUsername(username string) (User, error) {
	var user User
	q := `SELECT * FROM users WHERE username=$1`
	row := as.db.QueryRowx(q, username)
	err := row.StructScan(&user)
	if err == sql.ErrNoRows {
		err = errors.Join(err, fmt.Errorf("no such user"))
	}
	return user, err
}
func (as *UserStorageImpl) SaveUser(user User) (uuid.UUID, error) {
	q := `INSERT INTO users(id, username, hashed_password, coins) VALUES ($1, $2, $3, 1000) RETURNING id`
	var id uuid.UUID
	err := as.db.QueryRow(
		q,
		uuid.New().String(),
		user.Username,
		user.HashedPassword).Scan(&id)
	return id, err
}

func (as *UserStorageImpl) CheckEnoughCoins(amount int, userId uuid.UUID) bool {
	q := `SELECT coins FROM users WHERE id = $1`
	var coins int
	err := as.db.Get(&coins, q, userId.String())
	if err != nil {
		slog.Debug(fmt.Sprintf("error getting user coins with id : %v, err: %v", userId.String(), err.Error()))
		return false
	}
	return coins >= amount
}

func (as *UserStorageImpl) GetUserById(userId uuid.UUID) (User, error) {
	q := `SELECT * FROM users WHERE id = $1`
	var user User
	err := as.db.Get(&user, q, userId.String())
	return user, err
}

func NewUserStorage(db *sqlx.DB) UserStorage {
	return &UserStorageImpl{
		db: db,
	}
}

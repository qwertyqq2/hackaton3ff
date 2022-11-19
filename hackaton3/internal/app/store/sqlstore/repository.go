package sqlstore

import (
	"fmt"
	"log"
	"test_go/internal/app/model"
)

type Repository struct {
	store *Store
}

func (r *Repository) Create(e *model.User) error {
	queryUsers := fmt.Sprintf("INSERT INTO Users (name, role, balance, password) VALUES('%s', '%s', '%d','%s') RETURNING id;",
		e.Name,
		e.Role,
		0,
		e.Password,
	)
	return r.store.db.QueryRow(queryUsers).Scan(&e.ID)
}

func (r *Repository) GetBalanceForId(userId uint) (uint, error) {
	query := fmt.Sprintf("SELECT * FROM Users WHERE id='%d'", userId)

	u := &model.User{}

	if err := r.store.db.QueryRow(query).Scan(
		&u.ID,
		&u.Name,
		&u.Role,
		&u.Balance,
		&u.Password,
	); err != nil {
		return 0, err
	}

	return uint(u.Balance), nil
}

func (r *Repository) IncBalance(userId uint, value uint) error {
	balance, err := r.GetBalanceForId(userId)
	if err != nil {
		log.Fatal(err)
	}
	queryUser := fmt.Sprintf(
		"UPDATE Users SET balance = '%d' WHERE id = '%d';",
		balance+value,
		userId,
	)
	r.store.db.QueryRow(queryUser)
	return nil

}

func (r *Repository) DecrBalance(userId uint, value uint) error {
	balance, err := r.GetBalanceForId(userId)
	if err != nil {
		log.Fatal(err)
	}
	queryUser := fmt.Sprintf(
		"UPDATE Users SET balance = '%d' WHERE id = '%d';",
		balance-value,
		userId,
	)
	r.store.db.QueryRow(queryUser)
	return nil
}

func (r *Repository) BuyToken(userId uint, pairNumber uint, value uint, refund uint) error {
	if err := r.DecrBalance(userId, value); err != nil {
		log.Fatal(err)
	}
	queryTx := fmt.Sprintf(
		"INSERT INTO Transactions (type, userId, value, refund, to) VALUES('%d', '%d', '%d', '%d', '%d') RETURNING id;",
		pairNumber,
		userId,
		value,
		refund,
		1,
	)
	var txId int
	return r.store.db.QueryRow(queryTx).Scan(&txId)
}

func (r *Repository) SellToken(userId uint, pairNumber uint, value uint, refund uint) error {
	if err := r.IncBalance(userId, refund); err != nil {
		log.Fatal(err)
	}
	queryTx := fmt.Sprintf(
		"INSERT INTO Transactions (type, userId, value, refund, to) VALUES('%d', '%d', '%d', '%d', '%d') RETURNING id;",
		pairNumber,
		userId,
		value,
		refund,
		0,
	)
	var txId int
	return r.store.db.QueryRow(queryTx).Scan(&txId)
}

func (r *Repository) GetUsers() ([]model.User, error) {
	query := "SELECT * FROM Users"
	resp, err := r.store.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var users []model.User
	for resp.Next() {
		var e model.User
		if err := resp.Scan(&e.ID, &e.Name, &e.Role, &e.Balance, &e.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, e)
	}
	return users, err
}

func (r *Repository) FindUser(req *model.User) (*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM Users WHERE name='%s' AND password='%s'", req.Name, req.Password)

	u := &model.User{}

	if err := r.store.db.QueryRow(query).Scan(
		&u.ID,
		&u.Name,
		&u.Role,
		&u.Password,
	); err != nil {
		return nil, err
	}
	return u, nil
}

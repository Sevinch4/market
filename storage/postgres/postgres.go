package postgres

import (
	"database/sql"
	"fmt"
	"market/config"
	"market/storage"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host = %s port = %s user = %s password = %s database = %s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	return Store{
		DB: db,
	}, nil
}

func (s Store) Close() {
	s.DB.Close()
}

func (s Store) UserStorage() storage.IUserStorage {
	newUser := NewUserRepo(s.DB)
	return newUser
}

func (s Store) Basket() storage.IBasket {
	newBasket := NewBasketRepo(s.DB)
	return newBasket

}

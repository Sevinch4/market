package postgres

import (
	"database/sql"
	"fmt"
	"market/models"
	"market/storage"

	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return userRepo{db: db}
}

func (u userRepo) Create(user models.CreateUser) (models.User, error) {
	id := uuid.New()
	userC := models.User{}
	if _, err := u.db.Exec(`insert into users(id, full_name, phone, password, cash, user_type)  values($1, $2, $3, $4, $5, $6)`,
		id,
		user.FullName,
		user.Phone,
		user.Password,
		user.Cash,
		user.UserType); err != nil {
		fmt.Println("error is while creating user", err.Error())
		return models.User{}, err
	}

	err := u.db.QueryRow(`SELECT id, full_name, phone, password, cash, user_type FROM users WHERE id = $1`, id).Scan(&userC.ID, &userC.FullName,
		&userC.Phone, &userC.Password, &userC.Cash, &userC.UserType)

	if err != nil {
		fmt.Println("error while selecting:", err.Error())
		return models.User{}, err
	}

	return userC, nil
}

func (u userRepo) GetByID(key models.PrimaryKey) (models.User, error) {
	user := models.User{}

	if err := u.db.QueryRow(`select * from users where id = $1`, key.ID).Scan(&user.ID, &user.FullName,
		user.Phone,
		user.Password,
		user.Cash,
		user.UserType); err != nil {
		fmt.Println("error is while selecting user", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u userRepo) GetList() (models.UserResponse, error) {
	users := models.UserResponse{}

	//var count int
	rows, err := u.db.Query(`select id, full_name, phone, password, cash, user_type, count(*) as count from users group by id, full_name, phone, password, cash, user_type `)
	if err != nil {
		fmt.Println("error is while is selecting users", err.Error())
		return models.UserResponse{}, err
	}

	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(&user.ID, &user.FullName, &user.Phone, &user.Password, &user.Cash, &user.UserType, &users.Count); err != nil {
			fmt.Println("error is while scanning", err.Error())
			return models.UserResponse{}, err
		}

		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (u userRepo) Update(models.UpdateUSer) (models.User, error) {
	user := models.UpdateUSer{}

	if _, err := u.db.Exec(`update users set full_name = $1, phone = $2, password = $3, cash = $4 where id = $5 `,
		&user.FullName, &user.Phone, &user.Password, &user.Cash, user.ID); err != nil {
		fmt.Println("error is while updating users", err.Error())
		return models.User{}, err
	}

	userUp := models.User{}

	if err := u.db.QueryRow(`select * from users where id = $1`, user.ID).Scan(&userUp.ID, &userUp.FullName, &userUp.Phone, &userUp.Password, &userUp.Cash, &userUp.UserType); err != nil {
		fmt.Println("error is while selecting", err.Error())
		return models.User{}, err
	}

	return userUp, nil
}

func (u userRepo) Delete(key models.PrimaryKey) error {
	if _, err := u.db.Exec(`delete from users where id = $1`, key.ID); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return err
	}
	return nil
}

/*
GetByID(models.PrimaryKey) (models.User, error)
GetList(request models.GetListRequest) (models.UserResponse, error)
Update(models.UpdateUSer) (models.User, error)
Delete(key models.PrimaryKey) error*/

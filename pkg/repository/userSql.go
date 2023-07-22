package repository

import (
	"fmt"
	"sarkor/test/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UserSql struct{
	db *sqlx.DB
}

func NewUserSql(db *sqlx.DB) *UserSql{
	return &UserSql{db: db}
}

func (us *UserSql) RegisterUser(user model.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password, name, age) values($1, $2, $3, $4) RETURNING id", usersTable)
	row := us.db.QueryRow(query, user.Login, user.Password, user.Name, user.Age)

	if err := row.Scan(&id); err != nil{
		return 0, err
	}

	return id, nil
}

func (us *UserSql) GetUser(login, password string) (model.User, error){
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := us.db.Get(&user, query, login, password)

	return user, err
}

func (us *UserSql) GetUserByName(name string) (model.User, error){
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", usersTable)
	err := us.db.Get(&user, query, name)

	return user, err
}

func(us *UserSql) FindPhoneDuplicate(phone string) (string, error){
	var hasDuplicate string
	query := fmt.Sprintf("SELECT phone FROM %s WHERE phone=$1", phoneTable)
	err := us.db.Get(&hasDuplicate, query, phone)

	return hasDuplicate, err
}

func(us *UserSql) AddUserPhone(phone model.Phone, userId int) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (phone, description, is_fax, user_id) values($1, $2, $3, $4) RETURNING id", phoneTable)
	row := us.db.QueryRow(query, phone.Phone, phone.Description, phone.Is_fax, userId)

	if err := row.Scan(&id); err != nil{
		return 0, err
	}

	return id, nil
}

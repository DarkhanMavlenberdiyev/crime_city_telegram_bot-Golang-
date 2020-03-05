package t_bot

import (
	"fmt"

	"github.com/go-pg/pg"
)

type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
}

type postgreStore struct {
	db *pg.DB
}

func NewPostgreBot(config PostgreConfig) CrimeEvents {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     "postgres",
		Password: config.Password,
	})
	return &postgreStore{db: db}
}

func PostgreUser(config PostgreConfig) UserInfo {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     "postgres",
		Password: config.Password,
	})
	return &postgreStore{db: db}
}

func (p postgreStore) GetUser(id int) (*Users, error) {
	user := &Users{ID: id}
	err := p.db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (p postgreStore) GetAllUser() ([]*Users, error) {
	var users []*Users
	err := p.db.Model(&users).Select()

	if err != nil {
		return nil, err
	}
	return users, nil
}
func (p postgreStore) UpdateUser(id int, user *Users) (*Users, error) {
	panic("")
}
func (p postgreStore) DeleteUser(id int) error {
	user := &Users{ID: id}
	err := p.db.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
func (p postgreStore) CreateUser(user *Users) (*Users, error) {

	re := p.db.Insert(user)
	return user, re
}

func (p postgreStore) GetAllCrimes() ([]*Crime, error) {
	var crimes []*Crime
	err := p.db.Model(&crimes).Select()

	if err != nil {
		return nil, err
	}
	return crimes, nil
}

func (p postgreStore) CreateCrime(crime *Crime) (*Crime, error) {
	err := p.db.Insert(crime)

	return crime, err
}

func (p postgreStore) GetCrime(id int) (*Crime, error) {
	crime := &Crime{ID: id}
	err := p.db.Select(crime)
	if err != nil {
		return nil, err
	}

	fmt.Println(crime.ID)
	return crime, nil
}

func (p postgreStore) UpdateCrime(id int, crime *Crime) (*Crime, error) {
	crime.ID = id
	err := p.db.Update(crime)
	return crime, err
}

func (p postgreStore) DeleteCrime(id int) error {
	crime := &Crime{ID: id}
	err := p.db.Delete(crime)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreStore) GetAllCrime() ([]*Crime, error) {
	var crimes []*Crime
	err := p.db.Model(&crimes).Select()

	if err != nil {
		return nil, err
	}
	return crimes, nil
}

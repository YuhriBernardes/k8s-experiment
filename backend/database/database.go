package database

import (
	"fmt"
	"mykube/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

type Db struct {
	db *gorm.DB
}

func NewDb(config DbConfig) (*Db, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.User, config.Password, config.DbName, config.Port)
	pg := postgres.Open(connString)
	db, err := gorm.Open(pg, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Db{
		db: db,
	}, nil
}

func rollback(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
	}
}

func (d *Db) Migrate() error {
	return d.db.AutoMigrate(&model.User{})
}

func (d *Db) Create(user *model.User) error {
	tx := d.db.Begin()
	defer rollback(tx)

	if result := tx.Create(user); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	err := tx.Commit().Error

	return err
}
func (d *Db) GetAll() ([]*model.User, error) {
	users := make([]*model.User, 0)

	if err := d.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

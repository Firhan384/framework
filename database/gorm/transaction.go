package gorm

import (
	"gorm.io/gorm"

	"github.com/Firhan384/framework/contracts/config"
	"github.com/Firhan384/framework/contracts/database/orm"
)

type Transaction struct {
	orm.Query
	instance *gorm.DB
}

func NewTransaction(tx *gorm.DB, config config.Config) *Transaction {
	return &Transaction{Query: NewQueryWithWithoutEvents(tx, false, config), instance: tx}
}

func (r *Transaction) Commit() error {
	return r.instance.Commit().Error
}

func (r *Transaction) Rollback() error {
	return r.instance.Rollback().Error
}

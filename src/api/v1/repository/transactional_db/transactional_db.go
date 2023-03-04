package transactional_db

import (
	"gorm.io/gorm"
	"test-bpjs/src/api/v1/repository"
)

type TransactionalDBRepoStruct struct {
	db *gorm.DB
}

func NewTransactionalDBRepoImpl(db *gorm.DB) repository.TransactionalDBRepoInterface {
	return &TransactionalDBRepoStruct{db}
}

func (e TransactionalDBRepoStruct) BeginTrans() *gorm.DB {
	return e.db.Begin()
}

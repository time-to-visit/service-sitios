package repository_test

import (
	"errors"
	"fmt"
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/repository"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_RepositoryConcact_RegisterConctact(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	gormDB.Logger.LogMode(logger.Info)
	if err != nil {
		t.Fatalf("failed to open gorm database: %s", err)
	}
	createdAt := time.Now()
	contact := entity.Contact{
		Model: entity.Model{
			ID:        0x1,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		},
	}

	repo := repository.NewRepositoryContact(gormDB)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO `contacts` (.+) VALUES (.+)").WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := repo.InsertContact(contact)
		if err := mock.ExpectationsWereMet(); err != nil {
			fmt.Printf("Unfulfilled expectations: %s \n", err)
		}
		assert.NoError(t, err)
		assert.Equal(t, &contact, result)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO `contacts` (.+) VALUES (.+)").WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnError(errors.New("failed to insert user"))
		mock.ExpectRollback()

		_, err := repo.InsertContact(contact)
		assert.Error(t, err)
	})
}

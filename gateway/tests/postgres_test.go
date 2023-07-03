package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	database "gateway/database/postgres"
)

var (
	usr     string = "user"
	pwd     string = "password"
	isAdmin bool   = false
	dbConn         = database.PGConnProps{
		DB_HOST:     "db",
		DB_PORT:     "5432",
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "user_info",
	}

	serviceDB database.IDatabase = database.Initial(dbConn)
)

func TestCreateAccount(t *testing.T) {
	db, mock, errMock := sqlmock.New()
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}
	defer db.Close()

	mock.
		ExpectQuery(`INSERT INTO user_info(.+) returning uid`).
		WithArgs(usr, pwd, isAdmin).
		WillReturnRows(sqlmock.NewRows([]string{""}).AddRow(1))

	serviceDB.SetDB(db)

	if _, err := serviceDB.CreateAccount(usr, pwd, isAdmin); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

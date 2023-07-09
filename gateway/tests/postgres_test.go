package tests

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	database "gateway/database/postgres"
)

var (
	usr              string = "user"
	pwd              string = "password"
	mock_uid         string = "92299bbe-c1a1-400b-8d0e-59b082cd2cb9"
	mock_user_name   string = "PatharaNor"
	mock_encrypt_pwd string = "$2a$06$4A3WFWvVYRMeUR0SK2pRietesT989retpeb2FUYTV3ZntETRGXE.K"
	is_admin         bool   = false
	dbConn                  = database.PGConnProps{
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
		WithArgs(usr, pwd, is_admin).
		WillReturnRows(sqlmock.NewRows([]string{""}).AddRow(1))

	serviceDB.SetDB(db)

	if _, err := serviceDB.CreateAccount(usr, pwd, is_admin); err != nil {
		t.Errorf("error was not expected while create account: %s", err)
	}

	serviceDB.Close()
}

func TestGetAccount(t *testing.T) {
	db, mock, errMock := sqlmock.New()
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"uid", "username", "password", "is_admin"}).
		AddRow(mock_uid, mock_user_name, mock_encrypt_pwd, is_admin)

	mock.
		ExpectQuery("SELECT (.+) FROM user_info WHERE (.+)").
		WithArgs(usr, pwd).
		WillReturnRows(rows)

	serviceDB.SetDB(db)

	if _, err := serviceDB.GetAccount(usr, pwd); err != nil {
		fmt.Printf("TestGetAccount: %v", err.Error())
		t.Errorf("error was not expected while get account: %s", err)
	}

	serviceDB.Close()
}

func TestGetAccountByID(t *testing.T) {
	db, mock, errMock := sqlmock.New()
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"uid", "username", "password", "is_admin"}).
		AddRow(mock_uid, mock_user_name, mock_encrypt_pwd, is_admin)

	mock.
		ExpectQuery("SELECT (.+) FROM user_info WHERE (.+)").
		WithArgs(mock_uid).
		WillReturnRows(rows)

	serviceDB.SetDB(db)

	if _, err := serviceDB.GetAccountByID(mock_uid); err != nil {
		fmt.Printf("TestGetAccountByID: %v", err.Error())
		t.Errorf("error was not expected while get account by id: %s", err)
	}

	serviceDB.Close()
}

func TestUpdateAccount(t *testing.T) {
	var lastInsertID, affected int64
	db, mock, errMock := sqlmock.New()
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}
	defer db.Close()

	mock.
		ExpectExec("UPDATE user_info SET (.+) WHERE (.+) RETURNING uid").
		WithArgs(mock_user_name, is_admin, mock_uid).
		WillReturnResult(sqlmock.NewResult(lastInsertID, affected))

	serviceDB.SetDB(db)

	if err := serviceDB.UpdateAccount(mock_uid, mock_user_name, is_admin); err != nil {
		fmt.Printf("TestUpdateAccount: %v", err.Error())
		t.Errorf("error was not expected while update account: %s", err)
	}

	serviceDB.Close()
}

func TestDeleteAccountByID(t *testing.T) {
	var lastInsertID, affected int64
	db, mock, errMock := sqlmock.New()
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}
	defer db.Close()

	mock.
		ExpectExec("DELETE FROM user_info WHERE (.+)").
		WithArgs(mock_uid).
		WillReturnResult(sqlmock.NewResult(lastInsertID, affected))

	serviceDB.SetDB(db)

	if err := serviceDB.DeleteAccountByID(mock_uid); err != nil {
		fmt.Printf("TestDeleteAccountByID: %v", err.Error())
		t.Errorf("error was not expected while delete account: %s", err)
	}

	serviceDB.Close()
}

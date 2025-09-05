package database

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"auth/internal/model"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestHealth_Up(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}

	mock.ExpectPing() // Expect a successful ping

	stats := s.Health()

	if stats["status"] != "up" {
		t.Errorf("Expected status 'up', got %s", stats["status"])
	}
	if stats["message"] != "It's healthy" {
		t.Errorf("Expected message 'It's healthy', got %s", stats["message"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// func TestHealth_Down(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	s := &service{db: db}
//
// 	mock.ExpectPing().WillReturnError(errors.New("db connection error")) // Expect a failed ping
//
// 	stats := s.Health()
//
// 	if stats["status"] != "down" {
// 		t.Errorf("Expected status 'down', got %s", stats["status"])
// 	}
// 	if stats["error"] != "db down: db connection error" {
// 		t.Errorf("Expected error 'db down: db connection error', got %s", stats["error"])
// 	}
//
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestClose(t *testing.T) {
	dbName := "testdb"

	t.Run("Successful Close", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		s := &service{db: db}

		mock.ExpectClose().WillReturnError(nil)

		err = s.Close(dbName)
		if err != nil {
			t.Errorf("Close() error = %v, wantErr %v", err, false)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Failed Close", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		s := &service{db: db}

		mock.ExpectClose().WillReturnError(errors.New("close error"))

		err = s.Close(dbName)
		if err == nil {
			t.Errorf("Close() error = %v, wantErr %v", err, true)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestCountUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Count - isGetAll true", func(t *testing.T) {
		filter := "test"
		isGetAll := true
		expectedCount := 5

		mock.ExpectQuery(`
		select count\(\*\) from users
		where email ilike '%' \|\| \$1 \|\| '%'
		`).WithArgs(filter).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

		count, err := s.CountUsers(ctx, filter, isGetAll)
		if err != nil {
			t.Errorf("CountUsers() error = %v, wantErr %v", err, false)
		}
		if count != expectedCount {
			t.Errorf("CountUsers() got = %v, expected %v", count, expectedCount)
		}
	})

	t.Run("Successful Count - isGetAll false", func(t *testing.T) {
		filter := "test"
		isGetAll := false
		expectedCount := 3

		mock.ExpectQuery(`
		select count\(\*\) from users
		where email ilike '%' \|\| \$1 \|\| '%'
		and is_active = true
		`).WithArgs(filter).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

		count, err := s.CountUsers(ctx, filter, isGetAll)
		if err != nil {
			t.Errorf("CountUsers() error = %v, wantErr %v", err, false)
		}
		if count != expectedCount {
			t.Errorf("CountUsers() got = %v, expected %v", count, expectedCount)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		filter := "test"
		isGetAll := true

		mock.ExpectQuery(`
		select count\(\*\) from users
		where email ilike '%' \|\| \$1 \|\| '%'
		`).WithArgs(filter).WillReturnError(errors.New("db error"))

		_, err := s.CountUsers(ctx, filter, isGetAll)
		if err == nil {
			t.Errorf("CountUsers() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Select - isGetAll true", func(t *testing.T) {
		filter := "test"
		limit := 2
		offset := 0
		isGetAll := true
		expectedUsers := []*model.UserDTO{
			{Id: "1", Email: "test1@example.com", IsActive: true, Role: model.RoleUser},
			{Id: "2", Email: "test2@example.com", IsActive: true, Role: model.RoleAdmin},
		}

		rows := sqlmock.NewRows([]string{"id", "email", "is_active", "role"}).
			AddRow("1", "test1@example.com", true, model.RoleUser).
			AddRow("2", "test2@example.com", true, model.RoleAdmin)

		mock.ExpectQuery(`
		select id, email, is_active, role from users
		where email ilike '%' \|\| \$1 \|\| '%'
		limit \$2 offset \$3
		`).WithArgs(filter, limit, offset).WillReturnRows(rows)

		users, err := s.SelectUsers(ctx, limit, offset, filter, isGetAll)
		if err != nil {
			t.Errorf("SelectUsers() error = %v, wantErr %v", err, false)
		}
		if len(users) != len(expectedUsers) {
			t.Errorf("SelectUsers() got %d users, expected %d", len(users), len(expectedUsers))
		}
		for i, user := range users {
			if *user != *expectedUsers[i] {
				t.Errorf(
					"SelectUsers() at index %d got %v, expected %v",
					i,
					*user,
					*expectedUsers[i],
				)
			}
		}
	})

	t.Run("Successful Select - isGetAll false", func(t *testing.T) {
		filter := "test"
		limit := 1
		offset := 0
		isGetAll := false
		expectedUsers := []*model.UserDTO{
			{Id: "1", Email: "test1@example.com", IsActive: true, Role: model.RoleUser},
		}

		rows := sqlmock.NewRows([]string{"id", "email", "is_active", "role"}).
			AddRow("1", "test1@example.com", true, model.RoleUser)

		mock.ExpectQuery(`
		select id, email, is_active, role from users
		where email ilike '%' \|\| \$1 \|\| '%'
		and is_active = true
		limit \$2 offset \$3
		`).WithArgs(filter, limit, offset).WillReturnRows(rows)

		users, err := s.SelectUsers(ctx, limit, offset, filter, isGetAll)
		if err != nil {
			t.Errorf("SelectUsers() error = %v, wantErr %v", err, false)
		}
		if len(users) != len(expectedUsers) {
			t.Errorf("SelectUsers() got %d users, expected %d", len(users), len(expectedUsers))
		}
		for i, user := range users {
			if *user != *expectedUsers[i] {
				t.Errorf(
					"SelectUsers() at index %d got %v, expected %v",
					i,
					*user,
					*expectedUsers[i],
				)
			}
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		filter := "test"
		limit := 1
		offset := 0
		isGetAll := true

		mock.ExpectQuery(`
		select id, email, is_active, role from users
		where email ilike '%' \|\| \$1 \|\| '%'
		limit \$2 offset \$3
		`).WithArgs(filter, limit, offset).WillReturnError(errors.New("db error"))

		_, err := s.SelectUsers(ctx, limit, offset, filter, isGetAll)
		if err == nil {
			t.Errorf("SelectUsers() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Select", func(t *testing.T) {
		id := "123"
		expectedUser := &model.User{
			Id:       "123",
			Email:    "test@example.com",
			IsActive: true,
			Role:     model.RoleUser,
			Password: "hashedpassword",
		}

		rows := sqlmock.NewRows([]string{"id", "email", "is_active", "role", "password"}).
			AddRow(expectedUser.Id, expectedUser.Email, expectedUser.IsActive, expectedUser.Role, expectedUser.Password)

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where id = \$1
		`).WithArgs(id).WillReturnRows(rows)

		user, err := s.SelectUserById(ctx, id)
		if err != nil {
			t.Errorf("SelectUserById() error = %v, wantErr %v", err, false)
		}
		if *user != *expectedUser {
			t.Errorf("SelectUserById() got = %v, expected %v", *user, *expectedUser)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		id := "nonexistent"

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where id = \$1
		`).WithArgs(id).WillReturnError(sql.ErrNoRows)

		_, err := s.SelectUserById(ctx, id)
		if err == nil || err != sql.ErrNoRows {
			t.Errorf("SelectUserById() error = %v, expected %v", err, sql.ErrNoRows)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		id := "123"

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where id = \$1
		`).WithArgs(id).WillReturnError(errors.New("db error"))

		_, err := s.SelectUserById(ctx, id)
		if err == nil {
			t.Errorf("SelectUserById() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Select", func(t *testing.T) {
		email := "test@example.com"
		expectedUser := &model.User{
			Id:       "123",
			Email:    "test@example.com",
			IsActive: true,
			Role:     model.RoleUser,
			Password: "hashedpassword",
		}

		rows := sqlmock.NewRows([]string{"id", "email", "is_active", "role", "password"}).
			AddRow(expectedUser.Id, expectedUser.Email, expectedUser.IsActive, expectedUser.Role, expectedUser.Password)

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where email = \$1
		`).WithArgs(email).WillReturnRows(rows)

		user, err := s.SelectUserByEmail(ctx, email)
		if err != nil {
			t.Errorf("SelectUserByEmail() error = %v, wantErr %v", err, false)
		}
		if *user != *expectedUser {
			t.Errorf("SelectUserByEmail() got = %v, expected %v", *user, *expectedUser)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		email := "nonexistent@example.com"

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where email = \$1
		`).WithArgs(email).WillReturnError(sql.ErrNoRows)

		_, err := s.SelectUserByEmail(ctx, email)
		if err == nil || err != sql.ErrNoRows {
			t.Errorf("SelectUserByEmail() error = %v, expected %v", err, sql.ErrNoRows)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		email := "test@example.com"

		mock.ExpectQuery(`
		select id, email, is_active, role, password
		from users
		where email = \$1
		`).WithArgs(email).WillReturnError(errors.New("db error"))

		_, err := s.SelectUserByEmail(ctx, email)
		if err == nil {
			t.Errorf("SelectUserByEmail() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Insert", func(t *testing.T) {
		user := &model.User{
			Email:    "newuser@example.com",
			IsActive: true,
			Role:     model.RoleUser,
			Password: "Password123!",
		}
		expectedID := "generated-id-123"

		mock.ExpectQuery(`
		insert into users\(email, is_active, role, password\)
		values\(\$1, \$2, \$3, \$4\)
		returning id
		`).WithArgs(user.Email, user.IsActive, user.Role, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

		err := s.InsertUser(ctx, user)
		if err != nil {
			t.Errorf("InsertUser() error = %v, wantErr %v", err, false)
		}
		if user.Id != expectedID {
			t.Errorf("InsertUser() got ID %v, expected %v", user.Id, expectedID)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		user := &model.User{
			Email:    "error@example.com",
			IsActive: true,
			Role:     model.RoleUser,
			Password: "Password123!",
		}

		mock.ExpectQuery(`
		insert into users\(email, is_active, role, password\)
		values\(\$1, \$2, \$3, \$4\)
		returning id
		`).WithArgs(user.Email, user.IsActive, user.Role, sqlmock.AnyArg()).WillReturnError(errors.New("db error"))

		err := s.InsertUser(ctx, user)
		if err == nil {
			t.Errorf("InsertUser() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Update", func(t *testing.T) {
		id := "123"
		newEmail := "updated@example.com"
		expectedUserDTO := &model.UserDTO{
			Id:       "123",
			Email:    "updated@example.com",
			IsActive: true,
			Role:     model.RoleUser,
		}

		rows := sqlmock.NewRows([]string{"id", "role", "email", "is_active"}).
			AddRow(expectedUserDTO.Id, expectedUserDTO.Role, expectedUserDTO.Email, expectedUserDTO.IsActive)

		mock.ExpectQuery(`
		update users
		set email = \$1
		where id = \$2
		returning id, role, email, is_active
		`).WithArgs(newEmail, id).WillReturnRows(rows)

		userDTO, err := s.UpdateUser(ctx, id, newEmail)
		if err != nil {
			t.Errorf("UpdateUser() error = %v, wantErr %v", err, false)
		}
		if *userDTO != *expectedUserDTO {
			t.Errorf("UpdateUser() got = %v, expected %v", *userDTO, *expectedUserDTO)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		id := "123"
		newEmail := "updated@example.com"

		mock.ExpectQuery(`
		update users
		set email = \$1
		where id = \$2
		returning id, role, email, is_active
		`).WithArgs(newEmail, id).WillReturnError(errors.New("db error"))

		_, err := s.UpdateUser(ctx, id, newEmail)
		if err == nil {
			t.Errorf("UpdateUser() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Update", func(t *testing.T) {
		id := "123"
		newPassword := "NewPassword123!"

		mock.ExpectExec(`
		update users
		set password = \$1
		where id = \$2
		`).WithArgs(sqlmock.AnyArg(), id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.UpdateUserPassword(ctx, id, newPassword)
		if err != nil {
			t.Errorf("UpdateUserPassword() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		id := "nonexistent"
		newPassword := "NewPassword123!"

		mock.ExpectExec(`
		update users
		set password = \$1
		where id = \$2
		`).WithArgs(sqlmock.AnyArg(), id).WillReturnResult(sqlmock.NewResult(1, 0))

		err := s.UpdateUserPassword(ctx, id, newPassword)
		if err == nil || err != sql.ErrNoRows {
			t.Errorf("UpdateUserPassword() error = %v, expected %v", err, sql.ErrNoRows)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		id := "123"
		newPassword := "NewPassword123!"

		mock.ExpectExec(`
		update users
		set password = \$1
		where id = \$2
		`).WithArgs(sqlmock.AnyArg(), id).WillReturnError(errors.New("db error"))

		err := s.UpdateUserPassword(ctx, id, newPassword)
		if err == nil {
			t.Errorf("UpdateUserPassword() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUserStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Update - Activate", func(t *testing.T) {
		id := "123"
		isActive := true

		mock.ExpectExec(`
		update users
		set is_active = \$1
		where id = \$2
		`).WithArgs(isActive, id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.UpdateUserStatus(ctx, id, isActive)
		if err != nil {
			t.Errorf("UpdateUserStatus() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("Successful Update - Deactivate", func(t *testing.T) {
		id := "123"
		isActive := false

		mock.ExpectExec(`
		update users
		set is_active = \$1
		where id = \$2
		`).WithArgs(isActive, id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.UpdateUserStatus(ctx, id, isActive)
		if err != nil {
			t.Errorf("UpdateUserStatus() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		id := "nonexistent"
		isActive := true

		mock.ExpectExec(`
		update users
		set is_active = \$1
		where id = \$2
		`).WithArgs(isActive, id).WillReturnResult(sqlmock.NewResult(1, 0))

		err := s.UpdateUserStatus(ctx, id, isActive)
		if err == nil || err != sql.ErrNoRows {
			t.Errorf("UpdateUserStatus() error = %v, expected %v", err, sql.ErrNoRows)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		id := "123"
		isActive := true

		mock.ExpectExec(`
		update users
		set is_active = \$1
		where id = \$2
		`).WithArgs(isActive, id).WillReturnError(errors.New("db error"))

		err := s.UpdateUserStatus(ctx, id, isActive)
		if err == nil {
			t.Errorf("UpdateUserStatus() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := &service{db: db}
	ctx := context.Background()

	t.Run("Successful Delete", func(t *testing.T) {
		id := "123"

		mock.ExpectExec(`
		delete from users
		where id = \$1
		`).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := s.DeleteUserById(ctx, id)
		if err != nil {
			t.Errorf("DeleteUserById() error = %v, wantErr %v", err, false)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		id := "nonexistent"

		mock.ExpectExec(`
		delete from users
		where id = \$1
		`).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 0))

		err := s.DeleteUserById(ctx, id)
		if err == nil || err != sql.ErrNoRows {
			t.Errorf("DeleteUserById() error = %v, expected %v", err, sql.ErrNoRows)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		id := "123"

		mock.ExpectExec(`
		delete from users
		where id = \$1
		`).WithArgs(id).WillReturnError(errors.New("db error"))

		err := s.DeleteUserById(ctx, id)
		if err == nil {
			t.Errorf("DeleteUserById() error = %v, wantErr %v", err, true)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

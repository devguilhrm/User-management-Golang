package repository

import (
	"database/sql"

	"user-management/models"
)

// Insert User
func InsertUser(db *sql.DB, user *models.User) error {
	stmt, err := db.Prepare("INSERT INTO users(id, name, email, age) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Age)
	return err
}

// Read User
func ReadUser(db *sql.DB) ([]*models.User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		u := &models.User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Update User
func UpdateUser(db *sql.DB, user *models.User) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete User
func DeleteUser(db *sql.DB, name string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE name = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"ShareGo/api/models"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	createFilesTableQuery := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL,
		filepath TEXT NOT NULL,
		user_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createFilesTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(username, password string) (*models.User, error) {
	stmt, err := DB.Prepare("INSERT INTO users (username, password_hash) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, password)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        int(id),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	row := DB.QueryRow("SELECT id, username, password_hash, created_at FROM users WHERE username = ?", username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func CreateFile(filename, filepath string, userID int) (*models.File, error) {
	stmt, err := DB.Prepare("INSERT INTO files (filename, filepath, user_id) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(filename, filepath, userID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.File{
		ID:        int(id),
		Filename:  filename,
		Filepath:  filepath,
		UserID:    userID,
		CreatedAt: time.Now(),
	}, nil
}

func GetFileByID(id int) (*models.File, error) {
	row := DB.QueryRow("SELECT id, filename, filepath, user_id, created_at FROM files WHERE id = ?", id)

	var file models.File
	err := row.Scan(&file.ID, &file.Filename, &file.Filepath, &file.UserID, &file.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &file, nil
}

func DeleteFile(id int) error {
	stmt, err := DB.Prepare("DELETE FROM files WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func GetAllFiles() ([]models.File, error) {
	rows, err := DB.Query("SELECT id, filename, filepath, user_id, created_at FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.ID, &file.Filename, &file.Filepath, &file.UserID, &file.CreatedAt); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

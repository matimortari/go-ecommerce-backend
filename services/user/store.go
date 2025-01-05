package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/matimortari/go-backend/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
    _, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)", 
        user.FirstName, user.LastName, user.Email, user.Password)
    if err != nil {
        log.Println("Error creating user:", err)
        return err
    }
    return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
    if err != nil {
        log.Println("Error querying user by email:", err)
        return nil, err
    }
    defer rows.Close()

    u := new(types.User)
    if rows.Next() {
        u, err = scanRowsIntoUser(rows)
        if err != nil {
            return nil, err
        }
    } else {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
    if err != nil {
        log.Println("Error querying user by ID:", err)
        return nil, err
    }
    defer rows.Close()

    u := new(types.User)
    if rows.Next() {
        u, err = scanRowsIntoUser(rows)
        if err != nil {
            return nil, err
        }
    } else {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
    user := new(types.User)

    err := rows.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )
    if err != nil {
        return nil, err
    }

    return user, nil
}

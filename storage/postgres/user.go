package postgres

import (
	"fmt"
	"time"

	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(u *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			phone_number,
			email,
			password,
			username,
			image_url
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at
	`
	err := ur.db.QueryRow(
		query,
		u.FirstName,
		u.LastName,
		u.PhoneNumber,
		u.Email,
		u.Password,
		u.Username,
		u.ImageUrl,
	).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil 
} 

func (ur *userRepo) Get(user_id int64) (*repo.User, error) {
	var u repo.User
	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			username,
			image_url,
			created_at,
			updated_at
		FROM users WHERE id = $1
	`
	err := ur.db.QueryRow(
		query,
		user_id,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.Email,
		&u.Username,
		&u.ImageUrl,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return nil, nil 
}

func (ur *userRepo) Update(u *repo.User) (*repo.User, error) {
	query := `
		UPDATE users SET
			first_name = $1,
			last_name = $2,
			phone_number = $3,
			email = $4,
			username = $6,
			image_url = $7,
			updated_at = $8
		WHERE id = $9
		RETURNING
			id,
			first_name,
			last_name,
			phone_number,
			email,
			username,
			image_url,
			created_at,
			updated_at
	`
	err := ur.db.QueryRow(
		query,
		u.FirstName,
		u.LastName,
		u.PhoneNumber,
		u.Email,
		u.Username,
		u.ImageUrl,
		time.Now(),
		u.ID,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.Email,
		&u.Username,
		&u.ImageUrl,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (ur *userRepo) Delete(user_id int64) error {
	query := "UPDATE users SET deleted_at = $1 WHERE id = $2"

	_, err := ur.db.Exec(query, time.Now(), user_id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult,error) {
	var res repo.GetAllUsersResult
	res.Users = make([]*repo.User, 0)
	filter := ""
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(` 
		WHERE first_name ILIKE '%s' OR 
		last_name ILIKE '%s' OR 
		phone_number ILIKE '%s' OR 
		email ILIKE '%s' username ILIKE '%s' `, str, str, str, str, str)
	}

	orderBy := " ORDER BY DESC "
	if params.SortBy != "" {
		orderBy = fmt.Sprintf(" ORDER BY %s", params.SortBy)		
	}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			username,
			image_url,
			created_at,
			updated_at
		FROM users 
	` + filter + limit + orderBy

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var u repo.User
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Email,
			&u.Username,
			&u.ImageUrl,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		res.Users = append(res.Users, &u)
	}
	queryCount := "SELECT count(*) FROM users " + filter 

	err = ur.db.QueryRow(queryCount).Scan(&res.Count)
	if err != nil {
		return nil, err
	}

	return nil, nil 
}
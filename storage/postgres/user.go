package postgres

import (
	"database/sql"
	"fmt"

	"github.com/SaidovZohid/note-taking/pkg/utils"
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
	var (
		updatedAt sql.NullTime
	)
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
		utils.NullString(u.PhoneNumber),
		u.Email,
		u.Password,
		utils.NullString(u.Username),
		utils.NullString(u.ImageUrl),
	).Scan(
		&u.ID,
		&u.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	u.UpdatedAt = updatedAt.Time

	return u, nil
}

func (ur *userRepo) Get(user_id int64) (*repo.User, error) {
	var (
		updatedAt                       sql.NullTime
		imageUrl, phoneNumber, username sql.NullString
		res                             repo.User
	)
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
		FROM users WHERE id = $1 and deleted_at is null
	`
	err := ur.db.QueryRow(
		query,
		user_id,
	).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&phoneNumber,
		&res.Email,
		&username,
		&imageUrl,
		&res.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	res.PhoneNumber = phoneNumber.String
	res.Username = username.String
	res.ImageUrl = imageUrl.String
	res.UpdatedAt = updatedAt.Time

	return &res, nil
}

func (ur *userRepo) Update(u *repo.User) (*repo.User, error) {
	query := `
		UPDATE users SET
			first_name = $1,
			last_name = $2,
			phone_number = $3,
			email = $4,
			username = $5,
			image_url = $6,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
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
	var (
		res                             repo.User
		imageUrl, phoneNumber, username sql.NullString
	)
	err := ur.db.QueryRow(
		query,
		u.FirstName,
		u.LastName,
		utils.NullString(u.PhoneNumber),
		u.Email,
		utils.NullString(u.Username),
		utils.NullString(u.ImageUrl),
		u.ID,
	).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&phoneNumber,
		&res.Email,
		&username,
		&imageUrl,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	res.PhoneNumber = phoneNumber.String
	res.Username = username.String
	res.ImageUrl = imageUrl.String

	return &res, nil
}

func (ur *userRepo) Delete(user_id int64) error {
	query := "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1"

	res, err := ur.db.Exec(query, user_id)
	if err != nil {
		return err
	}

	if result, _ := res.RowsAffected(); result == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	var (
		res                             repo.GetAllUsersResult
		updatedAt                       sql.NullTime
		imageUrl, phoneNumber, username sql.NullString
	)
	res.Users = make([]*repo.User, 0)
	filter := " WHERE deleted_at IS NULL "

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(` 
		 AND first_name ILIKE '%s' OR 
		last_name ILIKE '%s' OR 
		phone_number ILIKE '%s' OR 
		email ILIKE '%s' OR username ILIKE '%s' `, str, str, str, str, str)
	}

	orderBy := " ORDER BY created_at DESC "
	if params.SortBy != "" {
		orderBy = fmt.Sprintf(" ORDER BY created_at %s ", params.SortBy)
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
	` + filter + orderBy + limit

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
			&phoneNumber,
			&u.Email,
			&username,
			&imageUrl,
			&u.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		u.PhoneNumber = phoneNumber.String
		u.Username = username.String
		u.ImageUrl = imageUrl.String
		u.UpdatedAt = updatedAt.Time

		res.Users = append(res.Users, &u)
	}
	queryCount := "SELECT count(*) FROM users " + filter

	err = ur.db.QueryRow(queryCount).Scan(&res.Count)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (ur *userRepo) GetByEmail(email string) (*repo.User, error) {
	query := `
		SELECT 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			password,
			username,
			image_url,
			created_at,
			updated_at
		FROM users WHERE email = $1
	`
	var (
		res                             repo.User
		updatedAt                       sql.NullTime
		imageUrl, phoneNumber, username sql.NullString
	)
	err := ur.db.QueryRow(
		query,
		email,
	).Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&phoneNumber,
		&res.Email,
		&res.Password,
		&username,
		&imageUrl,
		&res.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	res.PhoneNumber = phoneNumber.String
	res.Username = username.String
	res.ImageUrl = imageUrl.String
	res.UpdatedAt = updatedAt.Time

	return &res, nil
}

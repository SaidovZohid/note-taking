package repo

import "time"

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	Username    *string
	Email       string
	Password    string
	PhoneNumber *string
	ImageUrl    *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(user_id int64) (*User, error)
	Update(u *User) (*User, error)
	Delete(user_id int64) error
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
}

type GetAllUsersParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}

type GetAllUsersResult struct {
	Users []*User
	Count int64
}

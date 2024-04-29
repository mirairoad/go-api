package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	Verified  bool               `json:"verified,omitempty" bson:"verified,omitempty"`
	Attempts  int                `json:"attempts,omitempty" bson:"attempts,omitempty"`
	Roles     []Role             `json:"roles,omitempty" bson:"roles,omitempty"`
	LastIp    string             `json:"lastIp,omitempty" bson:"lastIp,omitempty"`
	LastLogin time.Time          `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
	Profile   *string            `bson:"profile"`
	Meta      Meta               `bson:"meta"`
}

type Role string

const (
	user       Role = "USER"
	admin      Role = "ADMIN"
	superAdmin Role = "SUPER_ADMIN"
)

type Profile struct {
	FirstName *time.Time `bson:"firstName"`
	LastName  *string    `bson:"lastName"`
	FullName  *time.Time `bson:"fullName"`
	Avatar    *string    `bson:"avatar"`
}
type Meta struct {
	CreatedAt time.Time `bson:"createdAt"`
	CreatedBy string    `bson:"createdBy"`
	UpdatedAt time.Time `bson:"updatedAt"`
	UpdatedBy string    `bson:"updatedBy"`
	DeletedBy string    `bson:"deletedBy"`
	Deleted   bool      `bson:"deleted"`
}

func NewUser(data *User) User {
	now := time.Now()
	defaultCreatedBy := data.Email
	return User{
		ID:        primitive.NewObjectID(),
		Email:     data.Email,
		Verified:  false,
		Roles:     []Role{user},
		LastIp:    "",
		LastLogin: now,
		Profile:   nil,
		Meta: Meta{
			CreatedAt: now,
			CreatedBy: defaultCreatedBy,
			UpdatedAt: now,
			UpdatedBy: defaultCreatedBy,
			DeletedBy: "",
			Deleted:   false,
		},
	}
}

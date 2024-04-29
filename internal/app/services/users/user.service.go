package users

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type UserService interface {
	Create(data User, ctx context.Context) (User, error)
	Get(id string, ctx context.Context) (User, error)
	Find(ctx context.Context) ([]User, error)
	Patch(id string, data User, ctx context.Context) (User, error)
	Delete(id string, ctx context.Context) (User, error)
}

type UserServiceImpl struct{}

func (u *UserServiceImpl) Create(data *User, ctx context.Context) (User, error) {
	// Create a new validator
	validate := validator.New()

	// Validate the data struct
	if err := validate.Struct(data); err != nil {
		fmt.Println(err)
		return User{}, err
	}

	data.Email = strings.ToLower(data.Email)
	// data.Roles = []string{"USER"}

	// db := connectors.DB
	newUser := NewUser(data)

	fmt.Println(newUser)
	return newUser, nil
}

// func Get(id string, ctx context.Context) (User, error) {
// 	db := connectors.DB

// 	// Convert the string to an ObjectId
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result bson.M
// 	err = db.Collection("Users").FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// No document was found
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return result, nil
// }

// func Find(ctx context.Context) ([]User, error) {
// 	db := connectors.DB
// 	cursor, err := db.Collection("Users").Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []bson.M
// 	if err := cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func Patch(id string, data bson.M, ctx context.Context) (User, error) {
// 	db := connectors.DB

// 	// Convert the string to an ObjectId
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	data["meta"] = bson.M{
// 		"createdAt": time.Now(),
// 		"createdBy": data["email"],
// 		"updatedAt": time.Now(),
// 		"updatedBy": "system",
// 		"deleted":   false,
// 	}

// 	var result bson.M

// 	err = db.Collection("Users").FindOneAndUpdate(ctx, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: data}},
// 		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// No document was found
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return result, nil
// }

// func Delete(id string, ctx context.Context) (User, error) {
// 	db := connectors.DB

// 	// Convert the string to an ObjectId
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result bson.M
// 	err = db.Collection("Users").FindOneAndUpdate(ctx, bson.D{{Key: "_id", Value: objectId}}, bson.D{{Key: "$set", Value: bson.M{"meta.deleted": true, "meta.deletedAt": time.Now(), "meta.deletedBy": "system"}}},
// 		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// No document was found
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return result, nil
// }

// Decode the data map into a User struct
// var user User
// bsonBytes, _ := bson.Marshal(data)
// bson.Unmarshal(bsonBytes, &user)

// // Validate the User struct
// err := validate.Struct(user)
// validationErrors := err.(validator.ValidationErrors)

// email, ok := data["email"].(string)
// if !ok || email == "" {
// 	return nil, errors.New("email is required")
// }

// // Validate email
// _, err := mail.ParseAddress(email)
// if err != nil {
// 	return nil, errors.New("invalid email")
// }

// password, ok := data["password"].(string)
// if !ok || password == "" {
// 	return nil, errors.New("password is required")
// }

// err = validations.Password(password)
// if err != nil {
// 	return nil, err
// }

// email = strings.ToLower(email)

// to encrypt password and test it
// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// if err != nil {
// 	return nil, err
// }

// data["email"] = email
// data["password"] = string(hashedPassword)

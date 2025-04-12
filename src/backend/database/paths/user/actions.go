package database_user

import (
	"context"
	"errors"
	authentication "template_backend/infrastructure/authentication"

	"github.com/go-kivik/kivik"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// CreateUser creates a new user in the database
func CreateUser(ctx context.Context, email string, password string) (*UserProfile, error) {
	uniqueID := uuid.NewString()
	hash, salt, err := authentication.CreatePassword(&password)
	if err != nil {
		log.Error().
			Str("email", email).
			Msg(err.Error())
		return nil, err
	}

	user := UserProfile{
		ID:    uniqueID,
		Hash:  hash,
		Salt:  salt,
		Email: email,
		Roles: TemporaryUser,
	}

	// Insert the document into the database
	rev, err := DatabaseUser.Put(ctx, user.ID, user)
	if err != nil {
		log.Error().
			Str("email", email).
			Msg(err.Error())

		return nil, err
	}

	log.Info().
		Str("rev", rev).
		Str("email", email).
		Msg("Created User")

	dbUser := FindUserById(ctx, &uniqueID)
	if dbUser == nil {
		log.Error().
			Str("email", email).
			Msg("Couldn't find newly created account")

		return nil, errors.New("couldn't find newly created account")
	}

	return dbUser, nil
}

func ChangeUserEmail(ctx context.Context, id *string, newEmail *string) (*UserProfile, error) {
	user := FindUserById(ctx, id)
	if user == nil {
		return nil, errors.New("couldn't find user")
	}
	user.Email = *newEmail
	_, err := DatabaseUser.Put(ctx, user.ID, user, kivik.Options{"_rev": user.Rev})
	if err != nil {
		return nil, errors.New("couldn't update user")
	}

	return user, nil
}

// ChangeUserPassword changes the password of a user in the database
func ChangeUserPassword(ctx context.Context, id *string, newPassword *string) (*UserProfile, error) {
	user := FindUserById(ctx, id)
	if user == nil {
		return nil, errors.New("couldn't find user")
	}

	hash, salt, err := authentication.CreatePassword(newPassword)
	if err != nil {
		log.Error().
			Str("id", *id).
			Msg(err.Error())
		return nil, err
	}

	user.Hash = hash
	user.Salt = salt

	_, err = DatabaseUser.Put(ctx, user.ID, user, kivik.Options{"_rev": user.Rev})
	if err != nil {
		return nil, errors.New("couldn't update user")
	}

	return user, nil
}

package database_user

import (
	"context"
	database_common "template_backend/database/common"
	authentication "template_backend/infrastructure/authentication"

	"github.com/rs/zerolog/log"
)

// createFindUserQuery creates a query to find a user in the database
func createFindUserQuery(email *string, id *string, fields []interface{}) database_common.Query {
	query := database_common.Query{
		Selector: map[string]interface{}{},
		Fields:   fields,
		Limit:    1,
	}

	if email != nil {
		query.Selector["email"] = map[string]interface{}{"$eq": *email}
	}
	if id != nil {
		query.Selector["_id"] = map[string]interface{}{"$eq": *id}
	}

	return query
}

// ExistsEmail checks if a user exists in the database by email
func ExistsEmail(ctx context.Context, email string) bool {
	emailInCache := Cache.getUserByEmail(&email)
	if emailInCache != nil {
		return true
	}

	query := createFindUserQuery(&email, nil, []interface{}{"email"})
	rows, err := DatabaseUser.Find(ctx, query)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user UserProfile
		if err := rows.ScanDoc(&user); err != nil {
			log.Error().Msg(err.Error())
		} else if user.Email == email {
			return true
		}
	}
	return false
}

// FindUserById finds a user in the database by id
func FindUserById(ctx context.Context, id *string) *UserProfile {
	userInCache := Cache.getUserById(id)
	if userInCache != nil {
		return userInCache
	}

	options := createFindUserQuery(nil, id, []interface{}{})
	rows, err := DatabaseUser.Find(ctx, options)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user UserProfile
		if err := rows.ScanDoc(&user); err != nil {
			log.Error().Msg(err.Error())
			return nil
		} else {
			Cache.setUser(&user)
			return &user
		}
	}
	return nil
}

// FindUserByEmail finds a user in the database by email
func FindUserByEmail(ctx context.Context, email *string) *UserProfile {
	userInCache := Cache.getUserByEmail(email)
	if userInCache != nil {
		return userInCache
	}

	options := createFindUserQuery(email, nil, []interface{}{})
	rows, err := DatabaseUser.Find(ctx, options)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user UserProfile
		if err := rows.ScanDoc(&user); err != nil {
			log.Error().Msg(err.Error())
			return nil
		} else {
			Cache.setUser(&user)
			return &user
		}
	}
	return nil
}

// AuthenticateUser authenticates a user in the database by email and password
func AuthenticateUser(ctx context.Context, email string, password string) *UserProfile {
	user := FindUserByEmail(ctx, &email)
	if user == nil {
		return nil
	}
	passwordMatches := authentication.CheckPassword(&password, &user.Hash, &user.Salt)
	if !passwordMatches {
		return nil
	}

	return user
}

// AuthenticateUserById authenticates a user in the database by id
func AuthenticateUserById(ctx context.Context, id string, password string) *UserProfile {
	user := FindUserById(ctx, &id)
	if user == nil {
		return nil
	}
	passwordMatches := authentication.CheckPassword(&password, &user.Hash, &user.Salt)
	if !passwordMatches {
		return nil
	}

	return user
}

package models

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type User struct {
	UID      string
	Username string
	Password string
}

// CreateUser: creates an user and add it to redis
func CreateUser(username string, password string) (*User, error) {
	if username != "" && password != "" {
		u, err := uuid.NewV4()
		if err == nil {
			user := &User{
				Username: user,
				Password: password,
				UID:      u.String(),
			}
			conn := RedisPool.Get()
			defer conn.Close()

			_, err := conn.Do("SADD", "Users", user.UID)
			if err == nil {
				_, err := c.Do("HMSET", user.UID, "Username", user.Username, "Password", user.Password)
				if err == nil {
					return user, nil
				}
				return nil, fmt.Errorf("Cannot set the user name")
			}
			return nil, fmt.Errorf("Cannot add the user")
		}
		return nil, fmt.Errorf("Cannot create UID")
	}
	return nil, fmt.Errorf("Users's name or password cannot be empty")
}

// GetUser: returns an user if it exists
func GetUser(UID string) (*User, error) {
	if UID != "" {
		c := RedisPool.Get()
		defer c.Close()

		userExist, err := redis.Int(c.Do("SISMEMBER", "Users", UID))
		if err == nil && userExist == 1 {
			username, err := redis.String(c.Do("HGET", UID, "Username"))
			if err == nil {
				user := &User{
					UID:      UID,
					Username: username,
				}
				return user, nil
			}
			return nil, fmt.Errorf("Cannot get username")
		}
		return nil, fmt.Errorf("User does not exists")
	}
	return nil, fmt.Errorf("UID not provided")
}

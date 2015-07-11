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

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

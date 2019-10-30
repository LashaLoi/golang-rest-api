package config

import "../customtypes"

// Config ...
type Config struct {
	Port  string
	Users []customtypes.User
}

// NewConfig ..
func NewConfig() Config {
	users := []customtypes.User{
		{
			ID:    1,
			Name:  "Aliaksei",
			Email: "lashalo11409@gmail.com",
			Phone: "+375 33 603 80 02",
		},
		{
			ID:    2,
			Name:  "David",
			Email: "david@gmail.com",
			Phone: "+375 11 222 33 44",
		},
	}

	config := Config{
		Port:  ":8080",
		Users: users,
	}

	return config
}

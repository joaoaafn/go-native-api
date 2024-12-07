package main

// Fake in memory database file
type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "example@email.com",
		Id:    "user1",
		Name:  "Test Man",
		Token: "123",
	},
	"user2": {
		Email: "user2@email.com",
		Id:    "user2",
		Name:  "Bill",
		Token: "456",
	},
}

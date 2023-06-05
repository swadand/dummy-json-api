package userModel

type User struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Age     int    `json:"age"`
	Likes   string `json:"like"`
}

func UserList() []User {
	userList := []User{
		{Name: "Ash", Website: "ash.gg", Age: 10, Likes: "animals"},
		{Name: "Raj", Website: "raj.in", Age: 23, Likes: "Silence"},
		{Name: "Yami", Website: "yami.to", Age: 20, Likes: "skincare"},
		{Name: "Abhishek", Website: "abhishek.cc", Age: 26, Likes: "comedy"},
		{Name: "Victoria", Website: "cookies.ru", Age: 27, Likes: "baking"},
	}
	return userList
}

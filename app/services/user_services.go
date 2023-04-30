package services

// User - A representation of a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetAllUsers - Returns all users
func GetAllUsers() ([]User, error) {
	// ユーザー一覧を取得するロジックを実装
	// ここでは、ダミーデータを返す
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
	}

	return users, nil
}
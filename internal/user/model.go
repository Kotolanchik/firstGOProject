package user

type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

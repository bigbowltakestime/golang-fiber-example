package models

// User는 사용자 정보를 나타내는 구조체입니다.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// 추가 필드들을 필요에 따라 정의할 수 있습니다.
}

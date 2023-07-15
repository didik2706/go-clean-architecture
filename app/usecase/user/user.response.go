package usecase

import (
	"latihan-restful-api-2/app/entity"
	"time"
)

type UserUsecaseResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(user entity.User) UserUsecaseResponse {
	return UserUsecaseResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUsersResponse(users []entity.User) []UserUsecaseResponse {
	responses := []UserUsecaseResponse{}
	for _, user := range users {
		responses = append(responses, UserUsecaseResponse{
			Id: user.Id,
			Name: user.Username,
			Username: user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return responses
}

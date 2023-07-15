package usecase

type UserUsecaseRequestCreate struct {
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required,min=8,max=20"`
}

type UserUsecaseRequestUpdate struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required,min=8,max=20"`
}

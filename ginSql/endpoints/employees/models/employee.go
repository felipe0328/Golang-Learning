package models

type GetEmployeeInput struct {
	ID int `uri:"id" binding:"required"`
}

package models

type Employee struct {
	ID                   int    `json:"id"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	FullName             string `json:"fullName"`
	JobTitle             string `json:"jobTitle"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Image                string `json:"image"`
	Country              string `json:"country"`
	City                 string `json:"city"`
	OnboardingCompletion int    `json:"onboardingCompletion"`
}

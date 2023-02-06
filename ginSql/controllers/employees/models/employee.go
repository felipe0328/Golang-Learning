package models

type Employee struct {
	ID                   int    `json:"id,omitempty"`
	FirstName            string `json:"firstName,omitempty"`
	LastName             string `json:"lastName,omitempty"`
	FullName             string `json:"fullName,omitempty"`
	JobTitle             string `json:"jobTitle,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Email                string `json:"email,omitempty"`
	Image                string `json:"image,omitempty"`
	Country              string `json:"country,omitempty"`
	City                 string `json:"city,omitempty"`
	OnboardingCompletion int    `json:"onboardingCompletion,omitempty"`
}

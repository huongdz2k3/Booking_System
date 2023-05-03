// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ent

type ChangePasswordInput struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type Jwt struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name             string `json:"name"`
	PhoneNumber      string `json:"phoneNumber"`
	Email            string `json:"email"`
	LicenseID        string `json:"licenseID"`
	Address          string `json:"address"`
	MembershipNumber *int   `json:"membershipNumber,omitempty"`
	Password         string `json:"password"`
}

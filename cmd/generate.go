package main

//go:generate fastjson_builder -type=GenerateUser
type GenerateUser struct {
	ID          string      `json:"id"`
	Credentials *Credential `json:"-"`
	Profile     *Profile
	Contacts    Contacts `json:"user_contacts"`
}

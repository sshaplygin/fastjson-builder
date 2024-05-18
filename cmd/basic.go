package main

import (
	"github.com/valyala/fastjson"
)

type Profile struct {
	FirstName string
	LastName  string
}

type Credential struct {
	Login    string
	Password string
}

type Contacts struct {
	Emails []string
}

//go:generate fastjson_builder -type=GenerateUser
type User struct {
	ID          string      `json:"id"`
	Credentials *Credential `json:"-"`
	Profile     *Profile
	Contacts    Contacts `json:"user_contacts"`
	Tags        []string
}

func GetUserVal(arena *fastjson.Arena, user *User) *fastjson.Value {
	if user == nil {
		return arena.NewNull()
	}

	obj := arena.NewObject()
	// arena.Reset()
	obj.Set("id", arena.NewString(user.ID))

	profileObj := arena.NewNull()
	if user.Profile != nil {
		profileObj = arena.NewObject()
		profileObj.Set("first_name", arena.NewString(user.Profile.FirstName))
		profileObj.Set("last_name", arena.NewString(user.Profile.LastName))
	}

	obj.Set("Profile", profileObj)

	contactsObj := arena.NewObject()
	emailArray := arena.NewArray()

	for i, email := range user.Contacts.Emails {
		emailArray.SetArrayItem(i, arena.NewString(email))
	}

	if len(user.Contacts.Emails) == 0 {
		emailArray = arena.NewNull()
	}

	contactsObj.Set("Emails", emailArray)
	obj.Set("user_contacts", contactsObj)

	return obj
}

package models

import "time"

type User struct{
	Name string
	BirthDate time.Time
	Age int
	Profile Profile
}

type Profile struct {
  Work    string
  Hobbies []string
}
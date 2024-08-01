package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type UniqueID int

type Email string

type SignUpRequest struct {
	Email    Email
	Password string
}

func SignUpToUser(r SignUpRequest) User {
	return User{
		Email:    r.Email,
		Password: r.Password,
	}
}

type User struct {
	ID       UniqueID `json:"id" ` // validate:"required"
	Email    Email    `json:"email" validate:"required"`
	Name     string   `json:"name,omitempty"`
	Password string   `json:"-" validate:"required"`
	Age      int      `json:"age"`
}

func Validate(tv any) error {
	t := reflect.TypeOf(tv)

	v := reflect.ValueOf(tv)

	if k := t.Kind(); k != reflect.Struct {
		return fmt.Errorf("unsupported kind: %v", k)
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag, ok := f.Tag.Lookup("validate")
		if !ok {
			continue
		}

		if tag != "required" {
			return errors.New("invalid tag value")
		}

		switch k := f.Type.Kind(); k {
		case reflect.String:
			strV := strings.TrimSpace(v.Field(i).String())
			if strV == "" {
				return fmt.Errorf("field %v (%v) is empty", f.Name, f.Type)
			}
		default:
			return fmt.Errorf("unsupported kind: %v", k)
		}
	}

	return nil
}

func main() {
	u := User{
		ID:    1,
		Email: "test@example.com",
	}

	fmt.Println(Validate(u))

	u.Password = "123"

	fmt.Println(Validate(u))
}

func reflectExample() {
	u := User{
		ID: 1,
	}

	t := reflect.TypeOf(u)

	fmt.Println(t)
	fmt.Println(t.Kind())

	f, _ := t.FieldByName("ID")
	fmt.Printf("%+v", f)
}

func exampleMarshal() {
	u := User{
		ID: 1,
	}

	bytes, _ := json.Marshal(u)

	fmt.Println(string(bytes))
}

func exampleRecursion() {
	fmt.Println(st(2, 3))
}

func st(num int, s uint) int {
	if s == 0 {
		return 1
	}

	if s == 1 {
		return num
	}

	return num * st(num, s-1)
}

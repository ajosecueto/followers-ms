// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package customTypes

type Profile struct {
	UserID   string  `json:"userId"`
	Username string  `json:"username"`
	Names    string  `json:"names"`
	Photo    *string `json:"photo"`
	Verified *bool   `json:"verified"`
}

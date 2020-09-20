package models

type Command struct {
	Name        string
	Slug        string
	Description string
	Keyboard    string
	Aliases     []string
	Func        interface{}
	Disabled    bool
}

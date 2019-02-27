package entities

import "fmt"

type Language struct {
	Code string
	Name string
}

func NewLanguage() Language {
	// new structure with default values ( 'zero values' )
	return Language{}

//	// new structure with specific values 
//	return Language{
//		Code: "??",
//		Name: "?????",
//	}
}

func (this Language) String() string {
	return fmt.Sprintf(
		"[%s, %s]",
		this.Code,
		this.Name)
}

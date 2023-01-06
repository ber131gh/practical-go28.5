package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(text string) *Student {
	textSlice := strings.Split(text, " ")
	name := textSlice[0] // получаем поле name

	age, err := strconv.Atoi(textSlice[1])
	if err != nil {
		fmt.Println(err)
	}
	grade, err := strconv.Atoi(textSlice[2])
	if err != nil {
		fmt.Println(err)
	}

	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
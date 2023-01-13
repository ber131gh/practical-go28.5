package repo

import (
	"fmt"
	"strconv"
	"test/internal/student"
)

type studentName string

type StudentStorage map[studentName]*student.Student

func NewStudentStorage() StudentStorage {
	return make(map[studentName]*student.Student)
}

func (ss StudentStorage) Put(s *student.Student) {
	ss[studentName(s.Name)] = s
}

func (ss StudentStorage) Get() {
	for _, k := range ss {
		fmt.Println(k.Name + " " + strconv.Itoa(k.Age) + " " + strconv.Itoa(k.Grade))
	}
}

/*
func (ss StudentStorage) Get(stName string) (*Student, error) {
	studentName := studentName(stName)
	s, ok := ss[studentName]
	if !ok {
		return nil, fmt.Errorf("no such student")
	}
	return s, nil
}
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"test/pkg/repo"
	"test/pkg/student"
)

func main() {
	ss := repo.NewStudentStorage() // создаём хранилище студентов

	reader := bufio.NewReader(os.Stdin) // создаем reader (os.Stdin)
	for {
		fmt.Print("Введите через пробел данные студента: имя, возраст, уровень: ")
		text, err := reader.ReadString('\n') // считываем текст с командной строки
		if err != nil {                      // обрабатываем ошибку ввода
			if err == io.EOF { // если прерывание считывания (Ctrl+z)
				break // то выход из цикла
			}
			panic(err)
		}
		text = strings.TrimSpace(text) // удаляем пробельные символы с начала и конца строки
		s := student.NewStudent(text)
		ss.Put(s) // добавляем студента в хранилище
		//fmt.Println(ss.get(s.name))
	}
	fmt.Println("Студенты из хранилища:")
	ss.Get() // выводим все данные из хранилища
}

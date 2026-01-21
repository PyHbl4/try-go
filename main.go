package main

import "fmt"

// User — пример структуры с полями разного поведения при копировании
type User struct {
	ID   int
	Name string
	Tags []string
	Meta map[string]string
}

// RenameByValue получает User по значению — изменение не повлияет на оригинал
func RenameByValue(u User, name string) {
	u.Name = name
}

// RenameByPointer получает указатель — изменение затронет оригинал
func RenameByPointer(u *User, name string) {
	u.Name = name
}

// AddTagByValue получает копию User и добавляет тег в копию
func AddTagByValue(u User, tag string) {
	u.Tags = append(u.Tags, tag)
}

// AddTagByPointer получает указатель и добавляет тег в оригинал
func AddTagByPointer(u *User, tag string) {
	u.Tags = append(u.Tags, tag)
}

// AddMetaByValue получает копию User и добавляет/создаёт map в копии
func AddMetaByValue(u User, key, value string) {
	if u.Meta == nil {
		u.Meta = make(map[string]string)
	}
	u.Meta[key] = value
}

// AddMetaByPointer получает указатель и добавляет/создаёт map в оригинале
func AddMetaByPointer(u *User, key, value string) {
	if u.Meta == nil {
		u.Meta = make(map[string]string)
	}
	u.Meta[key] = value
}

func main() {
	// Начальное состояние
	user := User{ID: 1, Name: "Alice"}
	fmt.Println("Начальное состояние:", user)

	// 1) Поменяем имя через передачу по значению — оригинал не поменяется
	RenameByValue(user, "Bob")
	fmt.Println("\nПосле RenameByValue(user, \"Bob\")")
	fmt.Println("Ожидаем: имя НЕ изменилось (передача по значению)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// 2) Поменяем имя через указатель — оригинал поменяется
	RenameByPointer(&user, "Charlie")
	fmt.Println("\nПосле RenameByPointer(&user, \"Charlie\")")
	fmt.Println("Ожидаем: имя изменилось (передача по указателю)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// 3) Добавим тег через передачу по значению — в копии добавится, оригинал останется прежним
	AddTagByValue(user, "golang")
	fmt.Println("\nПосле AddTagByValue(user, \"golang\")")
	fmt.Println("Ожидаем: теги НЕ добавлены в оригинал (append в копии)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// 4) Добавим тег через указатель — оригинал получит тег
	AddTagByPointer(&user, "programmer")
	fmt.Println("\nПосле AddTagByPointer(&user, \"programmer\")")
	fmt.Println("Ожидаем: теги добавлены в оригинал (изменение через указатель)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// 5) Добавим meta через передачу по значению — если оригинал.Meta == nil, то map создастся в копии и не попадёт в оригинал
	AddMetaByValue(user, "role", "admin")
	fmt.Println("\nПосле AddMetaByValue(user, \"role\", \"admin\")")
	fmt.Println("Ожидаем: meta НЕ добавлена в оригинал (map создана в копии)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// 6) Добавим meta через указатель — map создаётся/обновляется в оригинале
	AddMetaByPointer(&user, "department", "engineering")
	fmt.Println("\nПосле AddMetaByPointer(&user, \"department\", \"engineering\")")
	fmt.Println("Ожидаем: meta добавлена в оригинал (через указатель)")
	fmt.Printf("Текущее состояние: %+v\n", user)

	// Итоговый вывод
	fmt.Println("\nИтог (подробно):")
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Tags: %v\n", user.Tags)
	fmt.Printf("Meta: %v\n", user.Meta)
}

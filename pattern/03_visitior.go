package pattern

import (
	"fmt"
	"log"
)

/*
Паттерн Посетитель – это ценный инструмент в арсенале любого программиста, работающего с объектно-ориентированными языками
программирования. В Golang он может быть реализован с использованием интерфейсов и классов-посетителей для выполнения различных
операций с разными типами объектов. Метод Accept, используемый в классах объектов, позволяет легко добавлять новые методы
обработки без изменения классов объектов. Если вы до сих пор не использовали паттерн Посетитель в своей работе,
то возможно, вам стоит обратить на него внимание.
*/

// Интерфейс животных
type Animal interface {
	Accept(visitor VisitorAnimal) string
}

// Интерфейс посетителя
type VisitorAnimal interface {
	DogSound(d *Dog) string
	CatSound(c *Cat) string
}

// Структура собаки
type Dog struct {
	name string
}

func (d *Dog) Accept(visitor VisitorAnimal) string {
	return visitor.DogSound(d)
}

// Структура кота
type Cat struct {
	name string
}

func (c *Cat) Accept(visitor VisitorAnimal) string {
	return visitor.CatSound(c)
}

// Реализация интерфейса VisitorAnimal
type VisitorSoundAnimal struct{}

func (v *VisitorSoundAnimal) DogSound(d *Dog) string {
	return fmt.Sprintf("%s Woof", d.name)
}

func (v *VisitorSoundAnimal) CatSound(c *Cat) string {
	return fmt.Sprintf("%s Moew", c.name)
}

func StartVisitorPattern() {
	dog := &Dog{"Ikar"}
	cat := &Cat{"Gosha"}

	v := &VisitorSoundAnimal{}

	log.Print(dog.Accept(v))
	log.Print(cat.Accept(v))
}

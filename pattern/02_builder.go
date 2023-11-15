package pattern

import (
	"log"
	"math/rand"
)

/*
Паттерн Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Когда стоит использовать?

    Если у Вас есть один конструктор с десятью опциональными параметрами. Его неудобно вызывать, поэтому
	Вы создали ещё десять конструкторов с меньшим количеством параметров.
	Всё, что они делают — это переадресуют вызов к базовому конструктору, подавая какие-то значения по умолчанию в параметры,
	которые пропущены в них самих.

	Если создание нескольких представлений объекта состоит из одинаковых этапов, которые отличаются в деталях.

	Если Вам нужно собирать сложные составные объекты, например, деревья Компоновщика.

Плюсы:

	Позволяет создавать объекты пошагово.
	Позволяет использовать один и тот же код для создания различных объектов.
	Изолирует сложный код сборки объекта от его основной бизнес-логики.

Минусы:

	Усложняет код программы из-за введения дополнительных классов.
	Клиент может оказаться привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть метода получения результата.
*/

// Создаем интерфейс строитель
type Builder interface {
	setSize()
	setName()
	setTypeFile()
	file() *File
}

func getBuilder(file string) Builder {
	if file == "video" {
		return &Video{}
	} else {
		return &Photo{}
	}
}

// Строитель 1
type Photo struct {
	sizePhoto int
	namePhoto string
	typePhoto string
}

func (p *Photo) setSize() {
	p.sizePhoto = rand.Intn(20)
}

func (p *Photo) setName() {
	p.namePhoto = "myFamily"
}

func (p *Photo) setTypeFile() {
	p.typePhoto = ".jpeg"
}

func (p *Photo) file() *File {
	return &File{
		p.sizePhoto,
		p.namePhoto,
		p.typePhoto,
	}
}

// Строитель 2
type Video struct {
	sizeVideo int
	nameVideo string
	typeVideo string
}

func (v *Video) setSize() {
	v.sizeVideo = rand.Intn(40)
}

func (v *Video) setName() {
	v.nameVideo = "MyFirstSteps"
}

func (v *Video) setTypeFile() {
	v.typeVideo = ".mpeg"
}

func (v *Video) file() *File {
	return &File{
		v.sizeVideo,
		v.nameVideo,
		v.typeVideo,
	}
}

// Продукт строительства
type File struct {
	sizeFile int
	nameFile string
	typeFile string
}

func (f *File) WriteFileData() {
	log.Printf(`
	Size: %d
	Name: %s
	Type: %s
	`, f.sizeFile, f.nameFile, f.typeFile)
}

// Директор
type Develop struct {
	builder Builder
}

func (d *Develop) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Develop) BuildFile() *File {
	d.builder.setSize()
	d.builder.setName()
	d.builder.setTypeFile()
	return d.builder.file()
}

func StartBuilderPattern() {
	d := Develop{}
	v := getBuilder("video")
	p := getBuilder("photo")

	d.SetBuilder(v)
	d.BuildFile()
	d.builder.file().WriteFileData()

	d.SetBuilder(p)
	d.BuildFile()
	d.builder.file().WriteFileData()
}

package pattern

import "log"

/*
	Паттерн Facade инкапсулирует сложную подсистему в единственный интерфейсный объект.
	Это позволяет сократить время изучения подсистемы, а также способствует уменьшению степени
	связанности между подсистемой и потенциально большим количеством клиентов.
	С другой стороны, если фасад является единственной точкой доступа к подсистеме,
	то он будет ограничивать возможности, которые могут понадобиться "продвинутым" пользователям.

	В данном примере реализовано включение музыки в музыкальном сервисе.
*/

// Создаем структуру Music
type Music struct {
	name       string
	playingNow bool
	username   string
}

// Включение музыки
func (m *Music) StartMusic(name, username string, flag bool) {
	m.name = name
	m.playingNow = flag
	m.username = username
	log.Printf("Start %s Thanks %s", name, username)
}

// Смена музыки
func (m *Music) NewMusicStart(name, namePerson string) {
	if m.playingNow {
		log.Printf("Stopped %s ... Sorry %s", m.name, m.username)
	}
	m.name = name
	m.username = namePerson
	m.playingNow = true
	log.Printf("Start %s! Thanks %s", name, namePerson)
}

// Создание структуры человек
type Person struct {
	namePerson string
	songName   string
}

// Определяем человека
func (p *Person) SetPerson(name string, songName string) {
	p.namePerson = name
	p.songName = songName
}

// Включение музыки человеком
func (p *Person) OnMusic() string {
	return p.songName
}

// Структура музыкального сервиса
type MusicService struct {
	*Music
	*Person
}

// Создание нового муз. сервиса
func NewMusicService() *MusicService {
	return &MusicService{
		&Music{},
		&Person{},
	}
}

// Точка запуска сервиса
func (ms *MusicService) Start() {
	ms.SetPerson("Steve", "Numb")
	ms.StartMusic(ms.songName, ms.namePerson, true)
	ms.SetPerson("Lora", "SAD!")
	ms.NewMusicStart(ms.songName, ms.namePerson)
}

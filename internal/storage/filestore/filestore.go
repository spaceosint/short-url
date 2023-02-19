package filestore

import (
	"bufio"
	"github.com/segmentio/encoding/json"
	"os"
)

type Event struct {
	ID          uint   `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
type FileStore struct {
}
type Memory interface {
	GetOriginalURL(identifier string, filePath string) (string, error)
	AddNewLink(newID uint, shortURL string, originalURL string, filePath string) error
	GetNewID(filePath string) uint
	GetAllByPath(filePath string) []Event
}
type Producer interface {
	WriteEvent(event *Event) // для записи события
	Close() error            // для закрытия ресурса (файла)
}

type Consumer interface {
	ReadEvent() (*Event, error) // для чтения события
	Close() error               // для закрытия ресурса (файла)
}

type producer struct {
	file *os.File
	// добавляем writer в Producer
	writer *bufio.Writer
}

func NewProducer(filename string) (*producer, error) {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return &producer{
		file: file,
		// создаём новый Writer
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *producer) WriteEvent(event *Event) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	// записываем событие в буфер
	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	// добавляем перенос строки
	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	// записываем буфер в файл
	return p.writer.Flush()
}

type consumer struct {
	file *os.File
	// добавляем reader в Consumer
	reader *bufio.Reader
}

func NewConsumer(filename string) (*consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return &consumer{
		file: file,
		// создаём новый Reader
		reader: bufio.NewReader(file),
	}, nil
}

func (c *consumer) ReadEvent() (*Event, error) {
	// читаем данные до символа переноса строки
	data, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	// преобразуем данные из JSON-представления в структуру
	event := Event{}
	err = json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (f *FileStore) AddNewLink(newID uint, shortURL string, originalURL string, filePath string) error {

	var evn = Event{
		ID: newID, OriginalURL: originalURL, ShortURL: shortURL,
	}
	file, err := NewProducer(filePath)
	if err != nil {
		return err
	}
	err = file.WriteEvent(&evn)
	if err != nil {
		return err
	}

	return nil
}
func (f *FileStore) GetOriginalURL(identifier string, filePath string) (string, error) {
	file, err := NewConsumer(filePath)
	if err != nil {
		return "", err
	}
	for {
		ev, err := file.ReadEvent()
		if err != nil {
			break
		}
		if ev.ShortURL == identifier {
			return ev.OriginalURL, nil
		}
	}
	return "", err
}

func (f *FileStore) GetNewID(filePath string) uint {
	var maxID uint = 10000

	cons, err := NewConsumer(filePath)
	if err != nil {
		return maxID
	}

	for {
		link, err := cons.ReadEvent()
		if err != nil {
			break
		}
		if link.ID > maxID {
			maxID = link.ID
		}
	}
	return maxID + 1
}

func (f *FileStore) GetAllByPath(filePath string) []Event {
	var users []Event
	cons, err := NewConsumer(filePath)
	if err != nil {
		return nil
	}

	for {
		link, err := cons.ReadEvent()
		if err != nil {
			break
		}
		users = append(users, *link)
	}
	return users
}

package filestore

import (
	"bufio"
	"fmt"
	"github.com/segmentio/encoding/json"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/pkg/shorten"
	"io"
	"os"
	"strconv"
)

type Event struct {
	ID          uint   `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
type FileStore struct {
	cfg config.Config
}

func NewInFileStore(config config.Config) *FileStore {
	return &FileStore{
		cfg: config,
	}
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

func (f *FileStore) GetShortURL(newUserURL string) (string, error) {

	newID := f.GetNewIDFile(f.cfg.FileStoragePath)

	shortURL := shorten.ShortenURL(newID)
	fmt.Println(shortURL)
	var evn = Event{
		ID: newID, OriginalURL: newUserURL, ShortURL: shortURL,
	}
	fmt.Println(evn)
	prod, err := NewProducer(f.cfg.FileStoragePath)
	if err != nil {
		return "", err
	}
	defer prod.file.Close()
	err = prod.WriteEvent(&evn)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return f.cfg.BaseURL + "/" + shortURL, nil
}
func (f *FileStore) GetOriginalURL(identifier string) (string, error) {

	cons, err := NewConsumer(f.cfg.FileStoragePath)
	defer cons.file.Close()
	if err != nil {
		return "", err
	}
	for {
		ev, err := cons.ReadEvent()
		if err != nil {
			break
		}
		if ev.ShortURL == identifier {
			return ev.OriginalURL, nil
		}
	}
	return "", err
}

func (f *FileStore) GetNewIDFile(filePath string) uint {
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

func (f *FileStore) GetAllByPathFile(filePath string) []Event {
	var users []Event
	cons, err := NewConsumer(filePath)
	defer cons.file.Close()
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

func (f *FileStore) GetAll() (map[string]string, error) {
	events := make(map[string]string)
	cons, err := NewConsumer(f.cfg.FileStoragePath)
	if err != nil {
		return events, err
	}
	defer cons.file.Close()

	for {
		event, err := cons.ReadEvent()
		if err != nil {
			if err == io.EOF {
				break
			}
			return events, err
		}
		key := strconv.Itoa(int(event.ID))
		events[key] = event.OriginalURL
	}
	fmt.Println(events)
	return events, nil
}

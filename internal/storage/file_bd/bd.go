package file_bd

//
//import (
//	"os"
//)
//
//type producer struct {
//	file *os.File // файл для записи
//}
//
//func NewProducer(filename string, data map[string]string) (any, error) {
//	// открываем файл для записи в конец
//	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
//	if err != nil {
//		return nil, err
//	}
//
//	return file, nil
//}
//
//func (p *producer) Close() error {
//	// закрываем файл
//	return p.file.Close()
//}
//
//type fileBd struct {
//	file *os.File // файл для чтения
//}
//
//func NewConsumer(filename string) (*file_bd, error) {
//	// открываем файл для чтения
//	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0777)
//	if err != nil {
//		return nil, err
//	}
//
//	return &file_bd{file: file}, nil
//}
//
//func (c *file_bd) Close() error {
//	// закрываем файл
//	return c.file.Close()
//}

package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type MySQLStorage struct {
	data string
}

func (s *MySQLStorage) Read() string      { return s.data }
func (s *MySQLStorage) Write(data string) { s.data = data }
func NewMySQL() *MySQLStorage             { return &MySQLStorage{data: ""} }

type BlackHole struct{}

func (s *BlackHole) Read() string      { panic("cannot get anything from blackhole") }
func (s *BlackHole) Write(data string) {}
func NewBlackHole() *BlackHole         { return &BlackHole{} }

func Store(data string, reader Reader, writer Writer) {
	writer.Write(data)
	fmt.Println(reader.Read())
}

func main() {
	reader := NewMySQL()
	writer := NewMySQL()
	Store("data", reader, writer)
}

func FixStore(data string) {
	reader := NewMySQL()
	writer := NewMySQL()
	writer.Write(data)
	fmt.Println(reader.Read())
}

func Store(data string, reader Reader, writer Writer) {
	writer.Write(data)

}

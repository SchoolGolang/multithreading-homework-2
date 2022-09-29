package http

import (
	"multithreading/multithreading-homework-2/message"
	"multithreading/multithreading-homework-2/message/http"
)

type Processor[U any, T http.Message[U]] struct {
	chGet  <-chan message.Message[U]
	chPost <-chan message.Message[U]
}

func (p Processor[U, T]) ProcessMassages() {

}

func NewHTTPProcessor[U any, T http.Message[U]](chGet, chPost <-chan message.Message[U]) Processor[U, T] {
	return Processor[U, T]{
		chGet:  chGet,
		chPost: chPost,
	}
}

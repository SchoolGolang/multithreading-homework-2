package http

import (
	"fmt"
	"multithreading/homework2/logger"
	"multithreading/homework2/message/http"
)

type Receiver[U any, T http.Message[U]] struct {
	ch chan<- http.Message[U]
}

func (r Receiver[U, T]) PutMessage(message http.Message[U]) {
	logger.GetInstance().Info(fmt.Sprintf("Received a message: %s", message))
	r.ch <- message
}

func NewHTTPReceiver[U any, T http.Message[U]](ch chan<- http.Message[U]) Receiver[U, T] {
	return Receiver[U, T]{
		ch: ch,
	}
}

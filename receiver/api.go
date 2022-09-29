package receiver

import "multithreading/homework2/message"

type Receiver[U any, T message.Message[U]] interface {
	PutMessage(T)
}

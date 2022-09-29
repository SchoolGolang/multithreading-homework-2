package receiver

import "multithreading/multithreading-homework-2/message"

type Receiver[U any, T message.Message[U]] interface {
	PutMessage(T)
}

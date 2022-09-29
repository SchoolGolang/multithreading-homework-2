package processor

import "multithreading/homework2/message"

type Processor[U any, T message.Message[U]] interface {
	ProcessMassage()
}

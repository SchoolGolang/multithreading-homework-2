package processor

import "multithreading/multithreading-homework-2/message"

type Processor[U any, T message.Message[U]] interface {
	ProcessMassage()
}

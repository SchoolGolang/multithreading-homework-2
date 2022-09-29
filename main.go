package main

import "multithreading/multithreading-homework-2/application"

// Agenda: We have application that processes some user requests.
// Receiver structs are needed for messages routing after they come.
// Processor structs are needed to process messages. After processing Processor must give log of success or error in console.
// mock.EmulateUser(receiver, MessageType, ctx) emulates user that sends some requests to your receiver
//
// Task: Implement processing with channels in http.Processor method ProcessMassages().
// Set different processing result for GET and POST messages, do not forget about error handling.
// Receiver-Processor pair must use same channel, where Receiver puts Messages
// in channel and Processor pops and processes it. In this case processing of GET and POST messages is separated.
// You can use two receivers at once.
// You can use channels, select, ctx and waitgroups.
//
// Disclaimer: You are allowed to edit application.go, processor.go
func main() {
	application.Run()
}

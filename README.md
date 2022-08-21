# go-kafka-messaging
A Sample GoLang application for learning purposes, understanding GoLang syntax and semantics. The application relies on Kafka for implementing messaging patterns. Please find the below sections to read more on the core functionalities of this application.

![](draw.png?raw=true)

## Prerequisites

* GoLang v1.18.5
* Kafka Cluster

## What are the basic functionalities of this application?
* The Sender application receives input from the command line and sends the message to the receiver application via Kafka topics.
* Receiver application contains multiple threads, each acting as a receiver which is connected to a kafka topic for receiving messages.
* Reciever application accepts a parameter - receiver count. Based on that it would spin up the receiver threads.
* User can type in to Sender application with messages like:
  * Broadcast message to all receivers: "@all: broadcast test message", so all receivers will receive the message and prints it out to the console.
  * Targeted message: "@receiver1: targeted test message", so only receiver1 would receive the message and print out.
  * If message format is invalid, the application should print out "invalid message"
  * If no receiver can be found, then application should print out "no receiver found for @receiverX".
  * When receiver receives message, it should print out messages to the console.
  * If user type in console "@receiver1#history" then application should print all received messages in the format mentioned before.
* If user type in console "exit" then application should stop and print a summary of amount of messages received per each receiver.

### How to Run the applications

You can run sender/receiver application by following command,

For starting sender application.

```bash
$ ./scripts/run.sh SENDER
```

For starting receiver application
  
```bash
$ ./scripts/run.sh RECEIVER <CONSUMER_THREAD_COUNT>
```



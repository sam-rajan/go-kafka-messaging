# go-kafka-messaging
Sample GoLang application which produces messages to and consumes from an Apache Kafka Cluster. And this sort of 'get your hands dirty' project or a project for getting myself acquainted with Golang and hence learning the basic and advanced features of Go Programming Language.  Please find the below sections to read more on the core functionalities of this application.

![](draw.png?raw=true)

## Prerequisites

* GoLang v1.18.5
* Kafka Cluster

## What are the basic functionalities of this application?
* There are mainly three components for this application, the producer(sender), the consumer(receiver) and the kafka Broker
* The Sender application receives input messages from the command line and sends the message to the receiver application via Kafka topics.
* Receiver application contains multiple threads, each acting as a receiver which is connected to a kafka topic for receiving messages.
* Reciever application accepts a parameter - receiver count. Based on that it would spin up the receiver threads.
* User can type in to Sender application with messages with predefined formats:
  * Broadcast message to all receivers: "@all: broadcast test message", so all receivers will receive the message and prints it out to the console.
  * Targeted message: "@receiver1: targeted test message", so only receiver1 would receive the message and writes that message to console.
  * If message format is invalid, the application should print "invalid message"
  * If no receiver can be found, then application should print "no receiver found for @receiverX".
  * When receiver receives message, it should print the messages to the console.
  * If user type in console "@receiver1#history" then application should print all received messages in the format mentioned before.
* If user type in console "exit" then application should stop and print a summary of amount of messages received per each receiver.
* Commands can be executed either by producer itself or it can send to other consumers. For example, "@receiver1 exit" will send the exit command to receiver1
* Application accepts predefined commands like history, exit etc.

### How to Run the applications

You can run sender/receiver application by following command,

For starting sender application.

```bash
$ ./scripts/run.sh SENDER <DATA_FORMAT>
```
Data format can be JSON, AVRO and PROTOBUF

For starting receiver application
  
```bash
$ ./scripts/run.sh RECEIVER <CONSUMER_THREAD_COUNT>
```



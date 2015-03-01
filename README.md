# Gluttony

Insasiatable tasks consumer and executor :)

Current state: **In development...**

# What is Gluttony?

Gluttony is a task executor that uses a broker to comunicate a client with a bunch of workers.

The main objective of this project is to let a client send a task to a queue, for example a mail
notification, and this task will be consumed and executed by Gluttony

Gluttony isn't bounded to any broker, right now there is just one connector: RabbitMQ, but it's
open to any other connectors like SQS, Redis, etc.

# What do i need to use Gluttony?

All you need is,

* Go (1.4)

and,

```go get github.com/ethervoid/gluttony```

# What brokers can i use?

Right know Gluttony only supports ```RabbitMQ```.

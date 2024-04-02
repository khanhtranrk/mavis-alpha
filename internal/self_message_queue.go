package handle

type Sef struct {
  amqpURL string
  group *Group // to send message to master
}

func NewBroker(group *Group) *Broker {
  return &Broker{group: group}
}

func ()

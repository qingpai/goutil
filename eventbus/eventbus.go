package eventbus

import "github.com/asaskevich/EventBus"

var bus EventBus.Bus

func Init() {
	bus = EventBus.New()
}

func GetBus() EventBus.Bus {
	return bus
}

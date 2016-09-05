package eventemitter

type EventEmitter struct {
	listeners map[string]*eventListener
}

type eventListener struct {
	ch        chan *Event
	callbacks []*EventCallback
}

type Event struct {
	Name string
	Data map[string]interface{}
}

type EventCallback func(event *Event)

func New() *EventEmitter {
	return NewEventEmitter()
}

func NewEventEmitter() *EventEmitter {
	eventEmitter := EventEmitter{}
	eventEmitter.listeners = make(map[string]*eventListener)
	return &eventEmitter
}

func newEventListener() *eventListener {
	listener := eventListener{}
	listener.ch = make(chan *Event)
	//listener.callbacks = make([]*EventCallback)
	return &listener
}

func newEvent(name string, data map[string]interface{}) *Event {
	event := Event{}
	event.Name = name
	event.Data = data
	return &event
}

func (this *EventEmitter) On(name string, handle EventCallback) {
	listener, ok := this.listeners[name]
	if !ok {
		listener = newEventListener()
		this.listeners[name] = listener
	}
	listener.callbacks = append(listener.callbacks, &handle)
	go func() {
		for {
			event := <-listener.ch
			if event == nil {
				break
			}
			handle(event)
		}
	}()
}

func (this *EventEmitter) Emit(name string, data map[string]interface{}) {
	event := newEvent(name, data)
	listener, ok := this.listeners[name]
	if !ok {
		return
	}
	listener.ch <- event
}

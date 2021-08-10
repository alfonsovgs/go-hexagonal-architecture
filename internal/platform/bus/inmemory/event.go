package inmemory

import (
	"context"

	"github.com/alfonsovgs/go-hexagonal-architecture/kit/event"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers map[event.Type][]event.Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, ev := range events {
		handlers, ok := b.handlers[ev.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, ev)
		}
	}

	return nil
}

func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evtType]

	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}

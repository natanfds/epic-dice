package ws

import "sync"

type hub struct {
	channels map[string]*Channel
	mu       sync.RWMutex
}

func (h *hub) GetOrCreateChannel(channelName string) *Channel {
	h.mu.Lock()
	defer h.mu.Unlock()

	if channel, exists := h.channels[channelName]; exists {
		return channel
	}

	channel := NewChannel(channelName)
	h.channels[channelName] = channel

	return channel
}

func (h *hub) RemoveChannel(channelName string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.channels, channelName)
}

func NewHub() *hub {
	return &hub{
		channels: make(map[string]*Channel),
	}
}

var Hub = NewHub()

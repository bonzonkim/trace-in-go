package network

type Router struct {
	*Network	
}

func newRouter(n *Network) {
	s := &Router{n}

	n.Router(GET, "/send", s.send)
	n.Router(GET, "/send-with-tag", s.sendWithTag)
	n.Router(GET, "/send-with-child", s.sendWithChild)

	n.Router(GET, "/receive-from-other-host", s.defaultHandler)
	n.Router(GET, "/send-other-host", s.defaultHandler)
	n.Router(GET, "/receive-two-from-other-host", s.defaultHandler)

	n.Router(GET, "/send-for-panic", s.defaultHandler)
	n.Router(GET, "/receive-for-error", s.defaultHandler)

	n.Router(GET, "/send-for-baggage", s.defaultHandler)
	n.Router(GET, "/receive-for-baggage", s.defaultHandler)
}

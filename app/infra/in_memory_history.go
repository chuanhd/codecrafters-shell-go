package infra

type HistoryStore interface {
	Add(line string)
	List() []string
}

type InMemoryHistory struct {
	lines []string
}

func NewInMemoryHistory() *InMemoryHistory {
	return &InMemoryHistory{
		lines: make([]string, 0),
	}
}

func (h *InMemoryHistory) Add(line string) {
	h.lines = append(h.lines, line)
}

func (h *InMemoryHistory) List() []string {
	return h.lines
}

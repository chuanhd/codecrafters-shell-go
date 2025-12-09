package infra

type HistoryStore interface {
	Add(line string)
	List() []string
	Get(idx int) string
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

func (h *InMemoryHistory) Get(idx int) string {
	if idx < 0 || idx >= len(h.lines) {
		return ""
	}

	return h.lines[idx]
}

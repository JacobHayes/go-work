package work

type Work struct {
	Done    map[string]bool
	Started map[string]bool
	Queued  map[string]bool
}

func (w *Work) Queue(item string) {
	if !w.Done[item] && !w.Started[item] {
		w.Queued[item] = true
	}
}

func (w *Work) Start(item string) {
	w.Started[item] = true
	delete(w.Queued, item)
}

func (w *Work) Finish(item string) {
	w.Done[item] = true
	delete(w.Started, item)
}

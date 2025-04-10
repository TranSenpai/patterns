package factory

type notifyFactory struct {
	lst map[string]Notify
}

func NewFactory() *notifyFactory {
	return &notifyFactory{
		lst: map[string]Notify{},
	}
}

func (n notifyFactory) GetNotify(kind string) Notify {
	return n.lst[kind]
}

func (n notifyFactory) SetNotify(kind string, not Notify) {
	n.lst[kind] = not
}

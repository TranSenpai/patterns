package builder

type FactoryBuilder struct {
	lst map[string]IBuilder
}

func NewFactoryBuilder() *FactoryBuilder {
	return &FactoryBuilder{
		lst: make(map[string]IBuilder),
	}
}

func (fb FactoryBuilder) RegisterBuilder(name string, builder IBuilder) {
	fb.lst[name] = builder
}

func (fb FactoryBuilder) GetBuilder(name string) IBuilder {
	return fb.lst[name]
}

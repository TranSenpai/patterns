package abstractfactory

type ComboName struct {
	lst map[string]ComboAbstractFactory
}

func (c ComboName) GetCombo(k string) ComboAbstractFactory {
	return c.lst[k]
}

func (c ComboName) SetCombo(k string, ca ComboAbstractFactory) {
	c.lst[k] = ca
}

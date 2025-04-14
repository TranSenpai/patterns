package prototype

type PrototypeRegistry struct {
	clone []*Prototype
}

func (pr *PrototypeRegistry) AddItem(p Prototype) {
	pr.clone = append(pr.clone, &p)
}

func (pr *PrototypeRegistry) GetItem(index int) Prototype {
	if index < 0 || index >= len(pr.clone) {
		return nil
	}
	return *pr.clone[index]
}

func (pr *PrototypeRegistry) GetAllItems() []*Prototype {
	return pr.clone
}

func (pr *PrototypeRegistry) RemoveItem(index int) {
	if index < 0 || index >= len(pr.clone) {
		return
	}
	pr.clone = append(pr.clone[:index], pr.clone[index+1:]...)
}

func (pr *PrototypeRegistry) Clear() {
	pr.clone = nil
}

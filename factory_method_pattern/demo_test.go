package factory_method_pattern

import "testing"

func TestProduct(t *testing.T) {
	var productFactoryA ProductFactory
	productFactoryA = &ProductAFactory{}
	p1 := productFactoryA.Create()
	p1.SetName("A 产品")
	name := p1.GetName()

	t.Log(name)
}

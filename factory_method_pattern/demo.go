package factory_method_pattern

// 定义一个抽象的产品
type Product interface {
	SetName(name string)
	GetName() string
}

// 实现具体的产品A
type ProductA struct {
	name string
}

func (this *ProductA)SetName(name string) {
	this.name = name
}

func (this *ProductA)GetName() string {
	return this.name
}

// 实现具体的产品B
type ProductB struct {
	name string
}

func (this *ProductB)SetName(name string) {
	this.name = name
}

func (this *ProductB)GetName() string {
	return this.name
}

// 实现一个抽象工厂
type ProductFactory interface {
	Create() Product
}

// 实现具体的工厂: 产品A的工厂
type ProductAFactory struct {

}

func (this *ProductAFactory) Create() Product {
	return &ProductA{}
}

// 实现具体的工厂: 产品B的工厂
type ProductBFactory struct {

}

func (this *ProductBFactory) Create() Product {
	return &ProductB{}
}

package pattern

import "fmt"

type Product1 struct{}

func (p Product1) info() {
	fmt.Println("i am product1")
}

type Product2 struct{}

func (p Product2) info() {
	fmt.Println("i am product2")
}

// тот самый объект, который мы создаем по частям
type Entity struct {
	P1 *Product1
	P2 *Product2
}

func (e *Entity) Info() {
	e.P1.info()
	e.P2.info()
}

type Builder interface {
	CreateEntity()
	BuildProduct1()
	BuildProduct2()
	GetProduct() *Entity
}

type ConcreteBuilder struct {
	ent *Entity
}

func (c *ConcreteBuilder) CreateEntity() {
	c.ent = new(Entity)
}

func (c *ConcreteBuilder) BuildProduct1() {
	fmt.Println("product1 created!")
	c.ent.P1 = new(Product1)
}

func (c *ConcreteBuilder) BuildProduct2() {
	fmt.Println("product2 created!")
	c.ent.P2 = new(Product2)
}

func (c *ConcreteBuilder) GetProduct() *Entity {
	return c.ent
}

type Director struct{}

func (d *Director) Construct(b Builder) *Entity {
	b.CreateEntity()
	b.BuildProduct1()
	b.BuildProduct2()
	return b.GetProduct()
}

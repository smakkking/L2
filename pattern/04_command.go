package pattern

import "fmt"

// объект системы, в котором реализовано поведение
type Receiver struct{}

func (r *Receiver) action1() {
	fmt.Println("action1")
}

func (r *Receiver) action2() {
	fmt.Println("action1")
}

type Command struct {
	r *Receiver
}

func newCommand(r *Receiver) *Command {
	return &Command{
		r: r,
	}
}

type Commander interface {
	Execute()
}

type Command1 struct {
	Command
}

func newCommand1(r *Receiver) *Command1 {
	return &Command1{
		Command: *newCommand(r),
	}
}
func (c *Command1) Execute() {
	c.r.action1()
}

type Command2 struct {
	Command
}

func newCommand2(r *Receiver) *Command2 {
	return &Command2{
		Command: *newCommand(r),
	}
}
func (c *Command2) Execute() {
	c.r.action2()
}

type Invoker struct {
	command_seq []Commander
}

func (i *Invoker) execute_command_queue() {
	for _, c := range i.command_seq {
		c.Execute()
	}
}

func (i *Invoker) insert_command(com Commander) {
	i.command_seq = append(i.command_seq, com)
}

func MainFunc4() {
	var t Invoker
	r := new(Receiver)

	t.insert_command(newCommand1(r))
	t.insert_command(newCommand2(r))
	t.insert_command(newCommand1(r))
	t.insert_command(newCommand1(r))

	t.execute_command_queue()
}

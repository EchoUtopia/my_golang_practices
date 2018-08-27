package main

type context struct {
	index int
	handlers [] func(*context)
}

func NewContext() *context{
	c := new(context)
	c.index = -1
	c.handlers = make([]func(*context), 0, 10)
	return c
}

func (c *context) Next(){
	c.index ++
	for l := len(c.handlers);c.index < l;c.index++{
		c.handlers[c.index](c)
	}
}

func (c *context)AddHandler(f func (*context)){
	c.handlers = append(c.handlers, f)
}

func middleware1(c *context){
	println("middleware1")
}

func middleware2(c *context){
	println("middleware2 pre")
	c.Next()
	println("middleware2 post")
}

func middleware3(c *context){
	c.Next()
	println("middleware3")
}

func handler(c *context){
	println("handler")
}

func main() {
	c := NewContext()
	c.AddHandler(middleware1)
	c.AddHandler(middleware2)
	c.AddHandler(middleware3)
	c.AddHandler(handler)
	c.Next()
}

//output:
//middleware1
//middleware2 pre
//handler
//middleware3
//middleware2 post

package promise

// Func is the handler of promise. The argument v
// is for passing all arguments.
type Func func(v interface{}) (ret interface{}, err error)

// Catcher catches error
type Catcher func(error)

// Promise executes its handlers and catches error
type Promise struct {
	v interface{}
	e error

	funcs []Func
	catch Catcher
}

// New creates a promise
func New() *Promise {
	return &Promise{}
}

// Init creates a promise with initiated value and error
func Init(v interface{}, e error) *Promise {
	return &Promise{
		v: v,
		e: e,
	}
}

// Then does the next step of callback fn to do
func (p *Promise) Then(fn Func) *Promise {
	if p.e != nil {
		return p
	}

	p.v, p.e = fn(p.v)
	return p
}

// Catch catches error for promise
func (p *Promise) Catch(c Catcher) {
	if p.e != nil {
		c(p.e)
	}
}

// Result gets the result of promise
func (p *Promise) Result() (interface{}, error) {
	return p.v, p.e
}

// Register registers a handler to promise
func (p *Promise) Register(fn Func) *Promise {
	p.funcs = append(p.funcs, fn)
	return p
}

// Catcher sets the error handler for promise
func (p *Promise) Catcher(c Catcher) *Promise {
	p.catch = c
	return p
}

// Done executes all registered handlers of promise
// and catch the error
func (p *Promise) Done() error {
	if len(p.funcs) == 0 {
		return nil
	}

	for i := range p.funcs {
		p.v, p.e = p.funcs[i](p.v)
		if p.e != nil {
			break
		}
	}

	if p.catch != nil {
		p.catch(p.e)
	}

	return p.e
}

// DoneAsync runs Done in a goroutine
func (p *Promise) DoneAsync() {
	go p.Done()
}

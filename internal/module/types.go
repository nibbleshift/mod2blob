package module

type Module struct {
	raw       []byte
	Functions []*Function
	Name      string
	Prefix    string
	Constants []Constant
	// map[method|function][]*Function
	Map map[string][]*Function
}

type Arg struct {
	Name string
	Type string
}

type Constant struct {
	Name  string
	Value string
}

type Function struct {
	Name        string
	Description string
	Args        []Arg
	Return      []Arg
}

package ccitt

func GetBytes(data []byte) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, b := range data {
			for i := 7; i >= 0; i-- {
				ch <- int((b >> i) & 1)
			}
		}
	}()
	return ch
}

type BitParserState []interface{}

type BitParser struct {
	state  BitParserState
	accept func(interface{}) BitParserState
	pos    int
}

func NewBitParser() *BitParser {
	return &BitParser{pos: 0}
}

func (bp *BitParser) Add(root BitParserState, v interface{}, bits string) {
	p := root
	var b interface{}
	for _, bit := range bits {
		if b != nil {
			if p[b.(int)] == nil {
				p[b.(int)] = make(BitParserState, 2)
			}
			p = p[b.(int)].(BitParserState)
		}
		if bit == '1' {
			b = 1
		} else {
			b = 0
		}
	}
	p[b.(int)] = v
}

func (bp *BitParser) FeedBytes(data []byte) {
	for bit := range GetBytes(data) {
		bp.parseBit(bit)
	}
}

func (bp *BitParser) parseBit(x int) {
	var v interface{}
	if x == 1 {
		v = bp.state[1]
	} else {
		v = bp.state[0]
	}
	bp.pos++
	if _, ok := v.(BitParserState); ok {
		bp.state = v.(BitParserState)
	} else {
		if bp.accept != nil {
			bp.state = bp.accept(v)
		}
	}
}

// func main() {
// 	// Sample usage of the BitParser class
// 	parser := NewBitParser()
// 	root := make(BitParserState, 2)
// 	parser.Add(root, "Value1", "01")
// 	parser.Add(root, "Value2", "10")
// 	parser.Add(root, "Value3", "11")
// 	parser.state = root
// 	data := []byte{170, 85}
// 	parser.FeedBytes(data)
// 	fmt.Println(parser.state)
// }

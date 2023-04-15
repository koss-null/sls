package printer

import (
	"fmt"
	"sync"
)

var emptyString = ""

type dot struct{ x, y int }

type Printer struct {
	mx             sync.Mutex
	buffer         []string
	cursorPosition dot
}

func NewPrinter() *Printer {
	return &Printer{
		buffer:         make([]string, 0, 20),
		cursorPosition: dot{0, 0},
	}
}

// TopCursor moves cursor to 0, 0 position
func (p *Printer) TopCursor() {
	// this one does not look consistent enough
	p.MoveUp(p.cursorPosition.y)
	p.MoveLeft(p.cursorPosition.x)

	p.mx.Lock()
	p.cursorPosition = dot{0, 0}
	p.mx.Unlock()
}

// Cursor returns cursor position
func (p *Printer) Cursor() (x, y int) {
	p.mx.Lock()
	defer p.mx.Unlock()

	return p.cursorPosition.x, p.cursorPosition.y
}

// MoveUp moves cursor up, if it is not possible returns false.
func (p *Printer) MoveUp(i int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	if p.cursorPosition.y-i < 0 {
		return false
	}

	fmt.Printf("\033[<%d>A", i)
	p.cursorPosition.y -= i
	return true
}

// MoveDown moves cursor down, if it is not possible returns false.
func (p *Printer) MoveDown(i int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	if p.cursorPosition.y+i >= len(p.buffer) {
		return false
	}

	fmt.Printf("\033[<%d>B", i)
	p.cursorPosition.y -= i
	return true
}

// MoveLeft moves cursor left, if it is not possible returns false.
func (p *Printer) MoveLeft(i int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	if p.cursorPosition.x-i < 0 {
		return false
	}

	fmt.Printf("\033[<%d>D", i)
	p.cursorPosition.y -= i
	return true
}

// MoveRight moves cursor right, if it is not possible returns false.
func (p *Printer) MoveRight(i int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	if p.cursorPosition.x+i >= len(p.buffer[p.cursorPosition.y]) {
		return false
	}

	fmt.Printf("\033[<%d>C", i)
	p.cursorPosition.y -= i
	return true
}

// clearLine clears line under the cursor.
func (p *Printer) clearLine() {
	fmt.Print("\033[2K")
}

// PutLine puts line to the cursor position in the buffer.
func (p *Printer) PutLine(line string) {
	p.mx.Lock()
	defer p.mx.Unlock()

	// this append does not perform well but who cares
	p.buffer = append(p.buffer[:p.cursorPosition.y], append([]string{line}, p.buffer[p.cursorPosition.y:]...)...)
}

func (p *Printer) RemoveLine(n int) bool {
	p.mx.Lock()
	defer p.mx.Unlock()

	if n >= len(p.buffer) {
		return false
	}

	p.buffer[n] = emptyString
	if p.cursorPosition.y == n {
		p.cursorPosition.x = 0
	}
	return true
}

// PrintBuffer clear all lines and prints current buffer, cursor is placed to 0, 0.
func (p *Printer) PrintBuffer() {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.MoveUp(p.cursorPosition.y)
	p.MoveLeft(p.cursorPosition.x)
	for _, s := range p.buffer {
		p.clearLine()
		fmt.Println(s)
	}

	p.MoveUp(len(p.buffer))
	p.cursorPosition = dot{0, 0}
}

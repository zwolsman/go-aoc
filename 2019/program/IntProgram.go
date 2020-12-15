package program

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type intcode []int
type Program struct {
	code              intcode
	Memory            intcode
	ptr, relativeBase int
	In, Out           chan int
	logger            *log.Logger
}

func Read(path string) Program {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var instructions intcode
	for _, code := range strings.Split(string(data), ",") {
		opCode, err := strconv.Atoi(code)
		if err != nil {
			log.Fatal("error reading opcode", err, code)
		}
		instructions = append(instructions, opCode)
	}

	return newProgram(instructions)
}

func (p *Program) SetPrefix(prefix string) {
	p.logger = log.New(os.Stderr, prefix, log.LstdFlags)
}

func newProgram(instructions intcode) Program {
	memory := make(intcode, len(instructions)*10)
	copy(memory, instructions)
	return Program{
		code:   instructions,
		Memory: memory,
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func Fix(noun, verb int, p Program) Program {
	p.Memory[1] = noun
	p.Memory[2] = verb
	return p
}

type Mode int

const (
	POSITION_MODE Mode = iota
	IMMEDIATE_MODE
	RELATIVE_MODE
)

type Argument struct {
	mode          Mode
	position, raw int
}

func (p *Program) read(argument Argument) int {
	switch argument.mode {
	case POSITION_MODE:
		return p.Memory[argument.raw]
	case IMMEDIATE_MODE:
		return argument.raw
	case RELATIVE_MODE:
		return p.Memory[p.relativeBase+argument.raw]
	}
	p.logger.Fatal("unknown argument type")
	return -1
}

func (p *Program) arg(argument Argument) int {
	switch argument.mode {
	case POSITION_MODE, IMMEDIATE_MODE:
		return argument.raw
	case RELATIVE_MODE:
		return p.relativeBase + argument.raw
	}
	log.Fatal("unknown argument type")
	return -1
}

func (a Argument) String() string {
	return fmt.Sprintf("Argument(%v, %d)", a.mode, a.raw)
}

func (m Mode) String() string {
	return [...]string{"POSITION_MODE", "IMMEDIATE_MODE", "RELATIVE_MODE"}[m]
}

func (p *Program) readArgs(num int) []Argument {
	mask := p.Memory[p.ptr-1]
	modes := readModes(mask, num)
	args := make([]Argument, num)
	for i := 0; i < num; i++ {
		mode := modes[len(modes)-i-1]
		args[i] = Argument{mode, p.ptr + i, p.Memory[p.ptr+i]}
	}
	p.ptr += num
	return args
}

func (p *Program) next() int {
	n := p.Memory[p.ptr]
	p.ptr += 1
	return n
}
func (p Program) Run() intcode {
	for p.Memory[p.ptr] != 99 {
		mask := p.next()
		opcode := readOpcode(mask)

		switch opcode {
		case 1:
			args := p.readArgs(3)
			x, y, target := p.read(args[0]), p.read(args[1]), p.arg(args[2])
			p.Memory[target] = x + y
			break
		case 2:
			args := p.readArgs(3)
			x, y, target := p.read(args[0]), p.read(args[1]), p.arg(args[2])
			p.Memory[target] = x * y
			break
		case 3: //INPUT
			target := p.arg(p.readArgs(1)[0])
			var input int
			if p.In != nil {
				//p.logger.Printf("receiving from channel")
				input = <-p.In
				//p.logger.Printf("received, %d", input)
			} else {
				//print("input: ")
				_, err := fmt.Scan(&input)
				if err != nil {
					log.Fatal(err)
				}
			}
			p.Memory[target] = input
		case 4: //OUTPUT
			arg := p.read(p.readArgs(1)[0])
			if p.Out != nil {
				//p.logger.Printf("sending to channel %v\n", arg)
				p.Out <- arg
			} else {
				fmt.Printf("%v\n", arg)
			}
		case 5:
			args := p.readArgs(2)
			x, jmp := p.read(args[0]), p.read(args[1])
			if x != 0 {
				p.ptr = jmp
			}
		case 6:
			args := p.readArgs(2)
			x, jmp := p.read(args[0]), p.read(args[1])
			if x == 0 {
				p.ptr = jmp
			}
		case 7:
			args := p.readArgs(3)
			x, y, target := p.read(args[0]), p.read(args[1]), p.arg(args[2])
			if x < y {
				p.Memory[target] = 1
			} else {
				p.Memory[target] = 0
			}
		case 8:
			args := p.readArgs(3)
			x, y, target := p.read(args[0]), p.read(args[1]), p.arg(args[2])
			if x == y {
				p.Memory[target] = 1
			} else {
				p.Memory[target] = 0
			}
		case 9:
			args := p.readArgs(1)
			p.relativeBase += p.read(args[0])
		default:
			p.logger.Fatal("unknown opcode at ", p.ptr, "opcode:", opcode)
		}
	}
	p.logger.Println("done")
	return p.Memory
}

func Run(program Program) intcode {
	return program.Run()
}

func readOpcode(mask int) int {
	if mask < 10 {
		return mask
	}
	str := fmt.Sprint(mask)
	opcode, err := strconv.Atoi(str[len(str)-2:])
	if err != nil {
		log.Fatal(err)
	}
	return opcode
}

type Modes []Mode

func readModes(mask int, num int) Modes {
	correct := make(Modes, num+2)

	str := fmt.Sprintf("%d", mask)
	for i, c := range str {
		correct[len(correct)-len(str)+i] = Mode(c)
	}
	correct = correct[0:num]
	for i, v := range correct {
		if v != 0 {
			correct[i] = v - '0'
		}
	}
	return correct
}

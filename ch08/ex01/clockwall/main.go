package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode/utf8"
)

type clock struct {
	index int
	name  string
	host  string
	port  int
}

type terminalSize struct {
	width, height int
	clockWidth    int
}

type clockTimeDto struct {
	c    *clock
	time string
}

func main() {
	clocks := parseClocks()
	t := getTerminalSize(clocks)
	printFrame(t, clocks)
	ch := make(chan *clockTimeDto)
	for _, c := range clocks {
		go handleConn(c, ch)
	}
	for dto := range ch {
		printTime(t, dto.c, dto.time)
	}
}

func parseClocks() []*clock {
	printErrorAndExit := func() {
		fmt.Fprintln(os.Stderr, "Specify <CLOCK_NAME>=<HOST>:<PORT>")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		printErrorAndExit()
	}
	clocks := []*clock{}
	for i, arg := range os.Args[1:] {
		c := &clock{}
		c.index = i
		tmp := strings.Split(arg, "=")
		if len(tmp) != 2 {
			printErrorAndExit()
		}
		c.name = tmp[0]
		tmp = strings.Split(tmp[1], ":")
		if len(tmp) != 2 {
			printErrorAndExit()
		}
		c.host = tmp[0]
		port, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			printErrorAndExit()
		}
		c.port = port
		clocks = append(clocks, c)
	}
	return clocks
}

func getTerminalSize(clocks []*clock) *terminalSize {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	vals := strings.Split(string(out), " ")
	t := &terminalSize{}
	t.height, _ = strconv.Atoi(vals[0])
	t.width, _ = strconv.Atoi(strings.TrimSpace(vals[1]))
	t.clockWidth = t.width / len(clocks)
	return t
}

func moveCursor(x, y int) {
	fmt.Print(fmt.Sprintf("%s[%d;%dH", "\x1b", y, x))
}

func printFrame(t *terminalSize, clocks []*clock) {
	numClocks := len(clocks)
	for h := 0; h < t.height; h++ {
		switch h {
		case 0:
			for i := 0; i < numClocks; i++ {
				fmt.Print("┌")
				fmt.Print(strings.Repeat("─", t.clockWidth-2))
				fmt.Print("┐")
			}
			fmt.Println()
		default:
			for i := 0; i < numClocks; i++ {
				fmt.Print("│")
				fmt.Print(strings.Repeat(" ", t.clockWidth-2))
				fmt.Print("│")
			}
			fmt.Println()
		case t.height - 1:
			for i := 0; i < numClocks; i++ {
				fmt.Print("└")
				fmt.Print(strings.Repeat("─", t.clockWidth-2))
				fmt.Print("┘")
			}
			fmt.Print()
		}
	}

	for _, c := range clocks {
		moveCursor(t.clockWidth*c.index+(t.clockWidth-utf8.RuneCountInString(c.name))/2+1, t.height/2-1)
		fmt.Print(c.name)
	}
	moveCursor(t.width, t.height)
}

func printTime(t *terminalSize, c *clock, time string) {
	moveCursor(t.clockWidth*c.index+(t.clockWidth-utf8.RuneCountInString(time))/2+1, t.height/2+1)
	fmt.Print(time)
	moveCursor(t.width, t.height)
}

func handleConn(c *clock, ch chan<- *clockTimeDto) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		time := scanner.Text()
		dto := &clockTimeDto{c, time}
		ch <- dto
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/h-matsuo/gopl/ch02/ex01/tempconv"
	"github.com/h-matsuo/gopl/ch02/ex02/lengthconv"
	"github.com/h-matsuo/gopl/ch02/ex02/weightconv"
)

// Extracted for unit tests
var (
	in     io.Reader = os.Stdin
	osArgs []string  = os.Args
)

func main() {
	quantity, unit := getInput()
	fmt.Println(convert(quantity, unit))
}

func getUsage() string {
	return fmt.Sprintf("Usage: %s <quantity> <unit>\n", osArgs[0])
}

func getInput() (quantity float64, unit string) {
	if len(osArgs) == 2 {
		// Invalid argument (only quantity is specified)
		fmt.Fprintln(os.Stderr, getUsage())
		os.Exit(1)
	}
	quantityStr := ""
	if len(osArgs) > 1 {
		quantityStr = osArgs[1]
		unit = osArgs[2]
	} else {
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		quantityStr = scanner.Text()
		scanner.Scan()
		unit = scanner.Text()
	}
	quantity, err := strconv.ParseFloat(quantityStr, 64)
	if err != nil {
		log.Fatal(err)
	}
	return quantity, unit
}

func convert(quantity float64, unit string) string {
	switch unit {
	case "°C", "tempC":
		c := tempconv.Celsius(quantity)
		f := tempconv.CToF(c)
		return fmt.Sprintf("%f °F", f)
	case "°F", "tempF":
		f := tempconv.Fahrenheit(quantity)
		c := tempconv.FToC(f)
		return fmt.Sprintf("%f °C", c)
	case "ft", "feet":
		f := lengthconv.Foot(quantity)
		m := lengthconv.FToM(f)
		return fmt.Sprintf("%f m", m)
	case "m", "meters":
		m := lengthconv.Meter(quantity)
		f := lengthconv.MToF(m)
		return fmt.Sprintf("%f ft", f)
	case "lb", "pounds":
		p := weightconv.Pound(quantity)
		k := weightconv.PToK(p)
		return fmt.Sprintf("%f kg", k)
	case "kg", "kilograms":
		k := weightconv.Kilogram(quantity)
		p := weightconv.KToP(k)
		return fmt.Sprintf("%f lb", p)
	default:
		log.Fatal(`unsupported unit "`, unit, `" specified`)
		return ""
	}
}

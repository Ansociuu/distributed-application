package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net/rpc"
    "os"
    "strconv"
    "strings"
)

type Operands struct {
    A, B float64
}

type Unary struct {
    A float64
}

func callBinary(client *rpc.Client, method string, a, b float64) (float64, error) {
    var reply float64
    args := Operands{A: a, B: b}
    err := client.Call("Calculator."+method, &args, &reply)
    return reply, err
}

func callUnary(client *rpc.Client, method string, a float64) (float64, error) {
    var reply float64
    arg := Unary{A: a}
    err := client.Call("Calculator."+method, &arg, &reply)
    return reply, err
}

func printHelp() {
    fmt.Println("Commands:")
    fmt.Println("  add a b    - addition")
    fmt.Println("  sub a b    - subtraction")
    fmt.Println("  mul a b    - multiplication")
    fmt.Println("  div a b    - division")
    fmt.Println("  pow a b    - power a^b")
    fmt.Println("  sqrt a     - square root")
    fmt.Println("  sin a      - sine (radians)")
    fmt.Println("  cos a      - cosine (radians)")
    fmt.Println("  tan a      - tangent (radians)")
    fmt.Println("  log10 a    - base-10 log")
    fmt.Println("  ln a       - natural log")
    fmt.Println("  help       - show this")
    fmt.Println("  exit       - quit")
}

func main() {
    addr := flag.String("addr", "127.0.0.1:1234", "RPC server address")
    flag.Parse()

    client, err := rpc.Dial("tcp", *addr)
    if err != nil {
        log.Fatalf("dialing: %v", err)
    }
    defer client.Close()

    fmt.Printf("Connected to %s\nType 'help' for commands.\n", *addr)

    binaryMethods := map[string]string{
        "add": "Add", "sub": "Sub", "mul": "Mul", "div": "Div", "pow": "Pow",
    }
    unaryMethods := map[string]string{
        "sqrt": "Sqrt", "sin": "Sin", "cos": "Cos", "tan": "Tan", "log10": "Log10", "ln": "Ln",
    }

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        parts := strings.Fields(line)
        cmd := strings.ToLower(parts[0])

        if cmd == "exit" || cmd == "quit" {
            break
        }
        if cmd == "help" {
            printHelp()
            continue
        }

        if method, ok := binaryMethods[cmd]; ok {
            if len(parts) < 3 {
                fmt.Println("Need two numeric arguments")
                continue
            }
            a, err1 := strconv.ParseFloat(parts[1], 64)
            b, err2 := strconv.ParseFloat(parts[2], 64)
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            res, err := callBinary(client, method, a, b)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }
            fmt.Println(res)
            continue
        }
        if method, ok := unaryMethods[cmd]; ok {
            if len(parts) < 2 {
                fmt.Println("Need one numeric argument")
                continue
            }
            a, err1 := strconv.ParseFloat(parts[1], 64)
            if err1 != nil {
                fmt.Println("Invalid number")
                continue
            }
            res, err := callUnary(client, method, a)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }
            fmt.Println(res)
            continue
        }
        fmt.Println("Unknown command. Type 'help'.")
    }
    if err := scanner.Err(); err != nil {
        log.Printf("input error: %v", err)
    }
    fmt.Println("Bye.")
}

package main

import (
    "errors"
    "fmt"
    "log"
    "math"
    "net"
    "net/rpc"
)

type Operands struct {
    A, B float64
}

type Unary struct {
    A float64
}

type Calculator struct{}

func (c *Calculator) Add(args *Operands, reply *float64) error {
    *reply = args.A + args.B
    return nil
}

func (c *Calculator) Sub(args *Operands, reply *float64) error {
    *reply = args.A - args.B
    return nil
}

func (c *Calculator) Mul(args *Operands, reply *float64) error {
    *reply = args.A * args.B
    return nil
}

func (c *Calculator) Div(args *Operands, reply *float64) error {
    if args.B == 0 {
        return errors.New("division by zero")
    }
    *reply = args.A / args.B
    return nil
}

func (c *Calculator) Pow(args *Operands, reply *float64) error {
    *reply = math.Pow(args.A, args.B)
    return nil
}

func (c *Calculator) Sqrt(arg *Unary, reply *float64) error {
    if arg.A < 0 {
        return errors.New("sqrt of negative number")
    }
    *reply = math.Sqrt(arg.A)
    return nil
}

func (c *Calculator) Sin(arg *Unary, reply *float64) error {
    *reply = math.Sin(arg.A)
    return nil
}

func (c *Calculator) Cos(arg *Unary, reply *float64) error {
    *reply = math.Cos(arg.A)
    return nil
}

func (c *Calculator) Tan(arg *Unary, reply *float64) error {
    *reply = math.Tan(arg.A)
    return nil
}

func (c *Calculator) Log10(arg *Unary, reply *float64) error {
    if arg.A <= 0 {
        return errors.New("log10 undefined for <= 0")
    }
    *reply = math.Log10(arg.A)
    return nil
}

func (c *Calculator) Ln(arg *Unary, reply *float64) error {
    if arg.A <= 0 {
        return errors.New("ln undefined for <= 0")
    }
    *reply = math.Log(arg.A)
    return nil
}

func main() {
    calc := new(Calculator)
    rpc.Register(calc)
    ln, err := net.Listen("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatalf("listen error: %v", err)
    }
    fmt.Println("Calculator RPC server listening on 127.0.0.1:1234")
    rpc.Accept(ln)
}

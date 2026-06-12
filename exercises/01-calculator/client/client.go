package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Operands struct {
	A, B float64
}
type Unary struct {
	A float64
}

func callExample(client *rpc.Client) {
	var reply float64
	args := Operands{A: 2, B: 3}
	if err := client.Call("Calculator.Add", &args, &reply); err != nil {
		log.Fatalf("Add error: %v", err)
	}
	fmt.Printf("2 + 3 = %v\n", reply)

	args = Operands{A: 7, B: 2}
	if err := client.Call("Calculator.Div", &args, &reply); err != nil {
		log.Fatalf("Div error: %v", err)
	}
	fmt.Printf("7 / 2 = %v\n", reply)

	args = Operands{A: 2, B: 8}
	if err := client.Call("Calculator.Pow", &args, &reply); err != nil {
		log.Fatalf("Pow error: %v", err)
	}
	fmt.Printf("2^8 = %v\n", reply)

	u := Unary{A: 16}
	if err := client.Call("Calculator.Sqrt", &u, &reply); err != nil {
		log.Fatalf("Sqrt error: %v", err)
	}
	fmt.Printf("Sqrt(16) = %v\n", reply)
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("dialing: %v", err)
	}
	defer client.Close()
	callExample(client)
}

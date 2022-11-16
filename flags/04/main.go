package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (addr *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", addr.Host, addr.Port)
}

// Set связывает переменную типа со значением флага
// и устанавливает правила парсинга для пользовательского типа.
func (addr *NetAddress) Set(flagValue string) error {
	value := strings.Split(flagValue, ":")

	addr.Host = value[0]

	port, err := strconv.Atoi(value[1])
	if err != nil {
		return err
	}

	addr.Port = port
	return nil
}

//--addr=example.com:60

// допишите код реализации методов интерфейса

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)

	flag := flag.Lookup("addr")
	fmt.Println(flag)
	
}

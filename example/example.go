package main

import (
	"charles/redisCliPool"

	"flag"
	"fmt"
	"os"

	"github.com/dlintw/goconf"
)

func main() {
	var err error
	conf_file := flag.String("config", "./config.ini", "set redis config file.")
	flag.Parse()

	l_conf, err := goconf.ReadConfigFile(*conf_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", conf_file, err)
		os.Exit(1)
	}
	redisCliPool.InitRedisPool(l_conf)
	defer redisCliPool.Clipool.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	conn := redisCliPool.Clipool.Get()
	defer conn.Close()
	//	v, err := c.Do("SET", "name", "red")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(v)
	var v string
	v, err = redisCliPool.String(conn.Do("GET", "pool"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	select {}
	return
}

package main

import (
	"flag"
	"fmt"
	"github.com/fabtestorg/gohfc"
	"github.com/spf13/viper"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if err := gohfc.InitSDK("./client.yaml"); err != nil {
		fmt.Println(err)
		return
	}
	//args := []string{"invoke", "a", "b", "20"}
	peers := []string{"peer0"}
	if args[0] == "invoke" {
		result, err := gohfc.GetHandler().Invoke(args, peers, "orderer0")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	} else if args[0] == "query" {
		result, err := gohfc.GetHandler().Query(args, peers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(result[0].Response.Response.GetPayload()))
	} else if args[0] == "listen" {
		ch, err := gohfc.GetHandler().ListenEvent("peer0", viper.GetString("other.localMspId"))
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			}
		}
	} else {
		result, err := gohfc.GetHandler().QueryByQscc(args, peers)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(result[0].Response.Response.GetPayload()))
	}
}

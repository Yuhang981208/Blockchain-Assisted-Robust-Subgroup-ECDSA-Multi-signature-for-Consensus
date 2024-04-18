package main

import (
	"flag"
	"log"

	"os"
	"os/signal"
	MulECDSA "signer/pkg/signer"

	"github.com/spf13/viper"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	commonConfig := "../configs/common.json"
	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	var config MulECDSA.Config

	// 公共配置
	viper.SetConfigFile(*&commonConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Read config: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Println("Unmarshal common config into struct, ", err)
	}

	// 节点配置
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Read config: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Println("Unmarshal config into struct, ", err)
	}
	log.Println("Loaded config file ", *configFile)

	node, err := MulECDSA.NewOracleNode(config) // 根据config初始化node
	if err != nil {
		log.Println("New oracle node: ", err)
	}

	go func() {
		if err := node.Run(); err != nil {
			log.Println("Run node:", err)
		} // 运行node
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	node.Stop()
}

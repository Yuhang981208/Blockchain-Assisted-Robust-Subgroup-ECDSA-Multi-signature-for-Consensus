package signer

type Config struct {
	BindAddress string
	Contracts   ContractsConfig
	Ethereum    EthereumConfig
	Ta          string
	Kafka       KafkaConfig
}

type ContractsConfig struct {
	RegistryContractAddress string
	SingerContractAddress   string
}

type EthereumConfig struct {
	Address    string
	PrivateKey string
	ChainID    int64
}
type KafkaConfig struct {
	IpAddress string
	Topic     string
	Partition int64
}

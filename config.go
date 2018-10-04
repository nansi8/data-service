package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	DataNodes     = "DATA_NODES"
	ChecksumNodes = "CHECKSUM_NODES"
)

type DataServiceConfig struct {
	dataNodes     int
	checksumNodes int
	init          bool
}

var config = &DataServiceConfig{0, 0, false}

func Config() *DataServiceConfig {
	if config.init {
		return config
	}
	dataNodes, err := strconv.Atoi(os.Getenv(DataNodes))
	if err != nil {
		fmt.Println("DataNodes variable is not set")
	}
	config.dataNodes = dataNodes

	checksumNodes, err := strconv.Atoi(os.Getenv(ChecksumNodes))
	if err != nil {
		fmt.Println("ChecksumNodes variable is not set")
	}
	config.checksumNodes = checksumNodes
	config.init = true
	return config
}

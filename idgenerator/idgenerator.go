package idgenerator

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

var node *snowflake.Node

func Init(nodeId int64) {
	var err error
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		log.Fatalf("snowflake init error: %v", err)
	}
}

func Generate() int64 {
	return node.Generate().Int64()
}

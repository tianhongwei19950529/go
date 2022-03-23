package idgen

import "github.com/bwmarrin/snowflake"

var idNode *snowflake.Node

func InitNode(node int64) {
	var err error
	if idNode == nil {
		idNode, err = snowflake.NewNode(node)
	}
	if err != nil {
		panic(err)
	}
}

func getNode() *snowflake.Node {
	return idNode
}

func GenId() string {
	return getNode().Generate().String()
}

func GenBatchId(nums int) []string {
	var res []string
	for i := 0; i < nums; i++ {
		res = append(res, getNode().Generate().String())
	}
	return res
}

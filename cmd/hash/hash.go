package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"hash/fnv"
	"math/big"
)

const contentNumTables = 4

func main() {
	// cms_content.t_content_details_0
	// cms_content.t_content_details_1
	// cms_content.t_content_details_2
	// cms_content.t_content_details_3

	//u := uuid.New().String()
	// 5104856713695633295
	getContentDetailsTable("c02fa56c-895a-4eb6-8573-c63a21aeca1d")
}

func getContentDetailsTable(contentID string) string {
	tableIndex := getContentTableIndex(contentID)
	table := fmt.Sprintf("cms_content.t_content_details_%d", tableIndex)
	log.Infof("content_id = %s, table = %s", contentID, table)

	return table
}

func getContentTableIndex(uuid string) int {
	hash := fnv.New64()
	_, _ = hash.Write([]byte(uuid))
	hashValue := hash.Sum64()
	fmt.Println("hashValue = ", hashValue)

	bigNum := big.NewInt(int64(hashValue))
	mod := big.NewInt(contentNumTables)
	tableIndex := bigNum.Mod(bigNum, mod).Int64()

	return int(tableIndex)
}

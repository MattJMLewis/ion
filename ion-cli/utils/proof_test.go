// Copyright (c) 2018 Clearmatics Technologies Ltd

package utils_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/clearmatics/ion/ion-cli/ionflow"
	"github.com/clearmatics/ion/ion-cli/utils"
	"github.com/stretchr/testify/assert"
)

var TEST_PATH = "13"
var TEST_TX_VALUE = "f86707843b9aca008257c39461621bcf02914668f8404c1f860e92fc1893f74c8084457094cc1ba07e2ebe15f4ece2fd8ffc9a49d7e9e4e71a30534023ca6b24ab4000567709ad53a013a61e910eb7145aa93e865664c54846f26e09a74bd577eaf66b5dd00d334288"
var TEST_TX_NODES = "f90235f871a0804f9c841a6a1d3361d79980581c84e5b4d3e4c9bf33951346775542d0ee0728a0edadb5e660118ea4323654191131b62c81fc00203a15a21c925f9f50d0e4b3e4808080808080a03eda2d64b94c5ed45026a29c75c99677d44c561ea5efea30c1db6299871d5c2e8080808080808080f90151a0bc285699e68d2fe18e7af2cdf7e7e6456e91a3fd31e3c9935bc5bef92e94bf4ba06eb963b2c3a3b6c07a7221aa6f6f86f7cb8ddb45ab1ff1a9dc781f34da1f081fa0deea5b5566e7a5634d91c5fb56e25f4370e3531e2fd71ee17ed6c4ad0be2ced3a0b4e9d14555f162e811cfbcbff9b98a271a197b75271565f693912c2ff75e2131a03b0bc2d764fbefd76848ee2da7b211eb230ede08d8c54e6a868be9f5e42122c1a0b6dd488ad4fb82b0a98dff81ac6766d1dec26b29dc06174de1d315b0ab0bdf0ca066c20ff06dc33777f53eec32b0b9a8d99872bec24bb3998bb520ae6897c21d7ea02db2a399f611ba7993efb4768938a6f61b4add8959ce4c89f201f41e882ff375a02e31051a9f938b9b342b8070db3dd829f62da8d0c83a6dff91a4e3b4cb2adb9ea090e75708e7dbf856b75ed126a960085419fcde0e6a0129a92dffc0cb83ac089680808080808080f86c20b869f86707843b9aca008257c39461621bcf02914668f8404c1f860e92fc1893f74c8084457094cc1ba07e2ebe15f4ece2fd8ffc9a49d7e9e4e71a30534023ca6b24ab4000567709ad53a013a61e910eb7145aa93e865664c54846f26e09a74bd577eaf66b5dd00d334288"
var TEST_RECEIPT_VALUE = "f901640183252867b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000010000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000f85af8589461621bcf02914668f8404c1f860e92fc1893f74ce1a027a9902e06885f7c187501d61990eae923b37634a8d6dda55a04dc7078395340a0000000000000000000000000279884e133f9346f2fad9cc158222068221b613e"
var TEST_RECEIPT_NODES = "f90335f871a012d378fe6800bc18f22e715a31971ef7e73ac5d1d85384f4b66ac32036ae43dea004d6e2678656a957ac776dbef512a04d266c1af3e2c5587fd233261a3d423213808080808080a05fac317a4d6d78181319fbc7e2cae4a9260f1a6afb5c6fea066e2308eed416818080808080808080f90151a03da235c6dd0fbdaf208c60cbdca0d609dee2ba107495aa7adaa658362616c8aaa09ebf378a9064aa4da0512c55c790a5e007ac79d2713e4533771cd2c95be47a4da0c06fed36ffe1f2ec164ba88f73b353960448d2decbb65355c5298a33555de742a0e057afe423ee17e5499c570a56880b0f5b5c1884b90ff9b9b5baa827f72fc816a093e06093cd2fdb67e0f87cfcc35ded2f445cc1309a0ff178e59f932aeadb6d73a0193e4e939fbc5d34a570bea3fff7c6d54adcb1c3ab7ef07510e7bd5fcef2d4b3a0a17a0c71c0118092367220f65b67f2ba2eb9068ff5270baeabe8184a01a37f14a03479a38e63123d497588ad5c31d781276ec8c11352dd3895c8add34f9a2b786ba042254728bb9ab94b58adeb75d2238da6f30382969c00c65e55d4cc4aa474c0a6a03c088484aa1c73b8fb291354f80e9557ab75a01c65d046c2471d19bd7f2543d880808080808080f9016b20b90167f901640183252867b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000010000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000f85af8589461621bcf02914668f8404c1f860e92fc1893f74ce1a027a9902e06885f7c187501d61990eae923b37634a8d6dda55a04dc7078395340a0000000000000000000000000279884e133f9346f2fad9cc158222068221b613e"

func Test_GenerateProof(t *testing.T) {
	ctx := context.Background()

	// CONTRACT_ADDR, _ := utils.StringToBytes32("61621bcf02914668f8404c1f860e92fc1893f74c")
	TXHASH, _ := utils.StringToBytes32("afc3ab60059ed38e71c7f6bea036822abe16b2c02fcf770a4f4b5fffcbfe6e7e")

	// Connect to the RPC Client
	client := ionflow.ClientRPC("https://rinkeby.infura.io")
	defer client.Close()

	PATH, TX_VALUE, TX_NODES, RECEIPT_VALUE, RECEIPT_NODES := utils.GenerateProof(ctx, client, TXHASH)
	assert.Equal(t, TEST_PATH, hex.EncodeToString(PATH))
	assert.Equal(t, TEST_TX_VALUE, hex.EncodeToString(TX_VALUE))
	assert.Equal(t, TEST_TX_NODES, hex.EncodeToString(TX_NODES))
	assert.Equal(t, TEST_RECEIPT_VALUE, hex.EncodeToString(RECEIPT_VALUE))
	assert.Equal(t, TEST_RECEIPT_NODES, hex.EncodeToString(RECEIPT_NODES))
}

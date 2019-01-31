package serialize

import (
	"strconv"
)

//returns a string with combined ip, public key, and the current blocknumber
func Serialize(ip , publicKey string, blockNumber int ) (res string){

	return ip+publicKey+strconv.Itoa(blockNumber)
}

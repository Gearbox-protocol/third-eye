package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFacade"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"os"
	"strings"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ToJson(obj interface{}) string {
	str, err := json.Marshal(obj)
	log.CheckFatal(err)
	return string(str)
}

func GetCreditManagerEventIds() []string {
	var ids []string
	if a, err := abi.JSON(strings.NewReader(creditManager.CreditManagerABI)); err == nil {
		for _, event := range a.Events {
			// log.Info(event.RawName, event.ID.Hex())
			// if event.RawName != "ExecuteOrder" {
			ids = append(ids, event.ID.Hex())
			// }
		}
	}
	return ids
}
func Method() {
	if a, err := abi.JSON(strings.NewReader(creditFacade.CreditFacadeABI)); err == nil {
		for name, method := range a.Methods {
			log.Info(name, method.RawName, method.Name, method.Sig, hex.EncodeToString(method.ID))
		}
	}
}

// for testing and json file read methods
func RandomAddr() string {
	return common.HexToAddress(random(20)).Hex()
}
func RandomHash() string {
	return "0x" + random(32)
}
func random(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func ChecksumAddr(addr string) string {
	return common.HexToAddress(addr).Hex()
}

func ReadJsonAndSet(fileName string) []map[string]interface{} {
	data := []map[string]interface{}{}
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
	return data
}
func ReadJsonAndSetInterface(fileName string, data interface{}) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
}

func ReadJson(fileName string) map[string]interface{} {
	data := map[string]interface{}{}
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func SetJson(byteValue []byte, data interface{}) {
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
}

func ConvertToListOfString(list interface{}) (accountAddrs []string) {
	switch list.(type) {
	case []interface{}:
		accountList, ok := list.([]interface{})
		if !ok {
			panic("parsing accounts list for token transfer failed")
		}
		for _, account := range accountList {
			accountAddr, ok := account.(string)
			if !ok {
				log.Fatalf("parsing single account for token transfer failed %v", account)
			}
			accountAddrs = append(accountAddrs, accountAddr)
		}
	case []string:
		accountList, ok := list.([]string)
		if !ok {
			panic("parsing accounts list for token transfer failed")
		}
		accountAddrs = accountList
	}
	return
}

func ReadFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

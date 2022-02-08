package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"os"
	"reflect"
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

func RandomAddr() string {
	return "0x" + random(20)
}
func RandomHash() string {
	return "0x" + random(32)
}
func random(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
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
	log.Info(reflect.TypeOf(data))
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

package service

import (
	"github.com/Kashyap23/query-store/src/utils"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Name of the dir to store the queries
var queryDir = "query-store"

type QueryStore struct {
	Dir string
}

var queryStore QueryStore

// Initialize query store
func InitQueryStore(dirPath string) QueryStore {
	dir := QueryStore{}
	err := utils.CreateDir(dirPath)
	if err != nil {
		log.Panic(err)
	}
	dir.Dir = dirPath
	return dir
}

// Store query
func (q *QueryStore) Set(query string) string {
	// Validate if dir exists
	if ok := utils.CheckIfPathExists(q.Dir); !ok {
		utils.CreateDir(queryDir)
	}
	// Store to file
	hash, err := utils.GenerateRandomHash()
	if err != nil {
		log.Println("Error generating random hash - ", err)
		panic(err)
	}
	filePath := q.Dir + "/" + hash

	err = ioutil.WriteFile(filePath, []byte(query), 0755)
	if err != nil {
		log.Println("Error writing to file - ", err)
		return ""
	}

	log.Println("Wrote query to file successfully - ", hash)
	return hash
}

// Retreive query
func (q *QueryStore) Get(key string) string {
	// Check if file exists
	filePath := q.Dir + "/" + key
	if ok := utils.CheckIfPathExists(filePath); !ok {
		log.Println("Path - ", filePath, " doesn't exist")
		return ""
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading file - ", filePath, " error - ", err)
		return ""
	}

	return string(data)
}

func EncodeByteToString(body []byte) string {
	text := ""
	for _, i := range body {
		c := strconv.Itoa(int(i))
		text = text + c + "_"
	}

	text = text[0 : len(text)-1]

	return text
}

func DecodeStringToByte(text string) []byte {
	_bytes := []byte{}
	splitText := strings.Split(text, "_")
	for _, s := range splitText {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		_bytes = append(_bytes, []byte(string(i))...)
	}

	log.Println("Original string - ", string(_bytes))

	return _bytes
}

func (q *QueryStore) StoreQuery(query []byte) string {
	encodedQuery := EncodeByteToString(query)
	return q.Set(encodedQuery)
}

func (q *QueryStore) GetQueryForHash(key string) []byte {
	query := q.Get(key)
	return DecodeStringToByte(query)
}

# query-store
Store and retrieve query using file based storage

## Install query-store
```
go get https://github.com/Kashyap23/query-store
```

## Usage
```
import (
    "encoding/json"
    "fmt"
    queryStore "github.com/Kashyap23/query-store"
)

func main(){
    //Initialize query store directory
    queryStoreInit := queryStore.InitQueryStore("/var/tmp/")
    byteQuery, _ := json.Marshal(map[string]interface{}{"jsonKey1" : "jsonValue1",})

    //Store json query
    id := queryStore.StoreQuery(byteQuery)
    fmt.Println("Got id for query - ", id)

    //Retrive query for id
    jsonQuery := map[string]interface{}
    byteQuery = queryStore.GetQueryForHash(id)

    err := json.Umarshal(byteQuery, &jsonQuery)
    if err != nil {
       fmt.Println("Error during unmarshal - ", err)
    }

    fmt.Println("Response from query store - ", jsonQuery)
}
```

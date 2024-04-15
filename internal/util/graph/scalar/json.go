package scalar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalJson(s interface{}) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		byteData, err := json.Marshal(s)
		if err != nil {
			log.Printf("fail when marshal json%v\n", string(byteData))
		}
		_, err = w.Write(byteData)
		if err != nil {
			log.Printf("fail when write data %v\n", string(byteData))
		}
	})
}

func UnmarshalJson(v interface{}) (interface{}, error) {
	byteData, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("field must be valid graphql query")
	}
	tmp := make(map[string]interface{})
	err = json.Unmarshal(byteData, &tmp)
	if err != nil {
		return nil, fmt.Errorf("field must be valid graphql query [json]")
	}
	return tmp, nil
}

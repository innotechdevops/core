package printx

import (
	"encoding/xml"
	"fmt"

	"github.com/goccy/go-json"
)

func JSON(label string, value any) {
	b, er := json.Marshal(value)
	if er != nil {
		fmt.Println("Error marshalling to JSON:", er)
	}
	fmt.Println(label, string(b))
}

func XML(value any) {
	b, er := xml.Marshal(value)
	if er != nil {
		fmt.Println("Error marshalling to XML:", er)
	}
	fmt.Println(string(b))
}

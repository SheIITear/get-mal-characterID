package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func UnmarshalMal(data []byte) (Mal, error) {
	var r Mal
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Mal) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Mal struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {

	num_args := len(os.Args)

	if num_args < 2 {
		fmt.Println("No character given")
		os.Exit(1)
	}

	characterq := url.QueryEscape(os.Args[1])
	resp, err := http.Get("https://myanimelist.net/search/prefix.json?keyword=" + characterq + "&type=all")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	kk, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	data1, err := UnmarshalMal(kk)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("1. Name: "+"\""+data1.Categories[0].Items[0].Name+"\""+" id:", data1.Categories[0].Items[0].ID)
	fmt.Println("2. Name: "+"\""+data1.Categories[0].Items[1].Name+"\""+" id:", data1.Categories[0].Items[1].ID)
	fmt.Println("3. Name: "+"\""+data1.Categories[0].Items[2].Name+"\""+" id:", data1.Categories[0].Items[2].ID)
	fmt.Println("4. Name: "+"\""+data1.Categories[0].Items[3].Name+"\""+" id:", data1.Categories[0].Items[3].ID)
	fmt.Println("5. Name: "+"\""+data1.Categories[0].Items[4].Name+"\""+" id:", data1.Categories[0].Items[4].ID)
	fmt.Println("6. Name: "+"\""+data1.Categories[0].Items[5].Name+"\""+" id:", data1.Categories[0].Items[5].ID)

}

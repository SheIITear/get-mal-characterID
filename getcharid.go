package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	Categories []Categoryk `json:"categories"`
}

type Categoryk struct {
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
	resp, err := http.Get("https://myanimelist.net/search/prefix.json?keyword=" + characterq + "&type=character")

	if err != nil {
		fmt.Println(err.Error())
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

	k := data1.Categories[0].Items

	var b bytes.Buffer
	var s int

	for i := range k {
		s = i + 1
		b.WriteString(strconv.Itoa(s) + ". Name: " + "\"" + data1.Categories[0].Items[i].Name + "\"" + " id: " + strconv.FormatInt(data1.Categories[0].Items[i].ID, 10) + "\n")
		if i == 5 {
			break
		}
	}

	outb := b.String()
	fmt.Println(outb)
}

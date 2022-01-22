package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"` //mashalの際の名前を設定できる
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`

	/*
		Name      string   `json:"name,omitempty"` //からの場合、intで０の場合に表示されなくなる
		Age       int      `json:"age"`
		Nicknames []string `json:"nickname"`
	*/
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

/*
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct { //新しいstructを作る
		Name      string
		Age       int
		Nicknames []string
	}{
		Name:      "Mr. " + p.Name,
		Age:       p.Age + 100,
		Nicknames: p.Nicknames,
	})
	return v, err
}
*/

func main() {
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil { //ネットワークで入ってきたものをそのままstructのkeyをみて入れてくれる,UnMarshalJSON()が呼ばれる
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames) // => mike 20 [a b c]
	v, _ := json.Marshal(p)                 //MarshalJSON()が呼ばれる
	fmt.Println(string(v))
}

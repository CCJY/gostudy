package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var todoJSON = []byte(`
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false,
  "comment" : {
	  "userId": 2,
	  "message": "just do it"
  }
}
`)

type Comment struct {
	UserID int    `json:"userId"`
	Msg    string `json:"message"`
}

type todo struct {
	UserID      int     `json:"userId"`
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Completed   bool    `json:"completed"`
	LastComment Comment `json:"comment"`
}

func getBodyBytesFromAPI(url string) []byte {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return bytes
}

func (o *todo) toString() string {
	return fmt.Sprintf(`UserID=%d ID=%d Title=%s Completed=%t
	LastCommentMessage=%s 
	`, o.UserID, o.ID, o.Title, o.Completed, o.LastComment.Msg)
}

func (o *todo) toJSONString() (string, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func todoFromBytes(bytes []byte) (*todo, error) {
	t := new(todo)
	err := json.Unmarshal(bytes, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func main() {
	// var url string = "https://jsonplaceholder.typicode.com/todos/1"
	// var bytes []byte = getBodyBytesFromAPI(url)
	// t, err := todoFromBytes(bytes)

	t, err := todoFromBytes(todoJSON)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(t.toJSONString())
	fmt.Println(t.toString())
}

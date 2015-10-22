package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
        "encoding/json"
            "io/ioutil"
        
    
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
type UserInput struct {
    Name  string 
    }

type User struct {
Greeting string
}

var users []User

func create(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {

body, err := ioutil.ReadAll(req.Body)
    
    var t UserInput
    err = json.Unmarshal(body, &t)
        fmt.Println("body unmarshall is ",t)
    
 if err != nil {
        fmt.Println("error is",err)
    }

 u := User{
        Greeting:  "Hello," + t.Name+"!",
        }
        
       // users = append(users, u)
          fmt.Println("Users are ",u)
      	 send(rw, http.StatusOK, u)
    }
func send(w http.ResponseWriter, code int, val interface{}) error {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(code)
     fmt.Println("Val is ",val)
     return json.NewEncoder(w).Encode(val)
    
}



func main() {
    router := httprouter.New()
    router.GET("/hello/:name", Hello)
    router.POST("/hello", create) 
  
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: router,
    
}
    server.ListenAndServe()
}




//C:\>curl -H POST -d "{\"name\":\"mohit\"}" http://localhost:8080/users

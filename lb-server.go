package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "flag"
    "math/rand"
    "strconv"
    "strings"
)

type Cookie struct {
    Name  string
    Value string

    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string

    MaxAge   int
    Secure   bool
    HttpOnly bool
    SameSite SameSite
    Raw      string
    Unparsed []string
}

type SameSite int
const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)

type ServerNodeNumber struct {
    SNNumber string
}
var ServerID = os.Args[1]

func main() {
	finish := make(chan bool)
	if len(os.Args) <= 1  {
    fmt.Println("You need to specify the servers id and node number in this order: " + os.Args[0] + " <integer> 8<integer(1-5)>80\n" )
		os.Exit(2)
	}
	argsPort, _ := strconv.Atoi(os.Args[2])
	switch serverNum := argsPort; serverNum {
	case 1:
                newServer := http.NewServeMux()
		            myHTTPHandler := &ServerNodeNumber{SNNumber: "1"}
                newServer.HandleFunc("/", myHTTPHandler.SidServer)
                go func() {
                        http.ListenAndServe(":8180", newServer)
                }()
	              <-finish

	case 2:
                myHTTPHandler := &ServerNodeNumber{SNNumber: "1"}
                newServer := http.NewServeMux()
                newServer.HandleFunc("/", myHTTPHandler.SidServer)

                myHTTPHandler2 := &ServerNodeNumber{SNNumber: "2"}
                newServer2 := http.NewServeMux()
                newServer2.HandleFunc("/", myHTTPHandler2.SidServer)

                go func() {
                        http.ListenAndServe(":8180", newServer)
                }()
                go func() {
                        http.ListenAndServe(":8280", newServer2)
                }()
                <-finish
	case 3:
                myHTTPHandler := &ServerNodeNumber{SNNumber: "1"}
                newServer := http.NewServeMux()
                newServer.HandleFunc("/", myHTTPHandler.SidServer)

                myHTTPHandler2 := &ServerNodeNumber{SNNumber: "2"}
                newServer2 := http.NewServeMux()
                newServer2.HandleFunc("/", myHTTPHandler2.SidServer)

                myHTTPHandler3 := &ServerNodeNumber{SNNumber: "3"}
                newServer3 := http.NewServeMux()
                newServer3.HandleFunc("/", myHTTPHandler3.SidServer)

                go func() {
                        http.ListenAndServe(":8180", newServer)
                }()
                go func() {
                        http.ListenAndServe(":8280", newServer2)
                }()
                go func() {
                        http.ListenAndServe(":8380", newServer3)
                }()
                <-finish
  case 4:
                myHTTPHandler := &ServerNodeNumber{SNNumber: "1"}
                newServer := http.NewServeMux()
                newServer.HandleFunc("/", myHTTPHandler.SidServer)

                myHTTPHandler2 := &ServerNodeNumber{SNNumber: "2"}
                newServer2 := http.NewServeMux()
                newServer2.HandleFunc("/", myHTTPHandler2.SidServer)

                myHTTPHandler3 := &ServerNodeNumber{SNNumber: "3"}
                newServer3 := http.NewServeMux()
                newServer3.HandleFunc("/", myHTTPHandler3.SidServer)

                myHTTPHandler4 := &ServerNodeNumber{SNNumber: "4"}
                newServer4 := http.NewServeMux()
                newServer4.HandleFunc("/", myHTTPHandler4.SidServer)

                go func() {
                        http.ListenAndServe(":8180", newServer)
                }()
                go func() {
                        http.ListenAndServe(":8280", newServer2)
                }()
                go func() {
                        http.ListenAndServe(":8380", newServer3)
                }()
                go func() {
                        http.ListenAndServe(":8480", newServer4)
                }()
                <-finish
  case 5:
                myHTTPHandler := &ServerNodeNumber{SNNumber: "1"}
                newServer := http.NewServeMux()
                newServer.HandleFunc("/", myHTTPHandler.SidServer)

                myHTTPHandler2 := &ServerNodeNumber{SNNumber: "2"}
                newServer2 := http.NewServeMux()
                newServer2.HandleFunc("/", myHTTPHandler2.SidServer)

                myHTTPHandler3 := &ServerNodeNumber{SNNumber: "3"}
                newServer3 := http.NewServeMux()
                newServer3.HandleFunc("/", myHTTPHandler3.SidServer)

                myHTTPHandler4 := &ServerNodeNumber{SNNumber: "4"}
                newServer4 := http.NewServeMux()
                newServer4.HandleFunc("/", myHTTPHandler4.SidServer)

                myHTTPHandler5 := &ServerNodeNumber{SNNumber: "5"}
                newServer5 := http.NewServeMux()
                newServer5.HandleFunc("/", myHTTPHandler5.SidServer)

                go func() {
                        http.ListenAndServe(":8180", newServer)
                }()
                go func() {
                        http.ListenAndServe(":8280", newServer2)
                }()
                go func() {
                        http.ListenAndServe(":8380", newServer3)
                }()
                go func() {
                        http.ListenAndServe(":8480", newServer4)
                }()
                go func() {
                        http.ListenAndServe(":8580", newServer5)
                }()
                <-finish
	default:
                fmt.Println("You need to specify the servers id and node number in this order: " + os.Args[0] + " <integer> 8<integer(1-5)>80\n" )
                fmt.Println("Please note, node number cannot be higher than 5 or lower than 1. Server ID can be any integer\n" )
                os.Exit(2)
  }
}


func (sN *ServerNodeNumber) SidServer(w http.ResponseWriter, r *http.Request) {
    aCookie, err := r.Cookie("regal.sid")
	  aCookieSid := strconv.Itoa(rand.Int()) + "." + ServerID + "0" + sN.SNNumber
    if err == nil {
        fmt.Fprintf(os.Stdout, "Found cookie: " + aCookie.Value + "\n")
        aCookieNodeNumber := strings.Split(aCookie.Value, ".")
        if aCookieNodeNumber[1] != ServerID + "0" + sN.SNNumber {
          expiration := time.Now().Add(24 * time.Hour)
          cookie := http.Cookie{Name: "regal.sid", Value: aCookieSid, Expires: expiration}
          http.SetCookie(w, &cookie)
          fmt.Fprintf(w, "Invalid Cookie found:" + aCookie.Value + "\nSet new cookie: " + aCookieSid + "\n")
          fmt.Fprintf(os.Stdout, "Invalid Cookie found: " + aCookie.Value + "\n")
        } else {
          fmt.Fprintf(w, "Found valid cookie: " + aCookie.Value + "\n")
        }
	  } else {
		expiration := time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{Name: "regal.sid", Value: aCookieSid, Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "Set cookie: " + aCookieSid + "\n")
		//fmt.Fprintf(os.Stdout," New client: " + aCookieSid + "\n")
	}
}

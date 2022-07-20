package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// bigyihsuan

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Now).Methods("GET")
	r.HandleFunc("/gettest", GetTasks).Methods("GET")
	r.HandleFunc("/fromjson", JsonHandler).Methods("GET")
	r.HandleFunc("/ip", IpHandler).Methods("GET")
	r.HandleFunc("/losangeles", LosAngelesHandler).Methods("GET")
	r.HandleFunc("/taipei", TaipeiHandler).Methods("GET")
	r.HandleFunc("/london", LondonHandler).Methods("GET")
	r.HandleFunc("/{area}/{location}", TimeHandler).Methods("GET").Name("time")

	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", nil)
}

type Task struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type AllTasks = []Task

var tasks = AllTasks{
	Task{1, "Task one", "some content"},
}

// http get
func GetTasks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(tasks)
}

// http post
func GetTask(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(tasks)
}

var lastVisit = make(map[string]time.Time)

func JsonHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	article := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`
	data := make(map[string]string)
	json.NewDecoder(strings.NewReader(article)).Decode(&data)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func IpHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	ip := getUserIp(request)
	rw.Header().Set("Content-Type", "text/html")
	visit, ok := lastVisit[ip]
	lastVisit[ip] = time.Now()
	if !ok {
		fmt.Fprintln(rw, `<p>This is your first visit.</p>`)
		fmt.Fprintf(rw, "<p>It is %v at %v</p>", lastVisit[ip].Format("January 02, 2006"), lastVisit[ip].Format(time.Kitchen))
		fmt.Fprintf(rw, `<p>Your ip is <pre>%s</pre></p>`, ip)
		log.Printf("IP %v first entry", ip)
	} else {
		fmt.Fprintf(rw, "<p>You last visted on date %v and time %v</p>", visit.Format("January 02, 2006"), visit.Format(time.Kitchen))
		fmt.Fprintf(rw, `<p>Your ip is <pre>%s</pre></p>`, ip)
		log.Printf("IP %v entering again\n", ip)
	}
}

func TimeHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	// fmt.Fprintln(rw, request.RequestURI)
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	location := vars["area"] + "/" + vars["location"]
	json.NewEncoder(rw).Encode(getTimeAt(location))
}
func Now(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(getTimeAt(time.Local.String()))
}

func LosAngelesHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(getTimeAt("America/Los_Angeles"))
}

func TaipeiHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(getTimeAt("Asia/Taipei"))
}

func LondonHandler(rw http.ResponseWriter, request *http.Request) {
	log.Println(request.Method, request.RequestURI)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(getTimeAt("Europe/London"))
}
func getTimeAt(location string) map[string]string {
	loc, _ := time.LoadLocation(location)
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	p["loc"] = loc.String()
	return p
}

func getUserIp(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// func HomeHandler(rw http.ResponseWriter, request *http.Request) {
// 	log.Println(request.Method, request.RequestURI)
// 	fmt.Fprintln(rw, "<h1>home</h1>")
// 	fmt.Fprintln(rw, `<p><a href="/newyork">New York City</a></p>`)
// 	fmt.Fprintln(rw, `<p><a href="/taipei">Taipei</a></p>`)
// 	fmt.Fprintln(rw, `<p><a href="/london">London</a></p>`)
// }

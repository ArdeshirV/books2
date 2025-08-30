// go-solati.go: My practices about "The Go programming language reference by Solati"
package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"container/list"
	"context"
	"log"
	"reflect"
	"slices"

	//"strconv"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/ArdeshirV/book/go-solati/colors"
)

// Main entry point
func main() {
	defer func() {
		//waiteForEnter()
		fmt.Print(NORMAL)
	}()
	//PerformTitle()
	//mainChapterOne()
	//mainChapterTwo()
	//mainChapterThree()
	//mainReviewChapterOneAndTwo()
	//mainChannels()
	//mainChannels2()
	//mainChannels3()//
	//UsingRecover()
	//mainANewStepForward()
	//mainAncientAlphabets()
	//mainWriterReader()
	//mainWriterReader2()
	//mainFiles()
	//mainSockets()
	//mainWebBySockets(os.Getenv("HOME") + "/Documents/Pictures/Mine/Myself/Me/me_final_2024.png")
	//mainWebServerBySockets()
	//ainWeb1()
	//mainWebServerByHandleFunc()
	//mainNewWebAfterMux()
	//mainNewWebLessonMux()
	//mainMiddleWareByMUX()
	//mainGet()
	//mainPost()
	//mainNewRequestSolati()
	//mainMySQLtest()
	//mainCreateTableByQuery()
	//mainPrepare()
	//mainMongodb()
	//mainReflection()
	//mainReadStruct()
	//mainReflectionNew()
	//mainStandardLib()
	//mainToUpper()
	//mainDateTime()
	//mainContext()
	//mainLog()
	//mainPlugin()
	//mainZipArchive()
	//mainNext()
	//mainDesignPattern()
	mainDataStructures()
}

func mainDataStructures() {
	fmt.Println(Prompt("Data Structures"))
	usingList()
}

func usingList() {
	l := list.New()
	l.PushBack(4)
	l.PushFront(1)
	fmt.Print(colors.MagentaBold, *l, colors.Normal)
}

func mainDesignPattern() {
	Factory()
	Builder()
	Singleton()
	Adapter()
	Observer()
}

func Adapter() {
	fmt.Println(Prompt("Adapter Design Pattern"))
}

func Observer() {
	fmt.Println(Prompt("Observer Design Pattern"))
}

func Singleton() {
	GetEmail()
	SetEmail()
}

type CacheManager struct {
	cache map[string]any
	mutex sync.RWMutex
}

var instance *CacheManager
var once sync.Once

func GetCacheManager() *CacheManager {
	once.Do(func() {
		instance = &CacheManager{
			cache: make(map[string]any),
		}
	})
	return instance
}

func (cm *CacheManager) Set(key string, value any) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.cache[key] = value
}

func (cm *CacheManager) Get(key string) (any, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	value, ok := cm.cache[key]
	return value, ok
}

func GetEmail() {
	cache := GetCacheManager()
	cache.Set("email", "ArdeshirV@protonmail.com")
}

func SetEmail() {
	cache := GetCacheManager()
	value, ok := cache.Get("email")
	email := value.(string)
	fmt.Println(Prompt("Email:"), Out(fmt.Sprintf("%v %v", email, ok)))
}

func Builder() {
	fmt.Println(Prompt("Builder Design Pattern"))
	car := NewCarBuilder().SetName("Pride").SetColor("White").SetYear(1404).Build()
	fmt.Println(Out(car))
}

func Prompt(text string) string {
	return colors.WhiteBoldText(fmt.Sprintf("%v", text))
}

func In(text string) string {
	return colors.GreenBoldText(fmt.Sprintf("%v", text))
}

func Out(text any) string {
	return colors.MagentaBoldText(fmt.Sprintf("%v", text))
}

type car struct {
	color string
	name  string
	year  int
}

func (c car) String() string {
	return fmt.Sprintf("Name: %s, Color: %s, Year: %d", c.name, c.color, c.year)
}

func (c *car) GetColor() string {
	return c.color
}

func (c *car) GetName() string {
	return c.name
}

func (c *car) GetYear() int {
	return c.year
}

type CarBuilder struct {
	car car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) SetName(name string) *CarBuilder {
	cb.car.name = name
	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *CarBuilder) SetYear(year int) *CarBuilder {
	cb.car.year = year
	return cb
}

func (cb *CarBuilder) Build() *car {
	return &cb.car
}

func Factory() {
	e := CreateE()
	fmt.Println(e)
}

func CreateE() E {
	a := A{'A'}
	b := B{a}
	c := C{b}
	d := D{c}
	return E{d}
}

type A struct {
	r rune
}
type B struct {
	A A
}
type C struct {
	B B
}
type D struct {
	C C
}
type E struct {
	D D
}

func mainNext() {
}

func mainZipArchive() {
	homeDir := os.Getenv("HOME")
	zipFileName := homeDir + "/Documents/Downloads/go.zip"
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	file, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("This is a sample text file that is created in Go programming language"))
}

func mainPlugin() {
	fmt.Println("mainPlugin")
}

func mainLog() {
	logFileName := os.Getenv("HOME") + "/Documents/Downloads/go-practice.log"

	output, err := os.Create(logFileName)
	if err != nil {
		log.Fatal("Failed to create log file")
	}
	defer output.Close()
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.Ldate)
	log.SetOutput(output)
	log.Print(colors.Red, "hello", "a", "b", "c", colors.Normal)
	log.Println("hello", "a", "b", "c")
	go func() {
		for {
			time.Sleep(time.Second * 1)
		}
	}()
	//log.Panic("Fatal error occured")
}

func mainContext() {
	background := context.Background()
	ctx, cancel := context.WithCancel(background)
	go func() {
		//scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Press <Enter> to finish. ")
		var ch rune
		_, err := fmt.Scanf("%c", &ch)
		if err != nil {
			panic(err)
		}
		cancel()
	}()
	time.Sleep(time.Millisecond * 300)
	heavyTask(ctx)
}

func heavyTask(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task canceled because", ctx.Err())
			return
		case <-time.After(time.Second * 100000):
			fmt.Println("Task finished")
			return
		default:
			fmt.Print("\rWorking on heavy task    \b\b\b\b", strings.Repeat(".", i))
			time.Sleep(time.Millisecond * 500)
			i++
			if i >= 4 {
				i = 0
			}
		}
	}
}

func mainDateTime() {
	location, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		panic(err)
	}
	t := time.Date(1985, 06, 04, 03, 0, 0, 0, location)
	fmt.Println(t)
	fmt.Println(time.Since(t))
}

func mainStandardLib() {
	word := "Iran"
	s := []string{"Hello", "Hi", "Ardeshir", "Iran", "Azadi"}
	index := slices.Index(s, word)
	fmt.Printf("Index of %s is %v\n", word, index)
}

func mainToUpper() {
	someString := "This is a string XyZ ایران"
	fmt.Println(ToUpper(someString))
	fmt.Println(ToLower(ToUpper(someString)))
}

func ToUpper(s string) string {
	var sb strings.Builder
	const dist = 'a' - 'A'
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			sb.WriteRune(ch - dist)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func ToLower(s string) string {
	const dist = 'a' - 'A'
	var sb strings.Builder
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			sb.WriteRune(ch + dist)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func testAdd(a, b int) int {
	return a + b
}

func testSub(a, b int) int {
	return a - b
}

func mainReflectionNew() {
	fmt.Print(colors.YellowBoldText("Reflection New\n"))
	someVar := 239.00434
	v := reflect.ValueOf(&someVar).Elem()
	t := reflect.TypeOf(someVar)

	if v.CanSet() {
		v.SetFloat(1000.0001)
	}
	fmt.Print("Value:", v, ", Type:", t, ", Kind:", t.Kind(), "\n")
}

type SomeUser struct {
	user User
	age  int
}

func mainReadStruct() {
	data := SomeUser{user: User{Name: "Ardeshir", Address: "something@somewhere.com"}, age: 19}
	res, err := ReadStruct(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", res)
}

func ReadStruct(s any) (string, error) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("error: the input value of ReadStruct is not a struct instance")
	}

	var sb strings.Builder
	sb.WriteString("Fields:")
	for i := range t.NumField() {
		field := t.Field(i)
		value := v.Field(i)
		sb.WriteString(fmt.Sprintf("%s (%s) = %v\n", field.Name, field.Type, value))
		//if reflect.TypeOf(vv).Kind() == reflect.Struct {
		//	res, err := ReadStruct(value)
		//			if err != nil {
		//				sb.WriteString(res)
		//			}
		//		}
	}

	return sb.String(), nil
}

func mainReflection() {
	fmt.Print(colors.MagentaBoldText("Reflection is here\n"))
	data := User{Name: "Ardeshir", Address: "something@somewhere.com"}
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	fmt.Println("Value:", v, "\nType:", t, "\nKind:", t.Kind())

	numbers := []int{1, 2, 3}
	v = reflect.ValueOf(numbers)
	numbers2 := v.Interface().([]int)
	s := sum(numbers2)
	fmt.Println(s)

	x := 90
	y := &x
	z := &y
	u := &z
	fmt.Printf("%v, %v, %v, %v\n", x, y, z, u)

	fmt.Println()
	if output, err := ReadStruct(&data); err == nil {
		fmt.Println(output)
	}
}

func sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

/*
// The glorious mongodb project!

var context Context

func init() {
	context := Context.GetBackgroundContext()
}

type Person struct {
	Name string
	Age int
	Email string
}

func createPerson(client *mongo.Client, person Person) error {
	collection := client.Database("testdb").Collection("people")
	_, err := collection.InsertOne(context.Background(), person)
	return err
}

func getPersonName(client *mongo.Client, name string) (*Person, error) {
	var person Person
	collection := client.Database("testdb").Collection("people")
	filter := bson.M{"name": name}
	if err := collection.FindOne(context.Background(), filter).Decode(&person); err != nil {
		return nil, err
	}
	return &person, nil
}

func updatePersonAge(client *mongo.Client, name string, age int) error {
	collection := client.Database("testdb").Collection("people")
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"age": age}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

func deletePerson(client *mongodb.Client, name string) error {
	collection := client.Database("testdb").Collection("people")
	filter := bson.M{"name": name}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

func mainMongodb() {
	client, err := connectDB()
	if err != nil {
		log.Fatal("Error connection to Mongodb:", err)
	}
	defer client.Disconnect(context.Background())

	person := Person{
		Name: "John Doe",
		Age: 40,
		Email: "someone@somewhere.com",
	}

	err := createPerson(client, person)
	if err != nil {
		log.Fatal("Error creating person:", err)
	}

	foundPerson, err := getPersonByName(client, "John Doe")
	if err != nil {
		log.Fatal("Error reading person:", err)
	}
	fmt.Println("Found person:", foundPerson)

	if err = updatePersonAge(client, "John Doe", 40000); err != nil {
		log.Fatal("Error updating person:", err)
	}

	if err = deletePerson(client, "John Doe"); err != nil {
		log.Fatal("Error deleting person:", err)
	}
}*/

func mainPrepare() {
	db, err := sql.Open("mysql", GetConnectionStringToMariadb())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Print(colors.YellowBoldText("Connected to database successfully.\n"))

	stm, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?) ")
	if err != nil {
		panic(err)
	}

	result, err := stm.Exec("Ardeshir", "ardeshir@somewhere.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("result: ", result)

	row := db.QueryRow("SELECT name, email FROM users WHERE id = ?", 1)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%sname: %s%s%s, %semail: %s%s%s\n",
		colors.Bold, colors.MagentaBold, name, colors.Normal,
		colors.Bold, colors.MagentaBold, email, colors.Normal)

	result, err = db.Exec("UPDATE users SET email = 'myemail@gmailx.com' WHERE id = 1")
	if err != nil {
		panic(err)
	}

	row = db.QueryRow("SELECT name, email FROM users WHERE id = ?", 1)
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}

	fmt.Println("After modifications:")
	fmt.Printf("%sname: %s%s%s, %semail: %s%s%s\n",
		colors.Bold, colors.MagentaBold, name, colors.Normal,
		colors.Bold, colors.MagentaBold, email, colors.Normal)
}

const mysqlQuery = `
CREATE TABLE users(
	id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
	email VARCHAR(64) CHARACTER set utf8 COLLATE utf8_general_ci DEFAULT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB`

func mainCreateTableQuery() {
	db, err := sql.Open("mysql", GetConnectionStringToMariadb())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to mariadb successfully")

	if _, err = db.Exec(mysqlQuery); err != nil {
		panic(err)
	}
	fmt.Print(colors.GreenBoldText("The table has beend created successfully.\n"))
}

type DataConnection struct {
	username string
	password string
	database string
	host     string
	port     string
}

func NewConnection(username, password, host, port, database string) *DataConnection {
	var c DataConnection
	c.username = username
	c.password = password
	c.host = host
	c.port = port
	c.database = database
	return &c
}

func (c DataConnection) GetPostgresqlConnection() string {
	return fmt.Sprintf("%s%s%s%s%s",
		c.username,
		c.password,
		c.host,
		c.port,
		c.database,
	)
}

func (c DataConnection) GetMySqlConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.username,
		c.password,
		c.host,
		c.port,
		c.database,
	)
}

func (c DataConnection) GetMariadbConnection() string {
	return c.GetMySqlConnection()
}

func GetConnectionStringToMariadb() string {
	const prefix = "LOCAL_MARIADB_"
	dotenv := os.Getenv("DOTENV")
	godotenv.Load(dotenv)
	return NewConnection(
		os.Getenv(prefix+"USERNAME"),
		os.Getenv(prefix+"PASSWORD"),
		os.Getenv(prefix+"HOST"),
		os.Getenv(prefix+"PORT"),
		os.Getenv(prefix+"DATABASE_TEST"),
	).GetMariadbConnection()
}

func mainMySQLtest() {
	db, err := sql.Open("mysql", GetConnectionStringToMariadb())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Printf("%sConnected to database successfully%s\n",
		colors.MagentaBold, colors.Normal)
}

type User struct {
	ID      int `json:"id"`
	Name    string
	Age     int
	Address string
}

func mainNewRequest() {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(User{Name: "AmirReza", Age: 28})
	req, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/posts", buf)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func mainNewRequestSolati() {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(User{Name: "AmirReza", Age: 28})
	req, err := http.NewRequest(
		http.MethodPost,
		"http://jsonplaceholder.typicode.com/posts",
		buf,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	user := new(User)
	json.NewDecoder(resp.Body).Decode(user)
	fmt.Println("User ID:", user.ID)
}

func mainPost() {
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(User{Name: "AmirReza", Age: 28})
	resp, err := http.Post(
		"http://jsonplaceholder.typicode.com/posts",
		"application/json",
		buff,
	)
	if err != nil {
		panic(err)
	}
	user := new(User)
	json.NewDecoder(resp.Body).Decode(user)
	fmt.Println("User ID:", user.ID)
}

func mainGet() {
	address := "https://typeo.top/assets/img/typeoHide.svg"
	resp, err := http.Get(address)
	if err != nil {
		panic(err)
	}
	fileName := os.Getenv("HOME") + "/Documents/Downloads/typo-logo.svg"
	dst, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	io.Copy(dst, resp.Body)
}

type MiddlewareFunc func(http.Handler) http.Handler

func Before(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Run before", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func After(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Run after", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func mainMiddleWareByMUX() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, World!")
		fmt.Fprint(w, "Hello, World!")
	})

	middlewares := []MiddlewareFunc{After, Before}

	var finalHandler http.Handler = mux
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}

	http.ListenAndServe("localhost:64640", finalHandler)
}

func mainNewWebAfterMux() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from server mux1")
	})

	mux2 := http.NewServeMux()
	mux2.HandleFunc("hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello from server mux 2")
	})
	go http.ListenAndServe("localhost:64640", mux1)
	http.ListenAndServe("localhost:64640", mux2)
}

func mainNewWebLessonMux() {
	const address = "localhost:64640"
	fmt.Print("\033[1;35mListen and serve: \033[1;34mhttp://", address, colors.Normal)

	mux := new(mymux)
	http.ListenAndServe(address, mux)
}

type mymux struct{}

func (m *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<b>/ router</b>")
	case "/a":
		fmt.Fprint(w, "<b>/a router</b>")
	case "/b":
		fmt.Fprint(w, "<b>/b router</b>")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<b>Page not found found 404</b>")
	}
}

func mainWebServerByHandleFunc() {
	address := "localhost:64640"
	fmt.Print("\033[1;35mListen and serve: \033[1;34mhttp://", address, "\033[0m\n")
	http.HandleFunc("/", someHandler)
	http.ListenAndServe(address, nil)
}

func someHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b>Server is up and running ...</b>")
}

func mainWebServerBySockets() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		now := time.Now()
		body := "<html><body>"
		body += "<h1>Time</h1>"
		body += fmt.Sprintf("<b>Today: </b> %d:%d:%d",
			now.Year(), now.Month(), now.Day())
		body += "</br>"
		body += fmt.Sprintf("<b>Time: </b>%d:%d", now.Hour(), now.Minute())
		body += "</html></body>"

		response := "HTTP/1.1 200 OK\r\n"
		response += "Content-Type: text/html; charset=UTF-8\r\n"
		response += "Connection: close\r\n"
		response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
		response += "\r\n"
		response += body

		fmt.Fprint(conn, response)
		conn.Close()
	}
}

func mainWebBySockets() {
	address := "localhost:64640"
	headers, body, err := Get(address)
	if err != nil {
		panic(err)
	}
	fmt.Println(headers)
	fmt.Println(body)
}

func Get(addr string) (string, string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rt := "GET / HTTP/1.1\r\n"
	rt += "Host: " + addr + "\r\n"
	rt += "Connection: close\r\n"
	rt += "\r\n"

	fmt.Fprint(conn, rt)
	response, err := io.ReadAll(conn)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(response), "\r\n\r\n")
	headers := parts[0]
	body := parts[1]
	return headers, body, nil
}

func mainWeb1() {
	fmt.Print("mainWeb1()\n\n")
	fmt.Print("The Golang Programming Language\n\n")

}

func mainSockets() {
	listner, err := net.Listen("tcp", "localhost:5060")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	fmt.Println("Server started ...")

	go func() {
		i := 1
		for {
			conn, err := listner.Accept()

			if err == io.EOF {
				fmt.Println(BGREEN, "Connection closed cleanly")
				return
			} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
				fmt.Println(BGREEN, "Read timeout")
				return
			} else if opErr, ok := err.(*net.OpError); ok {
				fmt.Println("Caught net.OpError:")
				fmt.Println(" Operation:", opErr.Op)
				fmt.Println(" Network:", opErr.Net)
				fmt.Println(" Addr:", opErr.Addr)
				fmt.Println(" Inner error:", opErr.Err)
				return
			} else if err != nil {
				fmt.Printf("%sUnexpected error: %v%T\n", BGREEN, err, err)
				return
			}

			go handleConnection(conn, i)
			i++
		}
	}()

	conn, err := net.Dial("tcp", "localhost:64640")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println(BYELLOW, "Connected ...")
	message := ""
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		message = scanner.Text()
		if message == "exit" {
			return
		}
		fmt.Fprint(conn, BYELLOW, message)
		fmt.Print(BYELLOW, ">>")
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println(BBLUE, "New connections number:", i)
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err == io.EOF {
			fmt.Println(BBLUE, "Connection closed cleanly")
			return
		} else if ne, ok := err.(net.Error); ok && ne.Timeout() {
			fmt.Println(BBLUE, "Read timeout")
			return
		} else if err != nil {
			fmt.Printf("%sUnexpected error: %v%T\n", BBLUE, err, err)
			return
		}
		fmt.Printf("%sClient %d: %s\n", BBLUE, i, string(buffer[:n]))
	}
}

func mainFiles() {
	fmt.Println("Files in Golang")

	filename := os.Getenv("HOME") + "/d/sample-file.txt"
	fmt.Println("GetPageSize() = ", os.Getpagesize())

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, "Hello World")
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buff, err := io.ReadAll(f)
	fmt.Println(string(buff))
}

func mainWriterReader2() {
	fmt.Println("Reader & Writer 2")
	reader := strings.NewReader("This is a sample text")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func mainWriterReader() {
	fmt.Println("Reader & Writer")
	buff := []byte("Hello, World!")
	//os.Stdin.Read(buff)
	var s StringX
	s.Write(buff)
	fmt.Println(s)

	var ss String
	ss.Write(buff)
	fmt.Println(ss)

	r := io.TeeReader(&ss, os.Stdout)
	io.ReadAll(r)
	fmt.Println()

	fmt.Println()
	rx := bufio.NewReader(&ss)
	wx := bufio.NewWriter(&ss)

	wx.WriteString("Hello, World!")
	wx.Flush()
	str, err := rx.ReadString('\n')
	fmt.Println(str, err)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	fmt.Println("Hello", scanner.Text())
}

func (s StringX) ToString() string {
	return s.data
}

func mainAncientAlphabets() {
	WriteAlphabets(0x103A0, 0x103DF)
	WriteAlphabets(0x10B00, 0x10B3F)
	WriteAlphabets(0x10B40, 0x10B5F)
	WriteAlphabets(0x10B60, 0x10B7F)
	WriteAlphabets(0xFB50, 0xFDFF)
}

func WriteAlphabets(a, b rune) {
	for alph := a; alph <= b; alph++ {
		fmt.Print(string(alph), " ")
	}
	fmt.Println()
	fmt.Println()
}

func mainANewStepForward() {
	fmt.Println("A new step forward")
	str := "اردشیرThe Go Programming Language"
	buff := make([]byte, 20)
	var s String
	s.Write([]byte(str))
	s.Read(buff)
	fmt.Println(string(buff))
	for i, v := range buff {
		fmt.Printf("[%d]=%v ", i, string(v))
	}
	fmt.Println(s)
	fmt.Println()
	r := rune(buff[4]) | rune(buff[5])<<16
	fmt.Println(string(buff[0:12]))
	fmt.Println(string(r))
	fmt.Printf("%T\n", r)
	rx := rune(194) | rune(169)<<8
	fmt.Println("[", int(rx), "] =", string(rx))
}

type String struct {
	pos  int
	data string
}

type StringX = String

func (s String) String() string {
	return s.data
}

func (s *String) Write(b []byte) (n int, err error) {
	s.data += string(b)
	return len(b), nil
}

func (s *String) Read(b []byte) (n int, err error) {
	n = copy(b, s.data[s.pos:])
	s.pos += n
	if s.pos >= len(s.data) {
		err = io.EOF
	}
	return n, err
}

func UsingRecover() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Recovered from:", rec)
		}
	}()
	panic("Panic happened")
}

func mainChannels3() {
	n := 1000000
	RegisteredUsers := make(chan int, n)
	limit := make(chan bool, runtime.NumCPU())
	now := time.Now()

	for i := 1; i <= n; i++ {
		go func(in int) {
			limit <- true
			RegisteredUsers <- in
			<-limit
		}(i)
	}

	for n > 0 {
		<-RegisteredUsers
		n--
	}

	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Durations:", time.Since(now))
	fmt.Printf("%T\n", time.Second)
}

func mainChannels2() {
	c := make(chan int, 7)
	go func() {
		for i := 100; i >= 0; i-- {
			time.Sleep(time.Millisecond * 300)
			c <- i
		}
	}()
	go func() {
		for i := 1000; i >= 0; i-- {
			time.Sleep(time.Millisecond * 100)
			c <- i
		}
	}()
	for data := range c {
		fmt.Print(data, " ")
		if data == 0 {
			//close(c)
			break
		}
	}
	fmt.Printf("\n%T, len(c) = %v\n", c, len(c))
	c <- 10
	fmt.Printf("\n%T, len(c) = %v\n", c, len(c))
	fmt.Println()
	close(c)
}

func mainChannels1() {
	c := make(chan int, 4)
	for i := range 100 {
		c <- i
		go DoHeavyTask(c)
	}
	close(c)
	fmt.Println()
}

func DoHeavyTask(c chan int) {
	time.Sleep(200 * time.Millisecond)
	data := <-c
	fmt.Printf("%v ", data)
}

// Chapter three
func mainChapterThree() {
	fmt.Printf("  Chapter Three: \n\n")
	fmt.Println("Hello World")
	Integrate("Hello from mainChapterThree +/-")
}

// Integrate ax^2+bx+c
func Integrate(expr string) float64 {
	fmt.Println(expr)
	return 0.0
}

func mainReviewChapterOneAndTwo() {
	printTitle("Review chapter one and two\n")

	const (
		a1 = (iota + 1) * 1024
		a2
		a3
		a4
		_
		a6
	)
	fmt.Printf("a1 = %v, a6 = %v\n", a1, a6)

	arr1 := []uint8{0: 1, 2: 2, 1: 3, 3: 4}
	fmt.Printf("%v\n", arr1)

	m1 := make(map[string]string)
	m1["شنبه"] = "saturday"
	m1["یکشنبه"] = "sunday"
	m1["دوشنبه"] = "monday"
	m1["سهشنبه"] = "tuesday"
	for k, v := range m1 {
		fmt.Printf("%v:%v\n", k, v)
	}
	k := "شنبه"
	fmt.Println(k, "=", m1[k])
	value, ok := m1["حمعه"]
	fmt.Printf("value = %v, ok = %v\n", value, ok)

	type MyType int
	var m MyType
	m = 9
	fmt.Println(m)

	type operator = func(int, int) int
	add := func(a, b int) int { return a + b }
	res := func(a, b int, op operator) int {
		return op(a, b)
	}(1, 2, add)
	fmt.Println("res = ", res)

	p1 := player{
		name: "nightworlf",
		hp:   100,
	}
	fmt.Println("before hit:", p1.hp)
	p1.hit()
	fmt.Println("After hit:", p1.hp)

	//zebel()
}
func zebel() {
	arr := []int{1, 20, 33, 1, 12, 332, 4, 239, 4, 23, 434, 9, 46, 90, 95, 439, 3, 3, 0, 99}
	fmt.Println(arr)
	total1 := 0
	now1 := time.Now()
	for range 1000 {
		total1 += addAllOne(arr)
	}
	fmt.Println("Total: ", total1)
	fmt.Println("Duration: ", time.Since(now1))
	total2 := 0
	now2 := time.Now()
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	go func() {
		for range 250 {
			total2 += addAllOne(arr)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Total: ", total2)
	fmt.Println("Duration: ", time.Since(now2))
}

func addAllOne(a []int) int {
	total := 0
	t := 0
	for range 1000000 {
		for _, v := range a {
			total += v
		}
		t += total
	}
	return total
}

type player struct {
	name string
	hp   int
}

func (p *player) hit() {
	p.hp--
}

// find n max of array
func findNMax(arr []int, n int) []int {
	if n > len(arr) {
		n = len(arr)
	}
	result := make([]int, n)
	copy(result, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[:n]
}

func findMaxPrime(n int) int {
	maxPrime := 2
	for i := 3; i <= n; i++ {
		if isPrime(i) && i > maxPrime {
			maxPrime = i
		}
	}
	return maxPrime
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func rev(text string) string {
	return text
}

func EvalMathematicExpr(expr string) float64 {
	fmt.Print("Hello")
	return 0.0
}

func integrate() {
	fmt.Println("&&&")
}

func mainChapterOne() {
	fmt.Printf("  Chapter One: \n\n")

	fmt.Println("Hello World")
	fmt.Println(`Hello, World!
in multilines.`)

	var age int
	var name string
	var movie string
	var score float64
	name = "Parmis"
	age = 13
	movie = "Hotel Transylvania"
	score = 7
	fmt.Println(name, "is a good student")
	fmt.Println(name, "is", age, "years old")
	fmt.Println(movie, name, "favorite movie score is", score)

	age = 39
	name = "اردشیر"
	fmt.Println(name, "برنامه نویس کامپیوتر است و ", age, "سال سن دارد")

	const (
		a = iota + 1
		b
		c
		_
		e
	)
	fmt.Println(a, b, c, e)

	var bit1 uint8 = 0b11101101
	fmt.Printf("%b, %b\n", bit1, ^bit1)

	str1 := "Hello"
	str2 := "World"
	str3 := fmt.Sprintf("%s, %s!", str1, str2)
	fmt.Println(str3)

	fmt.Printf("%v\n", str1 != str2)
	const bit = 24
	fmt.Println(bit >> 2)
	const b1, b2 = 0b00110000, 0b00100001
	fmt.Printf("%b\n", b1^b2)

	gender := "male"
	if gender == "female" {
		fmt.Println("You are a woman")
	} else {
		fmt.Println("You are a man")
	}

	const day = 2
	dayName := ""

	switch day {
	case 0:
		dayName = "Saturday"
	case 1:
		dayName = "Sunday"
	case 2:
		dayName = "Monday"
	case 3:
		dayName = "Tuesday"
	case 4:
		dayName = "Wednsday"
	case 5:
		dayName = "Tursday"
	case 6:
		dayName = "Friday"
	default:
		dayName = "Xday"
	}

	arrDayNames := [7]string{"Saturday", "Sunday",
		"Monday", "Tuesday", "Wednsday", "Friday"}

	fmt.Println(dayName)
	fmt.Println(arrDayNames[day])

	goto next_statement
	fmt.Println("Khameneee is a dirty pig")
next_statement:
	fmt.Println("Yeah we are here")

	// Functions begins here:
	helloFunc()
	c1 := newCounter()
	c2 := newCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())

	expr := "2*3+4/2*2-1"
	operators := make(map[string]func(int, int) int)
	operators["+"] = add
	operators["-"] = sub
	operators["/"] = div
	operators["*"] = mul
	res := 0
	prev := 0
	newExpr := ""
	for i, v := range expr {
		item := string(v)
		if item == "*" || item == "/" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	for i, v := range expr {
		item := string(v)
		if item == "+" || item == "-" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	fmt.Println("\n", newExpr)

	// Stone 3
	fmt.Println("Stone 3")

}

func mul(a, b int) int {
	return a * b
}

func div(b, a int) int {
	return b / a
}

func add(a, b int) int {
	return a + b
}

func sub(b, a int) int {
	return b - a
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func helloFunc() {
	fmt.Println("helloFunc")
}

func mainChapterTwo() {
	fmt.Printf("  Chapter Two: \n\n")
	fmt.Println("Hello World")
}

func PrintTitle() {
	blnColor := true
	strAppName := "go-solati"
	strAppYear := "2025"
	strAppDescription := "The Go Programming Language"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV/book/go-solati"
	fmt.Print(FormatTitle(strAppName, strAppDescription, strVersion, blnColor))
	fmt.Print(FormatCopyright(strAppYear, strCopyright, strLicense, blnColor))
}

func FormatTitle(strAppName, strAppDescription, strVersion string, blnColor bool) string {
	NoneColored := "%v - %v Version %v\n"
	Colored := "\033[1;33m%v\033[0;33m - %v \033[1;33mVersion %v\033[0m\n"
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppName, strAppDescription, strVersion)
}

func FormatCopyright(strAppYear, strCopyright, strLicense string, blnColor bool) string {
	NoneColored := "Copyright (c) %v %v, Licensed under %v\n"
	Colored := ("\033[0;33mCopyright (c) \033[1;33m%v \033[1;34m%v" +
		"\033[0;33m, \033[1;33m%v\033[0m\n")
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppYear, strCopyright, strLicense)
}

func sprintf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func PerformTitle() {
	PrintTitle()
	bookName := "The Go programming language reference"
	title := "\n    %sMy Practices about \"%s%s%s\" %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, MAGENTA, BMAGENTA, bookName, MAGENTA, BGREEN, TEAL)
}

func printTitle(title string) {
	fmt.Printf("\n  \033[1;36m%s\033[0;36m\n", title)
}

const (
	NORMAL   = "\033[0m"
	BOLD     = "\033[1m"
	RED      = "\033[0;31m"
	TEAL     = "\033[0;36m"
	WHITE    = "\033[0;37m"
	BLUE     = "\033[0;34m"
	GREEN    = "\033[0;32m"
	YELLOW   = "\033[0;33m"
	MAGENTA  = "\033[0;35m"
	BRED     = "\033[1;31m"
	BBLUE    = "\033[1;34m"
	BTEAL    = "\033[1;36m"
	BWHITE   = "\033[1;37m"
	BGREEN   = "\033[1;32m"
	BYELLOW  = "\033[1;33m"
	BMAGENTA = "\033[1;35m"
)

// Problems in book
// P29: sudo -C /usr/local -xzf flle-name --> tar -xvf file-name
// P44: پیدا کردن کوحکترین عدد برای یافتن بزرکترین عدد
// P55: true || false ==> false
// P54: a < b, a < b
// P79: تعریف تابغ به خط فارسی برعکس شده است

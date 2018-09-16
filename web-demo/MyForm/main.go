package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
	"fmt"
	"os"
	"os/signal"
	"io"
	"strconv"
	"regexp"
	"crypto/md5"
)

const (
	Male = iota + 1
	Female
)

type Human struct {
	Name string
	Age int
	Gender int
	Interest []string
}
type User struct {
	Human
	Username string
	Password string
	Role string
	Phone int
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		index(w)
	case "/login":
		login(w, r)
	}
}

func index(w http.ResponseWriter) {
	if tmpl, err := template.ParseFiles("index.html"); err == nil {
		log.Println(tmpl.Execute(w, nil))
	} else {
		log.Fatalf("Parse: %v", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html")
		t.Execute(w, token)
	case "POST":
		formValidation(r)
	}
}

func formValidation(r *http.Request) bool {
	r.ParseForm()
	// 验证token
	token := r.Form.Get("token")
	if token != "" {
		fmt.Println(token)
	} else {
		fmt.Println("token不存在")
		return false
	}
	// 必填处理
	for k, v := range r.Form {
		if len(v[0]) == 0 && k != "interest"{
			log.Println("请填对应写字段!")
			return false
		}
	}
	// 账号是否为 Email 的处理
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, template.HTMLEscapeString(r.Form.Get("username"))); !m {
		log.Println("账号输入错误!")
		return false
	} else {
		log.Println("账号:", r.Form.Get("username"))
	}

	// 判断密码是否一致
	if r.Form.Get("password") != r.Form.Get("confirm_password") {
		log.Println("两次密码不一致!")
		return false
	} else {
		log.Println("密码:", r.Form.Get("password"))
	}

	// 判断年龄是否在正常范围内
	if age, _ := strconv.Atoi(r.Form.Get("age")); age < 0 && age > 120 {
		log.Println("年龄超出正常范围!")
		return false
	} else {
		log.Println("年龄:", r.Form.Get("age"))
	}

	// 判断手机号码是否正确
	if m, _ := regexp.MatchString(`^(1[3|4|5|7|8]\d{9})$`, r.Form.Get("phone")); !m {
		log.Println("手机号码输入错误!")
		return false
	} else {
		log.Println("手机号码:", r.Form.Get("phone"))
	}

	// 判断性别
	if v, _ := strconv.Atoi(r.Form.Get("gender")); v == 1 {
		log.Println("性别: 男")
	} else {
		log.Println("性别: 女")
	}

	// 输出兴趣爱好
	if len(r.Form["interest"]) == 0 {
		log.Println("兴趣爱好:", r.Form["interest"])
	}

	return true
}

// 文件上传
func upload(w http.ResponseWriter, r *http.Request) {
	// 打印请求方法
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	case "POST":
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	server := &http.Server{
		Addr: ":9090",
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.Handle("/", &User{})
	mux.HandleFunc("/upload", upload)
	server.Handler = mux
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func () {
		<-quit
		
		if err := server.Close(); err != nil {
			log.Fatal("Closed Server:", err)
		}
	}()

	log.Println("Start Server ...... ")
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Println("Server closed unexpected")
		}
	}
	log.Println("Server exit")
}
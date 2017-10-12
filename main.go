package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//设置访问的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	//设置监听的端口
	err := http.ListenAndServe(":3389", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//------------------ 开始：路由函数 ------------------

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w, nil)
	}

	t, err := template.ParseFiles("template/404.html")
	if err != nil {
		log.Println(err)
	}

	type errType struct {
		ErrorType int
	}
	tmp := errType{ErrorType: 404}
	// t.Execute(w, nil)
	t.Execute(w, tmp)
}

//------------------ 结束：路由函数 ------------------

// func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		http.Redirect(w, r, "/index", http.StatusFound) //重定向
// 	}

// 	t, err := template.ParseFiles("template/404.html")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	type errType struct {
// 		ErrorType int
// 	}

// 	tmp := errType{ErrorType: 404}
// 	// t.Execute(w, nil)
// 	t.Execute(w, tmp)
// }

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
// 	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
// 	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello world! 你好，世界！") //这个写入到w的是输出到客户端的
// }

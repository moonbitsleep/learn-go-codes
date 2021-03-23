package main

import "net/http"

// Handler 是一个实现了ServeHTTP方法的接口
type helloHandler struct {}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Hello World!</h1>"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>Welcome Index!</h1>"))
}

func about(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("<h1>About me!</h1>" +
		"<p>I am a post graduate student in CNU.</p>"))
}

func main() {
	server := http.Server{
		// 打开端口
		Addr: "localhost:8080",
		// 使用DefaultServeMux
		Handler: nil,
	}

	he := helloHandler{}
	// 向DefaultServeMux注册一个Handler，以供调用
	// 路径， Handler
	http.Handle("/hello", &he)

	// HandleFunc 调用DefaultServeMux.HandleFunc(pattern, handler)
	// 本质调用了mux.Handle(pattern, HandlerFunc(handler))
	// HandlerFunc 实现了 ServeHTTP方法，本身其实是一个Handler，是一种类型
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("<h1>My Home!</h1>"))
	})

	// welcome后不能加()，否则函数会执行
	http.HandleFunc("/", welcome)

	// 强制类型转换
	http.HandleFunc("/about", http.HandlerFunc(about))
	// 运行服务器
	_ = server.ListenAndServe()

}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"net/http/pprof"

	"github.com/felixge/httpsnoop"
	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	flag.Parse()
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":80", logRequestHandler(mux))
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	VERSION := os.Getenv("VERSION")
	w.Header().Add("VERSION", VERSION)

	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	io.WriteString(w, fmt.Sprintf("%s=%s\n", "VERSION", VERSION))
}

//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(h, w, r)
		fmt.Printf("remote request %s, %s, %s, %d\n", RemoteIp(r), r.Method, r.URL.String(), m.Code)
	}
	return http.HandlerFunc(fn)
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
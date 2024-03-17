// MIT License
//
// Copyright (c) 2024 Marcel Joachim Kloubert (https://marcel.coffee)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed ui/*
var webUI embed.FS

func main() {
	TGF_PORT := strings.TrimSpace(
		os.Getenv("TGF_PORT"),
	)
	if TGF_PORT == "" {
		TGF_PORT = "4000"
	}

	// this will provide all files into `ui/` subfolder
	http.Handle(
		"/ui/",
		http.FileServer(
			http.FS(webUI),
		),
	)

	// out backend endpoint
	http.HandleFunc("/api/times/now", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte(r.Method + " method not supported"))

			return
		}

		nowStr := now.Format(time.RFC3339)
		nowStrData := []byte(nowStr)

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.Header().Add("Content-Length", fmt.Sprint(len(nowStrData)))

		w.Write(nowStrData)
	})

	fmt.Println("Server will run on port: " + TGF_PORT)

	err := http.ListenAndServe(":"+TGF_PORT, nil)
	if err != nil {
		panic(err)
	}
}

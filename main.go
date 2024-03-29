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

	api "github.com/mkloubert/the-gitfahther-blog--go-cli-with-vue-ui-demo1/api"
	"github.com/spf13/cobra"
)

//go:embed ui/*
var webUI embed.FS

func main() {
	var port int32

	var rootCmd = &cobra.Command{
		Use:   "go-cli-with-vue-ui-demo1",
		Short: "Demo how to ship Vue app with Go CLI",
		Long:  "This demo shows how to integrate a Vuetify web app into a Go CLI application with implemented web server.",
		Run: func(cmd *cobra.Command, args []string) {
			tcpPort := fmt.Sprint(port)

			// this will provide all files into `ui/` subfolder
			http.Handle(
				"/ui/",
				http.FileServer(
					http.FS(webUI),
				),
			)

			// out backend endpoint
			http.HandleFunc("/api/times/now", api.GetCurrentTime)

			fmt.Println("Server will run on port: " + tcpPort)

			err := http.ListenAndServe(":"+tcpPort, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(666)
			}
		},
	}

	rootCmd.Flags().Int32VarP(&port, "port", "p", 4000, "custom TCP port for the HTTP service")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

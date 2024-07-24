// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
  "io"
)
  


func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting = flag.String("g", "Hello", "Greet with `greeting`")
	addr     = flag.String("addr", "0.0.0.0:8080", "address to serve")
	//addr     = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	// Register handlers.
	// All requests not otherwise mapped with go to greet.
	// /version is mapped specifically to version.
	http.HandleFunc("/", greet)
	http.HandleFunc("/version", version)

  //logfile, err := os.OpenFile(conf.LOGDIR+"/"+conf.AppName+".log", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  //logfile, err := os.OpenFile("./example.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	log.Printf("serving http://%s\n", *addr)
  log.Print("ASDFASDFASDF")
  logfile, err := os.OpenFile("./output.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
  if err != nil {
      log.Fatalf("error opening file: %v", err)
  }
  log.New(logfile, "asdfasdfasdfYAYAYAY", log.Ldate|log.Ltime|log.Lshortfile)
  wrt := io.MultiWriter(os.Stdout, logfile)
  log.SetOutput(wrt)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	log.Printf("asdf")

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
  fmt.Fprint(w, "ASDFASDFASDF")
	fmt.Fprintf(w, "%s, %s!\n", *greeting, html.EscapeString(name))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// read from connection
	request(conn)
}

func request(conn net.Conn) {
	ctr := 0
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
		if ctr == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		ctr++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	// GET /index.html HTTP/1.1
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "about" {
		about(conn)
	}
	if m == "GET" && u == "contact" {
		contact(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Index</title>
	</head>
	<body>
		<h1>Index</h1>
		<a href="/">Index</a><br>
		<a href="/about">About</a><br>
		<a href="/contact">Contact</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>About</title>
	</head>
	<body>
		<h1>About</h1>
		<a href="/">Index</a><br>
		<a href="/about">About</a><br>
		<a href="/contact">Contact</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Contact</title>
	</head>
	<body>
		<h1>Contact</h1>
		<a href="/">Index</a><br>
		<a href="/about">About</a><br>
		<a href="/contact">Contact</a><br>
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

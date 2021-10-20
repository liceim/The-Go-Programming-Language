package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

func execCmd(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			execCmd(c, cmds[0], cmds[1:]...)
		case "cd":
			if err := os.Chdir(cmds[1]); err != nil {
				log.Print(err)
			}
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("file %s: %v", cmds[1], err)
				continue
			}
			mustCopy(c, file)
		case "close":
			return
		default:
			help := "ls: list content\ncd: change directory\nget: get file content\n" +
				"close: close connection\n"
			mustCopy(c, strings.NewReader(help))
		}
	}

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}

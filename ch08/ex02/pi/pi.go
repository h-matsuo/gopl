package pi

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func StartProtocolInterpreter(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "220 Service ready for new user.\n")

	var dataConn net.Conn
	workDir, err := os.Getwd()
	if err != nil {
		io.WriteString(conn, "421 Service not available, closing control connection.\n")
		return
	}
	who := conn.RemoteAddr().String()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		fmt.Printf("[%s] %s\n", who, command)

		args := strings.Split(command, " ")
		switch args[0] {

		case "USER": // USER <username>
			if len(args) != 2 {
				io.WriteString(conn, "501 Syntax error in parameters or arguments.\n")
				continue
			}
			username := args[1]
			u, err := user.Lookup(username)
			if err != nil {
				switch err.(type) {
				case user.UnknownUserError:
					io.WriteString(conn, "530 Not logged in.\n")
				default:
					io.WriteString(conn, "451 Requested action aborted. Local error in processing.\n")
				}
				continue
			}
			// Currently PASS not supported
			// io.WriteString(conn, "331 User name okay, need password.\n")
			// scanner.Scan()
			// command := scanner.Text()
			// fmt.Println(command)
			// args := strings.Split(command, " ")
			// if args[0] != "PASS" {
			// 	io.WriteString(conn, "550 Requested action not taken.\n")
			// 	continue
			// }
			// password := args[1]
			uid, _ := strconv.Atoi(u.Uid)
			if err := syscall.Setuid(uid); err != nil {
				io.WriteString(conn, "530 Not logged in.\n")
				continue
			}
			io.WriteString(conn, "230 User logged in, proceed.\n")

		case "PORT": // PORT <host,port in special format>
			if len(args) != 2 {
				io.WriteString(conn, "500 Syntax error, command unrecognized.\n")
				continue
			}
			options := strings.Split(args[1], ",")
			if len(options) != 6 {
				io.WriteString(conn, "501 Syntax error in parameters or arguments.\n")
				continue
			}
			p1, _ := strconv.Atoi(options[4])
			p2, _ := strconv.Atoi(options[5])
			address := fmt.Sprintf("%s.%s.%s.%s:%d", options[0], options[1], options[2], options[3], p1*256+p2)
			dConn, err := net.Dial("tcp", address)
			if err != nil {
				io.WriteString(conn, "421 Service not available, closing control connection.\n")
				break
			}
			defer dConn.Close()
			dataConn = dConn
			io.WriteString(conn, "200 Command okay.\n")

		case "PWD": // PWD
			io.WriteString(conn, fmt.Sprintf("257 \"%s\".\n", workDir))

		case "LIST", "NLIST": // LIST [<pathname>]
			if len(args) > 2 {
				io.WriteString(conn, "500 Syntax error, command unrecognized.\n")
				continue
			}
			if dataConn == nil {
				io.WriteString(conn, "426 Connection closed; transfer aborted.\n")
				continue
			}
			pathname := workDir
			if len(args) == 2 {
				pathname = args[1]
			}
			fileInfoList, err := ioutil.ReadDir(pathname)
			if err != nil {
				io.WriteString(conn, "550 Requested action not taken.\n")
				continue
			}
			io.WriteString(conn, "125 Data connection already open; transfer starting.\n")
			for _, fileInfo := range fileInfoList {
				io.WriteString(dataConn, fmt.Sprintf("%s\r\n", fileInfo.Name()))
			}
			dataConn.Close()
			dataConn = nil
			io.WriteString(conn, "226 Closing data connection.\n")

		case "CWD": // CWD <pathname>
			if len(args) != 2 {
				io.WriteString(conn, "500 Syntax error, command unrecognized.\n")
				continue
			}
			pathname := args[1]
			if filepath.IsAbs(pathname) {
				workDir = pathname
			} else {
				workDir = filepath.Join(workDir, pathname)
			}
			io.WriteString(conn, fmt.Sprintf("200 directory changed to %s.\n", workDir))

		// case "TYPE":
		// TODO: 実装？

		case "RETR": // RETR <pathname>
			if len(args) != 2 {
				io.WriteString(conn, "500 Syntax error, command unrecognized.\n")
				continue
			}
			if dataConn == nil {
				io.WriteString(conn, "426 Connection closed; transfer aborted.\n")
				continue
			}
			pathname := args[1]
			f, err := os.Open(pathname)
			if err != nil {
				io.WriteString(conn, "550 Requested action not taken.\n")
				continue
			}
			io.WriteString(conn, "125 Data connection already open; transfer starting.\n")
			io.Copy(dataConn, bufio.NewReader(f))
			dataConn.Close()
			dataConn = nil
			io.WriteString(conn, "226 Closing data connection.\n")

		// case "STOR":
		// TODO: 実装

		case "QUIT": // QUIT
			io.WriteString(conn, "221 Service closing control connection.\n")
			break

		default:
			io.WriteString(conn, "502 Command not implemented.\n")
		}
	}
}

package main
import ("fmt";"sync";"os/exec";"os";"io/ioutil"; "net";"strings")


func main() {

	if len(os.Args)<2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <port num>\n", os.Args[0]);
		return;
	}

	port := fmt.Sprintf(":%s", os.Args[1])	
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening: %s\n", err)
		return
	}
	
	for {
		conn, err := ln.Accept()
		if err!=nil {
			fmt.Fprintf(os.Stderr, "Error accepting: %s", err)
			continue;
		}

		fmt.Println("Got a new connection on port ", port)
		processConnection(conn)
	}
}

func processConnection(conn net.Conn) {
	sendHeader(conn)
	sendFile("firstHalf.html", conn)
	m := mapIP(conn.RemoteAddr())
	sendFile("secondHalf.html", conn)
	sendString(m, conn)
	sendFile("thirdHalf.html",conn)
	conn.Close()
}

func sendHeader(conn net.Conn) {
	sendString("HTTP/1.0 200 OK\r\n", conn)
	sendString("fsociety: yes\r\n", conn)
	sendString("vertigo: yes\r\n", conn)
	sendString("fear-mode: yes\r\n", conn)
	sendString("HackMode: enabled\r\n", conn)
	sendString("Date: February 30, 2018\r\n", conn)
	sendString("served: revenge\r\n", conn)
	sendString("XSS: Enabled. Do not erase\r\n", conn)
	sendString("Content-Type: text/html\r\n", conn)
	sendString("\r\n", conn)
}

func sendFile(fn string, conn net.Conn) {
	dat, err := ioutil.ReadFile(fn)
	if err!=nil {fmt.Println("Couldn't open file: ", err);return}
	sendString(string(dat), conn)
}

func sendString(s string, conn net.Conn) {
	conn.Write([]byte(s))
}


func mapIP(ip net.Addr) string {

	fmt.Fprint(os.Stderr,".")

	ip_str := strings.Split(ip.String(), ":")[0]
	cmd :=exec.Command("nmap", "-Pn", "-T5", ip_str)
	out, e := cmd.Output()
	if e==nil {
		fmt.Fprintf(os.Stderr, "mapped IP: %s", ip)
		output(string(out))
	}else {
		fmt.Fprintf(os.Stderr, "err: %s    %s", out, e)
	}
	
	return string(out)
}

var outputMutex sync.Mutex
func output(s string) {
	outputMutex.Lock()

	fmt.Println(s)
	
	outputMutex.Unlock()
}





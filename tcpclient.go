package main
import (
    "net"
    "os"
    "time"
)

func main() {
    strEcho := "Halo"
    servAddr := "localhost:5000"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write([]byte(strEcho))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("write to server = ", strEcho)

    reply := make([]byte, 1024)

    _, err = conn.Read(reply)
    //time.Sleep(2 * time.Second)
    if err != nil {
        println("Read from server failed:", err.Error())
        os.Exit(1)
    }
    time.Sleep(2 * time.Second)
    println("reply from server=", string(reply))

    conn.Close()
}

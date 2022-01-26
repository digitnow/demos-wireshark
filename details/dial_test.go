// Basert på koden og fri oversettelse fra
//   Woodbeck, A. (2021). Network programming with Go :
//   code secure and reliable network services from scratch. 
//   No Starch Press.
//   Chapter 3: Reliable TCP Data Streams 

package details

import (
    "io"
    "net"
    "testing"
)

func TestDial(t *testing.T) {
    // Lage en "listener" på en tilfeldig port
    listener, err := net.Listen("tcp", "127.0.0.1:")
    if err != nil {
        t.Fatal(err)
    }
    
    done := make(chan struct{})
    go func() {
        defer func() { done <- struct{}{} }()

        for {
            conn, err := listener.Accept()
            if err != nil {
                t.Log(err)
                return
            } 
           
            go func(c net.Conn) {
                defer func() {
                    c.Close()
                    done <- struct{}{}
                }()
          

                buf := make([]byte, 1024)
                for {
                    n, err := c.Read(buf)
                    if err != nil {
                        if err != io.EOF {
                            t.Error(err)
                        }
                        return
                    } 
               
                    t.Logf("received: %q", buf[:n])
                }
            }(conn)
        }
    }()
    conn, err := net.Dial("tcp", listener.Addr().String())
    conn.Close() 
    <-done
    listener.Close()
    <-done
}

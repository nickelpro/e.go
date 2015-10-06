package main

import (
  "log"
  "net"
)

const (
  RESPONSE = " : USERID : UNIX : IDENT.HERE\r\n"
  CONN_PORT = "113"
  CONN_TYPE = "tcp"
)

func main() {
  l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
  if err != nil {
    log.Fatal(err)
  }
  defer l.Close()
  for {
    conn, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go func(c net.Conn) {
      buf := make([]byte, 4096)
      i, err := c.Read(buf)
      if err != nil {
        log.Print(err)
        c.Close()
        return
      }
      for buf[i - 1] != "\n" {
        j, err := c.Read(buf)
        if err != nil {
          log.Print(err)
          c.Close()
          return
        }
        i += j
      }
      if buf[i - 2] == "\r" {
        i -= 1
      }
      c.Write(append(buf[:i-1], []byte(RESPONSE)...))
      c.Close()
    }(conn)
  }
}

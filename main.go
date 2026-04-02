package main

import (
    "io"
    "net"
    "os"
    "log"
)

func handleClient(clientConn net.Conn, targetAddr string) {
    defer clientConn.Close()

    remoteConn, err := net.Dial("tcp", targetAddr)
    if err != nil {
        log.Printf("Failed to connect to backend %s: %v", targetAddr, err)
        return
    }
    defer remoteConn.Close()

    go func() {
        _, _ = io.Copy(remoteConn, clientConn)
    }()

    _, _ = io.Copy(clientConn, remoteConn)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    listenAddr := ":" + port
    
    targetAddr := os.Getenv("V2RAY_SERVER_IP")
    if targetAddr == "" {
        log.Fatal("V2RAY_SERVER_IP environment variable is required")
    }

    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatalf("Failed to listen on %s: %v", listenAddr, err)
    }
    defer listener.Close()

    log.Printf("Listening on %s, proxying to %s", listenAddr, targetAddr)

    for {
        clientConn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go handleClient(clientConn, targetAddr)
    }
}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/database"
	"github.com/myustanzah/GO-Fiber.git/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:               false, // true untuk otomatis menjalankan beberapa instance/child memaksilamlakn performa server
		CaseSensitive:         true,  // true untuk membedakan huruf besar dan kecil pada URL
		StrictRouting:         true,  // true untuk mengharuskan trailing slash pada URL
		DisableStartupMessage: true,  // true untuk menonaktifkan pesan startup
		// DisableKeepalive: true, // true untuk menonaktifkan koneksi keep-alive
		// DisableDefaultDate: true, // true untuk menonaktifkan header tanggal default
		// DisableHeaderNormalizing: true, // true untuk menonaktifkan normalisasi header
		// DisableCache: true, // true untuk menonaktifkan cache
		// DisableRequestLogging: true, // true untuk menonaktifkan logging request
		// DisableErrorHandler: true, // true untuk menonaktifkan error handler default
		ServerHeader: "Fiber",
		AppName:      "MyFiberApp-v1.0",
	})
	db, err := database.InitDB() // Initialize the database connection
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	// Check if the process is a child process
	if !fiber.IsChild() {
		fmt.Println("is the parent process")
	} else {
		fmt.Println("is the child process")
	}

	routes.SetupRoutes(app.Group("/dummy/v1"))             // Setup routes with a versioned API group
	routes.SetupUserRoutes(app.Group("/api/v1/users"), db) // Setup user routes with a versioned API group

	fmt.Println("Server is running on http://localhost:3040")
	app.Listen(":3040")

	// Uncomment the line below to enable HTTPS with TLS
	// app.ListenTLS(":443", "./cert.pem", "./cert.key");

	// Uncomment the line below to enable HTTPS with a custom certificate
	// app.ListenTLSWithCertificate(":443", cert);

	// Custom host
	// app.Listen("127.0.0.1:8080")

	// Uncomment the line below to enable mutual TLS
	// app.ListenMutualTLS(":443", "./cert.pem", "./cert.key", "./ca-chain-cert.pem");

	// Uncomment the line below to enable mutual TLS with a certificate and client certificate pool
	// app.ListenMutualTLSWithCertificate(":443", cert, clientCertPool);

	// Uncomment the line below to enable TLS with a custom listener
	// This requires a valid TLS certificate and key files
	// ln, _ := net.Listen("tcp", ":3000")
	// cer, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	// ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})
	// app.Listener(ln)
}

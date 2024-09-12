package main

import (
    "fmt"
    "log"
    "os"
    "github.com/jackc/pgx/v4"
    "github.com/joho/godotenv"
    "net/url"
    "context"
)

func main() {
    // Cargar variables de entorno desde .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Obtener las credenciales de conexión
    host := os.Getenv("SUPABASE_HOST")
    dbname := os.Getenv("SUPABASE_DB")
    user := url.QueryEscape(os.Getenv("SUPABASE_USER"))
    password := url.QueryEscape(os.Getenv("SUPABASE_PASSWORD"))
    port := os.Getenv("SUPABASE_PORT")

    // Crear la URL de conexión
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

    // Conectar a Supabase
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer conn.Close(context.Background())

    fmt.Println("Conexión exitosa con Supabase")

    // Ejemplo de consulta
    var greeting string
    err = conn.QueryRow(context.Background(), "SELECT 'Hello, Supabase!'").Scan(&greeting)
    if err != nil {
        log.Fatalf("Query failed: %v\n", err)
    }

    fmt.Println(greeting) // Debería imprimir "Hello, Supabase!"
}

package main

import (
    "bytes"
    "fmt"
    "os/exec"
)

func main() {
    // Palabra a buscar
    palabra := "Go"

    // Bloque de texto (entrada simulada)
    texto := `Go es un lenguaje de programación.
La concurrencia en Go es sencilla.
Java es otro lenguaje.`

    // Preparamos el comando grep (Windows → usar findstr)
    // En Linux/Mac: exec.Command("grep", palabra)
    cmd := exec.Command("findstr", palabra)

    // Conectamos stdin (entrada) al texto
    cmd.Stdin = bytes.NewBufferString(texto)

    // Capturamos la salida
    var out bytes.Buffer
    cmd.Stdout = &out

    // Ejecutamos y verificamos errores
    if err := cmd.Run(); err != nil {
        fmt.Println("❌ Error ejecutando el comando:", err)
        return
    }

    // Mostramos el resultado
    fmt.Println("Líneas encontradas:")
    fmt.Println(out.String())
}

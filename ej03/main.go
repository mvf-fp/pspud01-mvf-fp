package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
)

func main() {
    // Leer argumento de búsqueda o usar "go" por defecto
    term := "go"
    if len(os.Args) > 1 {
        term = os.Args[1]
    }

    // En Linux/Mac sería:
    // cmd1 := exec.Command("ps", "-edf")
    // cmd2 := exec.Command("grep", "-i", term)

    // En Windows usamos tasklist y findstr (equivalente)
    cmd1 := exec.Command("tasklist")
    cmd2 := exec.Command("findstr", "/i", term)

    // Conectar stdout de cmd1 a stdin de cmd2
    pipeReader, err := cmd1.StdoutPipe()
    if err != nil {
        fmt.Println("Error creando pipe:", err)
        return
    }
    cmd2.Stdin = pipeReader

    // Capturar salida final de cmd2
    cmd2.Stdout = os.Stdout

    // Iniciar primer comando
    if err := cmd1.Start(); err != nil {
        fmt.Println("Error iniciando cmd1:", err)
        return
    }

    // Iniciar segundo comando
    if err := cmd2.Start(); err != nil {
        fmt.Println("Error iniciando cmd2:", err)
        return
    }

    // Esperar a que terminen ambos
    if err := cmd1.Wait(); err != nil {
        fmt.Println("Error esperando cmd1:", err)
    }

    if err := cmd2.Wait(); err != nil {
        // findstr devuelve código 1 si no hay coincidencias → no es fallo real
        if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
            fmt.Println("  No se encontraron coincidencias.")
        } else {
            fmt.Println("Error esperando cmd2:", err)
        }
    }

    // Cerrar pipe explícitamente
    io.Copy(os.Stdout, pipeReader)
}

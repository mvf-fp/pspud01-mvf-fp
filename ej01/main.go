package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // En Windows usamos "cmd /C dir"
    cmd := exec.Command("cmd", "/C", "dir")

    salida, err := cmd.Output()
    if err != nil {
        fmt.Println("Error al ejecutar el comando:", err)
        return
    }

    fmt.Println("La salida del comando es:\n")
    fmt.Println(string(salida))
}

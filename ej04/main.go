package main

import (
    "fmt"
    "os/exec"
)

func main() {
    var cmd *exec.Cmd

    // En Windows usamos "cmd /C dir" con un nombre inexistente para generar error
    // En Linux/Mac sería: exec.Command("find", "/etc", "-name", "hosts")
    cmd = exec.Command("cmd", "/C", "dir", "C:\\Windows\\System32\\drivers\\etc\\hosts")

    // Capturamos stdout + stderr combinados
    salida, err := cmd.CombinedOutput()

    fmt.Println("Salida combinada del comando:\n")
    fmt.Println(string(salida))

    if err != nil {
        fmt.Println("Error de ejecución:", err)
    } else {
        fmt.Println("Comando ejecutado correctamente ")
    }
}

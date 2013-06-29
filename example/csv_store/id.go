package main

import (
    "math/rand"
    "time"
    "crypto/md5"
    "io"
    "strconv"
    "fmt"
)

var idLen = 6

func random (min, max int) int {
    rand.Seed (time.Now().Unix())
    return rand.Intn(max - min) + min
}


func main () {
    h := md5.New ()
    io.WriteString(h, strconv.Itoa(random (1, 100)))

    str := fmt.Sprintf ("%x", h.Sum(nil))
    startPoint := random(1, len(str) - idLen)
    println()
    print(str[startPoint:startPoint + idLen])
    println()
}

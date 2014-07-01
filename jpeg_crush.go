package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "time"
  "flag"
)

const HEADER = `
     ____._____________________ ________    _________   _____      _____    _________ ___ ___
    |    |\______   \_   _____//  _____/   /   _____/  /     \    /  _  \  /   _____//   |   \
    |    | |     ___/|    __)_/   \  ___   \_____  \  /  \ /  \  /  /_\  \ \_____  \/    ~    \
/\__|    | |    |    |        \    \_\  \  /        \/    Y    \/    |    \/        \    Y    /
\________| |____|   /_______  /\______  / /_______  /\____|__  /\____|__  /_______  /\___|_  /
                            \/        \/          \/         \/         \/        \/       \/
`

const SUBHEADER = `
By JD ...this shiz just keeps rolling yo!

`

var (
  JPEG_START = []byte{ 0xff, 0xda }
  JPEG_END   = []byte{ 0xff, 0xd9 }
  JAMES      = []byte{ 0x4a, 0x61, 0x6d, 0x65, 0x73 }
)

var inputFile, outputFile string

func init() {
  flag.StringVar(&inputFile, "f", "", "Input filename of the JPEG to destroy")
  flag.StringVar(&outputFile, "o", "", "Output filename of the JPEG output blah blah blah...")
}

func main() {
  flag.Parse()

  if flag.NFlag() != 2 {
    red()
    fmt.Printf(HEADER)
    blue()
    fmt.Printf(SUBHEADER)
    reset()
    flag.PrintDefaults()
    fmt.Println()
    os.Exit(0)
  }

  rand.Seed(time.Now().UnixNano())

  f, err := ioutil.ReadFile(inputFile)
  if err != nil {
    panic("Can't open the file yo!")
  }
  start := bytes.Index(f, JPEG_START)
  end   := bytes.Index(f, JPEG_END)

  fmt.Println("Start of JPEG: ", start, "End of JPEG: ", end)

  // flick over each one and replace the bytes
  for _, j := range(JAMES) {
    randy := rand.Intn(end - start)
    reset()
    fmt.Printf("Byte at offset: ")
    red()
    fmt.Printf("%v ", randy)
    fmt.Printf(" with Hex: ")
    blue()
    fmt.Printf("%x \n", j)
    f[start:end][randy] = j
  }

  // Write the output
  err2 := ioutil.WriteFile(outputFile, f, os.FileMode(0755))
  if err2 != nil {
    panic("Boom!! Something very bad happened in the file creation...")
  }
  fmt.Println("Written file ;)")

}

// Little helpers for colour output in the terminal

func red() {
  fmt.Print("\x1b[31;1m")
}

func reset() {
  fmt.Print("\x1b[0m")
}

func blue() {
  fmt.Print("\x1b[34;1m")
}

package main

import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "YOLO"
  })
  m.Run()
}

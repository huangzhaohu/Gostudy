package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Skills struct {
	Passive string `xml:"passive"`
	One string `xml:"one"`
	Two string `xml:"two"`
	Three string `xml:"three"`
}
type Hero struct {
	Name string `xml:"name"`
	Skills Skills `xml:"skills"`
}
type Heros struct {
	Hero []Hero `xml:"hero"`
}
type Game struct {
	Name string `xml:"name"`
	Type string `xml:"type"`
	Heros Heros `xml:"heros"`
}
type Games struct {
	Game []Game `xml:"game"`
}
func main()  {
	file, err := ioutil.ReadFile("C:/Alldata/codePlace/go/src/awesomeProject/demo01/game.xml")
	if err != nil {
		fmt.Fprintln(os.Stderr,err.Error())
		os.Exit(9)
	}
	var games Games
	xml.Unmarshal(file,&games)
	fmt.Println(games)
}

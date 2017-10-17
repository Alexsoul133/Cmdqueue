package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alexsoul133/cmdqueue"
)

var path string = "\\" + filepath.Join("\\", "Rackstation3", "Yandex", "Ton")

type Tfflite struct {
	srcname, dstname, vCh, crop, scale, afilter string
}

func Newfflite(o *Tfflite) {
	queue := []string{"fflite", "-i", o.srcname, "-map_metadata -1 -map_chapters -1 -map", o.vCh, "-vcodec libx264 -preset medium -crf 17 -pix_fmt yuv420p -g 0 -map", o.afilter, "-acodec aac -ab 256k -metadata:s:a:0 language=rus -metadata:s:a:0 \"handler=Russian 2.0\" -disposition:1 default -n", path + "\\" + o.dstname}
	cmdqueue.Newcmd(queue)
}
func main() {

	inFile, _ := os.Open(os.Args[1])
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		name := scanner.Text()
		strings.TrimSuffix(name, ".mov")
		Newfflite(&Tfflite{
			srcname: "\"" + name + ".mov\"",
			dstname: name + ".mp4",
			vCh:     "0:v",
			crop:    "",
			scale:   "",
			afilter: "0:a"})
	}
	cmdqueue.Start()
}

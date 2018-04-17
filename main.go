package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./imgconv"
)

var dstFmt string
var srcFmt string

func init() {
	flag.StringVar(&dstFmt, "dst", ".png", "set a converted image format. default value is .png")
	flag.StringVar(&srcFmt, "src", ".jpg", "set a image format to be converted. default value is .jpg")
}

func main() {
	flag.Parse()

	// 指定できるディレクトリは１つ
	if len(flag.Args()) != 1 {
		fmt.Println("Input only one directory.")
		os.Exit(1)
	}

	// 対象のディレクトリ
	dir := flag.Arg(0)

	// jpgファイルをimgディレクトリからさがす
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == srcFmt {
			err := imgconv.Convert(path, dstFmt)
			if err != nil {
				fmt.Printf("convertion error. file : %s, err : %s\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		os.Exit(1)
	}
}

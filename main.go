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
	// TODO: 無名関数のerrを処理する
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == srcFmt {
			fmt.Println(path)
			imgconv.Convert(path, dstFmt)
		}
		return nil
	})

	if err != nil {
		// TODO: エラー処理
	}
}

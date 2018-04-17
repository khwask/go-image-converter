package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./imgconv"
)

// 変換後の拡張子（例：.png）
var dstFmt string

// 変換前の拡張子（例：.jpg）
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

	// 変換前後のフォーマット名が同じならエラー
	if dstFmt == srcFmt {
		fmt.Println("Set differect file format.")
		os.Exit(1)
	}

	// 対象のディレクトリ
	dir := flag.Arg(0)

	// jpgファイルをimgディレクトリからさがす
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == srcFmt {
			file := new(imgconv.ImageFile)
			file.Path = path
			err := imgconv.Convert(file, dstFmt)
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

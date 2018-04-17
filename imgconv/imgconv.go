package imgconv

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

/*
Convert ファイル形式をJPGからPNGに変換する
*/
func Convert(srcfilename string, dstformat string) error {

	// ファイルオープン
	file, err := os.Open(srcfilename)
	if err != nil {
		return err
	}
	defer file.Close()

	// JPEG前提でdecode
	// TODO: ファイル拡張子なんでもうけれるようにしたい
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// 出力ファイル作成
	out, err := os.Create(setfilename(srcfilename, dstformat))
	if err != nil {
		return err
	}
	defer out.Close()

	// PNG形式でencodeして出力
	// TODO: ファイル拡張子なんでもうけれるようにしたい
	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}

// 指定の拡張子のファイル名を作る
func setfilename(srcfilename string, extention string) string {
	return strings.Replace(srcfilename, filepath.Ext(srcfilename), extention, 1)
}

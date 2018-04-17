package imgconv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

/*
Image タイプ
*/
type Image image.Image

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

	// Imageをdecode
	var img Image
	switch filepath.Ext(srcfilename) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return errors.New("Unsupported file decode format")
	}
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
	switch dstformat {
	case ".jpg", ".jpeg":
		option := &jpeg.Options{Quality: 100}
		err = jpeg.Encode(out, img, option)
	case ".png":
		err = png.Encode(out, img)
	default:
		return errors.New("Unsupported file encode format")
	}
	if err != nil {
		return err
	}

	return nil
}

// 指定の拡張子のファイル名を作る
func setfilename(srcfilename string, extention string) string {
	return strings.Replace(srcfilename, filepath.Ext(srcfilename), extention, 1)
}

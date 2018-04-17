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
ImageFile は、ファイルパスを持つユーザ定義型
*/
type ImageFile struct {
	Path string
}

// 拡張子を指定のものに変えたファイル名を返す
func (file *ImageFile) changeExt(ext string) string {
	return strings.Replace(file.Path, filepath.Ext(file.Path), ext, 1)
}

/*
Convert ファイル形式を指定のものに変換する
*/
func Convert(file *ImageFile, dstformat string) error {

	// ファイルオープン
	f, err := os.Open(file.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Imageをdecode
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// 出力ファイル作成
	out, err := os.Create(file.changeExt(dstformat))
	if err != nil {
		return err
	}
	defer out.Close()

	// encodeして出力
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

package imgconv

import (
	"image/jpeg"
	"image/png"
	"os"
)

/*
Convert ファイル形式をJPGからPNGに変換する
*/
func Convert(srcfilename string, dstfilename string) error {

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
	out, err := os.Create(dstfilename)
	if err != nil {
		return err
	}

	// PNG形式でencodeして出力
	// TODO: ファイル拡張子なんでもうけれるようにしたい
	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

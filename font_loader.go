package golaypdf

import (
	"fmt"
	"io"
	"io/fs"
)

type FontLoader interface {
	// TODO: support font query as in fontconfig-user.html
	OpenFont(familyStr, styleStr string) (string, io.ReadCloser, error)
}

func FSFontLoader(
	fsys fs.FS,
	fontNamer func(family, style string) string,
) FontLoader {
	return fsFontBytesProvider{
		fsys:      fsys,
		fontNamer: fontNamer,
	}
}

type fsFontBytesProvider struct {
	fsys      fs.FS
	fontNamer func(family, style string) string
}

func (f fsFontBytesProvider) OpenFont(family, style string) (string, io.ReadCloser, error) {
	filename := f.fontNamer(family, style)

	fontFile, err := f.fsys.Open(filename)
	if err != nil {
		return "", nil, fmt.Errorf("open font file: %w", err)
	}

	return filename, fontFile, nil
}

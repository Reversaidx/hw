package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	srt, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer srt.Close()
	srtStat, err := srt.Stat()
	if err != nil {
		return err
	}
	srtSize := srtStat.Size()
	if srtSize == 0 {
		return ErrUnsupportedFile
	}
	if offset > srtSize {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 {
		limit = srtSize
	}
	srt.Seek(offset, 0)
	fmt.Println(srtSize)
	dst, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func() {
		dst.Close()
	}()
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(srt)
	bar.Set(pb.Bytes, true)
	_, err = io.CopyN(dst, barReader, limit)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}
	bar.Finish()
	fmt.Println("kurwa")
	// Place your code here.
	return nil
}

package downloader

import (
	"time"

	"github.com/cheggaaa/pb"
)

func progressBar(size int) *pb.ProgressBar {
	bar := pb.New(size).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.ShowFinalTime = true
	bar.SetMaxWidth(1000)

	return bar
}

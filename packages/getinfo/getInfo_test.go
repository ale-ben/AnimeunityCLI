package getinfo_test

import (
	"AnimeunityCLI/packages/getinfo"
	"testing"
)

func benchmarkGetInfo(keyword string, b *testing.B) {
	// run the DownloadURL function b.N times
	for n := 0; n < b.N; n++ {
		getinfo.SetLogLevel("warn")
		getinfo.GetInfo(keyword)
	}
}

func BenchmarkGetInfoDxD(b *testing.B) {
	benchmarkGetInfo("DxD",b)
}

func BenchmarkGetInfoPL(b *testing.B) {
	benchmarkGetInfo("Punch Line",b)
}
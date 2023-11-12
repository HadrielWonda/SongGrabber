package commands

import (
	"strings"

	"github.com/hadrielwonda/songgrabber"
	"github.com/hadrielwonda/songgrabber/chiasenhac"
	"github.com/hadrielwonda/songgrabber/facebook"
	"github.com/hadrielwonda/songgrabber/nhaccuatui"
	"github.com/hadrielwonda/songgrabber/soundcloud"
	"github.com/hadrielwonda/songgrabber/vimeo"
	"github.com/hadrielwonda/songgrabber/youtube"
	"github.com/hadrielwonda/songgrabber/zing"
)

func getsonggrabber(link string) songgrabber.Source {
	switch {
	case strings.Contains(link, initNhacCuaTui):
		return &nhaccuatui.NhacCuaTui{}
	case strings.Contains(link, initZingMp3):
		return &zing.Zing{}
	case strings.Contains(link, initYoutube):
		return &youtube.Youtube{}
	case strings.Contains(link, initSoundCloud):
		return &soundcloud.SoundCloud{}
	case strings.Contains(link, initChiaSeNhac):
		return &chiasenhac.ChiaSeNhac{}
	case strings.Contains(link, initFacebook):
		return &facebook.Facebook{}
	case strings.Contains(link, initVimeo):
		return &vimeo.Vimeo{}
	}

	return nil
}

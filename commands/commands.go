package commands

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	initYoutube    string = "youtube"
	initSoundCloud string = "soundcloud"
	initFacebook   string = "facebook"
)

// Flag descriptions
var (
	Link   string
	Output string
)

// ObjectResponse stores the response body associate with a name of item
type ObjectResponse struct {
	Resp *http.Response
	Name string
}

// New ...
func New() *cobra.Command {
	cobra.OnInitialize()

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "songgrabber",
		Short: "CLI downloader",
		Long:  `A small cli written in Go to help download music/video from multiple sources.`,
	}

	rootCmd.AddCommand(initDownloadCommand())
	rootCmd.AddCommand(initPlayCommand())
	rootCmd.AddCommand(initVersionCommand())

	return rootCmd
}

func initVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of songgrabber",
		Long:  `All software has versions. This is just songgrabber's.`,
		Run: func(cmd *cobra.Command, args []string) {
			runVersion()
			os.Exit(0)
		},
	}

}

func initDownloadCommand() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "download command",
		Run: func(cmd *cobra.Command, args []string) {
			runDownload()
			os.Exit(0)
		},
	}

	downloadCmd.Flags().StringVarP(&Link, "link", "l", "", "Song, playlist link that want to download")
	downloadCmd.Flags().StringVarP(&Output, "output", "o", "", "The output directory")
	viper.BindPFlag("link", downloadCmd.Flags().Lookup("link"))
	viper.BindPFlag("output", downloadCmd.Flags().Lookup("output"))

	return downloadCmd
}

func initPlayCommand() *cobra.Command {
	playCmd := &cobra.Command{
		Use:   "play",
		Short: "play command",
		Run: func(cmd *cobra.Command, args []string) {
			runPlay()
			os.Exit(0)
		},
	}

	playCmd.Flags().StringVarP(&Link, "link", "l", "", "Song, playlist link that want to download")

	return playCmd
}

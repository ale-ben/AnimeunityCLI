package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"AnimeunityCLI/packages/downloadurl"
	"AnimeunityCLI/packages/getinfo"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "main")
	//File wide logger
	mainLog     = Log.WithField("File", "AnimeunityCLI.go")
	interactive = false
	keyword     = ""
	inputURL    = ""
	season      = ""
	logLevel    = ""
	version     = "v1.0"
)

func main() {
	app := &cli.App{
		Name:                 "Animeunity Unofficial Utility",
		Usage:                "Query Animeunity and get download links",
		Version:              version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "logLevel",
				Usage:       "Can be Trace, Debug, Info, Warn, Error",
				Value:       "Warn",
				Destination: &logLevel,
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "quickDownload",
				Category: "interactive",
				Aliases:  []string{"qd"},
				Usage:    "Search for a keyword and choose from a list to get download links",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "keyword",
						Aliases:     []string{"k"},
						Usage:       "Keyword to look for",
						Destination: &keyword,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "season",
						Usage:       "Season(s) to download, OVA, NOOVA, ALL, NO (Downloads only the season you pass as URL)",
						Value:       "NO",
						Destination: &season,
					},
				},
				Action: func(c *cli.Context) error {
					setLogLevel()
					log.WithFields(logrus.Fields{
						"command": "quickDownload",
						"keyword": keyword,
						"season":  season,
					}).Trace("Program Started")
					return quickDownload()
				},
			},
			{
				Name:     "search",
				Category: "batch",
				Aliases:  []string{"s"},
				Usage:    "Search for a keyword and displays a list of anime",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "keyword",
						Aliases:     []string{"k"},
						Usage:       "Keyword to look for",
						Destination: &keyword,
						Required:    true,
					},
				},
				Action: func(c *cli.Context) error {
					setLogLevel()
					log.WithFields(logrus.Fields{
						"command": "search",
						"keyword": keyword,
					}).Trace("Program Started")
					return getInfo()
				},
			},
			{
				Name:     "downloadURL",
				Category: "batch",
				Aliases:  []string{"d"},
				Usage:    "Prints the download links for an anime season",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "url",
						Aliases:     []string{"u"},
						Usage:       "Url of the anime page on AnimeUnity",
						Destination: &inputURL,
						Required:    true,
					},
					&cli.StringFlag{
						Name:        "season",
						Usage:       "Season(s) to download, OVA, NOOVA, ALL, NO (Downloads only the season you pass as URL)",
						Value:       "NO",
						Destination: &season,
					},
				},
				Action: func(c *cli.Context) error {
					setLogLevel()
					log.WithFields(logrus.Fields{
						"command":  "downloadURL",
						"inputURL": inputURL,
					}).Trace("Program Started")
					return getDownload()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setLogLevel() {
	switch strings.ToLower(logLevel) {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
		break
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		break
	case "info":
		log.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		log.SetLevel(logrus.WarnLevel)
		break
	case "error":
		log.SetLevel(logrus.ErrorLevel)
		break
	}
	mainLog.WithField("Level", strings.ToLower(logLevel)).Debug("Log Level Set")
	getinfo.SetLogLevel(logLevel)
	downloadurl.SetLogLevel(logLevel)
}

func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getInfo() error {
	getinfo.PrintFullAnimeList(getinfo.GetInfo(keyword))
	return nil
}

func getDownload() error {
	animePage := downloadurl.AnimePageStruct{"", inputURL, "", []string{}, false}
	animePageList := downloadurl.DownloadURL(animePage, season)
	downloadurl.PrintURLList(animePageList)
	return nil
}

func quickDownload() error {
	animeList := getinfo.GetInfo(keyword)
	getinfo.PrintSmallAnimeList(animeList)
	if len(animeList) == 0 {
		fmt.Println("No anime found, try changing the keyword")
		os.Exit(0)
	}
	key := -1
	for key == -1 {
		fmt.Printf("\n<- ID: ")
		id := scanInput()
		for k, anime := range animeList {
			if anime.AnimeID == id {
				key = k
				break
			}
		}
		if key == -1 {
			fmt.Println("Id not found")
		}
	}
	fmt.Printf("\nAnime Found :)\n")
	getinfo.PrintFullAnime(animeList[key])
	fmt.Printf("\nLooking for episodes\n")
	animePage := downloadurl.AnimePageStruct{animeList[key].AnimeID, inputURL, animeList[key].Titolo, []string{}, false}
	animePageList := downloadurl.DownloadURL(animePage, season)
	downloadurl.PrintURLList(animePageList)
	return nil
}
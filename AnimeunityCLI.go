package main

import (
	"AnimeunityCLI/packages/jdownloader"
	"AnimeunityCLI/packages/scraper"
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"AnimeunityCLI/packages/commonresources"
	"AnimeunityCLI/packages/downloadurl"
	"AnimeunityCLI/packages/getinfo"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "main")
	//File wide logger
	mainLog      = Log.WithField("File", "AnimeunityCLI.go")
	interactive  = false
	keyword      = ""
	inputURL     = ""
	season       = ""
	logLevel     = ""
	group        = true
	crawlPath    = ""
	downloadPath = ""
	version      = "v1.0"
	noPrintInfo  = false
)

//TODO Test Files
//TODO Comment
//TODO GoDoc

func main() {

	// ---- CLI ----

	// -- Flags --
	searchFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "keyword",
			Aliases:     []string{"k"},
			Usage:       "Keyword to look for",
			Destination: &keyword,
			Required:    true,
		},
	}

	downloadFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "url",
			Aliases:     []string{"u"},
			Usage:       "Url of the anime page on AnimeUnity",
			Destination: &inputURL,
			Required:    true,
		},
	}

	seasonFlag := &cli.StringFlag{
		Name:        "season",
		Usage:       "Season(s) to download, OVA, NOOVA, ALL, NO (Downloads only the season you pass as URL)",
		Value:       "NO",
		Destination: &season,
	}

	groupPrintFlag := &cli.BoolFlag{
		Name:        "group",
		Aliases:     []string{"g"},
		Usage:       "Boolean value determines if episode links are going to be grouped or not",
		Destination: &group,
		Value:       true,
	}

	jdownloaderFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "crawlpath",
			Usage:       "Specify the path for the CrawlJobs. IMPORTANT you must provide a value for both crawlpath and jdownloadpath for this to work",
			Destination: &crawlPath,
			Value:       "",
		},
		&cli.StringFlag{
			Name:        "jdownloadpath",
			Aliases:     []string{"jdp"},
			Usage:       "Specify the path where JDownloader should create the subdirectories. IMPORTANT you must provide a value for both crawlpath and jdownloadpath for this to work",
			Destination: &downloadPath,
			Value:       "",
		},
		&cli.BoolFlag{
			Name:        "noPrint",
			Usage:       "Should the program NOT print download info/url(s)",
			Destination: &noPrintInfo,
			Value:       false,
			Required:    false,
		},
	}

	// -- App --
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
				Flags:    append(searchFlags, append(jdownloaderFlags, seasonFlag, groupPrintFlag)...),
				Action: func(c *cli.Context) error {
					setGlobalLogLevel()
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
				Aliases:  []string{"se"},
				Usage:    "Search for a keyword and displays a list of anime",
				Flags:    searchFlags,
				Action: func(c *cli.Context) error {
					setGlobalLogLevel()
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
				Aliases:  []string{"dl"},
				Usage:    "Prints the download links for an anime season",
				Flags:    append(downloadFlags, append(jdownloaderFlags, seasonFlag, groupPrintFlag)...),
				Action: func(c *cli.Context) error {
					setGlobalLogLevel()
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

//scanInput returns the next line from the keyboard
func scanInput() string {
	mainLog.Trace("<scanInput/>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getInfo() error {
	mainLog.Trace("<getInfo/>")
	commonresources.PrintAnimeList(getinfo.GetInfo(keyword), 1)
	return nil
}

func getDownload() error {
	mainLog.Trace("<getDownload>")
	animePage := commonresources.AnimePageStruct{AnimeURL: inputURL, EpisodeList: []string{}} //Create an empty animePage only with the given URL
	mainLog.Debug("Launching link getter")
	animePageList := downloadurl.DownloadURL(animePage, season)
	if !noPrintInfo {
		mainLog.Debug("Printing URLs")
		commonresources.PrintURLList(animePageList, group)
	}
	if crawlPath != "" && downloadPath != "" {
		mainLog.Debug("Creating Crawlpaths")
		jdownloader.SendToJDownloader(animePageList, crawlPath, downloadPath)
	}
	mainLog.Trace("</getDownload>")
	return nil
}

func quickDownload() error {
	mainLog.Trace("<quickDownload>")
	mainLog.Debug("Getting Info")
	animeList := getinfo.GetInfo(keyword)
	mainLog.Debug("Printing Anime List")
	commonresources.PrintAnimeList(animeList, 0)
	if len(animeList) == 0 {
		fmt.Println("No anime found, try changing the keyword")
		os.Exit(0)
	}
	key := -1
	mainLog.Debug("Printing Anime List")
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
	commonresources.PrintAnime(animeList[key], 1)
	fmt.Printf("\nLooking for episodes\n")
	animePage := commonresources.AnimePageStruct{AnimeID: animeList[key].AnimeID, AnimeURL: inputURL, Title: animeList[key].Title, EpisodeList: []string{}}
	mainLog.Debug("Launching link getter")
	animePageList := downloadurl.DownloadURL(animePage, season)
	if !noPrintInfo {
		mainLog.Debug("Printing URLs")
		commonresources.PrintURLList(animePageList, group)
	}
	if crawlPath != "" && downloadPath != "" {
		mainLog.Debug("Creating Crawlpaths")
		jdownloader.SendToJDownloader(animePageList, crawlPath, downloadPath)
	}

	mainLog.Trace("</quickDownload>")
	return nil
}

func setLogLevel() {
	commonresources.SetLogLevel(log, logLevel, "main.go")
}

func setGlobalLogLevel() {
	mainLog.Trace("<setGlobalLogLevel/>")
	setLogLevel()
	getinfo.SetLogLevel(logLevel)
	downloadurl.SetLogLevel(logLevel)
	commonresources.SetOwnLogLevel(logLevel)
	jdownloader.SetLogLevel(logLevel)
	jdownloader.SetLogLevel(logLevel)
	scraper.SetLogLevel(logLevel)
}

package models

type Searx struct {
	Image   SearxImage   `json:"image"`
	It      SearxIt      `json:"it"`
	News    SearxNews    `json:"news"`
	Science SearxScience `json:"science"`
	Social  SearxSocial  `json:"social"`
	Videos  SearxVideos  `json:"videos"`
	Web     SearxWeb     `json:"web"`
}

type SearxImage struct {
	Answers             []string            `json:"answers"`
	Corrections         []string            `json:"corrections"`
	Infoboxes           []SearxImageInfobox `json:"infoboxes"`
	NumberOfResults     int                 `json:"number_of_results"`
	Query               string              `json:"query"`
	Results             []SearxImageResult  `json:"results"`
	Suggestions         []interface{}       `json:"suggestions"`
	UnresponsiveEngines [][]string          `json:"unresponsive_engines"`
}

type SearxImageInfobox struct {
	Attributes []SearxImageInfoboxAttribute `json:"attributes"`
	Content    string                       `json:"content"`
	Engine     string                       `json:"engine"`
	ID         string                       `json:"id"`
	ImgSrc     string                       `json:"img_src"`
	Infobox    string                       `json:"infobox"`
	Urls       []SearxImageInfoboxURL       `json:"urls"`
}

type SearxImageInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxImageInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxImageResult struct {
	Category     string   `json:"category"`
	Content      string   `json:"content"`
	Engine       string   `json:"engine"`
	Engines      []string `json:"engines"`
	ImgFormat    string   `json:"img_format"`
	ImgSrc       string   `json:"img_src"`
	ParsedURL    []string `json:"parsed_url"`
	Positions    []int    `json:"positions"`
	PrettyURL    string   `json:"pretty_url"`
	Score        float64  `json:"score"`
	Source       string   `json:"source"`
	Template     string   `json:"template"`
	ThumbnailSrc string   `json:"thumbnail_src"`
	Title        string   `json:"title"`
	URL          string   `json:"url"`
}

/*-----------------------------------------------*/
/*---------- JSON - IT results ------------------*/
/*-----------------------------------------------*/

type SearxIt struct {
	Answers             []string         `json:"answers"`
	Corrections         []string         `json:"corrections"`
	Infoboxes           []SearxItInfobox `json:"infoboxes"`
	NumberOfResults     int              `json:"number_of_results"`
	Query               string           `json:"query"`
	Results             []SearxItResult  `json:"results"`
	Suggestions         []string         `json:"suggestions"`
	UnresponsiveEngines [][]string       `json:"unresponsive_engines"`
}

type SearxItInfobox struct {
	Attributes []SearxItInfoboxAttribute `json:"attributes"`
	Content    string                    `json:"content"`
	Engine     string                    `json:"engine"`
	ID         string                    `json:"id"`
	ImgSrc     string                    `json:"img_src"`
	Infobox    string                    `json:"infobox"`
	Urls       []SearxItInfoboxURL       `json:"urls"`
}

type SearxItInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxItInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxItResult struct {
	Category     string          `json:"category"`
	CodeLanguage string          `json:"code_language"`
	Codelines    [][]interface{} `json:"codelines"`
	Content      string          `json:"content"`
	Engine       string          `json:"engine"`
	Engines      []string        `json:"engines"`
	ParsedURL    []string        `json:"parsed_url"`
	Positions    []int           `json:"positions"`
	PrettyURL    string          `json:"pretty_url"`
	Repository   string          `json:"repository"`
	Score        float64         `json:"score"`
	Template     string          `json:"template"`
	Title        string          `json:"title"`
	URL          string          `json:"url"`
}

/*-----------------------------------------------*/
/*--------- JSON - News results -----------------*/
/*-----------------------------------------------*/

type SearxNews struct {
	Answers             []string           `json:"answers"`
	Corrections         []string           `json:"corrections"`
	Infoboxes           []SearxNewsInfobox `json:"infoboxes"`
	NumberOfResults     int                `json:"number_of_results"`
	Query               string             `json:"query"`
	Results             []SearxNewsResult  `json:"results"`
	Suggestions         []interface{}      `json:"suggestions"`
	UnresponsiveEngines []interface{}      `json:"unresponsive_engines"`
}

type SearxNewsInfobox struct {
	Attributes []SearxNewsInfoboxAttribute `json:"attributes"`
	Content    string                      `json:"content"`
	Engine     string                      `json:"engine"`
	ID         string                      `json:"id"`
	ImgSrc     string                      `json:"img_src"`
	Infobox    string                      `json:"infobox"`
	Urls       []SearxNewsInfoboxURL       `json:"urls"`
}

type SearxNewsInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxNewsInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxNewsResult struct {
	Category      string      `json:"category"`
	Content       string      `json:"content"`
	Engine        string      `json:"engine"`
	Engines       []string    `json:"engines"`
	ImgSrc        interface{} `json:"img_src"`
	ParsedURL     []string    `json:"parsed_url"`
	Positions     []int       `json:"positions"`
	PrettyURL     string      `json:"pretty_url"`
	Pubdate       string      `json:"pubdate"`
	PublishedDate string      `json:"publishedDate"`
	Score         float64     `json:"score"`
	Title         string      `json:"title"`
	URL           string      `json:"url"`
}

/*-----------------------------------------------*/
/*----------- JSON - Science results ------------*/
/*-----------------------------------------------*/

type SearxScience struct {
	Answers             []string              `json:"answers"`
	Corrections         []string              `json:"corrections"`
	Infoboxes           []SearxScienceInfobox `json:"infoboxes"`
	NumberOfResults     int                   `json:"number_of_results"`
	Query               string                `json:"query"`
	Results             []SearxScienceResult  `json:"results"`
	Suggestions         []interface{}         `json:"suggestions"`
	UnresponsiveEngines [][]string            `json:"unresponsive_engines"`
}

type SearxScienceInfobox struct {
	Attributes []SearxScienceInfoboxAttribute `json:"attributes"`
	Content    string                         `json:"content"`
	Engine     string                         `json:"engine"`
	ID         string                         `json:"id"`
	ImgSrc     string                         `json:"img_src"`
	Infobox    string                         `json:"infobox"`
	Urls       []SearxScienceInfoboxURL       `json:"urls"`
}

type SearxScienceInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxScienceInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxScienceResult struct {
	Category      string   `json:"category"`
	Content       string   `json:"content"`
	Engine        string   `json:"engine"`
	Engines       []string `json:"engines"`
	ParsedURL     []string `json:"parsed_url"`
	Positions     []int    `json:"positions"`
	PrettyURL     string   `json:"pretty_url"`
	Pubdate       string   `json:"pubdate"`
	PublishedDate string   `json:"publishedDate"`
	Score         float64  `json:"score"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
}

/*-----------------------------------------------*/
/*------- JSON - Social results -----------------*/
/*-----------------------------------------------*/

type SearxSocial struct {
	Answers             []string             `json:"answers"`
	Corrections         []string             `json:"corrections"`
	Infoboxes           []SearxSocialInfobox `json:"infoboxes"`
	NumberOfResults     int                  `json:"number_of_results"`
	Query               string               `json:"query"`
	Results             []interface{}        `json:"results"`
	Suggestions         []interface{}        `json:"suggestions"`
	UnresponsiveEngines [][]string           `json:"unresponsive_engines"`
}

type SearxSocialInfobox struct {
	Attributes []SearxSocialInfoboxAttribute `json:"attributes"`
	Content    string                        `json:"content"`
	Engine     string                        `json:"engine"`
	ID         string                        `json:"id"`
	ImgSrc     string                        `json:"img_src"`
	Infobox    string                        `json:"infobox"`
	Urls       []SearxSocialInfoboxURL       `json:"urls"`
}

type SearxSocialInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxSocialInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxSocialResult struct {
	Author        string   `json:"author"`
	Category      string   `json:"category"`
	Content       string   `json:"content"`
	Embedded      string   `json:"embedded"`
	Engine        string   `json:"engine"`
	Engines       []string `json:"engines"`
	Length        string   `json:"length"`
	ParsedURL     []string `json:"parsed_url"`
	Positions     []int    `json:"positions"`
	PrettyURL     string   `json:"pretty_url"`
	Pubdate       string   `json:"pubdate"`
	PublishedDate string   `json:"publishedDate"`
	Score         float64  `json:"score"`
	Template      string   `json:"template"`
	Thumbnail     string   `json:"thumbnail"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
}

/*-----------------------------------------------*/
/*------- JSON -  Video results -----------------*/
/*-----------------------------------------------*/

type SearxVideos struct {
	Answers             []string             `json:"answers"`
	Corrections         []string             `json:"corrections"`
	Infoboxes           []SearxVideosInfobox `json:"infoboxes"`
	NumberOfResults     int                  `json:"number_of_results"`
	Query               string               `json:"query"`
	Results             []SearxVideosResult  `json:"results"`
	Suggestions         []interface{}        `json:"suggestions"`
	UnresponsiveEngines [][]string           `json:"unresponsive_engines"`
}

type SearxVideosInfobox struct {
	Attributes []SearxVideosInfoboxAttribute `json:"attributes"`
	Content    string                        `json:"content"`
	Engine     string                        `json:"engine"`
	ID         string                        `json:"id"`
	ImgSrc     string                        `json:"img_src"`
	Infobox    string                        `json:"infobox"`
	Urls       []SearxVideosInfoboxURL       `json:"urls"`
}

type SearxVideosInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxVideosInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     int      `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxVideosResult struct {
	Author        string   `json:"author"`
	Category      string   `json:"category"`
	Content       string   `json:"content"`
	Embedded      string   `json:"embedded"`
	Engine        string   `json:"engine"`
	Engines       []string `json:"engines"`
	Length        string   `json:"length"`
	ParsedURL     []string `json:"parsed_url"`
	Positions     []int    `json:"positions"`
	PrettyURL     string   `json:"pretty_url"`
	Pubdate       string   `json:"pubdate"`
	PublishedDate string   `json:"publishedDate"`
	Score         float64  `json:"score"`
	Template      string   `json:"template"`
	Thumbnail     string   `json:"thumbnail"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
}

/*-----------------------------------------------*/
/*-------- JSON - Web results -------------------*/
/*-----------------------------------------------*/

type SearxWeb struct {
	Answers             []string          `json:"answers"`
	Corrections         []string          `json:"corrections"`
	Infoboxes           []SearxWebInfobox `json:"infoboxes"`
	NumberOfResults     float64           `json:"number_of_results"`
	Query               string            `json:"query"`
	Results             []SearxWebResult  `json:"results"`
	Suggestions         []string          `json:"suggestions"`
	UnresponsiveEngines [][]string        `json:"unresponsive_engines"`
}

type SearxWebInfobox struct {
	Attributes []SearxWebInfoboxAttribute `json:"attributes"`
	Content    string                     `json:"content"`
	Engine     string                     `json:"engine"`
	ID         string                     `json:"id"`
	ImgSrc     string                     `json:"img_src"`
	Infobox    string                     `json:"infobox"`
	Urls       []SearxWebInfoboxURL       `json:"urls"`
}

type SearxWebInfoboxAttribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SearxWebInfoboxURL struct {
	Category  string   `json:"category"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     float64  `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

type SearxWebResult struct {
	Category  string   `json:"category"`
	Content   string   `json:"content"`
	Domain    string   `json:"domain"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Official  bool     `json:"official"`
	ParsedURL []string `json:"parsed_url"`
	Positions []int    `json:"positions"`
	PrettyURL string   `json:"pretty_url"`
	Score     float64  `json:"score"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
}

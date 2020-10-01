package omdb

const EndPoint = "http://www.omdbapi.com/"

type Response struct {
	Response string
	Error    string
	Title    string
	Poster   string
}

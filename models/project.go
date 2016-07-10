package models

type Project struct {
	Id     string   `gorethink:"id" json:"id"`
	Public bool     `gorethink:"public" json:"public"`
	Name   string   `gorethink:"name" json:"name"`
	Owners []string `gorethink:"owner" json:"owners"`
}

type ProjectContentTemp struct {
	Id      string `gorethink:"id" json "id"`
	Content string `gorethink:"content" json:"content"`
}

type ProjectContent struct {
	Id     string   `gorethink:"id" json:"id"`
	Tracks []Track  `gorethink:"tracks" json:"tracks"`
	Tempo  int      `gorethink:"tempo" json:"tempo"`
}

type Track struct {
	Id        	string `gorethink:"id" json:"id"`
	Name				string `gorethink:"name" json:"id"`
	Volume    	int8   `gorethink:"volume" json:"volume"`
	TrackType 	string `gorethink:"trackType" json:"tracktype"`
	Solo      	bool   `gorethink:"solo" json:"solo"`
	Mute      	bool   `gorethink:"mute" json:"mute"`
	Instrument  string `gorethink:"instrument" json:"instrument"`
	TrackNodes	[]TrackNode `gorethink:"tracknodes" json:"tracknodes"`
}

type TrackNode struct {
	StartTime   int8   `gorethink:"starttime" json:"starttime"`
	Duration    int8   `gorethink:"duration" json:"duration"`
	SnippetId   string `gorethink:"snippetid" json:"snippetid"`
	SnippetName string `gorethink:"snippetname" json:"snippetname"`
}

type Snippet struct {
	Id     string `gorethink:"id" json:"id"`
	Public bool   `gorethink:"public" json:"public"`
	Name   string `gorethink:"name" json:"name"`
	Owner  string `gorethink:"owner" json:"owner"`
	Duration int8 `gorethink:"duration" json:"duration"`
}

type SnippetContent struct {
	Id        string   `gorethink:"id" json:"id"`
	SoundFile string   `gorethink:"soundfile" json:"soundfile"`
	Notes     [][]Note `gorethink:"notes" json:"notes"`
}

type Note struct {
	Pitch     string `json:"pitch"`
	Duration  string `json:"duration"`
	StartTime int    `json:"starttime"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewDefaultProjectContent() ProjectContent {
	projectContent := ProjectContent{
		Tracks: []Track{},
		Tempo:  60,
	}
	return projectContent
}

func NewDefaultProject(uid string) Project {
	project := Project{
		Public: true,
		Name:   "untitled",
		Owners: []string{uid},
	}
	return project
}

func NewDefaultSnippet(uid string) Snippet {
	snippet := Snippet{
		Public: true,
		Name:   "untitled",
		Owner:  uid,
		Duration: 0,
	}
	return snippet
}

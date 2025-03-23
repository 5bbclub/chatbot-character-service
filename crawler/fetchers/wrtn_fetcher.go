package fetchers

import (
	"encoding/json"
	"fmt"
	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	_ Fetcher = (*WrtnFetcher)(nil)
)

type WrtnFetcher struct {
	BaseFetcher
	conf *config.Config
}

func NewWrtnFetcher(conf *config.Config) *WrtnFetcher {
	interval := time.Duration(conf.Services[1].Interval) * time.Second
	return &WrtnFetcher{
		BaseFetcher: BaseFetcher{
			Name:     "WrtnFetcher",
			Interval: interval,
		},
		conf: conf,
	}
}

type Icon struct {
	Dark  string `json:"dark"`
	Light string `json:"light"`
}

type Creator struct {
	Nickname           string `json:"nickname"`
	WrtnUid            string `json:"wrtnUid"`
	IsCertifiedCreator bool   `json:"isCertifiedCreator"`
	ProfileID          string `json:"profileId"`
}

type Category struct {
	ID                   string `json:"_id"`
	Name                 string `json:"name"`
	RecommendDescription string `json:"recommendDescription"`
}

type SuperChatModel struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Icon  Icon   `json:"icon"`
}

type ProfileImage struct {
	Origin string `json:"origin"`
	W200   string `json:"w200"`
	W600   string `json:"w600"`
}

type BadgeContent struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Badge struct {
	Slug    string       `json:"slug"`
	Name    string       `json:"name"`
	Content BadgeContent `json:"content"`
	Type    string       `json:"type"`
}

type PromptTemplate struct {
	Name     string `json:"name"`
	Template string `json:"template"`
	Icon     Icon   `json:"icon"`
}

type WrtnAPIResponse struct {
	Result      string `json:"result"`
	WrtnAPIData struct {
		Characters []WrtnCharacterDetail `json:"characters"`
		NextCursor string                `json:"nextCursor"`
	} `json:"Data"`
}

type WrtnCharacterDetail struct {
	ID                    string         `json:"_id"`
	UserID                string         `json:"userId"`
	WrtnUid               string         `json:"wrtnUid"`
	InitialMessages       []string       `json:"initialMessages"`
	Creator               Creator        `json:"creator"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Categories            []Category     `json:"categories"`
	DefaultSuperChatModel SuperChatModel `json:"defaultSuperChatModel"`
	ChatCount             int            `json:"chatCount"`
	ChatUserCount         int            `json:"chatUserCount"`
	LikeCount             int            `json:"likeCount"`
	ImageCount            int            `json:"imageCount"`
	Tags                  []string       `json:"tags"`
	HasImage              bool           `json:"hasImage"`
	IsLiked               bool           `json:"isLiked"`
	IsDisliked            bool           `json:"isDisliked"`
	CountryCode           string         `json:"countryCode"`
	Status                string         `json:"status"`
	Visibility            string         `json:"visibility"`
	ProfileImage          ProfileImage   `json:"profileImage"`
	ReplySuggestions      []string       `json:"replySuggestions"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	IsAdult               bool           `json:"isAdult"`
	IsConvertedToAdult    bool           `json:"isConvertedToAdult"`
	CommentCount          int            `json:"commentCount"`
	PromptTemplate        PromptTemplate `json:"promptTemplate"`
	Badges                []Badge        `json:"badges"`
}

func (w *WrtnFetcher) GetServiceName() string {
	return "wrtn"
}

func (w *WrtnFetcher) Start() {
	ticker := time.NewTicker(w.Interval)
	log.Printf("[%s] Fetcher started", w.Name)
	for {
		select {
		case <-ticker.C:
			dataBytes, err := w.Fetch()
			if err != nil {
				log.Printf("[%s] Failed to fetch data: %v", w.Name, err)
				continue
			}

			var wrtnAPIResponse WrtnAPIResponse
			var characters []WrtnCharacterDetail
			err = json.Unmarshal(dataBytes, &wrtnAPIResponse)
			if err != nil {
				fmt.Errorf("failed to parse JSON: %w", err)
				continue
			}
			if wrtnAPIResponse.Result != "SUCCESS" {
				log.Printf("Failed to fetch data: %v", wrtnAPIResponse)
				continue
			}
			if wrtnAPIResponse.WrtnAPIData.Characters == nil {
				log.Printf("No characters found in the response")
				continue
			}
			characters = wrtnAPIResponse.WrtnAPIData.Characters

			for _, character := range characters {
				w.OutputChannel <- character
			}
			w.OutputChannel <- dataBytes
		}
	}
}

func (w *WrtnFetcher) SetOutputChannel(c chan interface{}) {
	w.OutputChannel = c
}

func (w *WrtnFetcher) Fetch() ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}

	// HTTP 요청 보내기
	//FIXME: services 이름으로 config 인자를 받기
	resp, err := client.Get(w.conf.Services[1].Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Babechat: %w", err)
	}
	defer resp.Body.Close()

	// 응답 상태 확인
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// JSON 데이터 읽기 및 파싱
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

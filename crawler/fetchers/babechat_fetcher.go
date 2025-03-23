// crawler/fetchers/babechat_fetcher.go
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
	_ Fetcher = (*BabechatFetcher)(nil)
)

// BabechatFetcher는 Babechat 데이터를 크롤링하기 위한 구조체입니다.
type BabechatFetcher struct {
	BaseFetcher
	conf *config.Config
}

// NewBabechatFetcher는 BabechatFetcher 인스턴스를 생성합니다.
func NewBabechatFetcher(conf *config.Config) *BabechatFetcher {
	interval := time.Duration(conf.Services[0].Interval) * time.Second
	return &BabechatFetcher{
		BaseFetcher: BaseFetcher{
			Name:     "BabeChatFetcher",
			Interval: interval,
		},
		conf: conf,
	}
}

// BabechatCharacter는 Babechat API에서 반환되는 데이터 구조입니다.
type BabechatCharacter struct {
	ID              string   `json:"id"`
	CharacterID     string   `json:"characterId"`
	ChatCount       int      `json:"chatCount"`
	LikeCount       int      `json:"likeCount"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	CreatorNickname string   `json:"creatorNickname"`
	MainImage       string   `json:"mainImage"`
	Tags            []string `json:"tags"`
}

// GetServiceName은 서비스의 이름을 반환합니다.
func (b *BabechatFetcher) GetServiceName() string {
	return "babechat"
}

// Start: 데이터를 주기적으로 가져오는 프로세스 시작
func (f *BabechatFetcher) Start() {
	ticker := time.NewTicker(f.Interval)
	log.Printf("[%s] Fetcher started", f.Name)
	for {
		select {
		case <-ticker.C:
			dataBytes, err := f.Fetch()
			if err != nil {
				log.Printf("[%s] Failed to fetch data: %v", f.Name, err)
				continue
			}

			var characters []BabechatCharacter
			err = json.Unmarshal(dataBytes, &characters)
			if err != nil {
				fmt.Errorf("failed to parse JSON: %w", err)
			}

			for _, character := range characters {
				f.OutputChannel <- character
			}

			//log.Printf("[%s] Fetched data: %v", f.Name, data)
		}
	}
}

func (b *BabechatFetcher) SetOutputChannel(channel chan interface{}) {
	b.OutputChannel = channel
}

// FetchData는 Babechat API에서 데이터를 가져옵니다.
func (b *BabechatFetcher) Fetch() ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}

	// HTTP 요청 보내기
	//FIXME: services 이름으로 config 인자를 받기
	resp, err := client.Get(b.conf.Services[0].Endpoint)
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

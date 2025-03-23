// crawler/fetchers/babechat_fetcher.go
package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BabechatFetcher는 Babechat 데이터를 크롤링하기 위한 구조체입니다.
type BabechatFetcher struct {
	APIEndpoint string
}

// NewBabechatFetcher는 BabechatFetcher 인스턴스를 생성합니다.
func NewBabechatFetcher(apiEndpoint string) *BabechatFetcher {
	return &BabechatFetcher{
		APIEndpoint: apiEndpoint,
	}
}

// BabechatCharacter는 Babechat API에서 반환되는 데이터 구조입니다.
type BabechatCharacter struct {
	ID              string   `json:"id"`
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

// FetchData는 Babechat API에서 데이터를 가져옵니다.
func (b *BabechatFetcher) FetchData() ([]interface{}, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// HTTP 요청 보내기
	resp, err := client.Get(b.APIEndpoint)
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

	var characters []BabechatCharacter
	err = json.Unmarshal(body, &characters)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// 인터페이스 타입으로 변환
	var data []interface{}
	for _, character := range characters {
		data = append(data, character)
	}

	return data, nil
}

package processors

import (
	"fmt"
	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"github.com/5bbclub/chatbot-character-service/crawler/fetchers"
	"github.com/5bbclub/chatbot-character-service/models"
	"gorm.io/gorm"
	"log"
)

var (
	_ Processor = (*BabeChatProcessor)(nil)
)

type BabeChatProcessor struct {
	BaseProcessor
	conf *config.Config
	db   *gorm.DB
}

func NewBabeChatProcessor(conf *config.Config, db *gorm.DB) *BabeChatProcessor {
	return &BabeChatProcessor{
		BaseProcessor: BaseProcessor{
			Name:         "BabeChatProcessor",
			InputChannel: make(chan interface{}, 100),
		},
		conf: conf,
		db:   db,
	}
}

type BabeChatCreator struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CharacterInput struct {
	ServiceProviderName string          `json:"service_provider_name"` // 서비스 제공자 이름
	Name                string          `json:"name"`                  // 캐릭터 이름
	Description         string          `json:"description"`           // 캐릭터 설명
	ProfileImageURL     string          `json:"profile_image_url"`     // 캐릭터 프로필 이미지 URL
	TagNames            []string        `json:"tags"`                  // 태그 이름 목록
	Creator             BabeChatCreator `json:"creator"`               // 제작자 정보
}

func (b *BabeChatProcessor) GetServiceName() string {
	return "babechat"
}

func (b *BabeChatProcessor) Start() {
	log.Println("Processor: Waiting for data to process...")
	for data := range b.InputChannel {
		log.Println("Processor: Got data from channel, processing...")
		if err := b.Process(data); err != nil {
			log.Printf("Processor: Error processing data: %v\n", err)
		}
	}
}

func (b *BabeChatProcessor) SetInputChannel(channel chan interface{}) {
	b.InputChannel = channel
}

func (b *BabeChatProcessor) FormatData(data interface{}) (*CharacterInput, error) {
	character, ok := data.(fetchers.BabechatCharacter)
	if !ok {
		return nil, fmt.Errorf("failed to parse character data")
	}
	return &CharacterInput{
		ServiceProviderName: b.GetServiceName(),
		Name:                character.Name,
		Description:         character.Description,
		ProfileImageURL:     character.MainImage,
		TagNames:            character.Tags,
		Creator: BabeChatCreator{
			Name: character.CreatorNickname,
		},
	}, nil
}

func (b *BabeChatProcessor) Process(data interface{}) error {
	// 1. character 데이터를 CharacterInput으로 변환
	characterInput, err := b.FormatData(data)
	if err != nil {
		return fmt.Errorf("failed to parse character data: %v", err)
	}

	// 2. 트랜잭션 시작
	return b.db.Transaction(func(tx *gorm.DB) error {
		// ---------------------------
		// 1. 서비스 제공자 저장/조회
		// ---------------------------
		var serviceProvider models.ServiceProvider
		if err := tx.FirstOrCreate(&serviceProvider, models.ServiceProvider{
			Name: characterInput.ServiceProviderName,
		}).Error; err != nil {
			return fmt.Errorf("failed to find or create service provider: %v", err)
		}

		// ---------------------------
		// 2. 캐릭터 저장
		// ---------------------------
		character := models.Character{
			ServiceProviderID: serviceProvider.ID,
			Name:              characterInput.Name,
			Description:       characterInput.Description,
			ProfileImageURL:   characterInput.ProfileImageURL,
		}
		if err := tx.Create(&character).Error; err != nil {
			return fmt.Errorf("failed to create character: %v", err)
		}

		// ---------------------------
		// 3. 태그 저장 및 관계 설정
		// ---------------------------
		for _, tagName := range characterInput.TagNames {
			var tag models.Tag
			if err := tx.FirstOrCreate(&tag, models.Tag{
				Name: tagName,
			}).Error; err != nil {
				return fmt.Errorf("failed to find or create tag: %v", err)
			}

			// 캐릭터-태그 관계 저장
			if err := tx.Create(&models.CharacterTag{
				CharacterID: character.ID,
				TagID:       tag.ID,
			}).Error; err != nil {
				return fmt.Errorf("failed to create character-tag relation: %v", err)
			}
		}

		// ---------------------------
		// 4. 제작자 저장 및 관계 설정
		// ---------------------------
		if characterInput.Creator.Name == "" {
			return fmt.Errorf("creator name is required")
		}
		creatorData := characterInput.Creator
		var creator models.Creator
		if err := tx.FirstOrCreate(&creator, models.Creator{
			Name:  creatorData.Name,
			Email: creatorData.Email,
		}).Error; err != nil {
			return fmt.Errorf("failed to find or create creator: %v", err)
		}

		// 캐릭터-제작자 관계 저장
		if err := tx.Create(&models.CharacterCreator{
			CharacterID: character.ID,
			CreatorID:   creator.ID,
		}).Error; err != nil {
			return fmt.Errorf("failed to create character-creator relation: %v", err)
		}

		// 모든 작업 성공
		log.Println("Processor: Data processed successfully")
		return nil
	})
}

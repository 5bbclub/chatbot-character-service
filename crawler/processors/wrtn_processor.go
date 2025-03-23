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
	_ Processor = (*WrtnProcessor)(nil)
)

type WrtnProcessor struct {
	BaseProcessor
	conf *config.Config
	db   *gorm.DB
}

func NewWrtnProcessor(conf *config.Config, db *gorm.DB) *WrtnProcessor {
	return &WrtnProcessor{
		BaseProcessor: BaseProcessor{
			Name:         "WrtnProcessor",
			InputChannel: make(chan interface{}, 100),
		},
		conf: conf,
		db:   db,
	}
}

func (w *WrtnProcessor) GetServiceName() string {
	return "wrtn"
}

func (w *WrtnProcessor) SetInputChannel(c chan interface{}) {
	w.InputChannel = c
}

func (w *WrtnProcessor) Start() {
	log.Println("Processor: Waiting for data to process...")
	for data := range w.InputChannel {
		log.Println("Processor: Got data from channel, processing...")
		if err := w.Process(data); err != nil {
			log.Printf("Processor: Error processing data: %v\n", err)
		}
	}
}

func (w *WrtnProcessor) FormatData(data interface{}) (models.CharacterInput, error) {
	characterData, ok := data.(fetchers.WrtnCharacterDetail)
	if !ok {
		return models.CharacterInput{}, fmt.Errorf("failed to parse character data")
	}

	return models.CharacterInput{
		ServiceProviderName: w.GetServiceName(),
		InternalID:          characterData.ID,
		Name:                characterData.Name,
		Description:         characterData.Description,
		ProfileImageURL:     characterData.ProfileImage.Origin,
		TagNames:            characterData.Tags,
		Creator: models.Creator{
			Name: characterData.Creator.Nickname,
		},
	}, nil
}

func (w *WrtnProcessor) Process(data interface{}) error {
	log.Printf("[%s ]Processor: Processing data\n", w.GetServiceName())
	// 1. character 데이터를 CharacterInput으로 변환
	characterInput, err := w.FormatData(data)
	if err != nil {
		return fmt.Errorf("failed to parse character data: %v", err)
	}

	// 2. 트랜잭션 시작
	return w.db.Transaction(func(tx *gorm.DB) error {
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
			InternalID:        characterInput.InternalID,
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

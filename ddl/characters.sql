CREATE TABLE characters (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    service_provider_id BIGINT NOT NULL COMMENT '서비스 제공자 ID',
    name VARCHAR(255) NOT NULL COMMENT '캐릭터 이름',
    description TEXT COMMENT '캐릭터 설명',
    profile_image_url VARCHAR(255) COMMENT '캐릭터 프로필 이미지 URL',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 생성 날짜',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
    -- 외래 키 설정
    FOREIGN KEY (service_provider_id) REFERENCES service_providers(id) ON DELETE CASCADE
);
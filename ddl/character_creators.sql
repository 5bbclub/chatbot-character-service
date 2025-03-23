CREATE TABLE character_creators (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    character_id BIGINT NOT NULL COMMENT '캐릭터 ID',
    creator_id BIGINT NOT NULL COMMENT '제작자 ID',
    role VARCHAR(255) COMMENT '제작자 역할 (예: 스크립트 작성자, 디자이너)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 생성 날짜',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
    -- 외래 키 설정
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (creator_id) REFERENCES creators(id) ON DELETE CASCADE
);
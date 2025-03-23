CREATE TABLE character_tags (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    character_id BIGINT NOT NULL COMMENT '캐릭터 ID',
    tag_id BIGINT NOT NULL COMMENT '태그 ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 생성 날짜',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
    -- 외래 키 설정
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
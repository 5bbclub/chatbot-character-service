CREATE TABLE tags (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE COMMENT '태그 이름 (예: "가족형", "독창적")',
    description TEXT COMMENT '태그에 대한 설명',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '태그 생성 날짜'
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
);
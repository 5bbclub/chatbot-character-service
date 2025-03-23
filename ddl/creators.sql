CREATE TABLE creators (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '제작자 이름',
    email VARCHAR(255) NOT NULL UNIQUE COMMENT '제작자 이메일',
    linked_in_url VARCHAR(255) COMMENT 'LinkedIn 프로필 URL',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 생성 날짜'
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
);
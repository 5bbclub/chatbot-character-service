CREATE TABLE service_providers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE COMMENT '서비스 제공자 이름',
    description TEXT COMMENT '서비스 제공자에 대한 설명',
    website_url VARCHAR(255) COMMENT '서비스 제공자 웹사이트 URL',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 생성 날짜'
    deleted_at TIMESTAMP DEFAULT NULL COMMENT '데이터 삭제 날짜',
);
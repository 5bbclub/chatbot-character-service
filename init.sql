-- init.sql
CREATE DATABASE IF NOT EXISTS dipping;

USE dipping;

-- 서비스 데이터 테이블 생성
CREATE TABLE IF NOT EXISTS service_data (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL COMMENT '서비스 이름',
    data JSON NOT NULL COMMENT '크롤링된 데이터(JSON 형식)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '데이터 저장 시간'
);

-- 사용자 계정 및 권한 부여
CREATE USER IF NOT EXISTS 'crawler_user'@'%' IDENTIFIED BY 'crawler_password';
GRANT ALL PRIVILEGES ON dipping.* TO 'crawler_user'@'%';
FLUSH PRIVILEGES;
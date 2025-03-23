# Chatbot Character Service

🎨 **Chatbot Character Service**는 캐릭터 데이터를 관리하고, 크롤링 및 API 서버를 통해 데이터를 제공하는 프로젝트입니다. 이 서비스는 효율적인 캐릭터 데이터 관리와 외부 데이터 통합 작업을 목표로 합니다.

---

## 📂 프로젝트 구조

```plaintext
chatbot-character-service/
├── api/                    # API 서버 관련 코드
│   ├── controllers/        # HTTP 요청 처리 컨트롤러
│   ├── middleware/         # HTTP 미들웨어
│   ├── routes/             # API 라우트 정의
│   └── server.go           # API 서버 초기화 및 실행 코드
├── crawler/                # 크롤링 서버 관련 코드
│   ├── fetchers/           # 서비스별 데이터 가져오기 로직
│   ├── processors/         # 데이터 처리 로직
│   ├── scheduler/          # 크롤링 작업 스케줄링
│   └── server.go           # 크롤링 서버 초기화 및 실행 코드
├── cmd/                    # 실행 가능한 애플리케이션 진입점
│   ├── api/                # API 서버 진입점
│   │   └── main.go         
│   └── crawler/            # 크롤링 서버 진입점
│       └── main.go       
├── config/                 # 설정 관련 코드
│   └── config.go           # 설정 로드 및 관리
├── internal/               # 패키지 내부에서만 사용하는 코드
│   ├── database/           # 데이터베이스 연결 및 마이그레이션
│   └── logger/             # 로깅 유틸리티
├── models/                 # 데이터 모델 정의
│   ├── character.go        # 캐릭터 모델
│   ├── tag.go              # 태그 모델
│   └── image.go            # 이미지 모델
├── repository/             # 데이터 접근 계층
│   └── character_repo.go   # 캐릭터 데이터 접근 로직
├── services/               # 비즈니스 로직
│   ├── character_service.go # 캐릭터 관련 비즈니스 로직
│   └── import_service.go   # 데이터 가져오기 로직
├── utils/                  # 유틸리티 함수
│   ├── errors.go           # 에러 처리
│   └── helpers.go          # 기타 유틸리티 함수
├── docker/                 # Docker 구성 파일
│   ├── api/                # API 서버 Docker 구성
│   │   └── Dockerfile
│   └── crawler/            # 크롤링 서버 Docker 구성
│       └── Dockerfile
├── docker-compose.yml      # 전체 서비스 Docker Compose 구성
├── go.mod                  # Go 모듈 정의
├── go.sum                  # Go 모듈 의존성 해시
├── .env.example            # 환경 변수 예제
└── README.md               # 프로젝트 문서
```

---

## 🚀 주요 기능

1. **API 서버**: 캐릭터 데이터를 HTTP API로 제공.
2. **크롤러 서버**: 외부 데이터를 주기적으로 크롤링하고 처리.
3. **TOML 설정 파일 지원**: 설정 파일을 기반으로 시스템 파라미터를 관리.
4. **유연한 비즈니스 로직**: 데이터 가져오기와 처리 비즈니스 로직 분리.
5. **Docker 지원**: API 서버 및 크롤러 서버를 Docker로 쉽게 배포.

---

## ⚙️ 설치 및 실행

### 1️⃣ **의존성 설치**
Go 모듈을 초기화하고 필요한 라이브러리를 설치합니다:

```bash
go mod tidy
```

### 2️⃣ **설정 파일 작성**
프로젝트 루트에 `config/config.toml` 파일을 작성합니다. 다음은 예제입니다:

```toml
[server]
api_port = 8080
crawler_port = 9090

[database]
host = "localhost"
port = 5432
user = "your_user"
password = "your_password"
name = "chatbot"
```

또는 `.env.example` 파일을 `.env`로 복사하여 환경 변수를 설정할 수도 있습니다.

### 3️⃣ **API 서버 실행**
다음 명령어로 API 서버를 실행합니다:

```bash
go run ./cmd/api/main.go
```

### 4️⃣ **크롤러 서버 실행**
다음 명령어로 크롤러 서버를 실행합니다:

```bash
go run ./cmd/crawler/main.go
```

---

## 🐳 Docker로 실행

### 1️⃣ **Docker 및 Docker Compose 설치**
Docker와 Docker Compose가 설치되어 있어야 합니다. [Docker 설치](https://docs.docker.com/get-docker/) 문서를 참고하세요.

### 2️⃣ **Docker 이미지 빌드**
Docker 이미지를 빌드합니다:

```bash
docker-compose build
```

### 3️⃣ **컨테이너 시작**
Docker Compose를 사용해 컨테이너를 실행합니다:

```bash
docker-compose up
```

API 서버와 크롤러 서버는 각각 `localhost:8080`과 `localhost:9090`에서 실행됩니다.

---

## 📒 API 문서

| HTTP 메서드 | 경로       | 설명                |
|-------------|------------|---------------------|
| `GET`       | `/health`  | 서버의 상태 확인    |
| `GET`       | `/character` | 캐릭터 정보 가져오기 |

---

## 🛠️ 주요 라이브러리

- **[go-toml](https://github.com/pelletier/go-toml)**: TOML 설정 파일 파싱.
- **[logrus](https://github.com/sirupsen/logrus)**: 로깅.
- **Docker**: 컨테이너화.
- **[gorilla/mux](https://github.com/gorilla/mux)**: 고급 라우팅.

---

## 📜 기여 방법

1. 이 저장소를 포크합니다.
2. 새로운 브랜치를 생성합니다 (`git checkout -b feature/새로운기능`).
3. 변경 사항을 커밋합니다 (`git commit -m 'Add 새로운 기능'`).
4. 브랜치를 푸시합니다 (`git push origin feature/새로운기능`).
5. Pull Request를 생성합니다.

---

## ©️ 라이센스

MIT 라이센스를 따릅니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 확인하세요.
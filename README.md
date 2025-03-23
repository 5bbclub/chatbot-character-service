Chatbot Character Database Service
This service is designed to fetch, process, and store character data from various chat services into a unified database. It provides APIs for importing characters, querying the database, and managing character data.
Features

Import character data from multiple chat services
Store characters in a PostgreSQL database
RESTful API for character management
Docker and Docker Compose setup for easy deployment
Support for character metadata, tags, and images

Architecture
The service is built with:

Go (Golang) - Core programming language
Gin - Web framework
GORM - ORM library for database operations
PostgreSQL - Database for storing character data
Docker - Containerization

Data Model

Character: Main entity representing a chatbot character
Tag: Character tags for categorization
Image: Character emotion images

API Endpoints
MethodEndpointDescriptionGET/healthHealth check endpointPOST/api/characters/importImport characters from JSONGET/api/charactersGet all charactersGET/api/characters/Get character by IDPOST/api/charactersCreate a new characterPUT/api/characters/Update an existing characterDELETE/api/characters/Delete a characterPOST/api/fetch/allTrigger fetching from all services
Getting Started
Prerequisites

Go 1.21 or later
PostgreSQL database
Docker and Docker Compose (optional)

Setup

Clone the repository:
bash복사git clone https://github.com/yourusername/chatbot-character-service.git
cd chatbot-character-service

Set up environment variables:
bash복사cp .env.example .env
# Edit .env with your configuration

Build and run:
bash복사go build -o chatbot-service
./chatbot-service


Using Docker

Build and run using Docker Compose:
bash복사docker-compose up -d

The service will be available at http://localhost:8080

Importing Characters
To import characters from a JSON file:
bash복사curl -X POST http://localhost:8080/api/characters/import?provider=babechat \
-H "Content-Type: application/json" \
-d @characters.json
Development
Adding a New Service
To add support for a new chat service:

Add a new fetch method in the FetchService struct
Update the FetchAll method to include the new service
Add any necessary API credentials to the .env file

Database Migrations
The service uses GORM's AutoMigrate feature to handle schema changes. If you modify the data models, the changes will be applied automatically when the service starts.
License
This project is licensed under the MIT License - see the LICENSE file for details.
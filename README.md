# Genggam Makna Backend API

[![Go Version](https://img.shields.io/badge/Go-1.22.1-blue)](https://golang.org) [![Gin Framework](https://img.shields.io/badge/Gin-1.10.0-blue)](https://gin-gonic.com/) [![GORM](https://img.shields.io/badge/GORM-ORM-yellow)](https://gorm.io/) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-blue?logo=postgresql)](https://www.postgresql.org/) [![Redis](https://img.shields.io/badge/Redis-7.4.1-red?logo=redis&logoColor=white)](https://redis.io/) [![License](https://img.shields.io/badge/license-MIT-green)](LICENSE) [![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](#)

A backend API built with modern technologies to serve the Genggam Makna platform.

---

Genggam Makna Backend API is a backend service designed to handle user requests from the mobile application **Genggam Makna**, an application for translating **SIBI (Indonesian Sign Language System)**. This service authenticates users, forwards requests for image and video predictions to the Machine Learning (ML) service, and leverages a memory store for caching prediction results to minimize ML service usage and improve performance.

---

## Project Features

1. **User Authentication**: Ensures secure access to Genggam Makna resources.
2. **Prediction Request Handling**: Forwards image and video prediction requests to the ML service.
3. **Caching with Redis**: Caches prediction results to reduce the load on the ML service and improve response times.

---

## Tech Stack

- **Programming Language**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM
- **Caching**: Redis

---

## Project Architecture

The **Genggam Makna** project comprises three main services:

1. **Mobile Application**: User-facing application for interacting with the platform.
2. **Backend API**: (This repository) Handles authentication, communication with the ML service, and caching.
3. **Machine Learning Service**: Processes image and video predictions.

---

## Installation and Setup

### Prerequisites

Ensure the following are installed on your system:
- Go (Golang) 1.22.1 or higher
- Redis server
- A PostgreSQL database

### Steps to Run the Backend API

1. Clone the repository:
   ```bash
   git clone https://github.com/GenggamMakna/genggam-makna-backend-api
   cd genggam-makna-backend-api

2. Install dependencies:
   ```bash
   go mod tidy

3. Configure environment variables:
   Create a `.env` file in the project root with the following variables:
   ```
   DB_USER="your-db-username"
   DB_PASSWORD="your-db-password"
   DB_HOST="your-db-host"
   DB_PORT="your-db-port"
   DB_NAME="your-db-name"

   REDIS_ADDRESS="your-redis-server-address"

   PREDICT_BASE_API_URL="your-ml-service-api/predict"

   PORT="your-application port"
   ENVIRONMENT="[development/production]"   

   JWT_SECRET="your-jwt-secret"
   GOOGLE_APPLICATION_CREDENTIALS="path-to-your-gcloud-service-account-secret.json"
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The application will be available at `http://localhost:your-application-port/api/ping` (default).

---

## API Endpoints

### Authentication
- **POST** `/auth/login` - User login
- **POST** `/auth/register` - User registration
- **POST** `/auth/google/login` - User registration/login via Google OAuth

### Prediction
- **POST** `/predict/image` - Predict hand sign from image
- **POST** `/predict/video` - Predict hand sign from video

### Prediction
- **GET** `/ping` - Ping Backend API Service

### Caching
- Results from predictions are automatically cached using Redis to reduce latency and avoid repeated calls to the ML service.

---

## Contributing

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/your-feature-name`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- **Gin**: For providing a fast and lightweight web framework.
- **GORM**: For seamless ORM integration.
- **Redis**: For enabling high-performance caching.
- **Genggam Makna**: The larger project to which this backend contributes.

---

## Contact

For further questions or contributions, please contact:
- **Name**: Rama Diaz
- **Website**: [xann.my.id](https://xann.my.id)
- **Email**: ramadiaz221@gmail.com

Feel free to customize the placeholder values (e.g., repository URL, environment variables, email). Let me know if you need further adjustments!
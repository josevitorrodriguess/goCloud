<div align="center">
  <img src="https://i.imgur.com/peIic7g.png" alt="GoCloud Storage Logo" width="250">
</div>

<br>
<br>

<div align="center">

# GoCloud Storage

**Your personal, secure, and scalable cloud storage solution built with Go**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![AWS S3](https://img.shields.io/badge/AWS-S3-FF9900?style=for-the-badge&logo=amazon-aws)](https://aws.amazon.com/s3/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql)](https://postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/docs/latest/)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)


[Features](#-features) • [Tech Stack](#-tech-stack) • [Getting Started](#-getting-started) • [API Documentation](#-api-documentation) • [Contributing](#-contributing)

</div>

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| **📁 File Management** | Upload, download, list, and delete your files with ease |
| **🔐 Secure Authentication** | Robust user authentication system with OAuth 2.0 |
| **☁️ Scalable Storage** | Leverage AWS S3 for durability and high availability |
| **🚀 RESTful API** | Clean and organized interface for client integration |
| **⚡ High Performance** | Built with Go for optimal speed and efficiency |

## 🛠️ Tech Stack

<div align="center">

| Component | Technology |
|-----------|------------|
| **Backend** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white) ![Chi](https://img.shields.io/badge/Chi-Router-blue) |
| **Storage** | ![AWS S3](https://img.shields.io/badge/AWS-S3-FF9900?style=flat&logo=amazon-aws&logoColor=white) |
| **Database** | ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) |
| **Authentication** | ![OAuth](https://img.shields.io/badge/OAuth-2.0-green) |
| **Sessions** | ![Redis](https://img.shields.io/badge/-Redis-DC382D?logo=Redis&logoColor=FFF) |

</div>

## 📁 Project Structure

```bash
go-cloud-storage/
├── 📂 cmd/                    # Application executables
│   └── 📂 goCloud/            # Main application (entrypoint)
│       └── 📄 main.go         # Application entry point
├── 📂 internal/               # Core business logic and main services
│   └── 📂 api/                # HTTP handlers and Chi routes
│   ├── 📂 auth/               # OAuth authentication logic
│   ├── 📂 domain/             # Domain models (User, File, etc)
│   ├── 📂 storage/            # AWS S3 integration and storage logic
│   ├── 📂 encryption/         # File encryption utilities
│   ├── 📂 jsonutils/          # JSON utility functions
│   ├── 📂 logger/             # Custom logger
│   ├── 📂 repository/         # Data access repositories (DB/S3)
│   ├── 📂 session/            # Session management with Redis
│   ├── 📂 usecase/            # Business use cases (application logic)
├── 📂 config/                 # Environment configuration
└── 📄 go.mod                  # Go module and dependencies
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.3 
- PostgreSQL 13+
- AWS Account with S3 access
- Github and Google oauth credentials
- Redis

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/josevitorrodriguess/goCloud.git
   cd goCloud
   ```

2. **Create your environment file**
   - Copy the example file and edit it:
     ```bash
     cp .env.example .env
     ```
   - Open `.env` in your favorite editor and fill in the required variables:
     - **Database:**  
       `DB_HOST=postgres`  
       `DB_PORT=5432`  
       `DB_USER=postgres`  
       `DB_PASSWORD=your_password`  
       `DB_NAME=your_db_name`  
       `DB_SSL_MODE=disable`
     - **Redis:**  
       `REDIS_ADDR=redis:6379`  
       `REDIS_PASSWORD=`
     - **AWS S3:**  
       `AWS_S3_BUCKET=your-bucket`  
       `AWS_REGION=us-east-1`
     - **Google OAuth:**  
       `GOOGLE_CLIENT_ID=your_google_client_id`  
       `GOOGLE_CLIENT_SECRET=your_google_client_secret`
     - **GitHub OAuth:**  
       `GITHUB_CLIENT_ID=your_github_client_id`  
       `GITHUB_CLIENT_SECRET=your_github_client_secret`
     - **Session:**  
       `SESSION_KEY=your_session_key`
     - **Encryption:**  
       `ENCRYPT_KEY=your_encryption_key`

3. **Configure OAuth providers**
   - Create OAuth applications in Google Cloud Console and GitHub.
   - Set the callback URLs to:  
     `http://localhost:3050/auth/google/callback`  
     `http://localhost:3050/auth/github/callback`
   - Paste the generated credentials into your `.env` file.

4. **Start the application with Docker Compose**
   ```bash
   docker compose up --build -d
   ```

5. **Access the application**
   - Backend: [http://localhost:3050](http://localhost:3050)
   - Redis and Postgres are managed automatically by Docker Compose.

6. **Check logs (optional)**
   ```bash
   docker compose logs -f app
   ```


## 📊 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/auth/{provider}` | Initiate OAuth flow (redirects to provider) |
| `GET` | `/api/auth/{provider}/callback` | OAuth provider callback |
| `GET` | `/api/auth/logout/{provider}` | Logout and invalidate tokens |

**Supported OAuth Providers:**
- Google (`/api/auth/oauth/google`)
- GitHub (`/api/auth/oauth/github`)


### File Management Endpoints

| Method | Endpoint           | Description                |
|--------|--------------------|----------------------------|
| POST   | `/file/upload`     | Upload file (multipart, authenticated) |
| GET    | `/file/download`   | Download file (authenticated, query param: filename) |
| GET    | `/file/list`       | List user files (authenticated) |
| DELETE | `/file/delete`     | Delete file (authenticated, query param: filename) |

**Usage examples:**
- Upload: send a multipart file to `/file/upload` (authenticated)
- Download: `GET /file/download?filename=yourfile.ext` (authenticated)
- List: `GET /file/list` (authenticated)
- Delete: `DELETE /file/delete?filename=yourfile.ext` (authenticated)



## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

### How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style

- Follow Go best practices and conventions
- Use `gofmt` to format your code
- Update documentation as needed

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

<div align="center">

**Built with ❤️ using Go**

[![Email](https://img.shields.io/badge/Email-josevitorrodrigues17@gmail.com-red?style=flat&logo=gmail)](mailto:josevitorrodrigues17@gmail.com)

*If you found this project helpful, please consider giving it a ⭐!*

</div>

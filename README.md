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
| **🛡️ Data Security** | Secure file handling with proper access controls |

## 🛠️ Tech Stack

<div align="center">

| Component | Technology |
|-----------|------------|
| **Backend** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white) ![Chi](https://img.shields.io/badge/Chi-Router-blue) |
| **Storage** | ![AWS S3](https://img.shields.io/badge/AWS-S3-FF9900?style=flat&logo=amazon-aws&logoColor=white) |
| **Database** | ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) |
| **Authentication** | ![OAuth](https://img.shields.io/badge/OAuth-2.0-green) |
| **SDK** | ![AWS SDK](https://img.shields.io/badge/AWS-SDK%20v2-FF9900?style=flat&logo=amazon-aws&logoColor=white) |

</div>

## 📁 Project Structure

```bash
go-cloud-storage/
├── 📂 cmd/                    # Executable applications
│   └── 📂 go-cloud-storage/   # Main application
│       └── 📄 main.go         # Application entry point
├── 📂 api/                    # HTTP handlers and Chi routing
├── 📂 internal/               # Core business logic, services, models
│   ├── 📂 database/           # PostgreSQL connection and migrations
│   ├── 📂 auth/               # Authentication and JWT management
│   ├── 📂 file/               # File manipulation business logic
│   ├── 📂 storage/            # AWS S3 interaction layer
│   └── 📂 user/               # User business logic
├── 📂 pkg/                    # Shared utility packages
├── 📂 config/                 # Environment configuration
└── 📄 go.mod                  # Go module and dependencies
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 13+
- AWS Account with S3 access
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-cloud-storage.git
   cd go-cloud-storage
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Configure OAuth providers**
   - Create OAuth applications in Google Cloud Console, GitHub, etc.
   - Add callback URLs: `http://localhost:8080/api/auth/oauth/{provider}/callback`
   - Update `.env` with your OAuth credentials

5. **Run the application**
   ```bash
   go run cmd/go-cloud-storage/main.go
   ```

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `DATABASE_URL` | PostgreSQL connection string | ✅ |
| `AWS_ACCESS_KEY_ID` | AWS access key | ✅ |
| `AWS_SECRET_ACCESS_KEY` | AWS secret key | ✅ |
| `AWS_REGION` | AWS region | ✅ |
| `S3_BUCKET_NAME` | S3 bucket name | ✅ |
| `JWT_SECRET` | JWT signing secret | ✅ |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID | ✅ |
| `GOOGLE_CLIENT_SECRET` | Google OAuth client secret | ✅ |
| `GITHUB_CLIENT_ID` | GitHub OAuth client ID | ✅ |
| `GITHUB_CLIENT_SECRET` | GitHub OAuth client secret | ✅ |
| `MICROSOFT_CLIENT_ID` | Microsoft OAuth client ID | ❌ |
| `MICROSOFT_CLIENT_SECRET` | Microsoft OAuth client secret | ❌ |
| `OAUTH_REDIRECT_URL` | OAuth callback URL | ✅ |
| `PORT` | Server port (default: 8080) | ❌ |

## 📊 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/auth/oauth/{provider}` | Initiate OAuth flow (redirects to provider) |
| `GET` | `/api/auth/oauth/{provider}/callback` | OAuth provider callback |
| `POST` | `/api/auth/refresh` | Refresh access token (if supported by provider) |
| `POST` | `/api/auth/logout` | Logout and invalidate tokens |

**Supported OAuth Providers:**
- Google (`/api/auth/oauth/google`)
- GitHub (`/api/auth/oauth/github`)
- Microsoft (`/api/auth/oauth/microsoft`)

### File Management Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/files/upload` | Upload file |
| `GET` | `/api/files` | List user files |
| `GET` | `/api/files/{id}` | Download file |
| `DELETE` | `/api/files/{id}` | Delete file |

## 🗺️ Development Roadmap

| Status | Feature |
|--------|---------|
| ✅ | OAuth Authentication (Google, GitHub, Microsoft) |
| ✅ | Basic File Management (Upload, Download, List, Delete) |
| 🔄 | File Version Control |
| 📋 | Desktop/Web Frontend Client Development |
| 📋 | Data Encryption at Rest |
| 📋 | File Sharing & Permissions |
| 📋 | Advanced Search & Filtering |
| 📋 | Admin Dashboard |

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
- Add tests for new features
- Update documentation as needed

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

<div align="center">

**Built with ❤️ using Go**

[![Email](https://img.shields.io/badge/Email-josevitorrodrigues17@gmail.com-red?style=flat&logo=gmail)](mailto:josevitorrodrigues17@gmail.com)

*If you found this project helpful, please consider giving it a ⭐!*

</div>

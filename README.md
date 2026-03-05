# 🤖 Valentine ROBO - AI GitHub Bot

An **AI-powered GitHub App** that can **design, generate, and commit code automatically** using an LLM.

The bot listens to GitHub issue comments and executes commands like `/generate` to create code, push it to a new branch, and open a Pull Request automatically.

This project demonstrates how LLMs can act as **autonomous software engineers inside GitHub workflows**.

---

# 🚀 Features

* GitHub App authentication
* Webhook-based event handling
* AI-powered code generation
* Automatic branch creation
* Commit generated files
* Pull request creation
* File validation and safety checks
* Command-based interaction via GitHub comments

---

# ⚙️ How It Works

```
GitHub Issue Comment
        │
        ▼
Webhook Event
        │
        ▼
Command Parser
        │
        ▼
LLM Code Generation
        │
        ▼
Validate Generated Files
        │
        ▼
Create Branch + Commit Files
        │
        ▼
Open Pull Request
```

Example command:

```
/generate create a simple golang http server
```

The bot will:

1. Generate project files using the LLM
2. Create a new branch
3. Commit generated files
4. Open a Pull Request

---

# 📁 Project Structure

```
ai-github-bot
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── webhook/
│   │   └── handler.go
│   │
│   ├── github/
│   │   ├── auth.go
│   │   └── repo.go
│   │
│   ├── commands/
│   │   └── generate.go
│   │
│   └── llm/
│       ├── client.go
│       ├── prompt.go
│       ├── parser.go
│       └── validator.go
│
├── .env
├── go.mod
├── Makefile
└── Dockerfile
```

---

# 📦 Requirements

* Go **1.22+**
* A GitHub App
* LLM API key
* Docker (optional)

---

# 🔑 Environment Variables

Create a `.env` file in the project root:

```
APP_ID=your_github_app_id
GITHUB_PRIVATE_KEY="-----BEGIN RSA PRIVATE KEY-----..."
OPENAI_API_KEY=your_openai_api_key
```

---

# 🛠 Installation

### 1. Clone the repository

```
git clone https://github.com/yourusername/ai-github-bot.git
cd ai-github-bot
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the bot

Using Makefile:

```
make run
```

Or directly:

```
go run ./cmd/server
```

Server will start on:

```
http://localhost:8080
```

---

# 🌐 GitHub App Setup

Create a GitHub App and configure the following.

### Webhook URL

```
https://your-domain/webhook
```

If testing locally, use a tunneling tool like ngrok.

---

### Repository Permissions

```
Contents: Read & Write
Pull Requests: Read & Write
Issues: Read & Write
Metadata: Read
```

---

### Subscribe to Events

Enable the following webhook events:

```
Issue comments
```

---

# 🧪 Usage

Create an issue in a repository where the GitHub App is installed.

Then comment:

```
/generate create a simple golang http server
```

The bot will:

* Generate project files
* Create a branch
* Commit generated files
* Open a Pull Request

---

# 🔒 Safety Controls

To prevent unwanted changes, the bot blocks:

* `.github/workflows`
* `.env`
* infrastructure directories
* excessively large files
* too many generated files

---

# 🐳 Docker (Optional)

### Build Docker image

```
docker build -t ai-github-bot .
```

### Run container

```
docker run -p 8080:8080 --env-file .env ai-github-bot
```

---

# 🧠 Future Improvements

Planned enhancements:

* Multi-step AI workflows (design → code → tests)
* Automated code review
* CI failure auto-fix
* Repository context awareness
* Multi-model support
* Incremental code editing
* Code refactoring commands

---

# 🤝 Contributing

Contributions are welcome.

Steps:

1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Open a Pull Request

---

# 📜 License

MIT License

---

# ⭐ Project Vision

The long-term goal is to build a **fully autonomous software engineering assistant** capable of:

* Designing system architecture
* Writing production-ready code
* Reviewing pull requests
* Fixing CI failures
* Improving repositories continuously

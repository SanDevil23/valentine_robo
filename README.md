GitHub Repo
   |
   |  (Webhook events)
   ↓
GitHub App (Your Bot)
   |
   |-- Event Processor (Go)
   |-- Command Parser (/generate, /design, /test)
   |-- LLM Client (OpenAI / local LLM)
   |-- Code Generator
   |-- GitHub API Client
   |
   ↓
Creates Branch → Commits Code → Opens PR → Comments Results
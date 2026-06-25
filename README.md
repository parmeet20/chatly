<p align="center">
  <img width="400" height="150" alt="Chatly" src="https://github.com/user-attachments/assets/b123a90d-bba1-4dac-bcf0-95af5b13482d" />
</p>

<p align="center">
  <strong>Modern Real-Time Messaging Platform</strong>
</p>

<p align="center">
  Next.js • Golang • MongoDB • WebSockets • Turborepo
</p>

<p align="center">
  <img src="https://skillicons.dev/icons?i=nextjs,react,ts,go,mongodb,docker,kubernetes,github" />
</p>

<p align="center">
  <img src="https://img.shields.io/badge/WebSocket-Real_Time-0088CC?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Turborepo-Monorepo-EF4444?style=for-the-badge" />
</p>

<p align="center">
  <img src="https://img.shields.io/github/stars/parmeet20/chatly?style=flat-square" />
  <img src="https://img.shields.io/github/forks/parmeet20/chatly?style=flat-square" />
  <img src="https://img.shields.io/github/issues/parmeet20/chatly?style=flat-square" />
  <img src="https://img.shields.io/github/license/parmeet20/chatly?style=flat-square" />
</p>

<p align="center">
<strong>A scalable real-time messaging platform built with Next.js, Golang, MongoDB and WebSockets.</strong>
</p>

<p align="center">
<a href="#-overview">Overview</a> •
<a href="#-features">Features</a> •
<a href="#-architecture">Architecture</a> •
<a href="#-tech-stack">Tech Stack</a> •
<a href="#-project-structure">Project Structure</a> •
<a href="#-getting-started">Getting Started</a> •
<a href="#-deployment">Deployment</a> •
<a href="#-contributing">Contributing</a>
</p>

---

# 🎯 Overview

Chatly is a modern, full-stack real-time communication platform designed for high-performance messaging applications.

The platform consists of:

### 🌐 Chatly Web Client

A modern web application built with Next.js that provides users with a seamless chat experience.

Features include:

* User authentication
* Real-time messaging
* Responsive user interface
* Chat room management
* Message history
* Secure session handling
* Optimized performance and SEO

### ⚡ Chatly Server

A scalable Golang backend responsible for:

* Authentication & authorization
* JWT token management
* WebSocket communication
* Real-time message delivery
* Message persistence
* User management
* Chat room operations
* API services

The system is built using modern software engineering principles including clean architecture, dependency injection, modular design, and containerized deployment.

---

# ✨ Features

## 🔐 Authentication & Security

* JWT authentication
* Secure password hashing
* Protected API routes
* Refresh token support
* Input validation
* CORS protection
* Secure session management

## 💬 Real-Time Messaging

* WebSocket-powered communication
* Instant message delivery
* Room-based conversations
* Online user management
* Persistent message storage
* Message synchronization
* Scalable connection handling

## 🎨 Modern Frontend Experience

* Next.js App Router
* Responsive design
* Server-side rendering
* Optimized performance
* Type-safe development
* Component-based architecture

## ⚙️ Backend Infrastructure

* RESTful APIs
* Clean Architecture
* Dependency Injection
* Context propagation
* Structured logging
* MongoDB integration
* Docker support

## 🚀 Developer Experience

* Turborepo Monorepo
* TypeScript support
* Shared configurations
* CI/CD friendly
* Dockerized services
* Scalable codebase

---

# 🏗 Architecture

```text
┌─────────────────────────────┐
│      Chatly Web Client      │
│        Next.js App          │
└──────────────┬──────────────┘
               │
               │ REST API
               ▼
┌─────────────────────────────┐
│       Chatly Server         │
│       Golang + Chi          │
└──────────────┬──────────────┘
               │
      ┌────────┴────────┐
      │                 │
      ▼                 ▼

 MongoDB Atlas      WebSocket Hub
 Message Store     Real-Time Events
```

### Request Flow

```text
User
 │
 ▼
Chatly Web
 │
 ▼
REST API
 │
 ▼
Authentication Middleware
 │
 ▼
Business Services
 │
 ▼
MongoDB
```

### Real-Time Flow

```text
User A
  │
  ▼
WebSocket Connection
  │
  ▼
Chatly Server
  │
  ▼
WebSocket Hub
  │
  ▼
User B
```

---

# 🛠 Tech Stack

## Frontend

* Next.js 16
* React 19
* TypeScript
* Tailwind CSS
* Turborepo

## Backend

* Golang 1.25
* Chi Router
* Gorilla WebSocket
* JWT Authentication
* MongoDB Driver

## Database

* MongoDB Atlas

## DevOps

* Docker
* GitHub Actions
* Turborepo
* Kubernetes Ready

---

# 📁 Project Structure

```text
chatly
│
├── apps
│   │
│   ├── chatlyweb
│   │   ├── src
│   │   ├── public
│   │   ├── package.json
│   │   └── Dockerfile
│   │
│   └── chatlyserver
│       ├── cmd
│       ├── internal
│       ├── pkg
│       ├── go.mod
│       └── Dockerfile
│
├── packages
│   ├── eslint-config
│   ├── typescript-config
│   └── shared
│
├── .github
│   └── workflows
│
├── turbo.json
├── package.json
└── README.md
```

---

# 🚀 Getting Started

## Clone Repository

```bash
git clone https://github.com/yourusername/chatly.git

cd chatly
```

## Install Dependencies

```bash
pnpm install
```

## Run Frontend

```bash
pnpm --filter chatlyweb dev
```

Frontend runs on:

```text
http://localhost:3000
```

## Run Backend

```bash
cd apps/chatlyserver

go run cmd/server/main.go
```

Backend runs on:

```text
http://localhost:8080
```

---

# 🔌 Core Modules

### Authentication

* User Registration
* User Login
* JWT Access Tokens
* Refresh Tokens

### Messaging

* Create Rooms
* Join Rooms
* Send Messages
* Receive Messages
* Message History

### Real-Time Engine

* WebSocket Connections
* Connection Lifecycle Management
* Room Broadcasting
* Event Handling

---

# 🚢 Deployment

## Docker

Frontend

```bash
docker build -t chatly-web -f apps/chatlyweb/Dockerfile .
```

Backend

```bash
docker build -t chatly-server -f apps/chatlyserver/Dockerfile .
```

Run containers

```bash
docker run -p 3000:3000 chatly-web

docker run -p 8080:8080 chatly-server
```

---

# ⚡ CI/CD

GitHub Actions workflows can be configured to:

* Run tests
* Build frontend
* Build backend
* Build Docker images
* Push images to container registry
* Deploy to Kubernetes

Workflow location:

```text
.github/workflows
```

---

# 🌟 Why Chatly?

Chatly demonstrates modern full-stack engineering practices:

* Real-time communication architecture
* Production-grade Golang backend
* Modern Next.js frontend
* WebSocket scalability patterns
* Secure authentication workflows
* Dockerized deployments
* Monorepo development workflow

It serves as a strong foundation for building:

* Messaging platforms
* Team collaboration tools
* Community applications
* Customer support systems
* Real-time social applications

---

# 🤝 Contributing

1. Fork the repository
2. Create a feature branch

```bash
git checkout -b feat/amazing-feature
```

3. Commit changes

```bash
git commit -m "feat: add amazing feature"
```

4. Push branch

```bash
git push origin feat/amazing-feature
```

5. Open a Pull Request

---

# 📜 License

This project is licensed under the MIT License.

---

<p align="center">
  Built with ❤️ by <strong>Parmeet Singh</strong>
</p>

<p align="center">
  <strong>Chatly — Connecting Conversations in Real Time.</strong>
</p>

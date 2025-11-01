# Wallet-application
A scalable microservices-based Wallet &amp; Discount System built with Go, Fiber, GORM, and PostgreSQL, fully containerized using Docker.


# 🪙 Wallet & Discount System

> A high-performance, microservices-based Wallet & Discount platform built with **Go (Fiber + GORM)** and **PostgreSQL**, fully containerized.

---

## 🚀 Overview

This system simulates a **digital wallet & gift-code platform** — like when a football match halftime promo gives the *first 1000 users* one million tomans credit 🤑  

It’s built with a **clean microservice architecture** for scalability and separation of concerns.

### 🧩 Services
| Service | Description |
|----------|-------------|
| 🏦 **Wallet Service** | Manages users’ wallets, balances, and transactions. |
| 🎁 **Discount Service** | Handles promo codes and enforces “first 1000 users” logic. |
| 🌐 **API Gateway** | A single public entry point for clients to redeem codes and view balances. |

---

## ⚙️ Tech Stack

| Layer | Technology |
|-------|-------------|
| Language | Go (1.24+) |
| Framework | Fiber |
| ORM | GORM |
| Database | PostgreSQL |
| Containerization | Docker & Docker Compose |
| Deployment | ArvanCloud Container Registry |
| Logging | Fiber Middleware Logger |
| Concurrency Control | SQL transactions + Row Locking |


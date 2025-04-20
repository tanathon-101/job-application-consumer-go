# 🐹 job-application-consumer-go

Go microservice for consuming job application events from RabbitMQ queue `job_application`.  
This service is designed to work with [`JobApplicationTrackerAPI`]([https://github.com/yourusername/JobApplicationTrackerAPI](https://github.com/tanathon-101/JobApplicationTrackerAPI)) as the event publisher.

---

## 📦 Features

- Subscribe to `job_application` queue in RabbitMQ
- Parse and display job application details
- Ready for containerized deployment with Docker
- Lightweight and fast (written in Go)

---

## 🚀 Getting Started

### 1. Clone the project

```bash
git clone https://github.com/yourusername/job-application-consumer-go.git
cd job-application-consumer-go

# CorsaF1

**CorsaF1** is a cloud-native Formula 1 analytics platform built in Golang.

## Features
- 🚀 REST API for F1 standings, schedules, constructor stats, driver comparisons
- 🗂 PostgreSQL database
- 🕒 Scheduled data sync (cron)
- 🔁 Redis caching
- ☁️ Dockerized for Cloud Run

## Run Locally
```bash
docker-compose up --build
```

Visit:
- `http://localhost:8080/drivers`
- `http://localhost:8080/schedule`
- `http://localhost:8080/constructors`
- `http://localhost:8080/compare?a=Verstappen&b=Hamilton`

---
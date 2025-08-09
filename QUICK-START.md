# 🚀 VivaLearning - Quick Start Guide

## ⚡ One-Command Start
```powershell
.\start-all.ps1
```

## ⚡ One-Command Stop
```powershell
.\stop-all.ps1
```

## 🔧 Manual Start (Two Terminals)

**Terminal 1 - Backend:**
```powershell
.\vivaLearning.exe serve
```

**Terminal 2 - Frontend:**
```powershell
cd frontend
npm run dev
```

## 🌐 Access Points
- **Frontend**: http://localhost:3000
- **Backend**: http://localhost:8080/api/v1  
- **Health**: http://localhost:8080/ping

## 🧪 Quick Test
```powershell
.\test-api.ps1
```

## 🛠️ Build Commands
```powershell
# Backend
go build -o vivaLearning.exe .

# Frontend
cd frontend
npm run build
```

## 📝 Test Account
- **Email**: test@example.com
- **Password**: password123

---
**Need help? Check HOW-TO-RUN.md for detailed instructions**


```markdown
# FullStack Meeting App

This is a simple full-stack application for managing meetings, built with **Go (Gin)** for the backend and **React + Vite** for the frontend.

## 📁 Project Structure

```

FullStack/
├── backend/       # Go server using Gin and PostgreSQL
├── frontend/      # React app powered by Vite

````

## ⚙️ Technologies Used

- Go (Gin Framework)
- PostgreSQL
- React (Vite)
- GORM (ORM for Go)
- JWT (Authentication)

## 🛠️ Running the App (Local Development)

### 1. Backend

Make sure PostgreSQL is running and the connection details are correctly set in `backend/config/config.go`.

```bash

cd backend
go run main.go

```

### 2. Frontend

```bash
cd frontend
npm install
npm run dev
```

## 🔗 Connecting Frontend and Backend

Currently, the frontend and backend run on separate ports. You can serve the frontend build files via Gin if you want to deploy them together under a single server.

## 📝 TODOs

* [ ] Improve UI design
* [ ] Add dashboard for managing meetings
* [ ] Connect frontend to backend APIs
* [ ] Deploy on a production server

## ✍️ Author

* [@mjavadtavakoli](https://github.com/mjavadtavakoli)

---

Feel free to open issues for feature requests or bug reports.

```

---

Let me know if you'd like a version with deployment instructions or Docker support included!
```

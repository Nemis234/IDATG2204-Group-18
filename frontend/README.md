# Electromart

## Setup & Deployment (Windows)

### 📦 Requirements
- Node.js (LTS) — https://nodejs.org/
- Vite — installed via `npm install`
- Code editor like VS Code

### ▶️ Running Locally
1. Clone/unzip this project.
2. In terminal, run:
```bash
npm install
npm run dev
```
3. Visit `http://localhost:5173` in your browser.

### 📁 Project Structure
```
src/
  components/      # UI blocks
  pages/           # Route views
  data/            # JSON mock data
  context/         # Auth/Cart state (optional)
public/
  index.css        # Stylesheet
index.html
vite.config.js
tailwind.config.js
```

This project uses static mock data in JSON for frontend-only simulation. Backend integration expected to be handled by team members.
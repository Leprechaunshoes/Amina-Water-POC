Amina Water Backend — Step-by-Step

1. Copy .env.example and rename it to .env
2. Fill these:
   PROJECT_WALLET = your wallet address
   ADMIN_TOKEN = random string
3. On a laptop or Codespace, run:
   go mod tidy
   go run main.go
4. It starts on http://localhost:8080

Test it:
- Visit http://localhost:8080/stats → shows counter
- POST to /simulate-swap (each hit adds +1)

To connect to your site:
- In scripts.js change BACKEND URL:
  const BACKEND = "http://localhost:8080";

Later you’ll deploy it to Railway, Render, or Vercel for public use.

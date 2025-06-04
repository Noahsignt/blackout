# Summary
Frontend and backend for Blackout online game.
Probably the greatest card game ever made, have fun.

# Local Dev Environment

## Frontend
The frontend is a simple React app using TypeScript and Tailwind. Additionally, it leverages `boardgame.io` to help with turn-taking functionality.

### Starting
```bash
#!/bin/bash
cd fe
nvm use node
npm i
npm run dev
```

## Backend
The backend is written in Go. Should mean super performant (spend less $$$ on EC2 instances) for small number of players.

```bash
#!/bin/bash
cd be
go run main.go
```
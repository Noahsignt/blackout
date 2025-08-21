# Summary
Frontend and backend for Blackout online game.
Probably the greatest card game ever made, have fun.

# Structure

## Frontend
The frontend is a simple Vite React app using TypeScript and Tailwind. Additionally, it leverages `boardgame.io` to help with turn-taking functionality.

### Starting
```bash
#!/bin/bash
cd fe
nvm use node
npm i
npm run dev
```

### Deployment Info
Deployed using CloudFlare Pages.
`https://dash.cloudflare.com/feeec934d03d75f2325ddbf46af9c6bc/pages/view/blackout`

## Backend
The backend is written in Go. Should mean super performant (spend less $$$ on EC2 instances) for small number of players.

#### Local Dev

### Starting
```bash
#!/bin/bash
cd be
go run main.go
```

### Database
```bash
#!/bin/bash
mongosh "mongodb://localhost:27017/blackout"
show collections
db.users.find().pretty()
```

### Deployment Info
Deployed using an Akamai Nanode. Simple, cheap Virtual Private Server solution. Systemd used to boot the Go server up.
DuckDNS used for DNS routing.

#### Important Links
`http://blackout-be.duckdns.org:8080/`
`https://cloud.linode.com/linodes/78840807/metrics`

#### CI/CD
GitHub Actions is used for automated deployment.
Workflow SSHs in, pulls the merged code from main, rebuilds the Go binary and then restarts the server (managed via systemd).

### Testing
`go test ./integration`
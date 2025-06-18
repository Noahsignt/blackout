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

### Starting
```bash
#!/bin/bash
cd be
go run main.go
```

### Deployment Info
Deployed using an Akamai Nanode. Simple, cheap Virtual Private Server solution. Systemd used to boot the Go server up.
DuckDNS used for DNS routing.

`http://blackout-be.duckdns.org:8080/`
`https://cloud.linode.com/linodes/78840807/metrics`

### Testing
`go test ./integration`
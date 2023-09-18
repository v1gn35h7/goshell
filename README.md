# GoShell

Remote script management tool for IT written in go Lang.


## Set-up
### For node 18
curl -sL https://deb.nodesource.com/setup_18.x | bash -
apt-get install -y nodejs

## Build
npx tailwindcss build --output frontend/dist/css/output.css
npm run build

go build cmd/goshell/goshell.go

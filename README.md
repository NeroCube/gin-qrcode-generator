# Gin-QR-Code-Generator
Implement QR Code Generator in Docker with Gin framework
## Install
`docker-compose up --build`

## Hello Service !!
`curl -X GET "http://localhost:9000/ping"`

## Generate your own QR code
Use your browser with the below url, `content` is the URL you want to generate
`http://localhost:9000/qrcode\?content=https://github.com/NeroCube`

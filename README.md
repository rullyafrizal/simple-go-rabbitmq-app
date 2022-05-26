# RabbitMQ Work Queues

## Reference
- https://www.rabbitmq.com/tutorials/tutorial-two-go.html
## Installation
- Pastikan sudah install Go dan Docker
- Copy `.env.example` dan rename menjadi `.env`
- Run `go mod download`
- Jalankan command berikut:
```
docker compose up -d
```
tunggu hingga proses pull image dan build container selesai <br>
- Selanjutnya, untuk menyalakan consumer jalankan command berikut di beberapa terminal, hal ini berguna untuk mensimulasikan round-robin dispatching
```
make start
```
- Pindah ke terminal lain dan jalankan publisher dengan command berikut
```
make publish
```

Terminal-terminal yang kita gunakan untuk menyalakan consumer akan terbagi rata konsumsinya.
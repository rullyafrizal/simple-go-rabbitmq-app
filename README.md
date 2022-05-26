# RabbitMQ

## List of RabbitMQ Tutorials
Dalam repository ini, saya taruh metode penggunaan RabbitMQ dalam beberapa branch menganut pada tutorial di website resmi [RabbitMQ](https://www.rabbitmq.com/getstarted.html)
- https://github.com/rullyafrizal/simple-go-rabbitmq-app/tree/work-queues

## Installation
- Pastikan sudah install Go dan Docker
- Copy `.env.example` dan rename menjadi `.env`
- Jalankan command berikut:
```
docker compose up -d
```
<br>tunggu hingga proses pull image dan build container selesai <br>
- Selanjutnya, untuk menyalakan consumer jalankan command berikut:
```
make start
```
- Pindah ke terminal lain dan jalankan publisher dengan command berikut
```
make publish
```
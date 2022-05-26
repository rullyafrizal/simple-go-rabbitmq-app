# RabbitMQ Work Queues

## Reference
[- https://www.rabbitmq.com/tutorials/tutorial-two-go.html](https://www.rabbitmq.com/tutorials/tutorial-three-go.html)

## Konsep
Konsep utama dari messaging sendiri di RabbitMQ adalah publisher tidak mengirim data secara langsung ke queue, bahkan publisher tidak tahu-menahu apakah data dikirimkan ke queue. Publisher hanya bertanggung jawab untuk mengirimkan data ke exchange.<br>

Exchange hanyalah satu layer yang berguna untuk menerima data dari publisher dan mengirimkan ke queue yang sudah dipasangkan dengan exchange.

### Diagram Konsep
![konsep-exchange](https://www.rabbitmq.com/img/tutorials/exchanges.png)

#### Beberapa Type Exchange
- Direct Exchange
- Topic Exchange
- Headers Exchange
- Fanout Exchange

## Installation
- Pastikan sudah install Go dan Docker
- Copy `.env.example` dan rename menjadi `.env`
- Run `go mod download`
- Jalankan command berikut:
```
docker compose up -d
```
tunggu hingga proses pull image dan build container selesai <br>
- Selanjutnya, untuk menyalakan consumer jalankan command berikut di beberapa terminal, berbeda dengan work-queues, di dalam pub-sub ini semua consumer akan mengeluarkan output yang sama, mengapa hal ini bisa terjadi? karena publisher mengirim data ke sebuah exchange, dan data dari exchange akan dikirimkan ke queue tiap-tiap consumer.
```
make start
```
- Pindah ke terminal lain dan jalankan publisher dengan command berikut
```
make publish
```

Terminal-terminal yang kita gunakan untuk menyalakan consumer akan terbagi rata konsumsinya.
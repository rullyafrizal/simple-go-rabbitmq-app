# RabbitMQ Routing

## Reference
- https://www.rabbitmq.com/tutorials/tutorial-four-go.html

## Konsep
Konsep utama dari messaging sendiri di RabbitMQ adalah publisher tidak mengirim data secara langsung ke queue, bahkan publisher tidak tahu-menahu apakah data dikirimkan ke queue. Publisher hanya bertanggung jawab untuk mengirimkan data ke exchange.<br>

Exchange hanyalah satu layer yang berguna untuk menerima data dari publisher dan mengirimkan ke queue yang sudah dipasangkan dengan exchange. <br>

Di bagian routing ini, queue dari consumer hanya akan mengkonsumsi/subscribe ke message tertentu saja. Hal ini bisa kita define dengan memasukkan routing key ketika melakukan binding antara exchange dan queue.<br>

Kita akan menerapkan konsep FizzBuzz untuk routing di bagian ini, post yang memiliki ID kelipatan 3 akan dikirimkan ke queue Fizz, kelipatan 5 akan dikirimkan ke queue Buzz, kelipatan 3 dan 5 akan dikirimkan ke queue FizzBuzz. Untuk itu, kita perlu menyalakan 3 customer yang akan menerima data. <br>

Dalam praktek di real-world, konsep ini bisa digunakan untuk memisahkan logging berdasarkan statusnya (info, warning, critical, dsb.), dan juga lain-lain.

### Diagram Konsep
![konsep-route](https://www.rabbitmq.com/img/tutorials/direct-exchange.png)

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
- Selanjutnya, untuk menyalakan consumer jalankan command berikut di beberapa terminal.
```
make start route=fizz
```
```
make start route=buzz
```
```
make start route=fizzbuzz
```
- Pindah ke terminal lain dan jalankan publisher dengan command berikut
```
make publish
```
Dapat diperhatikan di terminal-terminal yang kita gunakan untuk menyalakan consumer akan mendapat hasil yang berbeda sesuai parameter route yang kita masukkan, jika fizz maka akan mendapatkan post dengan id kelipatan 3, jika buzz maka akan mendapatkan post dengan id kelipatan 5 dan default, dan jika fizzbuzz maka akan mendapatkan post dengan id kelipatan 3 dan 5.
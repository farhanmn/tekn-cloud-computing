# Tugas 12 - Docker Orchestration Hands-on Lab

## Konfigurasi Swarm Mode

1. Membuat container baru pada terminal node1
   ![newcon](new_container.png)

```
// Menggunakan container terbaru
$ docker run -dt ubuntu sleep infinity
```

2. Melihat daftar container
   ![dockerps](docker_ps.png)

```
$ docker ps
```

3. Inisiasi docker swarm
   ![dockerswarm](docker_swarm.png)

```
$ docker swarm init --advertise-addr $(hostname -i)

// akan mendapatkan token untuk menggabungkan dengan node lain
```

4. Melihat informasi docker swarm pada node1
   ![dockerinfo1](docker_info_1.png)
   ![dockerinfo2](docker_info_2.png)
   ![dockerinfo3](docker_info_3.png)

```
$ docker info
```

5. Menggabungkan node2 dan node3 kedalam swarm node1 menggunakan token yang telah didapatkan di node1.
   ![worker_2](worker_node2.png)
   ![worker_3](worker_node3.png)

```
// node2 dan node3 telah terhubung dengan swarm node1 sebagai worker
```

6. Melihat daftar node
   ![docker_ps1](docker_ps1.png)

## Deploy aplikasi antar hosts

7. Deploy komponen aplikasi sebagai service Docker
   ![deploy](deploy_component.png)

8. Cek seluruh servis docker yang berjalan
   ![check](service_ls.png)

## Eskalasi aplikasi

9. Mereplika layanan yang sama dengan skala 7
   ![replicas](replicas.png)
   ![replicas](replicas_2.png)

10. Melihat layanan yang sedang berjalan
    ![service_1](service_1.png)
    ![service_2](service_2.png)

11. Mereplika layanan yang sama dengan skala 4
    ![replicas](replicas_4.png)

12. Melihat layanan yang sedang berjalan
    ![service_3](service_3.png)
    ![service_4](service_4.png)

13. Melihat daftar node
    ![docker_ls](docker_node.png)

## Drain node dan lakukan jadwal ulang container

14. Melihat daftar container
    ![docker_ps](docker_ps3.png)

15. Melihat kembali daftar node yang ada.
    ![docker_ls](docker_node.png)

16. Mengubah availability node 2 menjadi drain
    ![drain](drain.png)

17. Hasil
    ![node](docker_node1.png)

18. Melihat container yang berjalan pada node2
    ![docker_ps2](docker_ps2.png)

```
// Container kosong karena node2 availability sebelumnya telah diubah menjadi drain
```

19. Melihat layanan yang sedang berjalan
    ![service_5](service_5.png)
    ![service_6](service_6.png)

```
// layanan yang menggunakan node2 terhenti
```

## Clean up

20. Menghapus layanan
    ![cleanup](cleanup.png)

21. Melihat daftar container yang ada
    ![docker_ps4](docker_ps4.png)

22. Menghapus container
    ![kill](kill.png)

23. Menghapus swarm pada node1, node2 dan node3
    ![leave1](leave_node1.png)
    ![leave2](leave_node2.png)
    ![leave3](leave_node3.png)

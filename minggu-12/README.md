# Tugas 12 - Docker Orchestration Hands-on Lab

## Konfigurasi Swarm Mode

1. Membuat container baru pada terminal node1
   ![newcon](assets/new_container.png)

```
// Menggunakan container terbaru
$ docker run -dt ubuntu sleep infinity
```

2. Melihat daftar container
   ![dockerps](assets/docker_ps.png)

```
$ docker ps
```

3. Inisiasi docker swarm
   ![dockerswarm](assets/docker_swarm.png)

```
$ docker swarm init --advertise-addr $(hostname -i)

// akan mendapatkan token untuk menggabungkan dengan node lain
```

4. Melihat informasi docker swarm pada node1
   ![dockerinfo1](assets/docker_info_1.png)
   ![dockerinfo2](assets/docker_info_2.png)
   ![dockerinfo3](assets/docker_info_3.png)

```
$ docker info
```

5. Menggabungkan node2 dan node3 kedalam swarm node1 menggunakan token yang telah didapatkan di node1.
   ![worker_2](assets/worker_node2.png)
   ![worker_3](assets/worker_node3.png)

```
// node2 dan node3 telah terhubung dengan swarm node1 sebagai worker
```

6. Melihat daftar node
   ![docker_ps1](assets/docker_ps1.png)

## Deploy aplikasi antar hosts

7. Deploy komponen aplikasi sebagai service Docker
   ![deploy](assets/deploy_component.png)

8. Cek seluruh servis docker yang berjalan
   ![check](assets/service_ls.png)

## Eskalasi aplikasi

9. Mereplika layanan yang sama dengan skala 7
   ![replicas](assets/replicas.png)
   ![replicas](assets/replicas_2.png)

10. Melihat layanan yang sedang berjalan
    ![service_1](assets/service_1.png)
    ![service_2](assets/service_2.png)

11. Mereplika layanan yang sama dengan skala 4
    ![replicas](assets/replicas_4.png)

12. Melihat layanan yang sedang berjalan
    ![service_3](assets/service_3.png)
    ![service_4](assets/service_4.png)

13. Melihat daftar node
    ![docker_ls](assets/docker_node.png)

## Drain node dan lakukan jadwal ulang container

14. Melihat daftar container
    ![docker_ps](assets/docker_ps3.png)

15. Melihat kembali daftar node yang ada.
    ![docker_ls](assets/docker_node.png)

16. Mengubah availability node 2 menjadi drain
    ![drain](assets/drain.png)

17. Hasil
    ![node](assets/docker_node1.png)

18. Melihat container yang berjalan pada node2
    ![docker_ps2](assets/docker_ps2.png)

```
// Container kosong karena node2 availability sebelumnya telah diubah menjadi drain
```

19. Melihat layanan yang sedang berjalan
    ![service_5](assets/service_5.png)
    ![service_6](assets/service_6.png)

```
// layanan yang menggunakan node2 terhenti
```

## Clean up

20. Menghapus layanan
    ![cleanup](assets/cleanup.png)

21. Melihat daftar container yang ada
    ![docker_ps4](assets/docker_ps4.png)

22. Menghapus container
    ![kill](assets/kill.png)

23. Menghapus swarm pada node1, node2 dan node3
    ![leave1](assets/leave_node1.png)
    ![leave2](assets/leave_node2.png)
    ![leave3](assets/leave_node3.png)

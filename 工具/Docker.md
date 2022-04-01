`docker image pull` 等价于 `docker pull`

`docker container run`等价于`docker run`  

`docker container ls`等价于`docker ps`  
默认只显示运行中的容器，加上`-a`显示全部容器

`docker container kill [containerID]` 手动终止不会自动终止的容器  
`docker container rm [containerID]`删除已终止的容器
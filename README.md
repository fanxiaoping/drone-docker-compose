# drone-docker-compose
drone 插件扩展

## Build

```console
make build
```

## Docker

Build the Docker images with the following commands:

```console
docker build \
  --file docker/Dockerfile.linux.amd64 --tag fanxp/drone-docker-compose:v1.0.0 .
```

## Usage

```yaml
kind: pipeline
type: docker
name: default

node:
  mode: release
  service: major

steps:
  - name: deploy
    image: fanxp/drone-docker-compose
    volumes:
      - name: docker
        path: /var/run/docker.sock
    settings:
      registry:
        from_secret: docker_registry
      username:
        from_secret: docker_username
      password:
        from_secret: docker_password
    commands:
      # 清理 docker 服务中不再使用的镜像、挂载、网络等资源
      - docker system prune --force --volumes || true
      # 执行编排文件
      - docker-compose -f docker-compose.yml up -d

trigger:
  event:
    - push

volumes:
  - name: docker
    host:
      path: /var/run/docker.sock
```
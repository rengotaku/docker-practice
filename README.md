# docker-practice
Docker初心者のために送るステップバイでDockerを動かす道場です。

Docker composeの利用を前提としております。

# 動かし方
```
$ cd sample-docker-comoses
$ docker compose -f docker-compose-1.yml up
```

# 利用環境

```
$ docker --version
Docker version 24.0.5, build ced0996
$ echo "$(sw_vers)\n$(sysctl -n machdep.cpu.brand_string)"
ProductName:            macOS
ProductVersion:         13.2.1
BuildVersion:           22D68
Apple M2
```
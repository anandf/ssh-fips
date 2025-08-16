# ssh-fips

Code changes for reproducing the SSH handshake problem with curve25519 algorithm which is not FIPS compliant. 
The binary must fail at runtime when run in a FIPS enabled cluster.

```
podman build -t quay.io/anjoseph/ssh_fips:latest .
podman push quay.io/anjoseph/ssh_fips:latest
```


# openvpn-lb

A wrapper around openvpn (2.x) to make it use multiple CPUs.

## Status

This is a work in progress. Use it at your own risk.

## License

GPL3

## Build

For portability, disable `CGO`:

```
CGO_ENABLED=0 go build
```

## Run

Invoke `openvpn-lb` with the arguments to be passed to individual `openvpn`
processes. Make sure that you remove `--port` and `--proto` from your config
file.

You can control how many cores to use with the `NCPU` environment variable:

```
NCPU=2 /usr/sbin/openvpn-lb --config server.conf
```

Also, be aware that the `CWD` is currently hardcoded to `/etc/openvpn/server`,
so your config files should be relative to that.

You should have one `tcp` and one `udp` port open per used core:

```
root@htp# netstat -putan | grep openvpn
tcp        0      0 0.0.0.0:1194            0.0.0.0:*               LISTEN      314493/openvpn
tcp        0      0 0.0.0.0:1195            0.0.0.0:*               LISTEN      314495/openvpn
udp        0      0 0.0.0.0:1194            0.0.0.0:*                           314491/openvpn
udp        0      0 0.0.0.0:1195            0.0.0.0:*                           314496/openvpn
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      314487/openvpn-lb
```

## Bugs

Likely. This is a sunday afternoon hack.

## References

From [this answer](https://serverfault.com/a/1024171) in Stack Overflow:

> There's no difference in this matter between Access Server and community edition, the engine is the same.
> OpenVPN Access Server just opens multiple OpenVPN instances, and load-balance between them. The same standard scaling method for community edition.


The [OpenVPN Access Server docs](https://openvpn.net/vpn-server-resources/advanced-option-settings-on-the-command-line/)
gives a little bit more detail:

> Because the OpenVPN 2 code base is single-thread, meaning that an OpenVPN
> process can run on only 1 CPU core and doesn't know how to make use of
> multi-core systems, the OpenVPN Access Server comes with the ability to launch
> multiple OpenVPN daemons at the same time. Ideally there would be one OpenVPN
> daemon for every CPU core. But there's more involved. To make it possible for
> OpenVPN clients to establish a connection via the UDP protocol and via the TCP
> protocol, there are additional OpenVPN daemons necessary. In the case of the
> OpenVPN Access Server this means we launch 1 TCP and 1 UDP daemon per CPU core.

* Maybe useful: [Turning IPTables into a TCP load balancer for fun and profit](https://scalingo.com/blog/iptables)
* [iptables load balancer tutorial](https://github.com/muzahid-c/iptables-loadbalancer)

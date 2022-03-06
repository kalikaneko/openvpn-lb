# openvpn-lb

A wrapper around openvpn (2.x) to make it use multiple CPUs.

## Status

This is a work in progress. Use it at your own risk.

## License

GPL3

## References

From [this answer](https://serverfault.com/a/1024171) in Stack Overflow:

> There's no difference in this matter between Access Server and community edition, the engine is the same.
> OpenVPN Access Server just opens multiple OpenVPN instances, and load-balance between them. The same standard scaling method for community edition.


The [OpenVPN Access Server docs](https://openvpn.net/vpn-server-resources/advanced-option-settings-on-the-command-line/)
give a little bit more detail:

> Because the OpenVPN 2 code base is single-thread, meaning that an OpenVPN
> process can run on only 1 CPU core and doesn't know how to make use of
> multi-core systems, the OpenVPN Access Server comes with the ability to launch
> multiple OpenVPN daemons at the same time. Ideally there would be one OpenVPN
> daemon for every CPU core. But there's more involved. To make it possible for
> OpenVPN clients to establish a connection via the UDP protocol and via the TCP
> protocol, there are additional OpenVPN daemons necessary. In the case of the
> OpenVPN Access Server this means we launch 1 TCP and 1 UDP daemon per CPU core.

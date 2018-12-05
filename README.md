# Tab Magic

Tab Magic is a shell-alias building tool for *NIX operating systems.

It utilizes doing a DNS zone-transfer and building bash aliases based on A, CNAME, and TXT records.

You will need to create a config.json with your information in it: 

```json
{
    "resolver": "172.19.0.5",
    "resolverPort": 53,
    "domains": [ "kc.linuxrocker.com"],
    "jumpHost": "jump01.kc.linuxrocker.com",
    "splitString": ".linuxrocker",
    "windowsGeometry": "1600x900"
}
```

For Tab_Magic to do proper host detection, The following TXT record options are available:

* SSH_PORT - SSH Port will allow you to specify a custom SSH Port for the remote host. Valid: 1-65536
* OS_FAMILY - Can be either "ESXi", "Ubiquiti", or "Windows". If ESXi, it will log in as root with no sudo opts. If Ubiquiti, it will not use sudo. If Windows, it will use the rdesktop command line tool.
* REMOTE_USER - You will be able to specify a custom username for the remote host.
version: 2
ethernets:
    ens3:
        dhcp4: no
        addresses:
            - ${ip_address}/${ip_block}
        gateway4: ${gateway}
        nameservers:
            addresses:
                - 8.8.8.8
                - 8.8.4.4

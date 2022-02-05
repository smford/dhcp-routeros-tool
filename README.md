# dhcp-routeros-tool

A simple tool to pull DHCP lease information out of Microtik RoutOS

## Example Output
**Default**
```
COMMENT     ADDRESS       MAC-ADDRESS       CLIENT-ID           ADDRESS-LISTS SERVER DHCP-OPTION STATUS  LAST-SEEN      HOST-NAME     RADIUS DYNAMIC BLOCKED DISABLED 
PC1         172.28.10.10  46:F5:20:E1:57:21                                   Guest              bound   1m33s          PC1           false  false   false   false    
Laptop1     172.28.10.51  7A:58:01:51:ED:ED 1:7a:58:1:51:ed:ed                Guest              bound   1m15s          Laptop1       false  false   false   false    
IP Camera   172.28.10.52  5A:58:34:52:6F:45 1:5a:58:34:52:6f:45               Guest              bound   1m12s          IPCAM         false  false   false   false    
Coffee Pot  172.28.10.53  9C:9C:F1:54:C8:5a                                   Guest              bound   4m43s          ESP_45C823    false  false   false   false    
Laptop2     172.28.10.61  B8:27:D5:B1:12:58                                   Guest              waiting 1w2d10h49m51s  Laptop2       false  false   false   false    
Desktop1    172.28.10.62  7A:27:EB:9B:7D:EF                                   Guest              waiting 11w3d15h34m33s desktop1      false  false   false   false    
Iphone1     172.28.10.81  E5:E2:7A:F1:BE:37 1:e5:e2:7a:f1:be:37               Guest              bound   20s            Marys Iphone  false  false   false   false    
Iphone2     172.28.10.55  50:C5:64:41:1F:65 1:50:c5:64:41:1f:65               Guest              bound   2m12s          Johns Iphone  false  false   false   false
```

**Simple**
```
ADDRESS       MAC-ADDRESS       STATUS  LAST-SEEN     HOST-NAME        COMMENT
172.28.10.10  40:F5:20:01:57:EA bound   2m2s          win7             Marys Laptop
172.28.10.53  9C:9C:1F:45:C8:A7 bound   2m15s         ringdoorbell     Ring DoorBell
172.28.10.61  B8:27:EB:B1:6A:58 waiting 1w3d14h27m23s roku             Roku
172.28.10.62  CE:27:EB:9B:C6:DE waiting 11w4d19h12m5s android-1        Johns Phone
172.28.10.63  34:27:EB:B2:66:BE waiting 8w1d15h59m49s appletv          Lounge ATV
172.28.10.64  56:27:EB:C3:00:EF bound   4m36s         octopi           Ender 5
172.28.10.137 78:F7:BE:D4:22:CA bound   4m8s          android-323asd0
```

## Installation

You can install a few ways:

1. Download the binary for your OS from https://github.com/smford/dhcp-routeros-tool/releases
1. or use `go install`
   ```
   go install -v github.com/smford/dhcp-routeros-tool@latest
   ```
1. or clone the git repo and build
   ```
   git clone git@github.com:smford/dhcp-routeros-tool.git
   cd dhcp-routeros-tool
   go get -v
   go build
   ```

### Configuring RouterOS
1. Configure the Service
1. Configure the User

### config.yaml Configuration
```
address: "192.168.10.1:8729"
usetls: true
async: true
username: "username"
password: "password"
padding: 2
simpledisplay: "address,mac-address,client-id,server,status,last-seen,host-name,disabled"
defaultdisplay: "comment,address,mac-address,client-id,address-lists,server,dhcp-option,status,last-seen,host-name,radius,dynamic,blocked,disabled"
```

| Setting | Default | Details |
|:--|:--|:--|
| `address` | none | The IP address or hostname of mikrotik router |
| `usetls` | true | Use TLS to connect to router |
| `async` | true | Execute commands asyncronously |
| `username` | none | Username |
| `password` | none | Password |
| `padding` | 2 | Number of spaces between columns |
| `simpledisplay` | address,mac-address,client-id,server,status,last-seen,host-name,disabled | Columns to display when --simple argument passed |
| `default` | comment,address,mac-address,client-id,address-lists,server,dhcp-option,status,last-seen,host-name,radius,dynamic,blocked,disabled | Columns to display by default |

#### DHCP Columns
```
comment
address
mac-address
client-id
address-lists
server
dhcp-option
status
last-seen
host-name
radius
dynamic
blocked
disabled
```

### SSL Configuration

If you are using TLS/SSL on your router you may need to install the certificate authority on the device running dhcp-routeros-tool.  If you are using a self signed certificate you can either place it in your operating systems central location; or alternatively store in a directory of your choosing and tell dhcp-routeros-tool (via environment variables) where to find the certificate authority.

If storing your self signed certificate in a directory of your choosing you can do either of:
- `export SSL_CERT_FILE="/path/to/your/cert.pem`
- `export SSL_CERT_DIR=/path/to/your/`

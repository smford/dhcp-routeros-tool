# dhcp-routeros-tool

A simple tool to pull DHCP lease information out of Microtik RouterOS

## Installation

### Installing dhcp-routeros-tool

### Configuring RouterOS
1. Configure the Service
1. Configure the User

### config.yaml Configuration


### SSL Configuration

If you are using TLS/SSL on your router you may need to install the certificate authority on the device running dhcp-routeros-tool.  If you are using a self signed certificate you can either place it in your operating systems central location; or alternatively store in a directory of your choosing and tell dhcp-routeros-tool (via environment variables) where to find the certificate authority.

If storing your self signed certificate in a directory of your choosing you can do either of:
- `export SSL_CERT_FILE="/path/to/your/cert.pem`
- `export SSL_CERT_DIR=/path/to/your/`

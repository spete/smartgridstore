#!/bin/bash

cat >devmachine-localhost.ca.crt << EOF
-----BEGIN CERTIFICATE-----
MIIDnzCCAoegAwIBAgIJAPRK/R59kVNbMA0GCSqGSIb3DQEBCwUAMGYxCzAJBgNV
BAYTAlVTMREwDwYDVQQIDAhEZXZUb3BpYTEXMBUGA1UEBwwOTXlMb2NhbE1hY2hp
bmUxGTAXBgNVBAoMEEJUckRCIERldk1hY2hpbmUxEDAOBgNVBAMMB1JPT1QgQ0Ew
HhcNMTgwNjA0MjM1NjI3WhcNMjgwNjAxMjM1NjI3WjBmMQswCQYDVQQGEwJVUzER
MA8GA1UECAwIRGV2VG9waWExFzAVBgNVBAcMDk15TG9jYWxNYWNoaW5lMRkwFwYD
VQQKDBBCVHJEQiBEZXZNYWNoaW5lMRAwDgYDVQQDDAdST09UIENBMIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwC/3taEZyDyatrH6ESGqVU4Kb4TAjWdZ
0b0EscohLhTsAxUPufUYKsuH6AroUqlidYFnLsiHW2AqIyQhX3agYKB9wCcfm3Qb
D6n/Z8zUULJ/YFpTTaUPY9ex9EvnlINc+sXwyOEojsL+Yq20Vz4NvjqR7GtE1GEr
fv36Y+hd2nanU2dlVcsv1dO7zKw+ttPRi6qTYXQ4WT+nz2Mk1xWhK8NqqPsQWL3X
p0lMaJXFzsQpr84ZbkfRjczLlO2gft4b96iBbC6lseuhWniE69Z48ZtZQR73vDCH
ZCvscG7ZuXcaW0r4nLsNaVS1WjuC3baD6SWKFiQvVNt49NXf4CxG1QIDAQABo1Aw
TjAdBgNVHQ4EFgQUaXMsIjbpaboHIT41m01rJyDA+r8wHwYDVR0jBBgwFoAUaXMs
IjbpaboHIT41m01rJyDA+r8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC
AQEAiyOdswyRDCmVlOICOjQGHNLDEOI4gbs6tN8KvfpJ5uZp7xorNKkV3peaTVpW
az97mrKRUdAZYEDoBY9oAvxX7enYheszG3vfgc1BXf+i3YdNKCUexhBZtetvkBFa
rg6Wi0nambU3tuZSmuUUCfskQVDSCRgnTsyMPYHfF7lOa8qJaqSpFYmhQAgdPMpe
piFhLHrAxqzMZXCfRSMP3Ji3lnjnXo2NwKu/GqIEXAIHxYZh7QnTA022Ws1FB0E4
alL9IOoTaR7TQRYf4TOATCw1m4nYe61pfXB1J/CQ0cLVPS5qnbs9m8heynn7066q
HRNq0S1ryc6KVDR6FaxnN2bcuQ==
-----END CERTIFICATE-----
EOF

docker exec -i devmachine-etcd /bin/bash << EOSCRIPT

export ETCDCTL_API=3

etcdctl --endpoints 172.29.0.20:2379 put api/hardcoded_pub << EOF
-----BEGIN CERTIFICATE-----
MIIDqzCCApOgAwIBAgIJAI2ZCK3cIv9tMA0GCSqGSIb3DQEBCwUAMGYxCzAJBgNV
BAYTAlVTMREwDwYDVQQIDAhEZXZUb3BpYTEXMBUGA1UEBwwOTXlMb2NhbE1hY2hp
bmUxGTAXBgNVBAoMEEJUckRCIERldk1hY2hpbmUxEDAOBgNVBAMMB1JPT1QgQ0Ew
HhcNMTgwNjA1MDAwMTA1WhcNMjgwNjAyMDAwMTA1WjB9MQswCQYDVQQGEwJVUzER
MA8GA1UECAwIRGV2VG9waWExFzAVBgNVBAcMDk15TG9jYWxNYWNoaW5lMRkwFwYD
VQQKDBBCVHJEQiBEZXZNYWNoaW5lMRMwEQYDVQQLDApEZXZNYWNoaW5lMRIwEAYD
VQQDDAlsb2NhbGhvc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCy
/w3AztG1VRYnGxYowEOHhlyC0Z5eu5avFozftsKv16hB9WkrkQyyicbv5+bSfeiJ
0EIEDXk7MMkuqhtkk3SOqgusLQUK05SCVKjtoBlagujds1W/X1E6kezjjSmOtnuc
+7uxo3VG4UNxphjKpXNCegWa0gw3A3QaF0ZCub7XdAPSbpwNaGB9TyB7HEP1GctB
0rqVnP7dG46bMSEfBs3UUHsUBEfv4akpf5i8Rk+igW2BUxHtbeQ+sLGvzpzAdLJF
kGceOuIUDkt4ZBCTv7mhx4u+W+A9eXEvGBEglDKPI//+gsdpU279QWwVSjhcxySW
iz61k43ao27CO/A0w0JbAgMBAAGjRTBDMAsGA1UdDwQEAwIEMDATBgNVHSUEDDAK
BggrBgEFBQcDATAfBgNVHREEGDAWgglsb2NhbGhvc3SCCTEyNy4wLjAuMTANBgkq
hkiG9w0BAQsFAAOCAQEAU40YmF7QcW8Z4UG9JRRHmMz1u4k8Vwb9pbsAeK1Yb6NO
PCtvv0s3jF5ZEDdGMVwmOvD4VloSrs4TPtA4GE49KDhHJ7eFNZSeCA3bin52vxhH
bSyqOnaMcDNqAY/nPwh262TCuCIWNG7MOXTWDYW+dXZ0uIqacrdx+xXBh5X3wFNA
HrRCbMsy4aWGvwN18uHcMPDC3gclk5RB+MbvmzIEyer1po5Ekiu4V1CHLGbVXdOE
o6Kir1OpmqtrxjhgtvwT4kC5TycE0FOAsh/v69dEIXT84pH7zvEYMBRNlmDbuz1u
oHuLW0nt3rUNKzpHIOYd8xJXhvI5N3ChfAh4MuBSKA==
-----END CERTIFICATE-----
EOF

etcdctl --endpoints 172.29.0.20:2379 put api/hardcoded_priv << EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAsv8NwM7RtVUWJxsWKMBDh4ZcgtGeXruWrxaM37bCr9eoQfVp
K5EMsonG7+fm0n3oidBCBA15OzDJLqobZJN0jqoLrC0FCtOUglSo7aAZWoLo3bNV
v19ROpHs440pjrZ7nPu7saN1RuFDcaYYyqVzQnoFmtIMNwN0GhdGQrm+13QD0m6c
DWhgfU8gexxD9RnLQdK6lZz+3RuOmzEhHwbN1FB7FARH7+GpKX+YvEZPooFtgVMR
7W3kPrCxr86cwHSyRZBnHjriFA5LeGQQk7+5oceLvlvgPXlxLxgRIJQyjyP//oLH
aVNu/UFsFUo4XMcklos+tZON2qNuwjvwNMNCWwIDAQABAoIBAQCFrVVoGQ0kj1br
/Z6e8Hd+TynnyQStWws912l96c+b40MVf2H712fULnET2ezYZo+z3IRw4l8XhWe6
IfAzPKxfnz74ZubNHxZZ/z/ptxc4MWwXpkbzlQvk4fY4OsQ+gKnwo0+Zaqm7NaBY
z+LT9JwPmXF2HkhhDYM3uQoP6whLfmiijZxTmigM/HQySiDYfXrCfcoh80xZZ9nh
tdOYpsww9/PQ1XwuFvW+spX2FxTyriRrv3XCpDuJBZ4gzsppS7Hn2biiNEl0utj7
FMyxvXVoRGPGjCPKLZqOPAdffD6grnC8QR2OS7y+veNNmlthFsy3oUw0A8Dhry/Q
gE4/GdUBAoGBAOUgePAMR5lYIQ1rU+0zJywSEgABtHX+XIkIv87TJYZh0GGnhvDI
9FUfCTvIQznoZeC59yL0/fxFWKGb4y15tOG5a9/a4NCP76Qlc77EybRFvF7kFNOr
R6EOVNEcCWk3WtcF+ec/kgcXJ+ZiZDlH2YMf2kd/FOcWQXPdTmryVaJzAoGBAMf9
acCVzokt3CNjJOF/bMneB3Zfs2bpBIyCiyc0j324Rimc+s238OgjW48JXimtXvcC
luzeBfejTqcEM9lu+JRw95XuAbeWrQSawRRDkXmSTjrFcJZ0em3RJDbrVVHv3HwE
PcG0Nax3vnWnh3ddTnf2fhBKlHffqhxe8YtJwB55AoGAJRAIoAPMfSCFUC9hRwg0
OOu/X6Lm9wMrIrt4k1MSSdd+pp07ta074J0BmFr/jNlryVsrf8sTXoA1IwcdS1jZ
in281lwIa5Qs1md8fopEelWhb9QDDm4xSvsPezfGye87UXbVArQEwgLb4GdgAOf/
Zjd7zn7e+bZe5ggRTDlg4sMCgYBQ77sLyNUEaX3tCGPVqvdBH01P191IKcfAgdiF
Ll1gGOK0Vqad+PJTUHPuiHEGVvbW6sJf7F7n4LylFStStPl/QdTBZchmH2G4OlUn
uUy3scFdQaiWC1+87+ZDH6yw82z8985yhVcvjGqVPQ6y/R0TqbtNJpG9jdRPlREW
OOu6qQKBgHfVG3gtHt1nLPgjErQGy2PeMrDZn9Ac5bOuXZ5WlR5G7lLzcK4Ew7rR
SjMiAAkzJMUfacx7USkOfONx6VT1ZfQeUbqW3v7ZqA9C6vt+oCl8ji60HBEZ6Bx6
llPVnHF+J/lkQsB8FCCKd8tOBZuHE0uuwlmstgg4/Ce5DFqte3uo
-----END RSA PRIVATE KEY-----
EOF

etcdctl --endpoints 172.29.0.20:2379 put api/certsrc hardcoded
EOSCRIPT

set -x
docker restart devmachine-console
set +x
echo "We have created a 'devmachine-localhost.ca.crt' file. Add this to your browser's Certificate Authorities and everything will just work"
echo "For non-browser use, copy the pem file to /usr/local/share/ca-certificates/btrdb/ and ensure it is world-readable. Then run sudo update-ca-certificates"
echo "Ensure that the API endpoint is given as localhost:4411. Using an IP address will not work"
echo "For python you also need to run export BTRDB_CA_BUNDLE=/etc/ssl/certs/ca-certificates.crt to make grpc use system ssl certificates"
FROM docker.io/library/golang:1.20.2-bullseye as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -x . 


FROM debian:11.6
WORKDIR /
COPY --from=build /app/cshell /cshell
RUN set -x \
 && dpkg --add-architecture i386 \
 && apt-get update -y\
 && apt-get install -y \
    bzip2 kmod expect iptables net-tools iputils-ping iproute2 curl wget bash \
    libstdc++5:i386 libpam0g:i386 libx11-6:i386 \
 && curl -sfLko /tmp/snx_install.sh 'https://ext.vpn.vtb.ru/sslvpn/SNX/INSTALL/snx_install.sh' \
 && bash /tmp/snx_install.sh && rm -rf /tmp/snx_install.sh \
 && apt install --yes opensc pcscd openssl libengine-pkcs11-openssl gnutls-bin \
 && curl -sfLko /tmp/librtpkcs11ecp.deb http://download.rutoken.ru/Rutoken/PKCS11Lib/2.7.1.0/Linux/x64/librtpkcs11ecp_2.7.1.0-1_amd64.deb \
 && dpkg -i /tmp/librtpkcs11ecp.deb && rm -rf /tmp/librtpkcs11ecp.deb

ENV CSHELL_PKCS11_LIB="/usr/lib/librtpkcs11ecp.so"
ENV CSHELL_PKCS11_ID=""
ENV CSHELL_PKCS11_PIN=""
ENV CSHELL_PKCS11_SELECTOR="example@localhost.local"
ENV CSHELL_SNX_PREFIX="/"
ENV CSHELL_SNX_GATEWAY="vpn.localhost.local"
ENV CSHELL_SNX_REALM=""

CMD ["/cshell"]
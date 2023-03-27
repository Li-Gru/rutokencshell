// Copy from https://github.com/francisco-anderson/snx-go
package main

import (
	"math/big"
	"net"
	"os/exec"
	"strconv"

	binary_pack "github.com/roman-kachanovsky/go-binary-pack/binary-pack"
)

type SNX struct {
	Params  map[string]string
	Debug   bool
	SnxPath string
	info    []byte
}

func (extender *SNX) generateSNXInfo() {
	params := extender.Params
	gwIP, err := net.LookupHost(params["host_name"])
	iferr(err)

	bp := new(binary_pack.BinaryPack)

	ip := net.ParseIP(gwIP[0])
	ipv4 := big.NewInt(0)
	ipv4.SetBytes(ip.To4())
	tmp := ip.To4()

	hwData, err := bp.UnPack([]string{"I"}, []byte{tmp[3], tmp[2], tmp[1], tmp[0]})
	iferr(err)

	gwInt := hwData[0].(int)

	magic := string([]byte{0x13, 0x11, 0x00, 0x00})
	length := 0x3d0

	port, err := strconv.Atoi(params["port"])
	iferr(err)

	format := []string{"4s", "L", "L", "64s", "L", "6s", "256s", "256s", "128s", "256s", "H"}

	values := []interface{}{
		magic,
		length,
		gwInt,
		params["host_name"],
		port,
		string([]byte{0}),
		params["server_cn"],
		params["user_name"],
		params["password"],
		params["server_fingerprint"],
		1,
	}

	data, err := bp.Pack(format, values)
	iferr(err)

	extender.info = data
}

func (extender *SNX) callSNX() {
	snxCmd := exec.Command(extender.SnxPath, "-Z")

	_, err := snxCmd.Output()
	iferr(err)

	connection, err := net.Dial("tcp", "localhost:7776")
	iferr(err)

	_, err = connection.Write(extender.info)
	iferr(err)

	buffer := make([]byte, 4096)

	_, err = connection.Read(buffer)
	iferr(err)

	connection.Read(buffer) //Block execution

}

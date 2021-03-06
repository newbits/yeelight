package yeelight

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Discover discovers a single device in the local network via SSDP
func Discover() (*Device, error) {
	var err error

	ssdp, _ := net.ResolveUDPAddr("udp4", ssdpAddr)
	c, _ := net.ListenPacket("udp4", ":0")
	socket := c.(*net.UDPConn)
	socket.WriteToUDP([]byte(discoverMSG), ssdp)
	socket.SetReadDeadline(time.Now().Add(timeout))

	rsBuf := make([]byte, 1024)
	size, _, err := socket.ReadFromUDP(rsBuf)

	defer socket.Close()

	if err != nil {
		return nil, errors.New("no devices found")
	}

	rs := rsBuf[0:size]
	addr := parseAddr(string(rs))

	fmt.Printf("Device with ip %s found\n", addr)

	return New(addr), nil
}

// DiscoverMany tries to discover many devices at once
func DiscoverMany() ([]*Device, error) {
	var devices []*Device

	skipped := 0

	for {
		device, err := Discover()
		if err != nil {
			return devices, err
		}

		newDevice := true
		for i := range devices {
			if devices[i].Address == device.Address {
				newDevice = false
			}
		}

		if newDevice {
			devices = append(devices, device)
			continue
		}

		skipped++

		if skipped >= 3 {
			break
		}
	}

	return devices, nil
}

// New creates new device instance for address provided
func New(address string) *Device {
	return &Device{
		Address: address,
		Random:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Listen connects to device and listens for NOTIFICATION events
func (d *Device) Listen() (<-chan *Notification, chan<- struct{}, error) {
	var err error
	notifCh := make(chan *Notification)
	done := make(chan struct{}, 1)

	conn, err := net.DialTimeout("tcp", d.Address, time.Second*3)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot connect to %s. %s", d.Address, err)
	}

	fmt.Println("Connection established")
	go func(c net.Conn) {
		//make sure connection is closed when method returns
		defer closeConnection(conn)

		connReader := bufio.NewReader(c)
		for {
			select {
			case <-done:
				return
			default:
				data, err := connReader.ReadString('\n')
				if nil == err {
					var rs Notification
					fmt.Println(data)
					json.Unmarshal([]byte(data), &rs)
					select {
					case notifCh <- &rs:
					default:
						fmt.Println("Channel is full")
					}
				}
			}

		}

	}(conn)

	return notifCh, done, nil
}

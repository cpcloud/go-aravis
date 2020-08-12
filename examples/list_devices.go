package main

import (
	"log"

	aravis "github.com/thinkski/go-aravis"
)

func main() {
	aravis.UpdateDeviceList()
	n, err := aravis.GetNumDevices()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Devices:", n)
	for i := uint(0); i < n; i++ {
        device_id, _ := aravis.GetDeviceId(i)
		log.Println(device_id)

        physical_id, _ := aravis.GetDevicePhysicalId(i)
        log.Println(physical_id)

        model, _ := aravis.GetDeviceModel(i)
        log.Println(model)

        serial_nbr, _ := aravis.GetDeviceSerialNbr(i)
        log.Println(serial_nbr)

        vendor, _ := aravis.GetDeviceVendor(i)
        log.Println(vendor)

        address, _ := aravis.GetDeviceAddress(i)
		log.Println(address)

        protocol, _ := aravis.GetDeviceProtocol(i)
		log.Println(protocol)

        interface_id, _ := aravis.GetInterfaceId(i)
        log.Println(interface_id)
	}
}

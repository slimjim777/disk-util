package service

import (
	"github.com/godbus/dbus"
	"log"
	"path"
)

const (
	udisksBlockFormat     = "org.freedesktop.DBus.UDisks2.Block.Format"
	udisksPartitionCreate = "org.freedesktop.UDisks2.PartitionTable.CreatePartition"
)

// FormatPartition use udisks to format a partition
func (db *DBus) FormatPartition(device, format string) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	call := obj.Call(udisksBlockFormat, 0, format, map[string]dbus.Variant{}, false)
	return call.Err
}

// CreatePartition creates a new partition on a block device
func (db *DBus) CreatePartition(device, format, name string, offset, size uint64) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	var resp map[string]dbus.Variant
	err := obj.Call(udisksPartitionCreate, 0, offset, size, format, name, map[string]dbus.Variant{}).Store(&resp)
	log.Println(resp)
	return err
}

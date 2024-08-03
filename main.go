package main

import (
	"fmt"
	"reflect"
	"syscall"
)

type Disk struct {
	Statfs    syscall.Statfs_t
	Total     uint64
	Used      uint64
	Free      uint64
	Available uint64
}

func (d *Disk) calculateSpace() error {
	err := syscall.Statfs("/", &d.Statfs)
	if err != nil {
		return err
	}

	d.Total = d.Statfs.Blocks * uint64(d.Statfs.Bsize)
	d.Free = d.Statfs.Bfree * uint64(d.Statfs.Bsize)
	d.Available = d.Statfs.Bavail * uint64(d.Statfs.Bsize)
	d.Used = d.Total - d.Free

	return nil
}

func (d *Disk) logDiskSpaceInfo() {
	fmt.Printf("Disk Info:\nTotal Space: %0.2f GB\nUsed Space: %0.2f GB\nFree Space: %0.2f GB\nAvailable Space: %0.2f GB",
		byteToGigabyte(d.Total),
		byteToGigabyte(d.Used),
		byteToGigabyte(d.Free),
		byteToGigabyte(d.Available),
	)
}

func (d *Disk) logDiskInfo() {
	reflectValue := reflect.ValueOf(d.Statfs)
    typeOfExample := reflectValue.Type()

    for i := 0; i < reflectValue.NumField(); i++ {
        field := reflectValue.Field(i)
        fmt.Printf("%-13s -> %v \n", typeOfExample.Field(i).Name, field.Interface())
    }
}

func byteToGigabyte(bytes uint64) float64 {
	return float64(bytes) / (1e+9)
}

func main() {
	disk := Disk{}
	err := disk.calculateSpace()
	if err != nil {
		fmt.Println(err.Error())
	}
	disk.logDiskSpaceInfo()
	// disk.logDiskInfo()
}

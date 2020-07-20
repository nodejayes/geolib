package raster

import (
	"bytes"
	"encoding/binary"
	"sync"
)

func WriteRasterWkb(raster Raster) ([]byte, error) {
	wkb := bytes.NewBuffer([]byte{})
	endiannessValue := raster.GetByteEndianness()

	endianErr := binary.Write(wkb, binary.LittleEndian, raster.Endianness)
	if endianErr != nil {
		return wkb.Bytes(), endianErr
	}

	header, err := writeHeader(wkb, endiannessValue, raster)

	if err != nil {
		return wkb.Bytes(), err
	}

	tmpWkb := make([]*bytes.Buffer, len(raster.Bands))
	in := make(chan int, len(raster.Bands))
	var wg sync.WaitGroup
	for i := 0; i < raster.Workers; i++ {
		wg.Add(1)
		go func() {
			for bandIdx := range in {
				tmpWkb[bandIdx] = bytes.NewBuffer([]byte{})
				_ = writeRasterBand(tmpWkb[bandIdx], endiannessValue, raster.Bands[bandIdx], header)
			}
			wg.Done()
		}()
	}
	for idx := 0; idx < len(raster.Bands); idx++ {
		in <- idx
	}
	close(in)
	wg.Wait()

	for i := 0; i < len(raster.Bands); i++ {
		wkb.Write(tmpWkb[i].Bytes())
	}

	return wkb.Bytes(), nil
}

func writeHeader(wkb *bytes.Buffer, endiannessValue binary.ByteOrder, raster Raster) (Header, error) {
	var header Header
	var header2 Header

	header.Version = raster.Version
	header.Bands = uint16(len(raster.Bands))
	header.ScaleX = raster.ScaleX
	header.ScaleY = raster.ScaleY
	header.IpX = raster.IpX
	header.IpY = raster.IpY
	header.SkewX = raster.SkewX
	header.SkewY = raster.SkewY
	header.Srid = raster.Srid
	header.Width = raster.Width
	header.Height = raster.Height

	err := binary.Write(wkb, endiannessValue, header)

	if err != nil {
		return header, err
	}

	wkb2 := bytes.NewReader(wkb.Bytes())
	err = binary.Read(wkb2, endiannessValue, &header2)
	return header, nil
}

func writeRasterBand(wkb *bytes.Buffer, endiannessValue binary.ByteOrder, rasterBand Band, header Header) error {
	err := binary.Write(wkb, endiannessValue, rasterBand.BandHeader)
	if err != nil {
		return err
	}

	err = binary.Write(wkb, endiannessValue, rasterBand.NoData)
	if err != nil {
		return err
	}

	for i := 0; i < int(header.Height); i++ {
		for j := 0; j < int(header.Width); j++ {
			err = binary.Write(wkb, endiannessValue, rasterBand.Data[i][j])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

package raster

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

func ReadRasterWkb(b []byte, workers int) (Raster, error) {
	wkb := bytes.NewReader(b)
	raster := Raster{}
	endiannesValue, endianessErr := readEndiannes(wkb)
	if endianessErr != nil {
		return raster, endianessErr
	}

	var endiannes binary.ByteOrder

	if endiannesValue == 0 {
		endiannes = binary.BigEndian
	} else if endiannesValue == 1 {
		endiannes = binary.LittleEndian
	}

	raster.Endianness = endiannesValue
	raster.Workers = workers

	return readHeader(wkb, endiannes, raster)
}

// Determine the endiannes of the definitions
//
// +---------------+-------------+------------------------------+
// | endiannes     | byte        | 1:ndr/little endian          |
// |               |             | 0:xdr/big endian             |
// +---------------+-------------+------------------------------+
func readEndiannes(wkb io.Reader) (uint8, error) {
	var endiannesValue uint8
	err := binary.Read(wkb, binary.LittleEndian, &endiannesValue)
	return endiannesValue, err
}

// Read the definitions header data.
//
// +---------------+-------------+------------------------------+
// | version       | uint16      | format version (0 for this   |
// |               |             | structure)                   |
// +---------------+-------------+------------------------------+
// | nBands        | uint16      | Number of bands              |
// +---------------+-------------+------------------------------+
// | scaleX        | float64     | pixel width                  |
// |               |             | in geographical units        |
// +---------------+-------------+------------------------------+
// | scaleY        | float64     | pixel height                 |
// |               |             | in geographical units        |
// +---------------+-------------+------------------------------+
// | ipX           | float64     | X ordinate of upper-left     |
// |               |             | pixel's upper-left corner    |
// |               |             | in geographical units        |
// +---------------+-------------+------------------------------+
// | ipY           | float64     | Y ordinate of upper-left     |
// |               |             | pixel's upper-left corner    |
// |               |             | in geographical units        |
// +---------------+-------------+------------------------------+
// | skewX         | float64     | rotation about Y-axis        |
// +---------------+-------------+------------------------------+
// | skewY         | float64     | rotation about X-axis        |
// +---------------+-------------+------------------------------+
// | srid          | int32       | Spatial reference id         |
// +---------------+-------------+------------------------------+
// | width         | uint16      | number of pixel columns      |
// +---------------+-------------+------------------------------+
// | height        | uint16      | number of pixel rows         |
// +---------------+-------------+------------------------------+
func readHeader(wkb io.Reader, endiannes binary.ByteOrder, raster Raster) (Raster, error) {
	var header Header
	err := binary.Read(wkb, endiannes, &header)

	if err != nil {
		return raster, err
	}

	raster.Version = header.Version
	raster.ScaleX = header.ScaleX
	raster.ScaleY = header.ScaleY
	raster.IpX = header.IpX
	raster.MidIpX = header.IpX + 1.5
	raster.IpY = header.IpY
	raster.MidIpY = header.IpY - 1.5
	raster.SkewX = header.SkewX
	raster.SkewY = header.SkewY
	raster.Srid = header.Srid
	raster.Width = header.Width
	raster.Height = header.Height

	for i := 0; i < int(header.Bands); i++ {
		raster, err = readBandHeaderData(wkb, endiannes, raster, header)
		if err != nil {
			return raster, err
		}
	}
	return raster, nil
}

// Read band header data
//
// +---------------+--------------+-----------------------------------+
// | isOffline     | 1bit         | If true, data is to be found      |
// |               |              | on the filesystem, trought the    |
// |               |              | path specified in RASTERDATA      |
// +---------------+--------------+-----------------------------------+
// | hasNodataValue| 1bit         | If true, stored nodata value is   |
// |               |              | a true nodata value. Otherwise    |
// |               |              | the value stored as a nodata      |
// |               |              | value should be ignored.          |
// +---------------+--------------+-----------------------------------+
// | isNodataValue | 1bit         | If true, all the values of the    |
// |               |              | band are expected to be nodata    |
// |               |              | values. This is a dirty flag.     |
// |               |              | To set the flag to its real value |
// |               |              | the function st_bandisnodata must |
// |               |              | must be called for the band with  |
// |               |              | 'TRUE' as last argument.          |
// +---------------+--------------+-----------------------------------+
// | reserved      | 1bit         | unused in this version            |
// +---------------+--------------+-----------------------------------+
// | pixtype       | 4bits        | 0: 1-bit boolean                  |
// |               |              | 1: 2-bit unsigned integer         |
// |               |              | 2: 4-bit unsigned integer         |
// |               |              | 3: 8-bit signed integer           |
// |               |              | 4: 8-bit unsigned integer         |
// |               |              | 5: 16-bit signed integer          |
// |               |              | 6: 16-bit unsigned signed integer |
// |               |              | 7: 32-bit signed integer          |
// |               |              | 8: 32-bit unsigned signed integer |
// |               |              | 9: 32-bit float                   |
// |               |              | 10: 64-bit float                  |
// +---------------+--------------+-----------------------------------+
//
// Requires reading a single byte, and splitting the bits into the
// header attributes
func readBandHeaderData(wkb io.Reader, endiannes binary.ByteOrder, raster Raster, header Header) (Raster, error) {
	band := Band{}
	var bandheader uint8
	err := binary.Read(wkb, endiannes, &bandheader)

	if err != nil {
		return raster, err
	}

	band.BandHeader = bandheader
	band.IsOffline = (int(bandheader) & 128) != 0
	band.HasNoDataValue = (int(bandheader) & 64) != 0
	band.IsNoDataValue = (int(bandheader) & 32) != 0

	// Read the pixel type
	pixType := (int(bandheader) & 15) - 1

	// +---------------+--------------+-----------------------------------+
	// | nodata        | 1 to 8 bytes | Nodata value                      |
	// |               | depending on |                                   |
	// |               | pixtype [1]  |                                   |
	// +---------------+--------------+-----------------------------------+

	// Read the nodata value
	noData, err := readFloat64OfType(wkb, endiannes, pixType)

	if err != nil {
		return raster, err
	}

	band.NoData = noData

	// Read the pixel values: width * height * size
	//
	// +---------------+--------------+-----------------------------------+
	// | pix[w*h]      | 1 to 8 bytes | Pixels values, row after row,     |
	// |               | depending on | so pix[0] is upper-left, pix[w-1] |
	// |               | pixtype [1]  | is upper-right.                   |
	// |               |              |                                   |
	// |               |              | As for endiannes, it is specified |
	// |               |              | at the start of WKB, and implicit |
	// |               |              | up to 8bits (bit-order is most    |
	// |               |              | significant first)                |
	// |               |              |                                   |
	// +---------------+--------------+-----------------------------------+

	for i := 0; i < int(header.Height); i++ {
		var row []float64

		for i := 0; i < int(header.Width); i++ {
			value, err := readFloat64OfType(wkb, endiannes, pixType)

			if err != nil {
				return raster, err
			}

			row = append(row, value)
		}

		band.Data = append(band.Data, row)
	}

	raster.Bands = append(raster.Bands, band)
	return raster, nil
}

func readFloat64OfType(reader io.Reader, endiannes binary.ByteOrder, valueType int) (float64, error) {
	switch valueType {
	case 0, 1, 2, 4:
		var value uint8

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 3:
		var value int8

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 5:
		var value int16

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 6:
		var value uint16

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 7:
		var value int32

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 8:
		var value uint32

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 9:
		var value float32

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	case 10:
		var value float64

		err := binary.Read(reader, endiannes, &value)

		if err != nil {
			return 0, err
		}

		return float64(value), nil
	}

	return 0, errors.New("Unknown value type")
}

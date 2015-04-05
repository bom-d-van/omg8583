package omg8583

import (
	"fmt"
	"strconv"
)

func unpack(msg string, id int) (*iso, error) {
	p := t[id]
	switch p.Type {
	case "a", "an", "ans":
		println(msg)
		println(p.Length)
		return &iso{Data: msg[0 : p.Length*2], RestData: msg[p.Length*2:]}, nil
	case "n":
		return &iso{Data: msg[0:p.Length], RestData: msg[p.Length:]}, nil
	case "lln":
		if length, err := strconv.ParseInt(msg[0:2], 10, 64); err == nil {
			if length > int64(p.Length) {
				length = int64(p.Length)
			}
			return &iso{Data: string(msg[2 : length+2]),
				Chunk:    string(msg[0 : length+2]),
				RestData: string(msg[length+2:])}, nil
		} else {
			return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
		}
	case "lllan":
		if length, err := strconv.ParseInt(msg[0:3], 10, 64); err == nil {
			if length > int64(p.Length) {
				length = int64(p.Length)
			}
			return &iso{
				Data:     string(msg[3 : length+3]),
				Chunk:    string(msg[0 : length+3]),
				RestData: string(msg[length+3:]),
			}, nil
		} else {
			return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
		}
	case "b":
		var data, bitmap string
		chunk := msg[0:p.Length]
		msg = msg[p.Length:]
		var ckb = 1
		for ckb > 0 {
			var chunkBitmap string
			for _, code := range chunk {
				if n, err := strconv.ParseInt(string(code), 16, 64); err == nil {
					if n == 0 {
						chunkBitmap += "0000"
					} else {
						code := strconv.FormatInt(n, 2)
						for len(code) < 4 {
							code = "0" + code
						}
						chunkBitmap += code
					}
				} else {
					return nil, fmt.Errorf("error %v\npacking data from bit %d\npackager: %v", err, id, t[id])
				}
			}
			for len(chunkBitmap) < (p.Length * 4) {
				chunkBitmap = "0" + chunkBitmap
			}
			ckb, _ = strconv.Atoi(string(chunkBitmap[0]))
			data += chunk
			bitmap += chunkBitmap
			chunk = msg[:p.Length]
			msg = msg[p.Length:]
		}
		fieldIds := make([]int, 0)
		for i, _ := range bitmap {
			b, _ := strconv.Atoi(string(bitmap[i]))
			if i > 0 && b == 1 {
				fieldIds = append(fieldIds, i+1)
			}
		}
		fmt.Printf("--> %+v\n", fieldIds)
		return &iso{
			Data:     data,
			Bitmap:   bitmap,
			FieldIds: fieldIds,
			RestData: msg,
		}, nil
	default:
		return nil, fmt.Errorf("error, type %d not implemented\n", id)
	}
}

package imageutils

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

const PNGSignature = "\x89PNG\r\n\x1a\n"

func checkSignature(r io.Reader) error {
	buf := make([]byte, len(PNGSignature))
	_, err := r.Read(buf)
	if err != nil {
		return err
	}

	if string(buf) != PNGSignature {
		return errors.New("not a PNG file")
	}

	return nil
}

func calculateCRC(data []byte) uint32 {
	// Calculate CRC
	checksum := crc32.ChecksumIEEE(data)
	return checksum
}

func readChunk(r io.Reader) (string, []byte, uint32, error) {
	var length uint32
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return "", nil, 0, err
	}

	if length > 1<<24 { // 16MB max to prevent overflow
		return "", nil, 0, errors.New("chunk too large")
	}

	chunkType := make([]byte, 4)
	_, err = r.Read(chunkType)
	if err != nil {
		return "", nil, 0, err
	}

	chunkData := make([]byte, length)
	_, err = r.Read(chunkData)
	if err != nil {
		return "", nil, 0, err
	}

	crc := make([]byte, 4)
	_, err = r.Read(crc)
	if err != nil {
		return "", nil, 0, err
	}

	givenCRC := binary.BigEndian.Uint32(crc)

	return string(chunkType), chunkData, givenCRC, nil
}

func writeChunk(w io.Writer, chunkType string, chunkData []byte) error {
	err := binary.Write(w, binary.BigEndian, uint32(len(chunkData)))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(chunkType))
	if err != nil {
		return err
	}

	_, err = w.Write(chunkData)
	if err != nil {
		return err
	}

	crc := calculateCRC(append([]byte(chunkType), chunkData...))
	return binary.Write(w, binary.BigEndian, crc)
}

func ValidatePNG(r io.Reader) error {

	// Check the PNG signature...
	buf := make([]byte, len(PNGSignature))
	_, err := r.Read(buf)
	if err != nil {
		fmt.Println("Error reading signature:", err)
		return err
	}

	if string(buf) != PNGSignature {
		fmt.Println("Not a PNG file")
		return errors.New("Not a PNG file")
	}

	for {
		chunkType, chunkData, _, err := readChunk(r)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading chunk:", err)
			return err
		}

		if chunkType == "tEXt" {
			str := string(chunkData)
			decodedBytes, err := base64.StdEncoding.DecodeString(str[6:])
			log.Println(str[6:])
			if err != nil {
				fmt.Println("Error decoding base64:", err)
				return err
			}
			var result map[string]interface{}

			// Unmarshal the JSON data
			err = json.Unmarshal([]byte(decodedBytes), &result)
			if err != nil {
				log.Fatal(err)
				return err
			}

			// Reset the position of the file to the beginning after validation
			if seeker, ok := r.(io.Seeker); ok {
				_, err = seeker.Seek(0, 0)
				if err != nil {
					return err
				}
			}
			return nil
		} else {
			fmt.Printf("Chunk type: %s, data: %x\n", chunkType, chunkData)
		}
	}
	return errors.New("no tEXt chunk")
}

func writePNG(file *os.File, byteData []byte) {

	// Create a new file to write the modified PNG to
	outFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Check the PNG signature...
	buf := make([]byte, len(PNGSignature))
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("Error reading signature:", err)
		return
	}

	if string(buf) != PNGSignature {
		fmt.Println("Not a PNG file")
		return
	}

	_, err = outFile.Write(buf)
	if err != nil {
		fmt.Println("Error writing signature:", err)
		return
	}

	for {
		chunkType, chunkData, _, err := readChunk(file)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading chunk:", err)
			return
		}

		if chunkType == "tEXt" {
			// Change chunkData here
			chunkData = append([]byte("chara\x00"), byteData...) // replace with your new data

			err = writeChunk(outFile, chunkType, chunkData)
			if err != nil {
				fmt.Println("Error writing chunk:", err)
				return
			}
		} else {
			err = writeChunk(outFile, chunkType, chunkData)
			if err != nil {
				fmt.Println("Error writing chunk:", err)
				return
			}
		}
	}
}

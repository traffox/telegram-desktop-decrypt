package tdata

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/atilaromero/telegram-desktop-decrypt/decrypt"
)

// TdataSettings reflects the streams contained in the tdata/settings0 file.
type TdataSettings struct {
	Salt      []byte
	Encrypted []byte
}

// ReadTdataSettings opens the tdata/settings0 or tdata/settings1
func ReadTdataSettings(f io.Reader) (TdataSettings, error) {
	result := TdataSettings{}
	tfile, err := ReadTdataFile(f)
	if err != nil {
		return result, fmt.Errorf("could not interpret file, error: %v", err)
	}
	mydata := bytes.NewReader(tfile.Data)
	result.Salt, err = ReadStream(mydata)
	if err != nil {
		return result, fmt.Errorf("could not read salt: %v", err)
	}
	result.Encrypted, err = ReadStream(mydata)
	if err != nil {
		return result, fmt.Errorf("could not read settingsEncrypted: %v", err)
	}
	return result, err
	// fmt.Printf("settingsKey (%d):\n", len(settingsKey))
	// fmt.Println(hex.EncodeToString(settingsKey[:]))
}
func (settings TdataSettings) GetKey(password string) []byte {
	settingsKey := decrypt.CreateLocalKey([]byte(password), settings.Salt)
	return settingsKey
}
func (settings TdataSettings) Decrypt(settingsKey []byte) ([]byte, error) {
	decrypted, err := decrypt.DecryptLocal(settings.Encrypted, settingsKey)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt settings file: %v", err)
	}
	return decrypted, nil
}

func (settings TdataSettings) Print() {
	fmt.Printf("salt\t%s\n", hex.EncodeToString(settings.Salt))
	fmt.Printf("encrypted\t%s\n", hex.EncodeToString(settings.Encrypted))
}
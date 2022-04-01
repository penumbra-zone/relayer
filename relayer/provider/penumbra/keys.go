package penumbra

import "github.com/cosmos/relayer/v2/relayer/provider"

func (pp *PenumbraProvider) CreateKeystore(path string) error {
	return nil
}
func (pp *PenumbraProvider) KeystoreCreated(path string) bool {
	return true
}
func (pp *PenumbraProvider) AddKey(name string, coinType uint32) (output *provider.KeyOutput, err error) {
	return nil, nil
}
func (pp *PenumbraProvider) RestoreKey(name, mnemonic string, coinType uint32) (address string, err error) {
	return "", nil
}
func (pp *PenumbraProvider) ShowAddress(name string) (address string, err error) {
	return "", nil
}
func (pp *PenumbraProvider) ListAddresses() (map[string]string, error) {
	return map[string]string{}, nil
}
func (pp *PenumbraProvider) DeleteKey(name string) error {
	return nil
}
func (pp *PenumbraProvider) KeyExists(name string) bool {
	return true
}
func (pp *PenumbraProvider) ExportPrivKeyArmor(keyName string) (armor string, err error) {
	return "", nil
}

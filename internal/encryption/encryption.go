package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

var MASTER_KEY = getMasterKey()

func getMasterKey() string {
	key := os.Getenv("ENCRYPT_KEY")
	if key == "" {
		panic("ENCRYPT_KEY não definida nas variáveis de ambiente")
	}
	return key
}

func EncryptFile(data []byte, password string) ([]byte, error) {
	// Combina chave master com senha do usuário para mais segurança
	key := generateKey(MASTER_KEY + password)

	// Cria cipher AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cipher: %v", err)
	}

	// Cria GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar GCM: %v", err)
	}

	// Gera nonce aleatório
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("erro ao gerar nonce: %v", err)
	}

	// Criptografa os dados
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

func DecryptFile(ciphertext []byte, password string) ([]byte, error) {
	// Combina chave master com senha do usuário
	key := generateKey(MASTER_KEY + password)

	// Cria cipher AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cipher: %v", err)
	}

	// Cria GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar GCM: %v", err)
	}

	// Verifica tamanho mínimo
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("arquivo inválido: muito pequeno")
	}

	// Separa nonce dos dados
	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Descriptografa
	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao descriptografar: senha incorreta ou arquivo corrompido")
	}

	return plaintext, nil
}

func generateKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

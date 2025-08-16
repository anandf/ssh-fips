package main

import (
	"crypto/fips140"
	"log"
	"os"
	"path/filepath"
)

import (
	"golang.org/x/crypto/ssh"
)

func main() {
	log.Printf("Fips enabled? %v\n", fips140.Enabled())
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error when getting user's home directory: %v", err)
	}
	sshKeyPath := filepath.Join(homeDir, ".ssh", "id_rsa")
	sshKey, err := os.ReadFile(sshKeyPath)
	if err != nil {
		log.Fatalf("error when reading private key from %s: %v", sshKeyPath, err)
	}
	sshKeySigner, err := ssh.ParsePrivateKey(sshKey)
	if err != nil {
		log.Fatalf("error when parsing private key: %v", err)
	}
	cfg := &ssh.ClientConfig{
		User:            "git",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(sshKeySigner)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	cfg.SetDefaults()

	//cfg.Ciphers = []string{"aes256-gcm@openssh.com"}
	cfg.KeyExchanges = []string{"curve25519-sha256"}
	conn, err := ssh.Dial("tcp", "github.com:22", cfg)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	log.Println("SSH connection was successful")
}

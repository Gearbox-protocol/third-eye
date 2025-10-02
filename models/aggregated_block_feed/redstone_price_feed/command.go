package redstone_price_feed

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
)

func GetRedStoneResult(data []byte, block int64, rpc string) []byte {
	// Define the command and its arguments.
	// The first argument to exec.Command is the command itself,
	// and subsequent arguments are passed to that command.
	cmdstr := []string{
		"cast",
		"call",
		"0xca11bde05977b3631167028862be2a173976ca11", // multicall
		"--data",
		hex.EncodeToString(data),
		"--block",
		fmt.Sprintf("%d", block),
		"-r",
		rpc,
		"--trace",
	}
	cmd := exec.Command(cmdstr[0], cmdstr[1:]...) // Example: list files in the current directory

	// Execute the command and capture its standard output.
	output, err := cmd.Output()
	if err != nil {
		log.Info(strings.Join(cmdstr, " "))
		log.Fatalf("Error executing command: %v", err)
	}

	// Convert output to string and get the last line
	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")
	if len(lines) == 0 {
		return nil
	}
	lastLine := ""
	for _, line := range lines {
		if strings.Contains(line, "0x") {
			lastLine = line
		}
	}

	// Extract hex string from the last line (remove "0x" prefix if present)
	spls := strings.Split(strings.TrimSpace(lastLine), "0x")
	// Convert hex string to bytes
	bytes, err := hex.DecodeString(spls[1])
	if err != nil {
		log.Errorf("Error decoding hex string: %v", err)
		return nil
	}

	return bytes[len(bytes)-5*32:]
}

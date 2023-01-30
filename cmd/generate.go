/*
Copyright © 2023 NAME HERE <nonsoamadi@aol.com>
*/
package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var envFile string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Convert environment variables in a .env file to a Kubernetes secret yaml file",
	Long:  `This tool converts environment variables in a .env file to a Kubernetes secret yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		secret := getSecret(envFile)
		writeYaml(secret)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&envFile, "file", "f", "", "Path to .env file")
	generateCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(generateCmd)
}

func getSecret(file string) map[string]string {
	secret := make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		secret[kv[0]] = kv[1]
	}

	return secret
}

func writeYaml(secret map[string]string) {
	yaml := "apiVersion: v1\nkind: Secret\nmetadata:\n  name: mysecret\ntype: Opaque\ndata:\n"
	for k, v := range secret {
		encoded := base64.StdEncoding.EncodeToString([]byte(v))
		yaml += fmt.Sprintf("  %s: %s\n", k, encoded)
	}

	err := os.WriteFile("secret.yaml", []byte(yaml), 0644)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Println("Generated secret yaml file: secret.yaml")
}

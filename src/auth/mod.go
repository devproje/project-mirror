package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-mirror/src/config"
	"github.com/devproje/project-mirror/src/util"
	"golang.org/x/term"
)

const (
	dirname  = ".tmp"
	infofile = "account.json"
)

var info = fmt.Sprintf("%s/%s", dirname, infofile)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	fmt.Println()

	acc := &Account{
		Username: strings.TrimSpace(username),
		Password: string(bytePassword),
	}

	err = acc.new()
	if err != nil {
		return err
	}

	return nil
}

func (acc *Account) new() error {
	acc.Password = HashPassword(acc.Password)

	data, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	err = os.WriteFile(info, data, 0655)
	if err != nil {
		return err
	}

	return nil
}

func get() (*Account, error) {
	data, err := util.ParseJSON[Account](info)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Init() {
	if _, err := os.Stat(dirname); err != nil {
		os.Mkdir(dirname, 0755)
	}

	if !config.Get().Auth {
		return
	}

	_, err := os.Stat(info)
	if err != nil {
		err = register()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (acc *Account) Login() bool {
	info, err := get()
	if err != nil {
		return false
	}

	if info.Username != acc.Username {
		return false
	}

	return CheckPasswordHash(acc.Password, info.Password)
}

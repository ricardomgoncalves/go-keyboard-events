package keylistener

import (
	"errors"
	"fmt"
	"log"

	"github.com/MarinX/keylogger"
	"github.com/ricardomgoncalves/go-keyboard-events/keycode"
)

type KeySequence []keycode.KeyCode

type KeyHandler func()

type KeyEvent interface {
	//The moment a key is pressed
	KeyDown()

	//The moment the key is released
	KeyUp()

	//Every tick the key is being pressed
	Key()
}

var keyRegister = map[string]KeyEvent{}

func RegisterNewKeyEvent(ev KeyEvent, keys ...keycode.KeyCode) {
	keyString := ""

	for _, key := range keys {
		keyString += string(key)
	}

	keyRegister[keyString] = ev
}

func GetKeyEvent(key string) (KeyEvent, error) {
	ke := keyRegister[key]

	if ke == nil {
		return nil, errors.New("NO KEY EVENT")
	}

	return ke, nil
}

func StartListenToKeyboard(keyboard string) {
	k, err := keylogger.New(keyboard)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {

		case keylogger.EvKey:
			keyEventHandle(e)
		}
	}
}

var handlingEvent KeyEvent = nil
var currentKeys []string

func keyEventHandle(e keylogger.InputEvent) {
	if e.KeyPress() {
		if handlingEvent != nil {
			handlingEvent.KeyUp()
		}

		currentKeys = append(currentKeys, e.KeyString())

		fmt.Println(currentKeys)

		k := getKeyString(currentKeys)

		he, err := GetKeyEvent(k)

		handlingEvent = he

		if err != nil {
			return
		}

		handlingEvent.KeyDown()
		return
	} else if e.KeyRelease() {
		filterKeys(e.KeyString())

		if handlingEvent == nil {
			return
		}

		handlingEvent.KeyUp()

		k := getKeyString(currentKeys)

		he, _ := GetKeyEvent(k)

		handlingEvent = he
	} else if handlingEvent != nil {
		handlingEvent.Key()
	}
}

func filterKeys(k string) {
	var aux []string
	for _, s := range currentKeys {
		if s != k {
			aux = append(aux, s)
		}
	}
	currentKeys = aux
}

func getKeyString(keys []string) string {
	k := ""

	for _, key := range keys {
		k += key
	}

	return k
}

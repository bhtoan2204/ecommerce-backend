package locale

import (
	"context"
	"fmt"
	"os"
	"payment/package/contxt"
	"payment/package/settings"
	"sync"

	"gopkg.in/yaml.v2"
)

var messages map[string]map[string]string

var _once sync.Once

const (
	VI = "vi"
	EN = "en"
)

func Init(cfg *settings.Config) error {
	_once.Do(func() {
		vimsgs, err := loadMessage(cfg.Locale.VIPath)
		if err != nil {
			panic(fmt.Errorf("local locale path=%s got err=%w", cfg.Locale.VIPath, err))
		}

		enmsgs, err := loadMessage(cfg.Locale.ENPath)
		if err != nil {
			panic(fmt.Errorf("local locale path=%s got err=%w", cfg.Locale.ENPath, err))
		}

		messages = map[string]map[string]string{
			VI: vimsgs,
			EN: enmsgs,
		}
	})

	return nil
}

func Get(ctx context.Context, key string, params ...interface{}) string {
	languageKey := getLanguageKeyFromContext(ctx)

	return GetByLocale(languageKey, key, params...)
}

func GetByLocale(locale string, key string, params ...interface{}) string {
	var message map[string]string
	if messages[locale] != nil {
		message = messages[locale]
	} else {
		message = messages[VI]
	}

	if message == nil {
		return key
	}

	val, has := message[key]
	if !has {
		return key
	}

	return fmt.Sprintf(val, params...)
}
func loadMessage(localePath string) (map[string]string, error) {
	locale, err := os.ReadFile(localePath)
	if err != nil {
		return nil, fmt.Errorf("read local file=%s got err=%w", localePath, err)
	}

	locale = []byte(os.ExpandEnv(string(locale)))
	message := map[string]string{}

	err = yaml.Unmarshal(locale, &message)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to local map got err=%w", err)
	}

	return message, nil
}

func getLanguageKeyFromContext(ctx context.Context) string {
	l, err := contxt.GetAcceptLanguage(ctx)
	if err != nil {
		return VI
	}

	return l
}

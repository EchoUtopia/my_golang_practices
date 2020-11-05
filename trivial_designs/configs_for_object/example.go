package configs

import (
	"database/sql"
	"encoding/json"
	"github.com/EchoUtopia/zerror"
	"sync"
)

type ExampleConfig struct {
	db *sql.DB
}

func (ec *ExampleConfig)GetName()string{return `ExampleConfig`}

func (ec *ExampleConfig)UnMarshalContent(content interface{}, c *Config) error {
	if err := json.Unmarshal([]byte(c.Content), content); err != nil {
		return zerror.BadRequest.Wrap(err)
	}
	return nil
}

func (ec *ExampleConfig)CheckAndSet(ctx interface{}, c *Config) error {
	list := []string{}
	if err := ec.UnMarshalContent(&list, c); err != nil {
		return err
	}
	object, ok := ctx.(int64)
	if ! ok {
		return zerror.BadRequest.New()
	}
	_ = object
	// any other checks
	return nil
}

var once sync.Once


func main(){

	once.Do(func() {
		// in actual scene, db not nil
		RegisterConfigurer(&ExampleConfig{db: nil})
	})
	input := &Config{
		Name:        `ExampleConfig`,
		Category:    `***`,
		Switch:      `on`,
		Description: `***`,
		ContentType: `***`,
		Content:     `***`,
	}
	objectID := 123
	if err := CheckAndSet(objectID, input);err != nil {
		panic(err)
	}

}



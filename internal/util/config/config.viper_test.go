package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/config"
)

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)
	var c config.Config

	// When try to get a env which does not exist, It should return an empty string
	//
	c = config.NewConfig()
	assert.Empty(c.GetString("INVALID"))

	// When try to get an env which does exist, It should return its value
	//
	os.Setenv("TEST", "value")
	c = config.NewConfig()
	assert.Equal("value", c.GetString("TEST"))

	// When try to get an env from file, It should return its value
	//
	// Create file
	f, err := os.Create("test.demoenv.env")
	assert.NoError(err)
	defer f.Close()
	_, err = f.WriteString("ENV_FROM_FILE=value_from_env_from_file\n")
	assert.NoError(err)
	// Check env
	c = config.NewConfig()
	err = c.LoadConfigFile(".", "env", "test.demoenv.env")
	assert.NoError(err)
	assert.Equal("value_from_env_from_file", c.GetString("ENV_FROM_FILE"))
	// Delete file
	err = os.Remove("test.demoenv.env")
	assert.NoError(err)

	// When try to get an env file and it does not exist, It should not throw an error
	//
	// Check env
	c = config.NewConfig()
	err = c.LoadConfigFile(".", "env", "env.nonexisted")
	assert.NoError(err)

	// When try to get an unsupported env file, It should throw an error
	//
	// Create file
	f, err = os.Create("test.demoenv.unsupported")
	assert.NoError(err)
	defer f.Close()
	_, err = f.WriteString("ENV_FROM_FILE=value_from_env_from_file\n")
	assert.NoError(err)
	// Check env
	c = config.NewConfig()
	err = c.LoadConfigFile(".", "unsupported", "test.demoenv.unsupported")
	assert.Error(err, "Unsupported Config Type \"unsupported\"")
	// Delete file
	err = os.Remove("test.demoenv.unsupported")
	assert.NoError(err)
}

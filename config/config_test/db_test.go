package config_test

import (
	"testing"
	"github.com/hftamayo/gingormrest/config"
	sqlmock "github.com/stretchr/testify/mock"
)

func TestConnect(t *testing.T) {
	dbMock := sqlmock.Open(`postgres`)

	dbMock.Expect().Ping().WillReturn(nil)

	config.DB = dbMock

	config.Connect()

	dbMock.AssertExpectations(t)
}

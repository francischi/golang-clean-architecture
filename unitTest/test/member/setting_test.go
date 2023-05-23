package member

import (
	"os"
	"testing"
	"golang/pkg/helpers"
)

func TestMain(m *testing.M) {
    // 讀取 .env 檔案並設定環境變數
    helpers.InitTestEnvSetting("../../.env.test")
	exitCode := m.Run()
	os.Exit(exitCode)
}
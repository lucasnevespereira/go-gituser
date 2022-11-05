package app

import (
	"encoding/json"
	"go-gituser/state"
	"go-gituser/utils"
	"go-gituser/utils/logger"
)

func Sync() {
	accountsFile, err := utils.GetAccountsDataFile()
	if err != nil {
		logger.PrintError(err)
	}

	data, err := utils.ReadFileData(accountsFile)
	if err != nil {
		logger.PrintErrorWithMessage(err, "Sync.ReadFileData")
	}


	err = json.Unmarshal(data, state.SavedAccounts)
	if err != nil {
		logger.PrintErrorWithMessage(err, "Sync.Unmarshal")
	}
}



package poj

import (
	bdconnect "MyServer/handler"
	"MyServer/mlog"
)

func Create() {
	//	decodeData := Parser()
	bdconnector, err := bdconnect.GetNewDBConnector()
	defer bdconnector.Connector.Close()
	if err != nil {
		mlog.Error("Database connect error: %v", err)
	}
	//	querry := ("INSERT INTO user (name,sirname,patronymic,number,email,birthdate) VALUES (?,?,?,?,?,?)")
	//	result, err := bdconnector.Connector.Exec(querry, decodeData.Name, decodeData.Sirname)
}

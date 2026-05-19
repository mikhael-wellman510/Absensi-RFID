package common

type DbInfo struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func NewDbInfo(dbInfo DbInfo) *DbInfo {

	if dbInfo.Host == "" {
		panic("Host Is Not Configured")
	}

	return &dbInfo
}

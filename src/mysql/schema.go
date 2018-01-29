package mysql

import "time"

type AndroidLog struct {
	//Id int64 `json:"id"`
	AppName    string `json:"appName"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	RedisData  string `json:"redisData"`
	ReturnData string `json:"returnData"`
	//Dateline time.Time `json:"dateline"`
}

type IOSLog struct {
	//Id int64 `json:"id"`
	AppName     string `json:"appName"`
	IOSMODE     string `json:"iosMode"`
	DEVICETOKEN string `json:"deviceToken"`
	Body        string `json:"body"`
	RedisData   string `json:"redisData"`
	ReturnData  string `json:"returnData"`
	//Dateline time.Time `json:"dateline"`
}

type PushApps struct {
	Id int64 `json:"id"`
	AppName           string `json:"appName"`
	Password          string `json:"password"`
	IPAddress         string `json:"IPAddress"`
	PemPassphrase     string `json:"pemPassphrase"`
	JPushAppKey       string `json:"jPushAppKey"`
	JPushMasterSecret string `json:"jPushMasterSecret"`
	BundleId          string `json:"bundleId"`
	Status            int    `json:"status"`
	Dateline time.Time `json:"dateline"`
}

type ConsumerRecord struct {
	Id int64 `json:"id"`
	Topic       string `json:"topic"`
	PartitionId int32 `json:"partitionId"`
	Offset      int64 `json:"offset"`
}

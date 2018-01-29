package mysql

import (
	"testing"
	"fmt"
	"time"
)



func TestFetchList(t *testing.T) {
	where := map[string]interface{}{}
	rows, err := FetchList("push_apps", "", where)
	if err != nil {
		t.Error(err)
	}
	pushApps := []PushApps{}
	for rows.Next() {
		var pushApp PushApps
		var Id int
		var AppName string
		var Password string
		var IPAddress string
		var PemPassphrase string
		var JPushAppKey string
		var JPushMasterSecret string
		var BundleId string
		var Status int
		var dateline time.Time

		//rows.Scan(&pushApp.Id, &pushApp.AppName, &pushApp.Password, &pushApp.IPAddress, &pushApp.PemPassphrase, &pushApp.JPushAppKey, &pushApp.JPushMasterSecret, &pushApp.BundleId, &pushApp.Status)
		//rows.Scan(&pushApp)
		rows.Scan(&Id, &AppName, &Password, &IPAddress, &PemPassphrase, &JPushAppKey, &JPushMasterSecret, &BundleId, &Status, &dateline)

		fmt.Println( Id, AppName, Password , IPAddress, PemPassphrase, JPushAppKey, JPushMasterSecret, BundleId, Status, dateline )
		pushApps = append(pushApps, pushApp)
	}
	fmt.Println( pushApps )
}

func TestFetchList1(t *testing.T) {
	where := map[string]interface{}{}
	rows, err := FetchList("consumer_record", "", where)
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var id int
		var topic string
		var partitionId int
		var offset int
		rows.Scan(&id, &topic, &partitionId, &offset)

		fmt.Println( id, topic, partitionId, offset )
	}
}

func TestCount(t *testing.T) {
	where := map[string]interface{}{}
	rows, err := Count("push_apps",  where)
	if err != nil {
		t.Error(err)
	}

	fmt.Println( rows )
}

func TestInsert(t *testing.T) {
	rowMap := map[string]interface{}{
		"topic":"rrr",
		"partitionId": 1,
		"offset":3,

	}
	id, err := Insert("consumer_record", rowMap)
	if err != nil{
		t.Error( err )
	}
	fmt.Println( id )
}
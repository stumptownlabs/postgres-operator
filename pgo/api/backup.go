package api

/*
 Copyright 2017 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	"github.com/crunchydata/postgres-operator/util"
	"net/http"
)

func ShowBackup(httpclient *http.Client, arg string, SessionCredentials *msgs.BasicAuthCredentials, ns string) (msgs.ShowBackupResponse, error) {

	var response msgs.ShowBackupResponse

	url := SessionCredentials.APIServerURL + "/backups/" + arg + "?version=" + msgs.PGO_VERSION + "&" + util.LABEL_NAMESPACE + "=" + ns

	log.Debugf("show backup called [%s]", url)

	action := "GET"
	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return response, err
	}

	req.SetBasicAuth(SessionCredentials.Username, SessionCredentials.Password)
	resp, err := httpclient.Do(req)
	if err != nil {
		fmt.Println("Error: Do: ", err)
		return response, err
	}
	defer resp.Body.Close()
	log.Debugf("%v", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err

}
func DeleteBackup(httpclient *http.Client, arg string, SessionCredentials *msgs.BasicAuthCredentials, ns string) (msgs.DeleteBackupResponse, error) {

	var response msgs.DeleteBackupResponse
	url := SessionCredentials.APIServerURL + "/backupsdelete/" + arg + "?version=" + msgs.PGO_VERSION + "&" + util.LABEL_NAMESPACE + "=" + ns

	log.Debug("delete backup called [%s]", url)

	action := "GET"
	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return response, err
	}

	req.SetBasicAuth(SessionCredentials.Username, SessionCredentials.Password)
	resp, err := httpclient.Do(req)
	if err != nil {
		fmt.Println("Error: Do: ", err)
		return response, err
	}
	defer resp.Body.Close()
	log.Debugf("%v", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err

}

func CreateBackup(httpclient *http.Client, SessionCredentials *msgs.BasicAuthCredentials, request *msgs.CreateBackupRequest) (msgs.CreateBackupResponse, error) {

	var response msgs.CreateBackupResponse

	jsonValue, _ := json.Marshal(request)
	url := SessionCredentials.APIServerURL + "/backups"

	log.Debugf("create backup called [%s]", url)

	action := "POST"
	req, err := http.NewRequest(action, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(SessionCredentials.Username, SessionCredentials.Password)

	resp, err := httpclient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	log.Debugf("%v", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err
}

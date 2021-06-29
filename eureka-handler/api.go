package eurekaHandler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/GreatLaboratory/go-sms/util"
	"github.com/xuanbo/requests"
)

var (
	ErrNotFound = errors.New("not found")
)

func Register(zone, app string, instance *Instance) error {
	type InstanceInfo struct {
		Instance *Instance `json:"instance"`
	}
	var info = &InstanceInfo{
		Instance: instance,
	}

	u := zone + "apps/" + app

	result := requests.Post(u).Json(info).Send().Status2xx()

	if result.Err != nil {
		if typeErr := util.IsErrorType(result.Err); typeErr != nil {
			return typeErr
		}
		return fmt.Errorf("register application instance failed, error: %s", result.Err)
	}
	return nil
}

func UnRegister(zone, app, instanceID string) error {
	u := zone + "apps/" + app + "/" + instanceID
	result := requests.Delete(u).Send().StatusOk()

	if result.Err != nil {
		if typeErr := util.IsErrorType(result.Err); typeErr != nil {
			return typeErr
		}
		return fmt.Errorf("unRegister application instance failed, error: %s", result.Err)
	}
	return nil
}

func Refresh(zone string) (*Applications, error) {
	type Result struct {
		Applications *Applications `json:"applications"`
	}
	apps := new(Applications)
	res := &Result{
		Applications: apps,
	}
	u := zone + "apps"
	err := requests.Get(u).Header("Accept", " application/json").Send().StatusOk().Json(res)

	if err != nil {
		if typeErr := util.IsErrorType(err); typeErr != nil {
			return nil, typeErr
		}
		return nil, fmt.Errorf("refresh failed, error: %s", err)
	}
	return apps, nil
}

func Heartbeat(zone, app, instanceID string) error {
	u := zone + "apps/" + app + "/" + instanceID
	params := url.Values{
		"status": {"UP"},
	}
	result := requests.Put(u).Params(params).Send()

	if result.Err != nil {
		if typeErr := util.IsErrorType(result.Err); typeErr != nil {
			return typeErr
		}
		return fmt.Errorf("heartbeat failed, error: %s", result.Err)
	}

	if result.Resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if typeErr := util.IsIntegerType(result.Resp.StatusCode); typeErr != nil {
		return typeErr
	}

	if result.Resp.StatusCode != http.StatusOK {
		return fmt.Errorf("heartbeat failed, invalid status code: %d", result.Resp.StatusCode)
	}
	return nil
}

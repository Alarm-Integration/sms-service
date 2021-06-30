package eurekaHandler

import (
	"errors"
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

	return util.ErrorHandle("register application instance failed, error: %s", result.Err)
}

func UnRegister(zone, app, instanceID string) error {
	u := zone + "apps/" + app + "/" + instanceID
	result := requests.Delete(u).Send().StatusOk()

	return util.ErrorHandle("unregister application instance failed, error: %s", result.Err)
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

	return apps, util.ErrorHandle("refresh failed, error: %s", err)
}

func Heartbeat(zone, app, instanceID string) error {
	u := zone + "apps/" + app + "/" + instanceID
	params := url.Values{
		"status": {"UP"},
	}
	result := requests.Put(u).Params(params).Send()

	return util.ErrorHandle("heartbeat failed, error: %s", result.Err, result.Resp.StatusCode)
}

package casbin

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type RoleType string

const (
	RoleUndefined RoleType = ""
	RoleAdmin     RoleType = "admin"
	RoleUser      RoleType = "user"
	RoleMerchant  RoleType = "merchant"
	RoleGuest     RoleType = "guest"
	RoleBlacklist RoleType = "blacklist"
)

type EffectType string

const (
	EffectUndefined EffectType = ""
	EffectAllow     EffectType = "allow"
	EffectDeny      EffectType = "deny"
)

type MethodType string

const (
	MethodUndefined  MethodType = ""
	MethodGet        MethodType = "GET"
	MethodPost       MethodType = "POST"
	MethodPut        MethodType = "PUT"
	MethodDelete     MethodType = "DELETE"
	MethodGetAndPost MethodType = "(GET)|(POST)"
	MethodAll        MethodType = "*"
)

var (
	enforcer *AsyncEnforcer
)

type AsyncEnforcer struct {
	sync.Mutex
	e  *casbin.Enforcer
	db *gorm.DB
}

type PolicyConf struct {
	Role   RoleType
	Object string
	Method MethodType
	Effect EffectType
}

func InitEnforcer(modelConf string, policyConf string) error {

	enforcer = &AsyncEnforcer{
		e: nil,
	}

	e, err := casbin.NewEnforcer(modelConf, policyConf)
	if err != nil {
		return err
	}
	enforcer.e = e
	// https://casbin.org/zh/docs/function/
	// https://casbin.org/zh/docs/rbac-with-pattern
	enforcer.e.AddNamedDomainMatchingFunc("g", "regexMatch", util.RegexMatch)
	return nil
}

func Check(subject, action, object string, explain bool) (bool, error) {
	if enforcer.e == nil {
		return false, errors.New("must init casbin enforcer")
	}
	enforcer.Lock()
	defer enforcer.Unlock()
	_ = enforcer.e.LoadModel()
	_ = enforcer.e.LoadPolicy()
	var ok bool
	var err error
	var ex []string
	if explain {
		ok, ex, err = enforcer.e.EnforceEx(subject, object, action)
		logrus.Infof("match result: %v, explain: %v", ok, ex)

	} else {
		ok, err = enforcer.e.Enforce(subject, object, action)
	}

	return ok, err

}

func AddRoleForUser(subject string, role RoleType) error {
	if enforcer.e == nil {
		return errors.New("must init casbin enforcer")
	}

	enforcer.Lock()
	defer enforcer.Unlock()
	ok, err := enforcer.e.AddRoleForUser(subject, string(role))
	if err != nil {
		return err
	}

	if ok {
		err = enforcer.e.SavePolicy()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteRoleForUser(subject string, role RoleType) error {
	if enforcer.e == nil {
		return errors.New("must init casbin enforcer")
	}
	enforcer.Lock()
	defer enforcer.Unlock()
	ok, err := enforcer.e.DeleteRoleForUser(subject, string(role))
	if err != nil {
		return err
	}
	if ok {
		err = enforcer.e.SavePolicy()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(subject string) error {
	if enforcer.e == nil {
		return errors.New("must init casbin enforcer")
	}

	enforcer.Lock()
	defer enforcer.Unlock()

	ok, err := enforcer.e.DeleteUser(subject)
	if err != nil {
		return err
	}

	if ok {
		err = enforcer.e.SavePolicy()
		if err != nil {
			return err
		}
	}
	return nil

}

func AddPolicy(conf PolicyConf) error {
	if enforcer.e == nil {
		return errors.New("must init casbin enforcer")
	}

	if conf.Role == RoleUndefined {
		logrus.Fatal("policy config 'role' is undefined")
	}

	if conf.Method == MethodUndefined {
		logrus.Fatal("policy config 'method' is undefined")
	}

	if conf.Effect == "" {
		logrus.Warnf("policy config 'effect' is undefined, default use `allow`")
		conf.Effect = EffectAllow
	}

	enforcer.Lock()
	defer enforcer.Unlock()

	ok, err := enforcer.e.AddPolicy(string(conf.Role), conf.Object, string(conf.Method), string(conf.Effect))
	if err != nil {
		return err
	}

	if ok {
		err = enforcer.e.SavePolicy()
		if err != nil {
			return err
		}
	}

	return nil
}

func AddPolicies(policies []PolicyConf) []error {
	if enforcer.e == nil {
		return []error{errors.New("must init casbin enforcer")}
	}

	e := make([]error, 0)
	for _, p := range policies {
		if err := AddPolicy(p); err != nil {
			e = append(e, err)
		}
	}
	return e
}

func DeletePolicy(conf PolicyConf) error {
	if enforcer.e == nil {
		return errors.New("must init casbin enforcer")
	}

	if "" == conf.Effect {
		return errors.New("must provide a specific effect(allow or deny)")
	}

	enforcer.Lock()
	defer enforcer.Unlock()
	ok, err := enforcer.e.RemovePolicy(string(conf.Role), conf.Object, string(conf.Method), string(conf.Effect))
	if err != nil {
		return err
	}

	if ok {
		err = enforcer.e.SavePolicy()
		if err != nil {
			return err
		}
	}
	return nil
}

func DeletePolicies(policies []PolicyConf) []error {
	if enforcer.e == nil {
		return []error{errors.New("must init casbin enforcer")}
	}

	e := make([]error, 0)
	for _, p := range policies {
		if err := DeletePolicy(p); err != nil {
			e = append(e, err)
		}
	}
	return e
}

func SaveToDB() {

}

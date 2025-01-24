package casbin

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func init() {
	err := InitEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestAddRoleForUser(t *testing.T) {
	_ = AddRoleForUser("alice", RoleAdmin)

	_ = AddRoleForUser("bob", RoleUser)

	_ = AddRoleForUser("candy", RoleUser)
	_ = AddRoleForUser("candy", RoleMerchant)

	_ = AddRoleForUser("david", RoleGuest)

	_ = AddRoleForUser("eva", RoleBlacklist)
}

func TestDeleteRoleForUser(t *testing.T) {
	_ = DeleteRoleForUser("candy", RoleMerchant)
}

func TestAddPolicy(t *testing.T) {
	conf := PolicyConf{
		Role:   RoleUser,
		Object: "/test/*",
		Method: MethodAll,
		Effect: EffectAllow,
	}

	_ = AddPolicy(conf)
}

func TestDeletePolicy(t *testing.T) {
	conf := PolicyConf{
		Role:   RoleUser,
		Object: "/test/*",
		Method: MethodAll,
		Effect: EffectAllow,
	}

	_ = DeletePolicy(conf)
}

func TestCheck(t *testing.T) {
	_ = AddRoleForUser("alice", RoleAdmin)
	_ = AddRoleForUser("bob", RoleUser)
	_ = AddRoleForUser("candy", RoleUser)
	_ = AddRoleForUser("candy", RoleMerchant)
	_ = AddRoleForUser("david", RoleGuest)
	_ = AddRoleForUser("eva", RoleBlacklist)
	_ = AddRoleForUser("flor", "other")

	testCases := [][]string{
		{"alice", "GET", "/product/id=123"},
		{"alice", "POST", "/product/id=123"},
		{"alice", "GET", "/cart/id=123"},
		{"alice", "POST", "/cart/id=123"},
		{"alice", "GET", "/order/id=123"},
		{"alice", "POST", "/order/id=123"},
		{"alice", "GET", "/payment/id=123"},
		{"alice", "POST", "/payment/id=123"},
		{"alice", "GET", "/user/id=123"},
		{"alice", "POST", "/user/id=123"},
		{"alice", "GET", "/admin/id=123"},
		{"alice", "POST", "/admin/id=123"},

		{"bob", "GET", "/product/id=123"},
		{"bob", "POST", "/product/id=123"},
		{"bob", "GET", "/cart/id=123"},
		{"bob", "POST", "/cart/id=123"},
		{"bob", "GET", "/order/id=123"},
		{"bob", "POST", "/order/id=123"},
		{"bob", "GET", "/payment/id=123"},
		{"bob", "POST", "/payment/id=123"},
		{"bob", "GET", "/user/id=123"},
		{"bob", "POST", "/user/id=123"},
		{"bob", "GET", "/admin/id=123"},
		{"bob", "POST", "/admin/id=123"},

		{"candy", "GET", "/product/id=123"},
		{"candy", "POST", "/product/id=123"},
		{"candy", "GET", "/cart/id=123"},
		{"candy", "POST", "/cart/id=123"},
		{"candy", "GET", "/order/id=123"},
		{"candy", "POST", "/order/id=123"},
		{"candy", "GET", "/payment/id=123"},
		{"candy", "POST", "/payment/id=123"},
		{"candy", "GET", "/user/id=123"},
		{"candy", "POST", "/user/id=123"},
		{"candy", "GET", "/admin/id=123"},
		{"candy", "POST", "/admin/id=123"},

		{"david", "GET", "/product/id=123"},
		{"david", "POST", "/product/id=123"},
		{"david", "GET", "/cart/id=123"},
		{"david", "POST", "/cart/id=123"},
		{"david", "GET", "/order/id=123"},
		{"david", "POST", "/order/id=123"},
		{"david", "GET", "/payment/id=123"},
		{"david", "POST", "/payment/id=123"},
		{"david", "GET", "/user/id=123"},
		{"david", "POST", "/user/id=123"},
		{"david", "GET", "/admin/id=123"},
		{"david", "POST", "/admin/id=123"},

		{"eva", "GET", "/product/id=123"},
		{"eva", "POST", "/product/id=123"},
		{"eva", "GET", "/cart/id=123"},
		{"eva", "POST", "/cart/id=123"},
		{"eva", "GET", "/order/id=123"},
		{"eva", "POST", "/order/id=123"},
		{"eva", "GET", "/payment/id=123"},
		{"eva", "POST", "/payment/id=123"},
		{"eva", "GET", "/user/id=123"},
		{"eva", "POST", "/user/id=123"},
		{"eva", "GET", "/admin/id=123"},
		{"eva", "POST", "/admin/id=123"},

		{"flor", "GET", "/product/id=123"},
		{"flor", "POST", "/product/id=123"},
		{"flor", "GET", "/cart/id=123"},
		{"flor", "POST", "/cart/id=123"},
		{"flor", "GET", "/order/id=123"},
		{"flor", "POST", "/order/id=123"},
		{"flor", "GET", "/payment/id=123"},
		{"flor", "POST", "/payment/id=123"},
		{"flor", "GET", "/user/id=123"},
		{"flor", "POST", "/user/id=123"},
		{"flor", "GET", "/admin/id=123"},
		{"flor", "POST", "/admin/id=123"},
	}

	testCasesResult := []bool{
		true, true, true, true, true, true, true, true, true, true, true, true,
		true, false, true, true, true, true, true, true, true, true, false, false,
		true, true, true, true, true, true, true, true, true, true, false, false,
		true, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false,
	}

	for i, c := range testCases {
		ok, err := Check(c[0], c[1], c[2], true)
		if err != nil {
			t.Fail()
			logrus.Error(err)
		}
		if ok != testCasesResult[i] {
			t.Fail()
			logrus.Infof("miss match at %d, %v\n", i, c)
		}
	}

	_ = DeleteRoleForUser("alice", RoleAdmin)
	_ = DeleteRoleForUser("bob", RoleUser)
	_ = DeleteRoleForUser("candy", RoleUser)
	_ = DeleteRoleForUser("candy", RoleMerchant)
	_ = DeleteRoleForUser("david", RoleGuest)
	_ = DeleteRoleForUser("eva", RoleBlacklist)
	_ = DeleteRoleForUser("flor", "other")
}

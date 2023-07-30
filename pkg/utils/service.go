package utils

import "QuickAuth/pkg/model"

func GetUserPoolId(tenants []model.Tenant) []int64 {
	res := make([]int64, 0, len(tenants))
	for _, tenant := range tenants {
		res = append(res, tenant.UserPoolID)
	}
	return res
}

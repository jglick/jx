// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	auth "github.com/jenkins-x/jx/pkg/auth"
)

func AnyAuthAuthConfigService() auth.AuthConfigService {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(auth.AuthConfigService))(nil)).Elem()))
	var nullValue auth.AuthConfigService
	return nullValue
}

func EqAuthAuthConfigService(value auth.AuthConfigService) auth.AuthConfigService {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue auth.AuthConfigService
	return nullValue
}

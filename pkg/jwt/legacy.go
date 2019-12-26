/*
Copyright 2014 The Kubernetes Authors.

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

package jwt

import (
	"github.com/flaxius/portia/pkg/api/core/v1"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

//PORTIA_ISSUER
const PortiaIssuer = "portia.io"

// time.Now stubbed out to allow testing
var now = time.Now

type legacyPrivateClaims struct {
	ServiceAccountName string `json:"portia.io/serviceaccount/service-account.name"`
	ServiceAccountUID  string `json:"portia.io/serviceaccount/service-account.uid"`
	SecretName         string `json:"portia.io/serviceaccount/secret.name"`
	Namespace          string `json:"portia.io/serviceaccount/namespace"`
}

func LegacyClaims(serviceAccount v1.ServiceAccount) (*jwt.Claims, interface{}) {
	now := now()
	jti, _ := uuid.NewRandom()

	return &jwt.Claims{
			Subject:  "Portia",
			Issuer:   PortiaIssuer,
			ID:       jti.String(),
			IssuedAt: jwt.NewNumericDate(now),
			Expiry:   jwt.NewNumericDate(now.Add(30 * time.Minute)),
		}, &legacyPrivateClaims{
			Namespace:          serviceAccount.Namespace,
			ServiceAccountName: serviceAccount.Name,
			ServiceAccountUID:  string(serviceAccount.UID),
			SecretName:         "secret.Namse",
		}
}

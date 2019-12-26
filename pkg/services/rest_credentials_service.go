package oauth

import (
	"context"
	"fmt"
	"github.com/flaxius/portia/internal/applications/credentials"
	"github.com/flaxius/portia/proto-gen/authentication/rest_credentials"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/util/keyutil"
	"sync"
)

const rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA249XwEo9k4tM8fMxV7zxOhcrP+WvXn917koM5Qr2ZXs4vo26
e4ytdlrV0bQ9SlcLpQVSYjIxNfhTZdDt+ecIzshKuv1gKIxbbLQMOuK1eA/4HALy
EkFgmS/tleLJrhc65tKPMGD+pKQ/xhmzRuCG51RoiMgbQxaCyYxGfNLpLAZK9L0T
ctv9a0mJmGIYnIOQM4kC1A1I1n3EsXMWmeJUj7OTh/AjjCnMnkgvKT2tpKxYQ59P
gDgU8Ssc7RDSmSkLxnrv+OrN80j6xrw0OjEiB4Ycr0PqfzZcvy8efTtFQ/Jnc4Bp
1zUtFXt7+QeevePtQ2EcyELXE0i63T1CujRMWwIDAQABAoIBAHJx8GqyCBDNbqk7
e7/hI9iE1S10Wwol5GH2RWxqX28cYMKq+8aE2LI1vPiXO89xOgelk4DN6urX6xjK
ZBF8RRIMQy/e/O2F4+3wl+Nl4vOXV1u6iVXMsD6JRg137mqJf1Fr9elg1bsaRofL
Q7CxPoB8dhS+Qb+hj0DhlqhgA9zG345CQCAds0ZYAZe8fP7bkwrLqZpMn7Dz9WVm
++YgYYKjuE95kPuup/LtWfA9rJyE/Fws8/jGvRSpVn1XglMLSMKhLd27sE8ZUSV0
2KUzbfRGE0+AnRULRrjpYaPu0XQ2JjdNvtkjBnv27RB89W9Gklxq821eH1Y8got8
FZodjxECgYEA93pz7AQZ2xDs67d1XLCzpX84GxKzttirmyj3OIlxgzVHjEMsvw8v
sjFiBU5xEEQDosrBdSknnlJqyiq1YwWG/WDckr13d8G2RQWoySN7JVmTQfXcLoTu
YGRiiTuoEi3ab3ZqrgGrFgX7T/cHuasbYvzCvhM2b4VIR3aSxU2DTUMCgYEA4x7J
T/ErP6GkU5nKstu/mIXwNzayEO1BJvPYsy7i7EsxTm3xe/b8/6cYOz5fvJLGH5mT
Q8YvuLqBcMwZardrYcwokD55UvNLOyfADDFZ6l3WntIqbA640Ok2g1X4U8J09xIq
ZLIWK1yWbbvi4QCeN5hvWq47e8sIj5QHjIIjRwkCgYEAyNqjltxFN9zmzPDa2d24
EAvOt3pYTYBQ1t9KtqImdL0bUqV6fZ6PsWoPCgt+DBuHb+prVPGP7Bkr/uTmznU/
+AlTO+12NsYLbr2HHagkXE31DEXE7CSLa8RNjN/UKtz4Ohq7vnowJvG35FCz/mb3
FUHbtHTXa2+bGBUOTf/5Hw0CgYBxw0r9EwUhw1qnUYJ5op7OzFAtp+T7m4ul8kCa
SCL8TxGsgl+SQ34opE775dtYfoBk9a0RJqVit3D8yg71KFjOTNAIqHJm/Vyyjc+h
i9rJDSXiuczsAVfLtPVMRfS0J9QkqeG4PIfkQmVLI/CZ2ZBmsqEcX+eFs4ZfPLun
Qsxe2QKBgGuPilIbLeIBDIaPiUI0FwU8v2j8CEQBYvoQn34c95hVQsig/o5z7zlo
UsO0wlTngXKlWdOcCs1kqEhTLrstf48djDxAYAxkw40nzeJOt7q52ib/fvf4/UBy
X024wzbiw1q07jFCyfQmODzURAx1VNT7QVUMdz/N8vy47/H40AZJ
-----END RSA PRIVATE KEY-----
`

type serviceTokenServer struct {
	m sync.Mutex
}

//RestAuthCredentials aa
func RestAuthCredentials() rest_credentials.RestAuthCredentialsServer {
	return &serviceTokenServer{}
}

func getPrivateKey(data string) interface{} {
	key, _ := keyutil.ParsePrivateKeyPEM([]byte(data))
	return key
}

func (s *serviceTokenServer) Create(ctx context.Context, req *rest_credentials.CreateRestRequest) (*rest_credentials.CreateRestResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	clientId, err := credentials.GenerateRandomAppSecret(35)
	clientSecret, err := credentials.GenerateRandomAppSecret(35)
	if err != nil {
		fmt.Println("Error to generate the application credentials")
	}
	// Related API objects
	/*serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-service-account",
			UID:       "pp.String()",
			Namespace: "test",
		},
	}*/
	//rsaGenerator, err := jwt.JWTTokenGenerator(jwt.PortiaIssuer,getPrivateKey(rsaPrivateKey))
	//_ , err:= rsaGenerator.GenerateToken(jwt.LegacyClaims(*serviceAccount))
	if err != nil {
		fmt.Println("Error to generate the application credentials")
	}
	//TODO generate portia validator grpc
	if req.ProjectId == "" || req.RedirectUris == "" {
		st := status.New(codes.InvalidArgument, "INVALID_NAME")
		// we add details here
		badRequest := &errdetails.BadRequest{}
		violations := make([]*errdetails.BadRequest_FieldViolation, 0)
		violation := errdetails.BadRequest_FieldViolation{Field: "N", Description: "unsupported value"}
		(*badRequest).FieldViolations = append(violations, &violation)
		det, err := st.WithDetails(badRequest)
		if err != nil {
			return nil, st.Err()
		}
		return nil, det.Err()
	}

	return &rest_credentials.CreateRestResponse{
		Rest: &rest_credentials.RestAppDefinition{
			ProjectId:    req.ProjectId,
			ClientId:     "P-" + clientId + "." + req.ProjectId,
			AuthUri:      "https://portia.io/oauth2/gen/token",
			ClientSecret: clientSecret,
			RedirectUris: []string{req.RedirectUris},
		},
	}, nil
}

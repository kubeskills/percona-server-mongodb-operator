package stub

import (
	"testing"

	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"

	motPkg "github.com/percona/mongodb-orchestration-tools/pkg"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetReplsetDialInfo(t *testing.T) {
	di := getReplsetDialInfo(
		&v1alpha1.PerconaServerMongoDB{
			ObjectMeta: metav1.ObjectMeta{
				Name:      t.Name(),
				Namespace: "default",
			},
			Spec: v1alpha1.PerconaServerMongoDBSpec{
				Mongod: &v1alpha1.MongodSpec{
					Net: &v1alpha1.MongodSpecNet{
						Port: 99999,
					},
				},
			},
		},
		&v1alpha1.ReplsetSpec{
			Name: defaultReplsetName,
		},
		[]corev1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "testPod",
				},
			},
		},
		&corev1.Secret{
			Data: map[string][]byte{
				motPkg.EnvMongoDBClusterAdminUser:     []byte("clusterAdmin"),
				motPkg.EnvMongoDBClusterAdminPassword: []byte("123456"),
			},
		},
	)
	assert.NotNil(t, di)
	assert.Equal(t, defaultReplsetName, di.ReplicaSetName)
	assert.Len(t, di.Addrs, 1)
	assert.Equal(t, "testPod."+t.Name()+"-"+defaultReplsetName+".default.svc.cluster.local:99999", di.Addrs[0])
	assert.Equal(t, "clusterAdmin", di.Username)
	assert.Equal(t, "123456", di.Password)
	assert.Equal(t, MongoDBTimeout, di.Timeout)
	assert.True(t, di.FailFast)
}

//func TestIsReplsetInitialized(t *testing.T) {}
//func TestHandlerHandleReplsetInit(t *testing.T) {}
//func TestEnsureReplsetStatefulSet(t *testing.T) {}
//func TestHandlerEnsureReplset(t *testing.T) {}
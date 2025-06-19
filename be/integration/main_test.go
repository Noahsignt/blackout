package integration

import (
    "context"
    "log"
    "os"
    "testing"
    "time"
	"fmt"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"
    "github.com/stretchr/testify/require"
)

var (
    mongoC testcontainers.Container
    client *mongo.Client
)

func TestMain(m *testing.M) {
    ctx := context.Background()

    var err error
    mongoC, client, err = setupMongoContainer(ctx)
    if err != nil {
        log.Fatalf("failed to start mongo container: %v", err)
    }

    code := m.Run()

    _ = mongoC.Terminate(ctx)

    os.Exit(code)
}

func SetupTest(t *testing.T) {
	err := client.Database("testdb").Drop(context.Background())
	require.NoError(t, err)
}


// func spins up a container to be used for testing
func setupMongoContainer(ctx context.Context) (testcontainers.Container, *mongo.Client, error) {
    req := testcontainers.ContainerRequest{
        Image:        "mongo:6",
        ExposedPorts: []string{"27017/tcp"},
        WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(20 * time.Second),
    }

    mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        return nil, nil, err
    }

    host, err := mongoC.Host(ctx)
    if err != nil {
        return mongoC, nil, err
    }

    port, err := mongoC.MappedPort(ctx, "27017")
    if err != nil {
        return mongoC, nil, err
    }

    uri := fmt.Sprintf("mongodb://%s:%s", host, port.Port())
    client, err := mongo.Connect(options.Client().ApplyURI(uri))
    if err != nil {
        return mongoC, nil, err
    }

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, nil, err
    }

	fmt.Println("Successfully connected to Test MongoDB Instance!")
    return mongoC, client, nil
}
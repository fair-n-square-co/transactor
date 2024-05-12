package authentication

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
	payloadHeader       = "payload"
)

type FirebaseAuthMiddleware struct {
	app *firebase.App
}

func (f *FirebaseAuthMiddleware) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Check if the service name is in a list of services that require authentication.
	// Replace "Service1" and "Service2" with the actual service names you want to authenticate.
	//requiredServices := []string{"pb.GrpcServerService"}

	//serviceName := info.FullMethod

	if methodRequiresAuthentication(info.FullMethod) {
		// Extract the metadata from the context.
		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "metadata not found")
		}

		// Get the authorization token from metadata.
		authTokens := md[authorizationHeader]
		if len(authTokens) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is missing")
		}

		authHeader := authTokens[0] // Assuming a single token is sent in the header.
		fields := strings.Fields(authHeader)

		if len(fields) != 2 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth header format: %v", fields)
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationBearer {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization type: %v", authType)
		}
		accessToken := fields[1]

		client, err := f.app.Auth(ctx)
		if err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, "internal error")
		}

		token, err := client.VerifyIDToken(ctx, accessToken)
		if err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
		}

		// TODO: fix lint warning
		ctx = context.WithValue(ctx, payloadHeader, token)
		return handler(ctx, req)
	}

	return handler(ctx, req)
}

func NewFirebaseAuthMiddleware() *FirebaseAuthMiddleware {
	// opt := option.WithCredentialsFile("sa.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: "fair-n-square-app",
	})
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firebaseAuth := &FirebaseAuthMiddleware{app}

	return firebaseAuth

}

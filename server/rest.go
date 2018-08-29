package server

import (
  "flag"
  "net/http"

  monitoring "gitlab-devops.totvs.com.br/lucas.martins/monitoring/monitoring"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
)

var (
  endpoint  = flag.String("endpoint", "localhost:9090", "endpoints")
)

func registerGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (*runtime.ServeMux, error) {
  var (
    mux       = runtime.NewServeMux(opts...)
    dialOpts  = []grpc.DialOption{grpc.WithInsecure()}
    err       error
  )

  if err = monitoring.RegisterMonitoringServiceHandlerFromEndpoint(ctx, mux, *endpoint, dialOpts); err != nil {
    return nil, err
  }

  return mux, nil
}

func setHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    h.ServeHTTP(w, r)
  })
}

func Gateway() error {
  var (
    mux     *http.ServeMux
    gw      *runtime.ServeMux
    ctx     = context.Background()
    cancel  context.CancelFunc
    opts    []runtime.ServeMuxOption
    err     error
  )

  mux = http.NewServeMux()

  ctx, cancel = context.WithCancel(ctx)
  defer cancel()

  if gw, err = registerGateway(ctx, opts...); err != nil {
    return err
  }

  mux.Handle("/", gw)

  return http.ListenAndServe(":8080", setHandler(mux))
}

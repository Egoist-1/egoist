package opentelemetry

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"net/http"
	"testing"
	"time"
)

func TestOtel(t *testing.T) {
	res, err := newResource2("demo", "v0.0.1")
	require.NoError(t, err)
	prop := newPropagator2()
	//在客户端和服务端之间传递 tracing相关信息
	otel.SetTextMapPropagator(prop)

	//初始化 trace provider
	//这个 provider 就是在打点的时候创建的trace

	exporter, err := zipkin.New(
		"http://localhost:9411/api/v2/spans")
	assert.NoError(t, err)
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second)),
		trace.WithResource(res),
	)
	defer tp.Shutdown(context.Background())
	otel.SetTracerProvider(tp)

	server := gin.Default()
	server.GET("/test", func(ginCtx *gin.Context) {
		//这个trace 的名字,最好设置为唯一的,比如说用户所在包名
		trace := otel.Tracer("opentelemetry")
		var ctx context.Context = ginCtx
		ctx, span := trace.Start(ctx, "top-span")
		defer span.End()
		span.AddEvent("event-1")
		time.Sleep(time.Second)
		ctx, subSpan := trace.Start(ctx, "sub-span")
		defer subSpan.End()
		time.Sleep(time.Millisecond * 300)
		subSpan.SetAttributes(attribute.String("key1", "value1"))
		subSpan.AddEvent("event_son_2")
		ginCtx.String(http.StatusOK, "OK")

	})
	server.Run(":8082")
}

func newResource2(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func newPropagator2() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

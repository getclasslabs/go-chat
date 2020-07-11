package main

import (
	"github.com/andrelrg/go-chat/internal"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"log"
	"net/http"
)

func main() {

	cfg := jaegercfg.Configuration{
		ServiceName: "go-chat",
		Sampler:     &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter:    &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Fatal("failed to initialize tracer")
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	s := internal.NewServer()
	log.Println("waiting routes...")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

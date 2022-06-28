package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace:   "golang_test",
			Name:        "my_counter",
			Help:        "This is my counter",
			ConstLabels: prometheus.Labels{"group": "runtime", "tag": "metric1"},
		})

	counter2 = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace:   "golang",
			Name:        "test2_my_counter",
			Help:        "This is my counter",
			ConstLabels: prometheus.Labels{"group": "cnnf"},
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace:   "golang",
			Name:        "my_gauge",
			Help:        "This is my gauge",
			ConstLabels: prometheus.Labels{"group": "runtime", "tag": "metric2", "other": "other2"},
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace:   "golang",
			Name:        "my_histogram",
			Help:        "This is my histogram",
			Buckets:     []float64{1, 4, 7, 8, 11, 23, 50},
			ConstLabels: prometheus.Labels{"group": "runtime", "tag": "metric3"},
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace:   "golang",
			Name:        "my_summary",
			Help:        "This is my summary",
			Objectives:  map[float64]float64{0.1: 0.05, 0.3: 0.005, 0.6: 0.008, 0.9: 0.003},
			ConstLabels: prometheus.Labels{"group": "runtime"},
		})
)

func main() {
	rand.Seed(time.Now().Unix())

	//histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
	//	Name: "prom_request_time",
	//	Help: "Time it has taken to retrieve the metrics",
	//}, []string{"time"})
	//
	//prometheus.Register(histogramVec)
	//
	//http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))

	prometheus.MustRegister(counter)
	prometheus.MustRegister(counter2)

	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			counter2.Add(rand.Float64() * 20)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			time.Sleep(time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9876", nil))
}

func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		status := http.StatusOK

		defer func() {
			histogram.WithLabelValues(fmt.Sprintf("%d", status)).Observe(time.Since(start).Seconds())
		}()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
		status = http.StatusBadRequest

		w.WriteHeader(status)
	})
}

//ConstLabels are used to attach fixed labels to this metric. Metrics
//with the same fully-qualified name must have the same label names in
//their ConstLabels.
//
//Due to the way a Summary is represented in the Prometheus text format
//and how it is handled by the Prometheus server internally, “quantile”
//is an illegal label name. Construction of a Summary or SummaryVec
//will panic if this label name is used in ConstLabels.
//
//ConstLabels are only used rarely. In particular, do not use them to
//attach the same labels to all your metrics. Those use cases are
//better covered by target labels set by the scraping Prometheus
//server, or by one specific metric (e.g. a build_info or a
//machine_role metric). See also

//Objectives defines the quantile rank estimates with their respective
//absolute error. If Objectives[q] = e, then the value reported for q
//will be the φ-quantile value for some φ between q-e and q+e.  The
//default value is an empty map, resulting in a summary without
//quantiles.

//
//NewCounterFunc creates a new CounterFunc based on the provided
//CounterOpts. The value reported is determined by calling the given function
//from within the Write method. Take into account that metric collection may
//happen concurrently. If that results in concurrent calls to Write, like in
//the case where a CounterFunc is directly registered with Prometheus, the
//provided function must be concurrency-safe. The function should also honor
//the contract for a Counter (values only go up, not down), but compliance will
//not be checked.
//
//Check out the ExampleGaugeFunc examples for the similar GaugeFunc.

//Namespace, Subsystem, and Name are components of the fully-qualified
//name of the Metric (created by joining these components with
//"_"). Only Name is mandatory, the others merely help structuring the
//name. Note that the fully-qualified name of the metric must be a
//valid Prometheus metric name.

//Write encodes the Metric into a "Metric" Protocol Buffer data
//transmission object.
//
//Metric implementations must observe concurrency safety as reads of
//this metric may occur at any time, and any blocking occurs at the
//expense of total performance of rendering all registered
//metrics. Ideally, Metric implementations should support concurrent
//readers.
//
//While populating dto.Metric, it is the responsibility of the
//implementation to ensure validity of the Metric protobuf (like valid
//UTF-8 strings or syntactically valid metric and label names). It is
//recommended to sort labels lexicographically. Callers of Write should
//still make sure of sorting if they depend on it.
//
//TODO(beorn7): The original rationale of passing in a pre-allocated
//dto.Metric protobuf to save allocations has disappeared. The
//signature of this method should be changed to "Write() (*dto.Metric,
//error)".

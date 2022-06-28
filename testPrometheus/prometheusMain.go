package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var opsProcessed1 = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "nnn",
	Name:      "_2myapp__2processed_ops1_total",
	Help:      "helloworld1",
})
var opsProcessed2 = promauto.NewCounter(prometheus.CounterOpts{
	Name: "myapp_processed_ops2_total",
	Help: "helloworld2",
})
var opsProcessed3 = promauto.NewCounter(prometheus.CounterOpts{
	Name: "myapp_processed_ops3_total",
	Help: "helloworld3",
})

const PORT = 9961

func recordMetrics() {
	go func() {
		for {
			opsProcessed1.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			opsProcessed2.Add(2)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			opsProcessed3.Add(3)
			time.Sleep(1 * time.Second)
		}
	}()
}
func main() {
	registry1 := prometheus.NewRegistry()
	for i := 0; i < 100; i++ {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "nnn" + strconv.Itoa(i),
			Name:      "myname" + strconv.Itoa(i),
			Help:      "helloworld" + strconv.Itoa(i),
		})
		registry1.MustRegister(counter)
		go func() {
			time.Sleep(2 * time.Second)
			counter.Add(rand.Float64() * 5)
		}()
	}
	//registry1.MustRegister(opsProcessed1)
	http.Handle("/metrics1", promhttp.HandlerFor(registry1, promhttp.HandlerOpts{
		Registry: registry1,
	}))
	registry2 := prometheus.NewRegistry()
	registry2.MustRegister(opsProcessed2)
	http.Handle("/metrics2", promhttp.HandlerFor(registry2, promhttp.HandlerOpts{
		Registry: registry2,
	}))
	registry3 := prometheus.NewRegistry()
	registry3.MustRegister(opsProcessed3)
	http.Handle("/metrics3", promhttp.HandlerFor(registry3, promhttp.HandlerOpts{
		Registry: registry3,
	}))
	recordMetrics()
	//http.ListenAndServe(":9099",nil)

	http.HandleFunc("/", handler)
	fmt.Println("start listening 9961")
	//if err:=http.ListenAndServeTLS(":"+strconv.Itoa(PORT), "D:/localGoProject/src/golearning/testPrometheus/prom.crt", "D:/localGoProject/src/golearning/testPrometheus/prom.key", nil); err!=nil{
	//	log.Println(err.Error())
	//}
	//log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(PORT), "D:/localGoProject/src/golearning/testPrometheus/prom1.crt", "D:/localGoProject/src/golearning/testPrometheus/prom1.key", nil))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

//
//ListenAndServeTLS acts identically to ListenAndServe, except that it
//expects HTTPS connections. Additionally, files containing a certificate and
//matching private key for the server must be provided. If the certificate
//is signed by a certificate authority, the certFile should be the concatenation
//of the server's certificate, any intermediates, and the CA's certificate.
//
//
//ListenAndServeTLS listens on the TCP network address srv.Addr and
//then calls ServeTLS to handle requests on incoming TLS connections.
//Accepted connections are configured to enable TCP keep-alives.
//
//Filenames containing a certificate and matching private key for the
//server must be provided if neither the Server's TLSConfig.Certificates
//nor TLSConfig.GetCertificate are populated. If the certificate is
//signed by a certificate authority, the certFile should be the
//concatenation of the server's certificate, any intermediates, and
//the CA's certificate.
//
//If srv.Addr is blank, ":https" is used.
//
//ListenAndServeTLS always returns a non-nil error. After Shutdown or
//Close, the returned error is ErrServerClosed.
//
//
//ServeTLS accepts incoming HTTPS connections on the listener l,
//creating a new service goroutine for each. The service goroutines
//read requests and then call handler to reply to them.
//
//The handler is typically nil, in which case the DefaultServeMux is used.
//
//Additionally, files containing a certificate and matching private key
//for the server must be provided. If the certificate is signed by a
//certificate authority, the certFile should be the concatenation
//of the server's certificate, any intermediates, and the CA's certificate.
//
//ServeTLS always returns a non-nil error.
//
//
//ServeTLS accepts incoming connections on the Listener l, creating a
//new service goroutine for each. The service goroutines perform TLS
//setup and then read requests, calling srv.Handler to reply to them.
//
//Files containing a certificate and matching private key for the
//server must be provided if neither the Server's
//TLSConfig.Certificates nor TLSConfig.GetCertificate are populated.
//If the certificate is signed by a certificate authority, the
//certFile should be the concatenation of the server's certificate,
//any intermediates, and the CA's certificate.
//
//ServeTLS always returns a non-nil error. After Shutdown or Close, the
//returned error is ErrServerClosed.

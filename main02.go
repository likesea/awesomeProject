package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func main0() {
	req, err := http.NewRequest("GET", "https://doc.rust-lang.org/std/cell/struct.Ref.html", nil)
	if err != nil {
		log.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), 5000*time.Millisecond)
	defer cancel()
	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := http.Client{
		Transport: &tr,
	}
	ch := make(chan error, 1)
	go func() {
		log.Println("Stratiing request")
		resp, err := client.Do(req)
		if err != nil {
			ch <- err
			return
		}
		defer resp.Body.Close()
		//io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()
	select {
	case <-ctx.Done():
		log.Println("timeout,cancel work...")
		tr.CancelRequest(req)
		log.Println(<-ch)
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
	}
}

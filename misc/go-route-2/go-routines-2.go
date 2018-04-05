package mytest

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/test", handlerTest)
	http.HandleFunc("/testwait10s", handlerTest10s)
	http.HandleFunc("/testwait5s", handlerTest5s)
}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctx, _ := context.WithTimeout(c, 30*time.Second)

	ctx1, ctx1Cancel := context.WithCancel(ctx)
	ctx2, ctx2Cancel := context.WithCancel(ctx)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Infof(ctx1, "Go1 begin ...")
		client1 := urlfetch.Client(ctx1)
		_, err := client1.Get("http://APP_NAME.appspot.com/testwait5s")
		if err != nil {
			log.Errorf(ctx1, "Go1 failed:  %v", err)
		}
		ctx2Cancel()
		log.Infof(ctx1, "Go1 over ...")
	}()

	go func() {
		defer wg.Done()
		log.Infof(ctx2, "Go2 begin ...")
		client2 := urlfetch.Client(ctx2)
		_, err := client2.Get("http://APP_NAME.appspot.com/testwait10s")
		if err != nil {
			log.Errorf(ctx2, "Go2 failed %v", err)
		}
		ctx1Cancel()
		log.Infof(ctx2, "Go2 over ...")
	}()

	wg.Wait()
	log.Infof(ctx1, "Go1 and GO2 over")
}

func handlerTest10s(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	return
}

func handlerTest5s(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	return
}
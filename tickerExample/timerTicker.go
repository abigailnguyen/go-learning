package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// func (connMgr *ConnectionManager) Init() {
// 	cfg := connMgr.Config
// 	defer log.Infof("Connector initiated")

// 	if cfg.Server.Name == "" {
// 		hostname, _ := os.Hostname()
// 		selfConn.Name = fmt.Sprintf("Connector [%s]", hostname)
// 	}

// 	if cfg.Server.ID == "" {
// 		selfConn.Id = randomIDGenerator()()
// 	} else {
// 		selfConn.Id = cfg.Server.ID
// 	}

// 	cfg.Server.Name = selfConn.Name
// 	cfg.Server.ID = selfConn.Id

// 	err := cfg.Persist()
// 	if err != nil {
// 		log.Errorf("Error saving the connector configuration. Please contact PaperCut support.", err)
// 	}

// 	if !connMgr.CheckPrimary() {
// 		//go func() {
// 		//	defer log.Infof("Quit monitor goroutine")
// 		//	var count int
// 		//	for {
// 		//		prim := UPConnector{
// 		//			Host: connMgr.Config.Server.PrimaryIPAddr,
// 		//		}
// 		//		if err := selfConn.SubscribeToPrimaryConnector(prim, connMgr.GetAuthToken()); err != nil {
// 		//			log.Errorf("Failed to subscribe to primary connector: %v", err)
// 		//			if count++; count > 10 {
// 		//				return
// 		//			}
// 		//			pollInterval := math.Min(math.Pow(2, float64(count)) + rand.Float64() + 30, 60.0) // seconds
// 		//			time.Sleep(time.Duration(pollInterval) * time.Second)
// 		//		} else {
// 		//			time.Sleep(40 * time.Second)
// 		//		}
// 		//	}
// 		//} ()

// 		connMgr.Monitor()

// 		//log.Info("Started HttpServer in Secondary mode.")
// 		//defer func() {connMgr.isrunning = true}()
// 		//abort := make(chan bool)
// 		//connErr := make(chan bool, 1) // let goroutine decide when to handle
// 		//updateTicker := make(chan float64, 1)
// 		//defaultPollInterval := 35.0 // as it takes 30s for a timeout error, we give 5s in between intervals to not block the goroutine with the ticker channel
// 		////ticker := time.NewTicker(time.Duration(pollPeriod) * time.Second)
// 		//
// 		////go func() {
// 		////	defer log.Info("Goroutine 1 terminated!") // condition to terminate? by a command to restart server?
// 		////	defer close(connErr)
// 		////	time.Sleep(5 * time.Minute)
// 		////	abort <- true
// 		////	return
// 		////}()
// 		//
// 		//// start polling to primary connector
// 		//// go func(tk *time.Ticker) {
// 		//go func() {
// 		//	defer log.Info("Goroutine 2 terminated!")
// 		//	prim := UPConnector{
// 		//		Host: cfg.Server.PrimaryIPAddr,
// 		//	}
// 		//	f, err := os.OpenFile("/tmp/ticker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 		//	if err != nil {
// 		//		log.Errorf("failed to open file")
// 		//	}
// 		//	ticker := time.NewTicker(time.Duration(defaultPollInterval) * time.Second)
// 		//	if err := selfConn.SubscribeToPrimaryConnector(prim, connMgr.GetAuthToken()); err != nil {
// 		//		//if true {
// 		//		log.Errorf("Failed to subscribe to primary connector: %v", err)
// 		//		connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
// 		//	} else {
// 		//		connErr <- false
// 		//	}
// 		//	//ticker := time.NewTicker(time.Second)
// 		//	defer ticker.Stop()
// 		//	defer f.Close()
// 		//	log.Info("Monitor started! \n")
// 		//	for {
// 		//		select {
// 		//		case t := <-ticker.C:
// 		//			log.Info("Subscribe to Primary connector")
// 		//			f.WriteString(fmt.Sprintf("Ticker time %v\n", t))
// 		//			if err := selfConn.SubscribeToPrimaryConnector(prim, connMgr.GetAuthToken()); err != nil {
// 		//			//if true {
// 		//				log.Errorf("Failed to subscribe to primary connector: %v", err)
// 		//				connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
// 		//			} else {
// 		//				connErr <- false
// 		//			}
// 		//		case p := <-updateTicker:
// 		//			ticker = time.NewTicker(time.Duration(p) * time.Second)
// 		//			log.Infof("Update time ticker to %v seconds", p)
// 		//		case <-abort:
// 		//			close(connErr)
// 		//			log.Info("Finished routine!")
// 		//			return
// 		//		default:
// 		//		}
// 		//	}
// 		//}()
// 		////}(ticker)
// 		//
// 		////go func(tk *time.Ticker) {
// 		//go func() {
// 		//	defer log.Info("Goroutine 3 terminated!")
// 		//	var (
// 		//		count        int
// 		//		pollInterval float64
// 		//		//defaultPollInterval = 35.0
// 		//	)
// 		//	// implement exponential back-off
// 		//	for c := range connErr {
// 		//		if c {
// 		//			count++
// 		//			if count > 10 {
// 		//				abort <- true
// 		//				continue
// 		//			}
// 		//			if pollInterval == 60 {
// 		//				continue
// 		//			}
// 		//			pollInterval = math.Min(math.Pow(2, float64(count)) + rand.Float64() +defaultPollInterval, 60.0) // seconds
// 		//			//updateTicker <- pollInterval
// 		//
// 		//			//log.Infof("Update time ticker to %v seconds", pollInterval)
// 		//		} else {
// 		//			pollInterval = defaultPollInterval
// 		//			count = 0
// 		//			//updateTicker <- pollInterval
// 		//			continue
// 		//		}
// 		//		updateTicker <- pollInterval
// 		//	}
// 		//}()
// 		//}(ticker)
// 	} else {
// 		log.Info("Started HttpServer in Primary mode.")
// 	}
// }

// func (cm *ConnectionManager) Monitor() {
// 	if cm.isrunning {
// 		log.Info("Server is already running")
// 		return
// 	}

// 	log.Info("Started HttpServer in Secondary mode.")
// 	defer func() {cm.isrunning = true}()
// 	abort := make(chan bool)
// 	connErr := make(chan bool, 1) // let goroutine decide when to handle
// 	updateTicker := make(chan float64, 1)
// 	//defaultPollInterval := 35.0
// 	//ticker := time.NewTicker(time.Duration(pollPeriod) * time.Second)

// 	//go func() {
// 	//	defer log.Info("Goroutine 1 terminated!") // condition to terminate? by a command to restart server?
// 	//	defer close(connErr)
// 	//	time.Sleep(5 * time.Minute)
// 	//	abort <- true
// 	//	return
// 	//}()

// 	// start polling to primary connector
// 	// go func(tk *time.Ticker) {
// 	go func() {
// 		defer log.Info("Stopped monitor to primary connector")
// 		defer func() { cm.isrunning = false } ()
// 		prim := UPConnector{
// 			Host: cm.Config.Server.PrimaryIPAddr,
// 		}
// 		//f, err := os.OpenFile("/tmp/ticker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 		//if err != nil {
// 		//	log.Errorf("failed to open file")
// 		//}
// 		ticker := time.NewTicker(time.Duration(35.0) * time.Second)
// 		if err := selfConn.SubscribeToPrimaryConnector(prim, cm.GetAuthToken()); err != nil {
// 			//if true {
// 			log.Errorf("Failed to subscribe to primary connector: %v", err)
// 			connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
// 		} else {
// 			connErr <- false
// 		}
// 		//ticker := time.NewTicker(time.Second)
// 		defer ticker.Stop()
// 		//defer f.Close()
// 		log.Info("Monitor started! \n")
// 		for {
// 			select {
// 			case t := <-ticker.C:
// 				log.Info("Subscribe to Primary connector")
// 				//f.WriteString(fmt.Sprintf("Ticker time %v\n", t))
// 				log.Infof("Ticker time %v\n", t)
// 				if err := selfConn.SubscribeToPrimaryConnector(prim, cm.GetAuthToken()); err != nil {
// 				//if false {
// 					log.Errorf("Failed to subscribe to primary connector: %v", err)
// 					connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
// 					ticker.Stop()
// 				} else {
// 					connErr <- false
// 				}
// 				runtime.Gosched()
// 			case p := <-updateTicker:
// 				//ticker = time.NewTicker(time.Duration(p) * time.Second)
// 				log.Infof("Update time ticker to %v seconds", p)
// 				runtime.Gosched()
// 			case <-abort:
// 				close(connErr)
// 				log.Info("Finished routine!")
// 				return
// 			default:
// 			}
// 		}
// 	}()
// 	//}(ticker)

// 	//go func(tk *time.Ticker) {
// 	go func() {
// 		defer log.Info("Stopped monitor error checking")
// 		var (
// 			count        int
// 			pollInterval float64
// 			defaultPollInterval = 35.0
// 		)
// 		// implement exponential back-off
// 		for c := range connErr {
// 			if c {
// 				count++
// 				if count > 10 {
// 					abort <- true
// 					continue
// 				}
// 				if pollInterval == 60.0 {
// 					continue
// 				}
// 				//pollInterval = math.Min(math.Pow(2, float64(count)) + rand.Float64() + defaultPollInterval, 60.0) // seconds
// 				pollInterval = 45.0
// 				//updateTicker <- pollInterval

// 				//log.Infof("Update time ticker to %v seconds", pollInterval)
// 			} else {
// 				if pollInterval == defaultPollInterval {
// 					continue
// 				}
// 				pollInterval = defaultPollInterval
// 				count = 0
// 				//updateTicker <- pollInterval
// 			}
// 			updateTicker <- pollInterval
// 			runtime.Gosched()
// 		}
// 	}()
// }

// Back up 3

//log.Info("Started HttpServer in Secondary mode.")
//defer func() {connMgr.isrunning = true}()
//abort := make(chan bool)
//connErr := make(chan bool, 1) // let goroutine decide when to handle
//updateTicker := make(chan float64, 1)
//defaultPollInterval := 35.0 // as it takes 30s for a timeout error, we give 5s in between intervals to not block the goroutine with the ticker channel
////ticker := time.NewTicker(time.Duration(pollPeriod) * time.Second)
//
////go func() {
////	defer log.Info("Goroutine 1 terminated!") // condition to terminate? by a command to restart server?
////	defer close(connErr)
////	time.Sleep(5 * time.Minute)
////	abort <- true
////	return
////}()
//
//// start polling to primary connector
//// go func(tk *time.Ticker) {
//go func() {
//	defer log.Info("Goroutine 2 terminated!")
//	prim := UPConnector{
//		Host: cfg.Server.PrimaryIPAddr,
//	}
//	f, err := os.OpenFile("/tmp/ticker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Errorf("failed to open file")
//	}
//	ticker := time.NewTicker(time.Duration(defaultPollInterval) * time.Second)
//	if err := selfConn.SubscribeToPrimaryConnector(prim, connMgr.GetAuthToken()); err != nil {
//		//if true {
//		log.Errorf("Failed to subscribe to primary connector: %v", err)
//		connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
//	} else {
//		connErr <- false
//	}
//	//ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//	defer f.Close()
//	log.Info("Monitor started! \n")
//	for {
//		select {
//		case t := <-ticker.C:
//			log.Info("Subscribe to Primary connector")
//			f.WriteString(fmt.Sprintf("Ticker time %v\n", t))
//			if err := selfConn.SubscribeToPrimaryConnector(prim, connMgr.GetAuthToken()); err != nil {
//			//if true {
//				log.Errorf("Failed to subscribe to primary connector: %v", err)
//				connErr <- true // should not use the same goroutine to deal with the change because the goroutine can only handle 1 job at a time
//			} else {
//				connErr <- false
//			}
//		case p := <-updateTicker:
//			ticker = time.NewTicker(time.Duration(p) * time.Second)
//			log.Infof("Update time ticker to %v seconds", p)
//		case <-abort:
//			close(connErr)
//			log.Info("Finished routine!")
//			return
//		default:
//		}
//	}
//}()
////}(ticker)
//
////go func(tk *time.Ticker) {
//go func() {
//	defer log.Info("Goroutine 3 terminated!")
//	var (
//		count        int
//		pollInterval float64
//		//defaultPollInterval = 35.0
//	)
//	// implement exponential back-off
//	for c := range connErr {
//		if c {
//			count++
//			if count > 10 {
//				abort <- true
//				continue
//			}
//			if pollInterval == 60 {
//				continue
//			}
//			pollInterval = math.Min(math.Pow(2, float64(count)) + rand.Float64() +defaultPollInterval, 60.0) // seconds
//			//updateTicker <- pollInterval
//
//			//log.Infof("Update time ticker to %v seconds", pollInterval)
//		} else {
//			pollInterval = defaultPollInterval
//			count = 0
//			//updateTicker <- pollInterval
//			continue
//		}
//		updateTicker <- pollInterval
//	}
//}()
//}(ticker)

// func (cm *ConnectionManager) MonitorA() {
// 	if cm.monitorrunning {
// 		fmt.Println("Server is already running")
// 		return
// 	}
// 	defer func() {
// 		cm.monitorrunning = true
// 		fmt.Println("Started HttpServer in Secondary mode.")
// 	}()

// 	abort := make(chan bool)
// 	connErr := make(chan bool, 1)
// 	updateTicker := make(chan float64, 1)
// 	defaultPollInterval := 35.0

// 	// start polling to primary connector
// 	go func() {
// 		prim := UPConnector{Host: cm.Config.Server.PrimaryIPAddr}
// 		ticker := time.NewTicker(time.Duration(defaultPollInterval) * time.Second)

// 		defer func() {
// 			cm.monitorrunning = false
// 			ticker.Stop()
// 			log.Infof("Stopped monitor to primary connector")
// 		}()

// 		if err := selfConn.SubscribeToPrimaryConnector(prim, cm.GetAuthToken()); err != nil {
// 			log.Errorf("Failed to subscribe to primary connector: %v", err)
// 			connErr <- true
// 		}
// 		log.Info("Monitor started! \n")
// 		for {
// 			select {
// 			case t := <-ticker.C:
// 				log.Info("Subscribe to Primary connector")
// 				log.Infof("Ticker time %v\n", t)
// 				if err := selfConn.SubscribeToPrimaryConnector(prim, cm.GetAuthToken()); err != nil {
// 					log.Errorf("Failed to subscribe to primary connector: %v", err)
// 					connErr <- true
// 				} else {
// 					connErr <- false
// 				}
// 			case p := <-updateTicker:
// 				ticker.Stop()
// 				ticker = time.NewTicker(time.Duration(p) * time.Second)
// 				log.Infof("Update time ticker to %v seconds", p)
// 			case <-abort:
// 				close(connErr)
// 				log.Info("Finished routine!")
// 				return
// 			}
// 		}
// 	}()

// 	// error handling goroutine
// 	go func() {
// 		defer log.Info("Stopped monitor error checking")
// 		var (
// 			count        		int
// 			pollInterval 		float64
// 			defaultPollInterval = 35.0
// 			maxPollInterval 	= 300.0 // 5 minutes
// 		)
// 		// implement exponential back-off
// 		for c := range connErr {
// 			if c {
// 				count++
// 				if count > 10 {
// 					abort <- true
// 					continue
// 				}
// 				if pollInterval == maxPollInterval {
// 					continue
// 				}
// 				pollInterval = math.Min(math.Pow(2, float64(count)) + rand.Float64() + defaultPollInterval, maxPollInterval) // seconds
// 			} else {
// 				if pollInterval == defaultPollInterval {
// 					continue
// 				}
// 				pollInterval = defaultPollInterval
// 				count = 0
// 			}
// 			updateTicker <- pollInterval
// 		}
// 	}()
// }

// func test_() {
// 	defer fmt.Println("main() terminated.")
// 	ticker := time.NewTicker(2*time.Second)
// 	defer ticker.Stop()
// 	abort := make(chan bool)
// 	update := make(chan bool)
// 	var wg sync.WaitGroup
// 	timeoutCh := time.After(20 * time.Second)
// 	for {
// 		select {
// 		case t := <-ticker.C:
// 			fmt.Println("Time is %v", t)
// 		case <-timeoutCh:
// 			fmt.Println("Finished routine!")
// 			return
// 		}
// 	}

// abort := make(chan bool)
// update := make(chan bool)
// // var wg sync.WaitGroup

// go func() {
// 	defer fmt.Println("goroutine 1 terminted")
// 	time.Sleep(60 * time.Second)  // read a single byte, keeps listening on till the first input
// 	abort <- true
// }()

// go func() {
// 	defer fmt.Println("goroutine 2 terminated")
// 	time.Sleep(5 * time.Second)
// 	update <- true
// }()

// wg.Add(1)

// 	go func() {
// 		defer fmt.Println("goroutine 3 terminated")
// 		fmt.Printf("Started\n")
// 		ticker := time.NewTicker(time.Second)
// 		f, err := os.OpenFile("ticker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 		check(err)
// 		defer ticker.Stop()
// 		defer f.Close()
// 		// defer wg.Done()
// 		for { // without for loop the select will only receive from the ticker once
// 			select {

// 			case <-update:
// 				ticker.Stop()
// 				ticker = time.NewTicker(3 * time.Second)
// 			case <-ticker.C:
// 				// fmt.Println("timer ticked - %v", t)
// 				// writeToFile(f, fmt.Sprintf("timer ticked - %v", t))
// 				// f.WriteString(fmt.Sprintf("timer ticked - %v\n", t))
// 				continue
// 			case <-abort:
// 				fmt.Println("Aborted!")
// 				return // return exit the goroutine
// 			// default:
// 			// 	continue
// 			}
// 		}
// 	}()

// 	go func() {
// 		select {
// 		case <-abort: // if receive an abort signal then stop looping
// 			fmt.Println("Launch aborted!")
// 			wg.Done()
// 			return   // when we return then exit and quit goroutine
// 		default:
// 		}
// 	}()
// 	wg.Wait()
// 	<-abort  // wait for next abort, the main will get the signal
// 	close(abort) // send signal to abort, if we don't send this, there will be a goroutine leak for goroutine 3
// }

// func writeToFilet(f *os.File, s string) {
// 	_, err := f.WriteString(s + "\n")
// 	check(err)
// }

// func checkt(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// 	ticker := time.NewTicker(2*time.Second)
// defer ticker.Stop()

// timeoutCh := time.After(20 * time.Second)
// for {
// 	select {
// 	case t := timeoutCh:
// 		fmt.Println("Time is %v", t)
// 	case <-timeoutCh:
// 		fmt.Println("Finished routine!")
// 		return
// 	}
// }

// func TestGoroutine(t *testing.T) {
// 	tick := time.NewTicker(2 * time.Second) // a channel that sends a signal every 5 minutes
// 	abort := make(chan struct{})
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1))  // read a single byte, keeps listening on till the first input
// 		close(abort)                    // as soon as there is an input, it will start to close the abort
// 	}()
// 	go func() {
// 		select {
// 		case <-tick.C:
// 			fmt.Printf("timer ticked - %v", time.Now())
// 			// close(abort)
// 		default:
// 		}
// 	}()

// 	go func() {
// 		select {
// 		case <-abort: // if receive an abort signal then stop looping
// 			tick.Stop()
// 			wg.Done()
// 			return   // when we return then exit and quit goroutine
// 		default:
// 		}
// 	}()
// 	wg.Wait()
// }

// func test() {
// 	tick := time.NewTicker(5 * time.Minute) // a channel that sends a signal every 5 minutes
// 	abort := make(chan struct{})
// 	go func() {
// 		primConn := UPConnector{
// 			Host: cfg.Server.PrimaryIPAddr,
// 		}
// 		select {
// 		case <-tick.C:
// 			if err := selfConn.SubscribeToPrimaryConnector(primConn, connMgr.GetAuthToken()); err != nil {
// 				log.Errorf("Failed to subscribe to primary connector: %v", err)
// 				close(abort)
// 			}
// 		default:
// 		}
// 	}()

// 	go func() {
// 		select {
// 		case <-abort: // if receive an abort signal then stop looping
// 			tick.Stop()
// 			return // when we return then exit and quit goroutine
// 		default:
// 		}
// 	}()

// 	// will return here on the main function
// 	// no other task requires to wait for goroutine to complete, so safe to return
// }

func main() {
	nestedTicker()
}
func tickerExample() {
	firstTime := 5
	t := time.NewTicker(time.Duration(firstTime) * time.Second)
	timeout := time.After(1 * time.Minute)
	for {
		select {
		case <-t.C:
			fmt.Fprintf(os.Stdout, "time is %v\n", time.Now())
			t.Stop() // you need to stop before reset.
			firstTime += 10
			t.Reset(time.Duration(firstTime) * time.Second)
		case <-timeout:
			return
		}
	}
}

func nestedTicker() {
	fmt.Println("Started monitor")
	maxPollInterval := 2048.0
	defaultTicker := time.NewTicker(1 * time.Second)
	maxCount := 3
	defaultInterval := 30 * time.Second

M:
	for {
		select {
		case <-defaultTicker.C:
			defaultTicker.Stop()
			if err := PerformATask(errors.New("error")); err != nil {
				fmt.Fprintf(os.Stdout, "Some task failed to perform %v \n", err)
				timeout := time.After(5 * time.Minute)
				retryTicker := time.NewTicker(5 * time.Second)
				count := 1
			L:
				for {
					select {
					case <-timeout:
						fmt.Fprintf(os.Stdout, "timeout exceeded, exit retry")
						return
					case <-retryTicker.C:
						retryTicker.Stop() // stop so that we can reset later
						if count > maxCount {
							fmt.Println("Cancelling. Max attempts reached.")
							break M
						}
						fmt.Fprintf(os.Stdout, "Retry again. Count %d\n", count)
						err = PerformATask(errors.New("test"))
						// err = PerformATask(nil)
						if err != nil {
							pollInterval := math.Min(math.Pow(2, float64(count))+rand.Float64(), maxPollInterval)
							retryTicker.Reset(time.Duration(pollInterval) * time.Second) // wait and retry again
							count++
						} else {
							fmt.Fprintf(os.Stdout, "Resume normal interval \n")
							defaultTicker.Reset(defaultInterval)
							break L
						}
					}
				}
			}

		}
	}

}

func PerformATask(err error) error {
	return err
}

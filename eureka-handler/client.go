package eurekaHandler

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/GreatLaboratory/go-sms/util"
)

type Client struct {
	// for monitor system signal
	signalChan   chan os.Signal
	mutex        sync.RWMutex
	Running      bool
	Config       *Config
	Applications *Applications
}

func (c *Client) Start() error {
	c.Running = true

	if err := c.doRegister(); err != nil {
		log.Println(err.Error())
		return errors.New("client registration failed")
	}
	log.Println("register application instance successful")

	go c.refresh()
	go c.heartbeat()
	go c.handleSignal()

	return nil
}

func (c *Client) refresh() {
	for {
		if !c.Running {
			break
		}
		if err := c.doRefresh(); err != nil {
			log.Println(err)
		} else {
			log.Println("refresh application instance successful")
		}
		sleep := time.Duration(c.Config.RegistryFetchIntervalSeconds)
		time.Sleep(sleep * time.Second)
	}
}

func (c *Client) heartbeat() {
	for {
		if !c.Running {
			break
		}
		if err := c.doHeartbeat(); err != nil {
			if err == ErrNotFound {
				log.Println("heartbeat Not Found, need register")

				err = c.doRegister()

				if typeErr := util.IsErrorType(err); typeErr != nil {
					log.Fatal(typeErr)
				}

				if err != nil {
					log.Printf("do register error: %s\n", err)
				}
				continue
			}
			log.Println(err)
		} else {
			log.Println("heartbeat application instance successful")
		}
		sleep := time.Duration(c.Config.RenewalIntervalInSecs)
		time.Sleep(sleep * time.Second)
	}
}

func (c *Client) doRegister() error {
	instance := c.Config.instance
	return Register(c.Config.DefaultZone, c.Config.App, instance)
}

func (c *Client) doUnRegister() error {
	instance := c.Config.instance
	return UnRegister(c.Config.DefaultZone, instance.App, instance.InstanceID)
}

func (c *Client) doHeartbeat() error {
	instance := c.Config.instance
	return Heartbeat(c.Config.DefaultZone, instance.App, instance.InstanceID)
}

func (c *Client) doRefresh() error {
	// get all applications
	applications, err := Refresh(c.Config.DefaultZone)
	if err != nil {
		return err
	}

	// set applications
	c.mutex.Lock()
	c.Applications = applications
	c.mutex.Unlock()
	return nil
}

func (c *Client) handleSignal() {
	if c.signalChan == nil {
		c.signalChan = make(chan os.Signal)
	}
	signal.Notify(c.signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		switch <-c.signalChan {
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGKILL:
			fallthrough
		case syscall.SIGTERM:
			log.Println("receive exit signal, client instance going to de-register")
			err := c.doUnRegister()
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println("unRegister application instance successful")
			}
			os.Exit(0)
		}
	}
}

func NewClient(config *Config) *Client {
	defaultConfig(config)
	ip := os.Getenv("HOST_IP")
	if ip == "" {
		ip = getLocalIP()
	}
	config.instance = NewInstance(ip, config)
	return &Client{Config: config}
}

func defaultConfig(config *Config) {
	if config.RenewalIntervalInSecs == 0 {
		config.RenewalIntervalInSecs = 30
	}
	if config.RegistryFetchIntervalSeconds == 0 {
		config.RegistryFetchIntervalSeconds = 15
	}
	if config.DurationInSecs == 0 {
		config.DurationInSecs = 90
	}
}

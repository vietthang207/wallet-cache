package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	persister "github.com/KyberNetwork/server-go/persister"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

const (
	MAX_PAGE_SIZE = 50
	DEFAULT_PAGE  = 1
)

type HTTPServer struct {
	persister persister.Persister
	host      string
	r         *gin.Engine
}

func (self *HTTPServer) GetRate(c *gin.Context) {
	rates := self.persister.GetRate()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": rates},
	)
	return
}

// func (self *HTTPServer) GetEvent(c *gin.Context) {
// 	if !self.persister.GetIsNewEvent() {
// 		c.JSON(
// 			http.StatusOK,
// 			gin.H{"success": false},
// 		)
// 		return
// 	}

// 	events := self.persister.GetEvent()
// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{"success": true, "data": events},
// 	)
// }

func (self *HTTPServer) GetLatestBlock(c *gin.Context) {
	if !self.persister.GetIsNewLatestBlock() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}
	blockNum := self.persister.GetLatestBlock()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": blockNum},
	)
}

func (self *HTTPServer) GetRateUSD(c *gin.Context) {
	if !self.persister.GetIsNewRateUSD() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}

	rates := self.persister.GetRateUSD()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": rates},
	)
}

func (self *HTTPServer) GetRateETH(c *gin.Context) {
	if !self.persister.GetIsNewRateUSD() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}

	ethRate := self.persister.GetRateETH()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": ethRate},
	)
}

func (self *HTTPServer) GetKyberEnabled(c *gin.Context) {
	if !self.persister.GetNewKyberEnabled() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}

	enabled := self.persister.GetKyberEnabled()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": enabled},
	)
}

func (self *HTTPServer) GetMaxGasPrice(c *gin.Context) {
	if !self.persister.GetNewMaxGasPrice() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}

	gasPrice := self.persister.GetMaxGasPrice()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": gasPrice},
	)
}

func (self *HTTPServer) GetGasPrice(c *gin.Context) {
	if !self.persister.GetNewGasPrice() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false},
		)
		return
	}

	gasPrice := self.persister.GetGasPrice()
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": gasPrice},
	)
}

// func (self *HTTPServer) GetTokenInfo(c *gin.Context) {
// 	tokenInfo := self.persister.GetTokenInfo()
// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{"success": true, "data": tokenInfo},
// 	)
// }

func (self *HTTPServer) GetErrorLog(c *gin.Context) {
	dat, err := ioutil.ReadFile("error.log")
	if err != nil {
		log.Print(err)
		c.JSON(
			http.StatusOK,
			gin.H{"success": false, "data": err},
		)
	}
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": string(dat[:])},
	)
}

func (self *HTTPServer) GetMarketInfo(c *gin.Context) {
	data := self.persister.GetMarketData()
	if self.persister.GetIsNewMarketInfo() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": true, "data": data, "status": "latest"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": data, "status": "old"},
	)
}

func (self *HTTPServer) GetRightMarketInfo(c *gin.Context) {
	data := self.persister.GetRightMarketData()
	if self.persister.GetIsNewMarketInfo() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": true, "data": data, "status": "latest"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": data, "status": "old"},
	)
}

func (self *HTTPServer) GetLast7D(c *gin.Context) {
	listTokens := c.Query("listToken")
	data := self.persister.GetLast7D(listTokens)
	if self.persister.GetIsNewMarketInfo() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": true, "data": data, "status": "latest"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": data, "status": "old"},
	)
}

func (self *HTTPServer) GetMarketInfoByTokens(c *gin.Context) {
	listTokens := c.Query("listToken")
	data := self.persister.GetMarketDataByTokens(listTokens)
	if self.persister.GetIsNewMarketInfo() {
		c.JSON(
			http.StatusOK,
			gin.H{"success": true, "data": data, "status": "latest"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"success": true, "data": data, "status": "old"},
	)
}

// func (self *HTTPServer) GetLanguagePack(c *gin.Context) {
// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{"success": true, "data": "get language pack"},
// 	)
// 	return
// }

func (self *HTTPServer) Run() {
	//self.r.GET("/getRate", self.GetRate)
	// self.r.GET("/getHistoryOneColumn", self.GetEvent)
	self.r.GET("/getLatestBlock", self.GetLatestBlock)

	self.r.GET("/getRateUSD", self.GetRateUSD)
	self.r.GET("/getRate", self.GetRate)
	// self.r.GET("/getTokenInfo", self.GetTokenInfo)

	self.r.GET("/getKyberEnabled", self.GetKyberEnabled)
	self.r.GET("/getMaxGasPrice", self.GetMaxGasPrice)
	self.r.GET("/getGasPrice", self.GetGasPrice)
	self.r.GET("/getMarketInfo", self.GetMarketInfo)
	self.r.GET("/getRightMarketInfo", self.GetRightMarketInfo)
	self.r.GET("/getLast7D", self.GetLast7D)
	self.r.GET("/getMarketInfoByTokens", self.GetMarketInfoByTokens)
	self.r.GET("/getRateETH", self.GetRateETH)

	//self.r.GET("/getLanguagePack", self.GetLanguagePack)
	if os.Getenv("KYBER_ENV") != "production" {
		self.r.GET("/9d74529bc6c25401a2f984ccc9b0b2b3", self.GetErrorLog)
	}

	self.r.Run(self.host)
}

func NewHTTPServer(host string, persister persister.Persister) *HTTPServer {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Use(cors.Default())

	return &HTTPServer{
		persister, host, r,
	}
}

// Rest Runtime Server for CB-Contributhon Edu.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//      * CB-Contributhon: https://github.com/cb-contributhon
//
// by powerkim, 2020.08.

package main

import (
        "fmt"

        // for log print
        "github.com/sirupsen/logrus"
        "github.com/cloud-barista/cb-log"

        // for REST Server (echo)
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        "net/http"
)

var logger *logrus.Logger
func init() {
        // cblog is a global variable.
        logger = cblog.GetLogger("CB-Contributhon")
}

// REST API Return struct for Test
//====================================================================
type MyInfo struct {
        Name string `json:"name"`
        Github string `json:"github"`
}

// ex) {"POST", "/driver", registerCloudDriver}
type route struct {
        method, path string
        function     echo.HandlerFunc
}

func main() {

        //======================================= setup routes
        routes := []route{

                //----------MyInfo
                {"GET", "/2020/myinfo", getMyInfo},
                /*
                {"POST", "/driver", registerCloudDriver},
                {"GET", "/driver", listCloudDriver},
                {"GET", "/driver/:DriverName", getCloudDriver},
                {"DELETE", "/driver/:DriverName", unRegisterCloudDriver},
                */
        }
        //======================================= setup routes

        // Run REST Server
        ApiServer(routes, ":8080")

        fmt.Println("\n[CB-Contributhon:Test REST Framework]")
        fmt.Println("\n   Initiialized REST Server....__^..^__....\n\n")
}

//================ REST Server: setup & start
func ApiServer(routes []route, strPort string) {
        e := echo.New()

        // Middleware
        e.Use(middleware.CORS())
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        for _, route := range routes {
                switch route.method {
                case "POST":
                        e.POST(route.path, route.function)
                case "GET":
                        e.GET(route.path, route.function)
                case "PUT":
                        e.PUT(route.path, route.function)
                case "DELETE":
                        e.DELETE(route.path, route.function)

                }
        }

        e.HideBanner = true
        if strPort == "" {
                strPort = ":8080"
        }
        e.Logger.Fatal(e.Start(strPort))
}

//================ Get Myinfo Service
func getMyInfo(c echo.Context) error {
        logger.Info("call getMhyInfo()")

        myInfo := MyInfo {"powerkim", "powerkimhub"}
        return c.JSON(http.StatusOK, &myInfo)
}

package riwayat

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	rwtctrl "ms-simaba-riwayat-donor/controller/riwayat"
	rwtmdl "ms-simaba-riwayat-donor/models/riwayat"

	"github.com/gin-gonic/gin"
)

const modul = "jadwal donor"

func RiwayatHandler(c *gin.Context) {
	var param rwtmdl.RiwayatRequest
	c.BindJSON(&param)

	if _, err := strconv.Atoi(param.KodePendonor); err != nil {
		result, err := rwtctrl.GetRiwayat(param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(http.StatusOK, result)
		} else {
			result.Code = 500
			result.Message = "internal server error"
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": fmt.Sprintf("kode %s must be a string", modul)})
	}
}

func RiwayatDetailHandler(c *gin.Context) {
	var param rwtmdl.RiwayatDetailRequest
	c.BindJSON(&param)

	if param.Ktp != "" && param.KuesionerId != "" {
		result, err := rwtctrl.GetRiwayatDetail(param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(http.StatusOK, result)
		} else {
			result.Code = 500
			result.Message = "internal server error"
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "ktp and kuesioner id cannot empty"})
	}

}

func CreateRiwayatHandler(c *gin.Context) {
	var param rwtmdl.RiwayatCreateRequest
	c.BindJSON(&param)

	var JWT, _ = c.Get("JWT_CLAIMS")
	jwtClaims := reflect.ValueOf(JWT).Elem()
	param.Role = jwtClaims.FieldByName("Role").Interface().(string)
	if _, err := strconv.Atoi(param.KodePendonor); err != nil && param.Role == "admin" {
		err = rwtctrl.CreateRiwayat(param)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success create data"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 500, "message": err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "message": "this account doesn't have access to create data"})
	}
}

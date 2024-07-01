package payment

import (
	"MelkOnline/internal/controller"
	"MelkOnline/internal/core/signup/payment"
	"os"
	"strconv"
	"strings"

	"github.com/hooklift/gowsdl/soap"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	ps payment.PaymentServiceContract
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		ps: payment.NewPaymentService(),
	}
}

func (ph *PaymentHandler) Pay(c echo.Context) error {
	var (
		TerminalID = os.Getenv("BANK_TERMINAL_ID")
		UserName   = os.Getenv("BANK_USERNAME")
		Password   = os.Getenv("BANK_PASSWORD")
		pReq       = &controller.PaymentRequest{}
	)
	if err := c.Bind(pReq); err != nil {
		return c.JSON(400, err)
	}
	TerminalIDint, err := strconv.Atoi(TerminalID)
	if err != nil {
		return c.JSON(400, err)
	}
	client := soap.NewClient("https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/services/pgw?wsdl=IPaymentGateway.wsdl")
	TerminalIDint64, ld, lt, ad, curl := int64(TerminalIDint), "14030411", "161231", "", "https://bing.com"
	oid, err := ph.ps.GetLastID()
	if err != nil {
		return c.JSON(400, err)
	}
	oid64, amount, pid := int64(oid), int64(pReq.Amount), int64(pReq.UserID)
	oid64++
	gw := payment.NewIPaymentGateway(client)
	req := payment.BpPayRequest{
		TerminalId:     TerminalIDint64,
		UserName:       UserName,
		UserPassword:   Password,
		OrderId:        10,
		Amount:         amount,
		LocalDate:      ld,
		LocalTime:      lt,
		AdditionalData: ad,
		CallBackUrl:    curl,
		PayerId:        pid,
	}
	resp, err := gw.BpPayRequest(&req)
	if err != nil {
		return c.JSON(400, err)
	}
	data := strings.Split(resp.Return_, ",")
	ph.ps.Pay(pReq.UserID, pReq.Amount, data[0])
	if data[0] != "0" {
		return c.JSON(400, resp.Return_)
	}
	token := data[1]
	return c.Redirect(302, "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/startpay.mellat?RefId="+token)
}

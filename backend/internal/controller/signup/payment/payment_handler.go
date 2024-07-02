package payment

import (
	"MelkOnline/internal/core/signup/payment"
	"MelkOnline/internal/infrastructure/auth"
	"os"
	"strconv"
	"strings"

	"github.com/hooklift/gowsdl/soap"
	"github.com/labstack/echo/v4"
)

const PAYMENT_AMOUNT = int64(200000)

type PaymentHandler struct {
	ps payment.PaymentServiceContract
	sr *auth.SessionRepository
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		ps: payment.NewPaymentService(),
		sr: auth.NewSessionRepository(),
	}
}

// Pay	 pay
//
//	@Summary		pay
//	@Description	pay
//	@Tags			payment
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		400	{string}	string
//	@Router			api/v1/signup/payment/ [get]

func (ph *PaymentHandler) Pay(c echo.Context) error {
	var (
		TerminalID = os.Getenv("BANK_TERMINAL_ID")
		UserName   = os.Getenv("BANK_USERNAME")
		Password   = os.Getenv("BANK_PASSWORD")
	)
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
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, err)
	}
	UserIDint, err := ph.sr.ValidateSession(token.Value)
	if err != nil {
		return c.JSON(400, err)
	}
	oid64, pid := int64(oid), int64(UserIDint)
	oid64++
	gw := payment.NewIPaymentGateway(client)
	req := payment.BpPayRequest{
		TerminalId:     TerminalIDint64,
		UserName:       UserName,
		UserPassword:   Password,
		OrderId:        10,
		Amount:         PAYMENT_AMOUNT,
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
	ph.ps.Pay(UserIDint, int(PAYMENT_AMOUNT), data[0])
	if data[0] != "0" {
		return c.JSON(400, resp.Return_)
	}
	ref := data[1]
	return c.Redirect(302, "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/startpay.mellat?RefId="+ref)
}

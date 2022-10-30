package tcpservice

import "fmt"

const (
	TCP_OPERATION_VERIFY = "verify"
)

func (h *TCPService) VerifyHandler(req *Request, res *Response) {
	algo, err := h.findAlgo(req)
	if err != nil {
		h.c.Logger().Error("verifyHandler:error finding algorithm %v", err)
		res.SetError(fmt.Sprintf("error finding algorithm %v", err))
		if err = res.Send(); err != nil {
			h.c.Logger().Error("verifyHandler:send failed %v", err)
		}
		return
	}
	verified := algo.Verify(req.Hash, req.Password)
	res.SetVerified(verified)
	if err = res.Send(); err != nil {
		h.c.Logger().Error("verifyHandler:send failed %v", err)
	}
}

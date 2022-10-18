package tcpservice

import "fmt"

const (
	TCP_OPERATION_VERIFY = "verify"
)

func (h *TCPService) VerifyHandler(req Request, res Response) {
	algo, err := h.findAlgo(req)
	if err != nil {
		h.c.Logger().Error("error finding algorithm %v", err)
		res.SetError(fmt.Sprintf("error finding algorithm %v", err))
		if err = res.Send(); err != nil {
			h.c.Logger().Error("send failed %v", err)
		}
		return
	}
	verified := algo.Verify(req.hash, req.password)
	res.SetVerified(verified)
	if err = res.Send(); err != nil {
		h.c.Logger().Error("send failed %v", err)
	}
}

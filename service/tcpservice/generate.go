package tcpservice

import "fmt"

const (
	TCP_OPERATION_GENERATION = "generate"
)

func (h *TCPService) GenerateHandler(req Request, res Response) {
	algo, err := h.findAlgo(req)
	if err != nil {
		h.c.Logger().Error("error finding algorithm %v", err)
		res.SetError(fmt.Sprintf("error finding algorithm %v", err))
		if err = res.Send(); err != nil {
			h.c.Logger().Error("send failed %v", err)
		}
		return
	}
	hash, err := algo.Generate(req.password)
	if err != nil {
		h.c.Logger().Error("error generating hash %v", err)
		res.SetError(fmt.Sprintf("error generating hash %v", err))
		if err = res.Send(); err != nil {
			h.c.Logger().Error("send failed %v", err)
		}
		return
	}
	res.SetHash(hash)
	if err = res.Send(); err != nil {
		h.c.Logger().Error("send failed %v", err)
	}
}

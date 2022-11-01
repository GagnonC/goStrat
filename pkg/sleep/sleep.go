package sleep

import (
	internalTrans "awesomeProject/internal/translator"
	"awesomeProject/pkg/translator"
)

type ISleep interface {
	Sleep(incoming <-chan translator.ITranslator) error
}

type Sleep struct {
	translator *internalTrans.TerrierTranslator
}

func (s *Sleep) Sleep(incoming <-chan translator.ITranslator) error {
	var err error
	for t := range incoming {
		t.Translate()
		if err != nil {
			return err
		}
		go func(t translator.ITranslator) error {
			err = t.SendSuccessCallback()
			if err != nil {
				return err
			}
			return err
		}(t)
	}
	return err
}

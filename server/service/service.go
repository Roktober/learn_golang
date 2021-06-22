package service

import "server/text/processor"

type Service struct {
	TextProcessor processor.TextProcessor
}

func (s *Service) ProcessText(text string) {
	s.TextProcessor.ProcessText(text)
}

func (s *Service) GetTop(count int) string {
	return s.TextProcessor.Top(count)
}

package event

import (
    "aws_rekognition_demo/internal/domain"
    "aws_rekognition_demo/internal/domain/event"
    log "github.com/sirupsen/logrus"
)

type ServiceImpl struct {
    eventRepo event.Repository
}

func NewEventService(eventRepo event.Repository) *ServiceImpl {
    s := new(ServiceImpl)
    s.eventRepo = eventRepo
    return s
}

func (s *ServiceImpl) Search(page domain.Page) ([]event.Event, error) {
    return s.eventRepo.FindAll(page)
}

func (s *ServiceImpl) Create(event *event.CreateEventInputDto) (int, error) {
    return s.eventRepo.Store(event)
}

func (s *ServiceImpl) DeleteEvent(eventId int) error {
    err := s.eventRepo.CheckExist(eventId)
    if err != nil {
        log.Error(err)
        return err
    }
    return s.eventRepo.Delete(eventId)
}

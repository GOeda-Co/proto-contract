package convert

import (
	"fmt"

	cardv1 "github.com/GOeda-Co/proto-contract/gen/go/card"
	deckv1 "github.com/GOeda-Co/proto-contract/gen/go/deck"
	"github.com/google/uuid"

	model "github.com/GOeda-Co/proto-contract/model/card"
	modelDeck "github.com/GOeda-Co/proto-contract/model/deck"

	schemes "github.com/GOeda-Co/proto-contract/scheme/card"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoToModelCard(card *cardv1.Card) (*model.Card, error) {
	cardId, err := uuid.Parse(card.CardId)
	if err != nil {
		return nil, fmt.Errorf("cardId is invalid: %w", err)
	}
	createdBy, err := uuid.Parse(card.CreatedBy)
	if err != nil {
		return nil, fmt.Errorf("createdBy is invalid: %w", err)
	}
	deckId, err := uuid.Parse(card.DeckId)
	if err != nil {
		return nil, fmt.Errorf("DeckId is invalid: %w", err)
	}

	return &model.Card{
		CardId:           cardId,
		CreatedBy:        createdBy,
		CreatedAt:        card.CreatedAt.AsTime(),
		Word:             card.Word,
		Translation:      card.Translation,
		Easiness:         card.Easiness,
		UpdatedAt:        card.UpdatedAt.AsTime(),
		Interval:         int(card.Interval),
		ExpiresAt:        card.ExpiresAt.AsTime(),
		RepetitionNumber: int(card.RepetitionNumber),
		DeckID:           deckId,
		Tags:             card.Tags,
		IsPublic:         card.IsPublic,
	}, nil
}

func FromModelToProtoCard(card *model.Card) *cardv1.Card {
	return &cardv1.Card{
		CardId:           card.CardId.String(),
		CreatedBy:        card.CreatedBy.String(),
		CreatedAt:        timestamppb.New(card.CreatedAt),
		Word:             card.Word,
		Translation:      card.Translation,
		Easiness:         card.Easiness,
		UpdatedAt:        timestamppb.New(card.UpdatedAt),
		Interval:         int32(card.Interval),
		ExpiresAt:        timestamppb.New(card.ExpiresAt),
		RepetitionNumber: int32(card.RepetitionNumber),
		DeckId:           card.DeckID.String(),
		Tags:             card.Tags,
		IsPublic:         card.IsPublic,
	}
}

func FromProtoToUpdateSchemeCard(card *cardv1.UpdateCardRequest) *schemes.UpdateCardScheme {
	return &schemes.UpdateCardScheme{
		Word:             card.Word,
		Translation:      card.Translation,
		Easiness:         card.Easiness,
		Interval:         int(card.Interval),
		ExpiresAt:        card.ExpiresAt.AsTime(),
		RepetitionNumber: int(card.RepetitionNumber),
		Tags:             card.Tags,
	}
}

func FromProtoToAnswerSchemeCard(answer *cardv1.Answer) (*schemes.AnswerScheme, error) {
	cardId, err := uuid.Parse(answer.CardId)
	if err != nil {
		return nil, err
	}
	return &schemes.AnswerScheme{
		CardId: cardId,
		Grade:  int(answer.Grade),
	}, nil
}

func FromAnswerSchemeToProtoCard(answer *schemes.AnswerScheme) (*cardv1.Answer, error) {
	cardId, err := uuid.Parse(answer.CardId.String())
	if err != nil {
		return nil, err
	}
	return &cardv1.Answer{
		CardId: cardId.String(),
		Grade:  int32(answer.Grade),
	}, nil
}

func FromAnswerSchemesToProtosCard(answers []*schemes.AnswerScheme) ([]*cardv1.Answer, error) {
	var result []*cardv1.Answer
	for _, answer := range answers {
		converted, err := FromAnswerSchemeToProtoCard(answer)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}
	return result, nil
}

func FromProtoToModelDeck(deck *deckv1.Deck) (*modelDeck.Deck, error) {
	deckId, err := uuid.Parse(deck.DeckId)
	if err != nil {
		return nil, err
	}
	createdBy, err := uuid.Parse(deck.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &modelDeck.Deck{
		DeckId:        deckId,
		CreatedBy:     createdBy,
		CreatedAt:     deck.CreatedAt.AsTime(),
		Name:          deck.Name,
		CardsQuantity: uint(deck.CardsQuantity),
		Description:   deck.Description,
		IsPublic:      deck.IsPublic,
	}, nil
}
func FromModelToProtoDeck(deck *modelDeck.Deck) *deckv1.Deck {
	return &deckv1.Deck{
		DeckId:        deck.DeckId.String(),
		CreatedBy:     deck.CreatedBy.String(),
		CreatedAt:     timestamppb.New(deck.CreatedAt),
		Name:          deck.Name,
		Description:   deck.Description,
		CardsQuantity: uint32(deck.CardsQuantity),
		IsPublic:      deck.IsPublic,
	}
}
